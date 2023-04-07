package endpoint

import (
	"flytcp/endpoint/handler"
	"flytcp/endpoint/tcp"
)

func Init() {
	handler.InitHandler()
	tep := tcp.NewCSEndpoint()
	tep.Start()
}
