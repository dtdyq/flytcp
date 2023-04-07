package boot

import (
	logs "flytcp/common/logs"
	"flytcp/config"
	"flytcp/endpoint"
	"flytcp/endpoint/dispatch"
	pb "flytcp/pb/pbgen"
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func Boot(id int) {
	logs.LogInit()
	if err := config.Init("main/configuration", "common", "servers"); err != nil {
		panic(err)
	}
	endpoint.Init()
	dispatch.RegisterHandler(&pb.CSTestHandler{})
	logs.Info("server id %s", config.GetConfig().AllSettings())

	go func() {
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:606%d", id), nil)
		if err != nil {
			panic(err)
		}
	}()
}
