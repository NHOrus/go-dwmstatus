// bitfmt.go
package main

import (
	"fmt"
)

//KiB and others are binary powers of byte. They are used to format values from
//big amount of bytes into human-readable values
const (
	KiB float32 = 1024
	MiB         = 1024 * 1024
	GiB         = 1024 * 1024 * 1024
	TiB         = 1024 * 1024 * 1024 * 1024
	PiB         = 1024 * 1024 * 1024 * 1024 * 1024
)

//formats value in bytes into human-readable value in binary power of bytes.
func humanBytes(b float32) string {
	if b > 10*PiB {
		return fmt.Sprintf("%.0f PB", b/PiB)
	}
	if b > 10*TiB {
		return fmt.Sprintf("%.0f TB", b/TiB)
	}
	if b > 10*GiB {
		return fmt.Sprintf("%.0f GB", b/GiB)
	}
	if b > 10*MiB {
		return fmt.Sprintf("%.0f MB", b/MiB)
	}
	if b > 10*KiB {
		return fmt.Sprintf("%.0f KB", b/KiB)
	}
	return fmt.Sprintf("%.0f B", b)
}
