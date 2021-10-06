package main

import (
	"./modules/config"
	"./modules/telegram"
	"./singleton"
	"./webserver"
	"fmt"
	"github.com/Psh777/gen_id"
	"os"
	"time"
)

func main() {
	_, _ = fmt.Fprintln(os.Stdout, "START")
	var version = "1.0.4"
	var node, _ = gen_id.New(6)

	singleton.SetData(version, node)
	singleton.SetUptime(time.Now().Unix())

	fmt.Println("v", version)
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())

	//config
	conf := config.GetConfig("prod")
	//tables

	go telegram.Init(conf.Env)

	webserver.Init(conf)
}
