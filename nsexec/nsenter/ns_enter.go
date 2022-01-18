//go:build linux && !gccgo
// +build linux,!gccgo

package nsenter

/*
#cgo CFLAGS: -Wall
extern void ns_exec();
void __attribute__((constructor)) init(void) {
	ns_exec();
}
*/
import "C"
