package main

import (
	"github.com/kiririx/krim/conf"
	"github.com/kiririx/krim/router"
	_ "github.com/kiririx/krim/systemx"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 设置端口号启动
	router.SetupRouter(conf.Ginner)
	if err := conf.Ginner.Run(":8080"); err != nil {
		panic(err)
	}
}
