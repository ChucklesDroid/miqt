package qt

/*

#cgo CFLAGS: -fPIC
#cgo pkg-config: Qt5Widgets
#include "gen_qbuffer.h"
#include <stdlib.h>

*/
import "C"

import (
	"unsafe"
)

type QBuffer struct {
	h *C.QBuffer
	*QIODevice
}

func (this *QBuffer) cPointer() *C.QBuffer {
	if this == nil {
		return nil
	}
	return this.h
}

func newQBuffer(h *C.QBuffer) *QBuffer {
	return &QBuffer{h: h, QIODevice: newQIODevice_U(unsafe.Pointer(h))}
}

func newQBuffer_U(h unsafe.Pointer) *QBuffer {
	return newQBuffer((*C.QBuffer)(h))
}

// NewQBuffer constructs a new QBuffer object.
func NewQBuffer() *QBuffer {
	ret := C.QBuffer_new()
	return newQBuffer(ret)
}

// NewQBuffer2 constructs a new QBuffer object.
func NewQBuffer2(buf *QByteArray) *QBuffer {
	ret := C.QBuffer_new2(buf.cPointer())
	return newQBuffer(ret)
}

// NewQBuffer3 constructs a new QBuffer object.
func NewQBuffer3(parent *QObject) *QBuffer {
	ret := C.QBuffer_new3(parent.cPointer())
	return newQBuffer(ret)
}

// NewQBuffer4 constructs a new QBuffer object.
func NewQBuffer4(buf *QByteArray, parent *QObject) *QBuffer {
	ret := C.QBuffer_new4(buf.cPointer(), parent.cPointer())
	return newQBuffer(ret)
}

func (this *QBuffer) MetaObject() *QMetaObject {
	ret := C.QBuffer_MetaObject(this.h)
	return newQMetaObject_U(unsafe.Pointer(ret))
}

func QBuffer_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QBuffer_Tr(s_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QBuffer_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QBuffer_TrUtf8(s_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QBuffer) Buffer() *QByteArray {
	ret := C.QBuffer_Buffer(this.h)
	return newQByteArray_U(unsafe.Pointer(ret))
}

func (this *QBuffer) Buffer2() *QByteArray {
	ret := C.QBuffer_Buffer2(this.h)
	return newQByteArray_U(unsafe.Pointer(ret))
}

func (this *QBuffer) SetBuffer(a *QByteArray) {
	C.QBuffer_SetBuffer(this.h, a.cPointer())
}

func (this *QBuffer) SetData(data *QByteArray) {
	C.QBuffer_SetData(this.h, data.cPointer())
}

func (this *QBuffer) SetData2(data string, lenVal int) {
	data_Cstring := C.CString(data)
	defer C.free(unsafe.Pointer(data_Cstring))
	C.QBuffer_SetData2(this.h, data_Cstring, (C.int)(lenVal))
}

func (this *QBuffer) Data() *QByteArray {
	ret := C.QBuffer_Data(this.h)
	return newQByteArray_U(unsafe.Pointer(ret))
}

func (this *QBuffer) Close() {
	C.QBuffer_Close(this.h)
}

func (this *QBuffer) Size() int64 {
	ret := C.QBuffer_Size(this.h)
	return (int64)(ret)
}

func (this *QBuffer) Pos() int64 {
	ret := C.QBuffer_Pos(this.h)
	return (int64)(ret)
}

func (this *QBuffer) Seek(off int64) bool {
	ret := C.QBuffer_Seek(this.h, (C.int64_t)(off))
	return (bool)(ret)
}

func (this *QBuffer) AtEnd() bool {
	ret := C.QBuffer_AtEnd(this.h)
	return (bool)(ret)
}

func (this *QBuffer) CanReadLine() bool {
	ret := C.QBuffer_CanReadLine(this.h)
	return (bool)(ret)
}

func QBuffer_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QBuffer_Tr2(s_Cstring, c_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QBuffer_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QBuffer_Tr3(s_Cstring, c_Cstring, (C.int)(n), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QBuffer_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QBuffer_TrUtf82(s_Cstring, c_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QBuffer_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QBuffer_TrUtf83(s_Cstring, c_Cstring, (C.int)(n), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QBuffer) Delete() {
	C.QBuffer_Delete(this.h)
}