package tcp

import (
	"encoding/binary"
	"flytcp/common/logs"
	"flytcp/endpoint/dispatch"
	"flytcp/pb"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	logs.LogInit()
	wg := sync.WaitGroup{}
	wg.Add(10)
	for j := 0; j < 10; j++ {
		go func(jj int) {
			addr, err := net.ResolveTCPAddr("tcp", "9.134.151.80:38001")
			conn, err := net.DialTCP("tcp", nil, addr)
			if err != nil {
				fmt.Println("err : ", err)
				return
			}

			est := pb.TcpEstablish{ID: strconv.Itoa(jj), EndpointType: 0}
			b, _ := proto.Marshal(&est)

			_ = binary.Write(conn, binary.LittleEndian, uint32(len(b)))
			conn.Write(b)
			fmt.Println("est success")
			cur := time.Now().UnixMilli()

			for i := 0; i < 1000; i++ {

				sc := time.Now().UnixMilli()
				//	s, _ := strconv.ParseUint(fmt.Sprintf("%d%d", jj+1, i), 32, 10)
				m := dispatch.Msg{
					H: &pb.MsgHead{
						MsgID: "cstest",
						ID:    "cli1",
						SEQ:   uint32(i),
					},
					B: &pb.CSTest_Req{Ping: strconv.Itoa(i)},
				}
				bys, _ := dispatch.EncodeMsg(m)
				conn.Write(bys)

				buf := make([]byte, dispatch.MsgLeadIngSize)
				n, err := io.ReadFull(conn, buf)
				if err != nil || n != dispatch.MsgLeadIngSize {
					fmt.Println(err)
				}
				leading, _ := dispatch.DecodeMsgLeading(buf)
				if err != nil {
					fmt.Println(err)
				}
				buf = make([]byte, leading.HL+leading.BL)
				n, err = io.ReadFull(conn, buf)
				if err != nil || n != len(buf) {
					fmt.Println(err)
				}
				mh := &pb.MsgHead{}
				err = proto.Unmarshal(buf[:leading.HL], mh)
				if err != nil {
					fmt.Println(err)
				}
				message := pb.CSTest_Rsp{}
				if err != nil {
					fmt.Println(err)
				}
				err = proto.Unmarshal(buf[leading.HL:leading.HL+leading.BL], &message)
				if err != nil {
					fmt.Println(err)
				}
				if err == nil && message.Pong != m.B.(*pb.CSTest_Req).Ping {
					logs.Error("msg err esx %s %s", m.B.(*pb.CSTest_Req).String(), message.String())
				}
				fmt.Println("fs cost ", time.Now().UnixMilli()-sc)
			}
			fmt.Println("total cost ", time.Now().UnixMilli()-cur)
			wg.Done()
		}(j)
	}
	wg.Wait()
}
