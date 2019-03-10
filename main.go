package main

import (
	"fmt"
	"github.com/blanccobb/go-mgo-girdfs-fileserver/app"
)

func main() {
	fmt.Print("hello Go")
	
	apps := &app.App{}
	apps.Init()
	apps.Run(":8082")
}

		