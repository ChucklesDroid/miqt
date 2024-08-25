package qt

/*

#cgo CFLAGS: -fPIC
#cgo pkg-config: Qt5Widgets
#include "gen_qstandardpaths.h"
#include <stdlib.h>

*/
import "C"

import (
	"unsafe"
)

type QStandardPaths struct {
	h *C.QStandardPaths
}

func (this *QStandardPaths) cPointer() *C.QStandardPaths {
	if this == nil {
		return nil
	}
	return this.h
}

func newQStandardPaths(h *C.QStandardPaths) *QStandardPaths {
	return &QStandardPaths{h: h}
}

func newQStandardPaths_U(h unsafe.Pointer) *QStandardPaths {
	return newQStandardPaths((*C.QStandardPaths)(h))
}

func QStandardPaths_FindExecutable(executableName string) string {
	executableName_Cstring := C.CString(executableName)
	defer C.free(unsafe.Pointer(executableName_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QStandardPaths_FindExecutable(executableName_Cstring, C.ulong(len(executableName)), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QStandardPaths_EnableTestMode(testMode bool) {
	C.QStandardPaths_EnableTestMode((C.bool)(testMode))
}

func QStandardPaths_SetTestModeEnabled(testMode bool) {
	C.QStandardPaths_SetTestModeEnabled((C.bool)(testMode))
}

func QStandardPaths_IsTestModeEnabled() bool {
	ret := C.QStandardPaths_IsTestModeEnabled()
	return (bool)(ret)
}

func QStandardPaths_FindExecutable2(executableName string, paths []string) string {
	executableName_Cstring := C.CString(executableName)
	defer C.free(unsafe.Pointer(executableName_Cstring))
	// For the C ABI, malloc two C arrays; raw char* pointers and their lengths
	paths_CArray := (*[0xffff]*C.char)(C.malloc(C.ulong(8 * len(paths))))
	paths_Lengths := (*[0xffff]C.size_t)(C.malloc(C.ulong(8 * len(paths))))
	defer C.free(unsafe.Pointer(paths_CArray))
	defer C.free(unsafe.Pointer(paths_Lengths))
	for i := range paths {
		single_cstring := C.CString(paths[i])
		defer C.free(unsafe.Pointer(single_cstring))
		paths_CArray[i] = single_cstring
		paths_Lengths[i] = (C.size_t)(len(paths[i]))
	}
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QStandardPaths_FindExecutable2(executableName_Cstring, C.ulong(len(executableName)), &paths_CArray[0], &paths_Lengths[0], C.ulong(len(paths)), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}