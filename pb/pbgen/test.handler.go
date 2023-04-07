package pb

import (
	"flytcp/endpoint/protocol"
	"github.com/golang/protobuf/proto"
)

type CSTestHandler struct {
}

func (h *CSTestHandler) MsgID() string {
	return "cstest"
}
func (h *CSTestHandler) Receiver() []proto.Message {
	return []proto.Message{&CSTest_Req{}, &CSTest_Rsp{}}
}
func (h *CSTestHandler) Handle(msg protocol.Msg) (proto.Message, error) {
	_ = msg.B.(*CSTest_Req)
	return &CSTest_Rsp{}, nil
}
