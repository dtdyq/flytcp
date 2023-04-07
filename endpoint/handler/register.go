package handler

import (
	"flytcp/endpoint/dispatch"
	pb "flytcp/pb/pbgen"
)

func InitHandler() {
	dispatch.RegisterHandler(&pb.CSTestHandler{})
}
