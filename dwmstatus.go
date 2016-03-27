package main

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lX11 -L/usr/lib -lkvm
// #include <X11/Xlib.h>
// #include <kvm.h>
import "C"

import (
	"fmt"
	"log"
	"time"

	"github.com/blabber/go-freebsd-sysctl/sysctl"
)

var pagesize int64
var pae bool
var kd C.kvm_t
var swapArr C.struct_kvm_swap

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

func init() {
	var err error
	pagesize, err = sysctl.GetInt64("hw.pagesize")
	if err != nil {
		panic(err)
	}

	_, err = sysctl.GetString("kern.features.pae")
	if err.Error() == "no such file or directory" {
		pae = false
	} else {
		pae = true
	}

}

func (m *memData) Update() error {
	if pae {
		mtemp, err := sysctl.GetInt64("hw.availpages")
		if err != nil {
			panic(err)
		}
		m.memTotal = uint64(mtemp * pagesize)
	}

	var mpage, mtemp int64

	mtemp, err := sysctl.GetInt64("hw.physmem")
	if err != nil {
		panic(err)
	}
	m.memTotal = uint64(mtemp)

	for _, str := range []string{"vm.stats.vm.v_cache_count", "vm.stats.vm.v_free_count"} {

		mpage, err = sysctl.GetInt64(str)
		if err != nil {
			panic(err)
		}
		mtemp = mpage * pagesize
	}
	m.memFree = uint64(mtemp)

	m.memUse = m.memTotal - m.memFree
	m.memPercent = int(m.memUse * 100 / m.memTotal)

	mtemp, err = sysctl.GetInt64("vm.swap_total")
	m.swapTotal = uint64(mtemp)

	err = nil

	i, _ := C.kvm_getswapinfo(&kd, &swapArr, C.int(1), C.int(0))
	if err != nil {
		panic(err)
	}
	if i >= 0 && swapArr.ksw_total != 0 {
		m.swapUse = uint64(swapArr.ksw_used) * uint64(pagesize)
	}
	m.swapPercent = int(m.swapUse * 100 / m.swapTotal)

	return nil
}
