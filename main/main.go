package main

import (
	"flag"
	logs "flytcp/common/logs"
	"flytcp/config"
	"flytcp/endpoint/dispatch"
	"flytcp/main/boot"
	pb "flytcp/pb/pbgen"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var cl sync.Mutex
var c int

const cfgF = `
svr_id = "%d"

endpoint.enabled = ["tcp"]
[endpoint.tcp]
addr = "9.134.151.80:3800%d"
`

var id int

func main() {
	flag.IntVar(&id, "id", 1, "id")
	flag.Parse()
	f, err := os.OpenFile("configuration/common.toml", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(fmt.Sprintf(cfgF, id, id))
	if err != nil {
		fmt.Println(err)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
	flag.Parse()
	boot.Boot(id)
	if config.GetString("svr_id") == "12" {
		wg := sync.WaitGroup{}
		wg.Add(100)
		for j := 0; j < 100; j++ {

			go func(int2 int) {
				logs.Info("new go func %d", int2)
				for i := 0; i < 1000; i++ {
					id := strconv.Itoa(rand.Intn(3) + 1)
					cur := time.Now().UnixMilli()
					m := dispatch.BuildReq("cstest", id, &pb.CSTest_Req{
						Ping: "gor " + strconv.Itoa(int2) + " " + id + strconv.Itoa(i),
					})

					logs.Info("new send go rt %d to svr  %s", int2, id)
					ret, err := dispatch.REQ(m)
					fmt.Println(err)
					if err == nil && ret.B.(*pb.CSTest_Rsp).Pong != m.B.(*pb.CSTest_Req).Ping {
						logs.Error("msg err esx %s ", id+strconv.Itoa(i))
					}
					logs.Info("svr cfg map r %d %d %s", int2, time.Now().UnixMilli()-cur, ret)
					cl.Lock()
					c += 1
					cl.Unlock()
				}
				wg.Done()
				logs.Info("cur finish %d", int2)
			}(j)
		}

		wg.Wait()
		logs.Info("finish count %d", c)
	}
	for {
		time.Sleep(time.Second)
		logs.Info("server running")
	}
}
