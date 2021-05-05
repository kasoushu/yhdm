package main

import (
	"yhdm/global"
	"yhdm/log"
)

func main() {
	global.LOG =  log.InitLogger()
	global.LOG.Info("gin run")
	global.InitGlobal("127.0.0.1:7999")
}