//go:build unix

package main

/*
#include <stdlib.h>
#include <unistd.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	//
	// sysconf, pathconfは syscallやunixパッケージでも直接呼び出せるようになっていないため
	// cgoを利用して呼び出すのが楽。cgoに依存するのが嫌な場合は
	//   https://github.com/tklauser/go-sysconf
	// のようなパッケージを利用するという手もある。
	//
	v := C.sysconf(C._SC_NPROCESSORS_ONLN)
	fmt.Printf("_SC_NPROCESSORS_ONLN=%v\n", v)

	{
		ptrPath := C.CString("/")
		defer C.free(unsafe.Pointer(ptrPath))

		v = C.pathconf(ptrPath, C._PC_PATH_MAX)
		fmt.Printf("_PC_PATH_MAX=%v\n", v)

		v = C.pathconf(ptrPath, C._PC_NAME_MAX)
		fmt.Printf("_PC_NAME_MAX=%v\n", v)
	}

	return nil
}
