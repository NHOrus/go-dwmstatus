package main

// #cgo LDFLAGS: -lX11 -lasound
// #include <X11/Xlib.h>
import "C"

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

var dpy = C.XOpenDisplay(nil)

func getLoadAverage(file string) (lavg string, err error) {
	loadavg, err := ioutil.ReadFile(file)
	if err != nil {
		return "Couldn't read loadavg", err
	}
	lavg = strings.Join(strings.Fields(string(loadavg))[:3], " ")
	return
}

func setStatus(s *C.char) {
	C.XStoreName(dpy, C.XDefaultRootWindow(dpy), s)
	C.XSync(dpy, 1)
}

func formatStatus(format string, args ...interface{}) *C.char {
	status := fmt.Sprintf(format, args...)
	return C.CString(status)
}

func main() {
	if dpy == nil {
		log.Fatal("Can't open display")
	}
	for {
		t := time.Now().Format("Mon 02 15:04")

		l, err := getLoadAverage("/proc/loadavg")
		if err != nil {
			log.Println(err)
		}
		s := formatStatus("%s :: %s", l, t)
		setStatus(s)
		time.Sleep(time.Second)
	}
}
