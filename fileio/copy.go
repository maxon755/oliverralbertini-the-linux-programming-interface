package main

/*
#cgo CFLAGS: -I../lib/
#cgo LDFLAGS: -L../lib/ -lerror_functions
#include <stdarg.h>
#include <stdlib.h>
#include "error_functions.h"

// Create a wrapper function to call the variadic C function
void callUsageErr(const char *msg) {
    usageErr("%s", msg);
}

void callErrExit(const char *msg) {
    errExit("%s", msg);
}
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {

	if (len(os.Args) == 2 && os.Args[1] == "--help") || len(os.Args) != 3 {

		msg := C.CString(fmt.Sprintf("%s old-file new-file\n", os.Args[0]))
		defer C.free(unsafe.Pointer(msg))

		C.callUsageErr(msg)
		return
	}

	sourceFile, err := os.Open(os.Args[1])
	if err != nil {
		msg := C.CString(fmt.Sprintf("opening file %s", os.Args[1]))
		defer C.free(unsafe.Pointer(msg))

		C.callErrExit(msg)
	}
	defer sourceFile.Close()

	destFile, err := os.OpenFile(os.Args[2], os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		msg := C.CString(fmt.Sprintf("opening file %s", os.Args[2]))
		defer C.free(unsafe.Pointer(msg))

		C.callErrExit(msg)
	}
	defer destFile.Close()

	bufferSize := 1024
	buffer := make([]byte, bufferSize)

	for {
		// Read data from source file
		bytesRead, err := sourceFile.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading from source file:", err)
			return
		}

		// If no more data to read, break the loop
		if bytesRead == 0 {
			break
		}

		// Write data to destination file
		_, err = destFile.Write(buffer[:bytesRead])
		if err != nil {
			fmt.Println("Error writing to destination file:", err)
			return
		}
	}
}
