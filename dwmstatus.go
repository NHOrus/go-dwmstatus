package main

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lX11
// #include <X11/Xlib.h>
import "C"

import (
	"fmt"
	"log"
	"time"
)

var dpy = C.XOpenDisplay(nil)

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
		t := time.Now().Format("Mon 2006-Jan-2 15:04:05")

		s := formatStatus("%s", t)
		setStatus(s)
		time.Sleep(time.Second)
	}
}
