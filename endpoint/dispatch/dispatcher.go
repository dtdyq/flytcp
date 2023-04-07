package dispatch

import (
	"bytes"
	"encoding/binary"
	"errors"
	"flytcp/endpoint/protocol"
	"fmt"
	"github.com/golang/protobuf/proto"
	"sync"
	"time"
)

var (
	eptMsg = protocol.Msg{}
)

func DecodeMsgLeading(buf []byte) (protocol.MsgLeading, error) {
	if len(buf) != protocol.MsgLeadIngSize {
		return protocol.MsgLeading{}, errors.New("invalid size")
	}
	r := protocol.MsgLeading{}
	r.HL = binary.LittleEndian.Uint32(buf[0:4])
	r.BL = binary.LittleEndian.Uint32(buf[4:8])
	return r, nil
}

type Handler interface {
	MsgID() string
	Receiver() []proto.Message
	Handle(msg protocol.Msg) (proto.Message, error)
}

type Endpoint interface {
	Start() error
	Stop() error
	NTF(msg protocol.Msg) error
	Request(msg protocol.Msg) (protocol.Msg, error)
}
type Dispatcher struct {
	hl      sync.RWMutex
	handler map[string]Handler
	cse     Endpoint
}

func (d *Dispatcher) AttachCSEndpoint(ep Endpoint) {
	d.cse = ep
}

func (d *Dispatcher) Create(msgID string, direct byte) (proto.Message, error) {
	d.hl.RLock()
	defer d.hl.RUnlock()
	c, ok := d.handler[msgID]
	if !ok {
		return nil, fmt.Errorf("can not find creator %s", msgID)
	}
	return c.Receiver()[direct], nil
}

func (d *Dispatcher) Dispatch(msg protocol.Msg) (protocol.Msg, error) {
	d.hl.RLock()
	defer d.hl.RUnlock()
	c, ok := d.handler[msg.H.MsgID]
	if !ok {
		return eptMsg, fmt.Errorf("can not find handler %s", msg.H.MsgID)
	}
	rt, err := c.Handle(msg)
	if err != nil {
		return eptMsg, err
	}
	return protocol.Msg{
		H: &protocol.MsgHead{
			MsgID: msg.H.MsgID,
		},
		B: rt,
	}, err
}

func (d *Dispatcher) SendNTF(msg protocol.Msg) {
	d.cse.NTF(msg)
}
func (d *Dispatcher) REQ(msg protocol.Msg) (protocol.Msg, error) {
	defer func(t time.Time) {
	}(time.Now())
	return d.cse.Request(msg)
}

func NTF(msg protocol.Msg) {
	msg.H.MT = protocol.MsgType_NTF
	dispatcher.SendNTF(msg)
}
func REQ(msg protocol.Msg) (protocol.Msg, error) {
	return dispatcher.REQ(msg)
}
func EncodeMsg(msg protocol.Msg) ([]byte, error) {
	buf := bytes.Buffer{}
	hb, err := proto.Marshal(msg.H)
	if err != nil {
		return nil, err
	}
	bb, err := proto.Marshal(msg.B)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buf, binary.LittleEndian, uint32(len(hb)))
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buf, binary.LittleEndian, uint32(len(bb)))
	if err != nil {
		return nil, err
	}
	buf.Write(hb)
	buf.Write(bb)
	return buf.Bytes(), nil
}

var dispatcher = Dispatcher{handler: make(map[string]Handler)}

func GetDispatcher() *Dispatcher {
	return &dispatcher
}

func RegisterHandler(handler Handler) {
	dispatcher.handler[handler.MsgID()] = handler
}

func BuildReq(msgID, id string, message proto.Message) protocol.Msg {
	return protocol.Msg{
		H: &protocol.MsgHead{
			MsgID: msgID,
			ID:    id,
			MT:    protocol.MsgType_SS,
		},
		B: message,
	}
}
func BuildNtf(msgID, id string, message proto.Message) protocol.Msg {
	return protocol.Msg{
		H: &protocol.MsgHead{
			MsgID: msgID,
			ID:    id,
			MT:    protocol.MsgType_NTF,
		},
		B: message,
	}
}
