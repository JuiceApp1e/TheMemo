package main

import (
	"fmt"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/routers"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	r := routers.SetupRouter()

	db.Get().AutoMigrate(&model.TodoModel{})

	r.Run(":80")
}
