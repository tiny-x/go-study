
package nsenter

/*
#cgo CFLAGS: -Wall
extern void ns_exec();
void __attribute__((constructor)) ex(void) {
	ns_exec();
}
*/
import "C"
