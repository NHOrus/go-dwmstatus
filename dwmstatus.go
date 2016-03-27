package main

import (
	"log"
	"time"
)

func main() {
	if dpy == nil {
		log.Fatal("Can't open display")
	}
	for {
		t := time.Now().Format("Mon 2006-Jan-2 15:04:05")

		m, _ := memUpdate()

		s := formatStatus("%s | %s", m, t)
		setStatus(s)
		time.Sleep(time.Second)
	}
}
