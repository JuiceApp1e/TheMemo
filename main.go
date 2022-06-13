package main

import (
	"fmt"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/routers"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	r := routers.SetupRouter()

	r.Run(":80")
}
