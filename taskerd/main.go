package main

import (
	"github.com/notdylanburns/tasker/taskerd/config"
	"log"
)

func main() {
	_, err := config.Create("/etc")
	if err != nil {
		log.Fatal(err)
	}

}
