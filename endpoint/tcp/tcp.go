package tcp

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"flytcp/common/collect"
	"flytcp/common/gort"
	"flytcp/common/logs"
	"flytcp/common/pool"
	"flytcp/config"
	"flytcp/endpoint/dispatch"
	"flytcp/endpoint/protocol"
	"fmt"
	"github.com/golang/protobuf/proto"
	cmap "github.com/orcaman/concurrent-map/v2"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var (
	acceptWait = 5 * time.Millisecond
	bufPool    = pool.NewPool("tcp.buffer", func() interface{} {
		return &bytes.Buffer{}
	})
)

const (
	bufSize            = 1024 * 100
	acceptTimeout      = 3 * time.Second
	idleTimeout        = 60 * time.Second
	sendQueueSize      = 256
	REQ           byte = 0
	ACK           byte = 1
	reqTimeout         = 60 * time.Second
)

type TcpEndpoint struct {
	ctx     context.Context
	cf      context.CancelFunc
	d       *dispatch.Dispatcher
	fc      collect.ConcurrentMap[string, *tcpCtx] // connections from client
	fs      collect.ConcurrentMap[string, *tcpCtx] //connections from server
	ts      cmap.ConcurrentMap[string, *tcpCtx]    //connections to server
	tcmLock sync.RWMutex
	l       *net.TCPListener
}

func (t *TcpEndpoint) NTF(msg protocol.Msg) error {
	t.tcmLock.RLock()
	defer t.tcmLock.RUnlock()
	uid := msg.H.ID
	if uid == "" {
		return errors.New("pls specify id")
	}
	ctx, ok := t.fc.Get(uid)
	if !ok {
		return fmt.Errorf("can not find ctx for %s", uid)
	}
	ctx.sendCh <- msg
	return nil
}

func (t *TcpEndpoint) Request(msg protocol.Msg) (ret protocol.Msg, err error) {
	if msg.H.ID == "" {
		return protocol.Msg{}, errors.New("no id specify")
	}
	id := msg.H.ID
	c, err := t.createTSTcpCtx(id)
	if err != nil {
		return protocol.Msg{}, err
	}
	//cl.Lock()
	//cnt += 1
	//msg.H.SEQ = cnt
	//cl.Unlock()
	msg.H.SEQ = c.newSeq(msg.B, id)
	c.sendCh <- msg
	to := time.After(reqTimeout)
	for {
		select {
		case <-c.curCtx.Done():
			return protocol.Msg{}, errors.New("error ctx close")
		case <-to:
			return protocol.Msg{}, errors.New("error request timeout")
		default:
		}
		r, ok := c.rcvPadding.Get(msg.H.SEQ)
		if ok {
			c.rcvPadding.Remove(msg.H.SEQ)
			return r, nil
		}
		time.Sleep(time.Microsecond * 100)
	}
}

func (t *TcpEndpoint) createTSTcpCtx(id string) (*tcpCtx, error) {
	fmt.Println("start to lock c", id)
	t.tcmLock.Lock()
	defer t.tcmLock.Unlock()
	c, ok := t.ts.Get(id)
	if ok { //return direct if exist
		return c, nil
	}
	svrCfg := config.GetSvrCfg(id)
	tcpAddr := svrCfg["tcpaddr"]
	if tcpAddr == "" {
		return nil, fmt.Errorf("no server cfg found for %s", id)
	}
	addr, err := net.ResolveTCPAddr("tcp", tcpAddr.(string))
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		logs.Info("err dial tcp : ", err)
		return nil, err
	}
	if err = setConnBuffer(conn); err != nil {
		logs.Info("err set tcp buffer : ", err)
		return nil, err
	}
	var tc = newTcpCtx(id, t, conn)
	if err = tc.establishTo(conn.LocalAddr().String()); err != nil {
		return nil, err
	}
	logs.Info("establish conn success src=%s dest=%s", config.GetConfig().GetString("svr_id"), id)
	t.ts.Set(tc.id, tc)
	gort.QuietGo(func() {
		tc.startReadMsg(ACK)
	})
	gort.QuietGo(func() {
		tc.selectOnSendChan()
	})
	return tc, nil
}

func newTcpCtx(id string, t *TcpEndpoint, conn *net.TCPConn) *tcpCtx {
	ctx2, cancel2 := context.WithCancel(t.ctx)
	return &tcpCtx{
		ept:        2,
		parCtx:     t.ctx,
		curCtx:     ctx2,
		cf:         cancel2,
		conn:       conn,
		lAddr:      conn.LocalAddr(),
		rAddr:      conn.RemoteAddr(),
		sendCh:     make(chan protocol.Msg, 128),
		d:          dispatch.GetDispatcher(),
		ext:        make(map[string]string),
		id:         id,
		rcvPadding: collect.NewUInt32Map[protocol.Msg](),
	}
}
func (t *TcpEndpoint) Stop() error {
	t.cf()
	return nil
}

// each tcp context represent a tcp connect
// ept = 0 specify a connection from client
// ept = 1 specify a connection from  other server
// ept = 2 specify a connection connect to other server
type tcpCtx struct {
	id         string                                      // tcp ctx identifier,unique
	parCtx     context.Context                             // parent context to control all sub go rt close
	curCtx     context.Context                             // tcp ctx content
	cf         context.CancelFunc                          // call to close current tcp ctx
	conn       *net.TCPConn                                // real connect
	ept        byte                                        // 0 cs 1 ss 2 client
	lAddr      net.Addr                                    // tcp local addr
	rAddr      net.Addr                                    //tcp remote addr
	co         sync.Once                                   // once used to close tcp ctx
	lrt        time.Time                                   //last read time
	lwt        time.Time                                   //last write time
	d          *dispatch.Dispatcher                        //service layer's msg handler
	sendCh     chan protocol.Msg                           //chanel to send pkg
	rcvPadding collect.ConcurrentMap[uint32, protocol.Msg] //if tcp ctx is some connect to another server, use to recv pkg
	ext        map[string]string                           //tcp establish param,see TcpEstablish
	seq        uint32                                      // current msg seq ,the recv package seq must equal to send pkg
	sl         sync.Mutex                                  //tcp ctx lock
}

var (
	tcpEndpoint TcpEndpoint
	inited      atomic.Bool
	cesLock     sync.RWMutex
)

func NewCSEndpoint() dispatch.Endpoint {
	cesLock.Lock()
	defer cesLock.Unlock()
	if !inited.Load() {
		ctx, cancel := context.WithCancel(context.Background())
		tcpEndpoint = TcpEndpoint{
			ctx: ctx,
			cf:  cancel,
			d:   dispatch.GetDispatcher(),
			fc:  collect.NewStrMap[*tcpCtx](),
			fs:  collect.NewStrMap[*tcpCtx](),
			ts:  cmap.New[*tcpCtx](),
		}
		dispatch.GetDispatcher().AttachCSEndpoint(&tcpEndpoint)
		inited.Store(true)
	}
	return &tcpEndpoint
}

func (t *TcpEndpoint) Start() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", config.GetString("endpoint.tcp.addr"))
	if err != nil {
		errMsg := fmt.Sprintf("tcp resolve failed %s %s", tcpAddr, err)
		logs.Error(errMsg)
		return errors.New(errMsg)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		errMsg := fmt.Sprintf("tcp listen up failed %s", err)
		logs.Error(errMsg)
		return errors.New(errMsg)
	}
	t.l = listener

	gort.QuietGo(func() { t.serve() })
	return nil
}

func (t *TcpEndpoint) serve() { // nolint:revive
	logs.Info("TcpEndpoint server start")
	var once sync.Once
	cl := func() {
		logs.Info("TcpEndpoint server end")
		if err := t.l.Close(); err != nil {
			logs.Error("listener close err:%s", err.Error())
		}
	}
	defer once.Do(cl)

	for {
		conn, err := t.l.AcceptTCP()
		if err != nil {
			logs.Error("tcp accept fail, err:%s", err.Error())
			var e net.Error
			if errors.As(err, &e) && e.Timeout() {
				acceptWait = acceptWait * 2
				if acceptWait > acceptTimeout {
					// accept time out return direct
					return
				}
				time.Sleep(acceptWait)
				continue
			}
			// some bad error, return
			return
		}
		logs.Info("new conn from: %s", conn.RemoteAddr().String())

		if err := setConnBuffer(conn); err != nil {
			logs.Error("set conn buffer error:%s", err)
			continue
		}

		acceptWait = 5 * time.Millisecond

		ctx2, cancel2 := context.WithCancel(t.ctx)
		var c = &tcpCtx{
			parCtx: t.ctx,
			curCtx: ctx2,
			cf:     cancel2,
			conn:   conn,
			lAddr:  conn.LocalAddr(),
			rAddr:  conn.RemoteAddr(),
			sendCh: make(chan protocol.Msg, sendQueueSize),
			d:      dispatch.GetDispatcher(),
			ext:    make(map[string]string),
		}

		gort.QuietGo(func() {
			if err := c.establishIncoming(); err != nil {
				logs.Error("tcp establish incoming err:%s", err)
				return
			}
			c.startReadMsg(REQ)
		})
	}
}

func setConnBuffer(conn *net.TCPConn) error {
	if err := conn.SetReadBuffer(bufSize); err != nil {
		logs.Error("Set read buffer error %s err:%s", conn.RemoteAddr().String(), err.Error())
		if err = conn.Close(); err != nil {
			logs.Error("Set read buffer, close err:%s", err.Error())
		}
		return err
	}
	if err := conn.SetWriteBuffer(bufSize); err != nil {
		logs.Error("Set write buffer error %s err:%s", conn.RemoteAddr().String(), err.Error())
		if err = conn.Close(); err != nil {
			logs.Error("Set write buffer, close err:%s", err.Error())
		}
		return err
	}
	return nil
}

// direct 0 request;1 response
func (tc *tcpCtx) startReadMsg(direct byte) {
	defer tc.stop()
	for {
		select {
		case <-tc.parCtx.Done():
			logs.Error("tcp server stop,close client conn:id %s local addr %s remote addr %s, current seq %d", tc.id, tc.lAddr.String(), tc.rAddr.String(), tc.seq)
			return
		case <-tc.curCtx.Done():
			logs.Error("tcp content stop,close client conn:id %s local addr %s remote addr %s, current seq %d", tc.id, tc.lAddr.String(), tc.rAddr.String(), tc.seq)
			return
		default:
		}
		q, err := tc.readOneMsg(direct)
		if err != nil {
			logs.Error("tcp read msg error id %s err %s,quit %t", tc.id, err, q)
		}
		if q {
			return
		}
	}
}
func (tc *tcpCtx) stop() {
	tc.co.Do(func() {
		if tc.ept == 0 {
			tcpEndpoint.fc.Remove(tc.id)
		} else if tc.ept == 1 {
			tcpEndpoint.fs.Remove(tc.id)
		} else {
			tcpEndpoint.ts.Remove(tc.id)
		}
		logs.Error("stop tcp conn id %s end point type %d local addr %s remote addr %s", tc.id, tc.ept, tc.lAddr.String(), tc.rAddr.String())
		tc.cf()
		err := tc.conn.Close()
		if err != nil {
			logs.Error("stop tcp ctx error %s", err)
		}
	})
}

func (tc *tcpCtx) setRdl() {
	// timeout control
	err := tc.conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(config.GetIntDef("endpoint.tcp.idletimeout", 60))))
	if err != nil {
		logs.Error("Set read deadline fail, err:%s", err.Error())
	}
}

func (tc *tcpCtx) setWdl() {
	// timeout control
	err := tc.conn.SetWriteDeadline(time.Now().Add(time.Second * time.Duration(config.GetIntDef("endpoint.tcp.idletimeout", 60))))
	if err != nil {
		logs.Error("Set write deadline fail, err:%s", err.Error())
	}
}

// GetEncodeBuf 取一个编码缓冲区.
func getBuf() *bytes.Buffer {
	buf, ok := bufPool.Get().(*bytes.Buffer)
	if !ok {
		logs.Error("_encodeBufPool type assert *bytes.Buffer")
		return &bytes.Buffer{}
	}
	return buf
}

// releaseBUf 将缓冲区放回池中.
func releaseBUf(buf *bytes.Buffer) {
	if buf == nil {
		return
	}
	buf.Reset()
	bufPool.Put(buf)
}
func (tc *tcpCtx) establishIncoming() error {
	tc.setRdl()
	lenBuf := make([]byte, 4)
	n, err := io.ReadFull(tc.conn, lenBuf)
	if err != nil || n != len(lenBuf) {
		logs.Error("read uid head error %s", err)
		return err
	}
	msgLen := binary.LittleEndian.Uint32(lenBuf)
	idCap := make([]byte, msgLen)

	n, err = io.ReadFull(tc.conn, idCap)
	if err != nil || n != int(msgLen) {
		logs.Error("read uid data error %s", err)
		return err
	}

	var establish = protocol.TcpEstablish{}
	err = proto.Unmarshal(idCap, &establish)
	if err != nil {
		return err
	}

	tc.id = establish.ID
	tc.ept = byte(establish.EndpointType)
	tc.ext = establish.Ext
	if tc.ept == 0 { // c 2 s
		tcpEndpoint.fc.Set(tc.id, tc)
	} else if tc.ept == 1 { // s 2 s
		tcpEndpoint.fs.Set(tc.id, tc)
	} else {
		logs.Error("unknown end point type remote addr %s,type %d", tc.rAddr.String(), tc.ept)
	}
	gort.QuietGo(func() {
		tc.selectOnSendChan()
	})
	return nil
}

func (tc *tcpCtx) readLeading() (protocol.MsgLeading, error) {
	tc.setRdl()
	buf := make([]byte, protocol.MsgLeadIngSize)
	n, err := io.ReadFull(tc.conn, buf)
	if err != nil || n != protocol.MsgLeadIngSize {
		return protocol.MsgLeading{}, err
	}
	return dispatch.DecodeMsgLeading(buf)
}

func (tc *tcpCtx) readOneMsg(direct byte) (bool, error) {
	leading, err := tc.readLeading()
	if err != nil {
		return true, err
	}
	buf := make([]byte, leading.HL+leading.BL)
	tc.setRdl()
	n, err := io.ReadFull(tc.conn, buf)
	if err != nil || n != len(buf) {
		return true, err
	}
	mh := &protocol.MsgHead{}
	err = proto.Unmarshal(buf[:leading.HL], mh)
	if err != nil {
		return true, err
	}
	message, err := tc.d.Create(mh.MsgID, direct)
	if err != nil {
		return true, err
	}
	err = proto.Unmarshal(buf[leading.HL:leading.HL+leading.BL], message)
	if err != nil {
		return true, err
	}
	if tc.ept != 2 {
		tc.setSeq(mh.SEQ)
	}
	if mh.ID == "" {
		logs.Warn("msg no id specify id %s addr %s seq %d", tc.id, tc.rAddr.String(), tc.seq)
		mh.ID = tc.id
	}
	msg := protocol.Msg{
		H: mh,
		B: message,
	}
	if tc.ept == 2 {
		tc.rcvPadding.Set(msg.H.SEQ, msg)
		return false, nil
	}
	ret, err := tc.d.Dispatch(msg)
	ret.H.SEQ = mh.SEQ
	ret.H.MsgID = mh.MsgID
	ret.H.ID = mh.ID
	if err != nil {
		return true, err
	}
	tc.sendCh <- ret
	return false, nil
}

func (tc *tcpCtx) selectOnSendChan() {
	defer tc.stop()
	for {
		select {
		case <-tc.parCtx.Done():
			return
		case <-tc.curCtx.Done():
			return
		case pkg := <-tc.sendCh:
			logs.Info("current send pkg info tcp ctx id %s tc addr %p,seq %d", tc.id, &tc.sl, pkg.H.SEQ)
			err := tc.send(pkg)
			if err != nil {
				logs.Error("client=%s UID=%s send pkg err:%s", tc.rAddr.String(), tc.id, err)
				return
			}
		}
	}
}

func (tc *tcpCtx) send(pkg protocol.Msg) error {
	p, err := dispatch.EncodeMsg(pkg)
	if err != nil {
		return fmt.Errorf("send UID=%s, err:%w", tc.id, err)
	}

	tc.setWdl()
	_, err = tc.conn.Write(p)
	if err != nil {
		return fmt.Errorf("SendToClient pkg fail, msgid %s, uid %s, err:%w", pkg.H.MsgID, pkg.H.ID, err)
	}
	logs.Info("sended new pkg MSGID=%s,CurSeq=%d", pkg.H.MsgID, tc.seq)
	return nil
}

func (tc *tcpCtx) establishTo(id string) error {
	esInfo := protocol.TcpEstablish{
		EndpointType: 1,
		ID:           id,
	}
	bys, err := proto.Marshal(&esInfo)
	if err != nil {
		return err
	}
	err = binary.Write(tc.conn, binary.LittleEndian, uint32(len(bys)))
	if err != nil {
		return err
	}
	_, err = tc.conn.Write(bys)
	if err != nil {
		return err
	}
	return nil
}

func (tc *tcpCtx) newSeq(b proto.Message, id string) uint32 {
	tc.sl.Lock()
	defer tc.sl.Unlock()
	logs.Info("new seq request gort %s id %s seq %d %p", b, id, tc.seq, &tc.sl)
	tc.seq += 1
	return tc.seq
}

func (tc *tcpCtx) setSeq(seq uint32) {
	tc.sl.Lock()
	defer tc.sl.Unlock()
	tc.seq = seq
}
