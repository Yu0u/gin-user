package main

import (
	"study07/model"
	"study07/routers"
)

func main() {
	model.InitDb()
	routers.InitRouter()
}
