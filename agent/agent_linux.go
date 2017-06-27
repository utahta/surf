// +build linux

package agent

import (
	"runtime"
)

// osName returns the name of the OS.
func osName() string {
	return runtime.GOOS
}

// osVersion returns the OS version.
func osVersion() string {
	return "0.0"
}

// charsToString converts a [65]int8 byte array into a string.
func charsToString(ca [65]int8) string {
	s := make([]byte, len(ca))
	var lens int
	for ; lens < len(ca); lens++ {
		if ca[lens] == 0 {
			break
		}
		s[lens] = uint8(ca[lens])
	}
	return string(s[0:lens])
}
