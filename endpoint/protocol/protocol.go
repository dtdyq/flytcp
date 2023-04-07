package protocol

import (
	"github.com/golang/protobuf/proto"
)

const (
	MsgLeadIngSize = 8
)

type MsgLeading struct {
	HL uint32
	BL uint32
}

type Msg struct {
	H *MsgHead
	B proto.Message
}
