package qt

/*

#cgo CFLAGS: -fPIC
#cgo pkg-config: Qt5Widgets
#include "gen_qtextstream.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QTextStream struct {
	h *C.QTextStream
}

func (this *QTextStream) cPointer() *C.QTextStream {
	if this == nil {
		return nil
	}
	return this.h
}

func newQTextStream(h *C.QTextStream) *QTextStream {
	return &QTextStream{h: h}
}

func newQTextStream_U(h unsafe.Pointer) *QTextStream {
	return newQTextStream((*C.QTextStream)(h))
}

// NewQTextStream constructs a new QTextStream object.
func NewQTextStream() *QTextStream {
	ret := C.QTextStream_new()
	return newQTextStream(ret)
}

// NewQTextStream2 constructs a new QTextStream object.
func NewQTextStream2(device *QIODevice) *QTextStream {
	ret := C.QTextStream_new2(device.cPointer())
	return newQTextStream(ret)
}

func (this *QTextStream) SetCodec(codec *QTextCodec) {
	C.QTextStream_SetCodec(this.h, codec.cPointer())
}

func (this *QTextStream) SetCodecWithCodecName(codecName string) {
	codecName_Cstring := C.CString(codecName)
	defer C.free(unsafe.Pointer(codecName_Cstring))
	C.QTextStream_SetCodecWithCodecName(this.h, codecName_Cstring)
}

func (this *QTextStream) Codec() *QTextCodec {
	ret := C.QTextStream_Codec(this.h)
	return newQTextCodec_U(unsafe.Pointer(ret))
}

func (this *QTextStream) SetAutoDetectUnicode(enabled bool) {
	C.QTextStream_SetAutoDetectUnicode(this.h, (C.bool)(enabled))
}

func (this *QTextStream) AutoDetectUnicode() bool {
	ret := C.QTextStream_AutoDetectUnicode(this.h)
	return (bool)(ret)
}

func (this *QTextStream) SetGenerateByteOrderMark(generate bool) {
	C.QTextStream_SetGenerateByteOrderMark(this.h, (C.bool)(generate))
}

func (this *QTextStream) GenerateByteOrderMark() bool {
	ret := C.QTextStream_GenerateByteOrderMark(this.h)
	return (bool)(ret)
}

func (this *QTextStream) SetLocale(locale *QLocale) {
	C.QTextStream_SetLocale(this.h, locale.cPointer())
}

func (this *QTextStream) Locale() *QLocale {
	ret := C.QTextStream_Locale(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQLocale(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QLocale) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextStream) SetDevice(device *QIODevice) {
	C.QTextStream_SetDevice(this.h, device.cPointer())
}

func (this *QTextStream) Device() *QIODevice {
	ret := C.QTextStream_Device(this.h)
	return newQIODevice_U(unsafe.Pointer(ret))
}

func (this *QTextStream) String() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QTextStream_String(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextStream) ResetStatus() {
	C.QTextStream_ResetStatus(this.h)
}

func (this *QTextStream) AtEnd() bool {
	ret := C.QTextStream_AtEnd(this.h)
	return (bool)(ret)
}

func (this *QTextStream) Reset() {
	C.QTextStream_Reset(this.h)
}

func (this *QTextStream) Flush() {
	C.QTextStream_Flush(this.h)
}

func (this *QTextStream) Seek(pos int64) bool {
	ret := C.QTextStream_Seek(this.h, (C.int64_t)(pos))
	return (bool)(ret)
}

func (this *QTextStream) Pos() int64 {
	ret := C.QTextStream_Pos(this.h)
	return (int64)(ret)
}

func (this *QTextStream) SkipWhiteSpace() {
	C.QTextStream_SkipWhiteSpace(this.h)
}

func (this *QTextStream) ReadLine() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QTextStream_ReadLine(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextStream) ReadAll() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QTextStream_ReadAll(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextStream) Read(maxlen int64) string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QTextStream_Read(this.h, (C.int64_t)(maxlen), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextStream) SetPadChar(ch QChar) {
	C.QTextStream_SetPadChar(this.h, ch.cPointer())
}

func (this *QTextStream) PadChar() *QChar {
	ret := C.QTextStream_PadChar(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQChar(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QChar) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextStream) SetFieldWidth(width int) {
	C.QTextStream_SetFieldWidth(this.h, (C.int)(width))
}

func (this *QTextStream) FieldWidth() int {
	ret := C.QTextStream_FieldWidth(this.h)
	return (int)(ret)
}

func (this *QTextStream) SetIntegerBase(base int) {
	C.QTextStream_SetIntegerBase(this.h, (C.int)(base))
}

func (this *QTextStream) IntegerBase() int {
	ret := C.QTextStream_IntegerBase(this.h)
	return (int)(ret)
}

func (this *QTextStream) SetRealNumberPrecision(precision int) {
	C.QTextStream_SetRealNumberPrecision(this.h, (C.int)(precision))
}

func (this *QTextStream) RealNumberPrecision() int {
	ret := C.QTextStream_RealNumberPrecision(this.h)
	return (int)(ret)
}

func (this *QTextStream) OperatorShiftRight(ch *QChar) *QTextStream {
	ret := C.QTextStream_OperatorShiftRight(this.h, ch.cPointer())
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithCh(ch *byte) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithCh(this.h, (*C.char)(unsafe.Pointer(ch)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithShort(i *int16) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithShort(this.h, (*C.int16_t)(unsafe.Pointer(i)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithUnsignedshort(i *uint16) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithUnsignedshort(this.h, (*C.uint16_t)(unsafe.Pointer(i)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithInt(i *int) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithInt(this.h, (*C.int)(unsafe.Pointer(i)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithUnsignedint(i *uint) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithUnsignedint(this.h, (*C.uint)(unsafe.Pointer(i)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithLong(i *int64) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithLong(this.h, (*C.long)(unsafe.Pointer(i)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithUnsignedlong(i *uint64) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithUnsignedlong(this.h, (*C.ulong)(unsafe.Pointer(i)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithQlonglong(i *int64) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithQlonglong(this.h, (*C.int64_t)(unsafe.Pointer(i)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithQulonglong(i *uint64) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithQulonglong(this.h, (*C.uint64_t)(unsafe.Pointer(i)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithFloat(f *float32) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithFloat(this.h, (*C.float)(unsafe.Pointer(f)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithDouble(f *float64) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithDouble(this.h, (*C.double)(unsafe.Pointer(f)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithQString(s string) *QTextStream {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	ret := C.QTextStream_OperatorShiftRightWithQString(this.h, s_Cstring, C.ulong(len(s)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithArray(array *QByteArray) *QTextStream {
	ret := C.QTextStream_OperatorShiftRightWithArray(this.h, array.cPointer())
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftRightWithChar(c string) *QTextStream {
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	ret := C.QTextStream_OperatorShiftRightWithChar(this.h, c_Cstring)
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeft(ch QChar) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeft(this.h, ch.cPointer())
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithCh(ch byte) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithCh(this.h, (C.char)(ch))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithShort(i int16) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithShort(this.h, (C.int16_t)(i))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithUnsignedshort(i uint16) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithUnsignedshort(this.h, (C.uint16_t)(i))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithInt(i int) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithInt(this.h, (C.int)(i))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithUnsignedint(i uint) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithUnsignedint(this.h, (C.uint)(i))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithLong(i int64) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithLong(this.h, (C.long)(i))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithUnsignedlong(i uint64) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithUnsignedlong(this.h, (C.ulong)(i))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithQlonglong(i int64) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithQlonglong(this.h, (C.int64_t)(i))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithQulonglong(i uint64) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithQulonglong(this.h, (C.uint64_t)(i))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithFloat(f float32) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithFloat(this.h, (C.float)(f))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithDouble(f float64) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithDouble(this.h, (C.double)(f))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithQString(s string) *QTextStream {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	ret := C.QTextStream_OperatorShiftLeftWithQString(this.h, s_Cstring, C.ulong(len(s)))
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithArray(array *QByteArray) *QTextStream {
	ret := C.QTextStream_OperatorShiftLeftWithArray(this.h, array.cPointer())
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) OperatorShiftLeftWithChar(c string) *QTextStream {
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	ret := C.QTextStream_OperatorShiftLeftWithChar(this.h, c_Cstring)
	return newQTextStream_U(unsafe.Pointer(ret))
}

func (this *QTextStream) ReadLine1(maxlen int64) string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QTextStream_ReadLine1(this.h, (C.int64_t)(maxlen), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextStream) Delete() {
	C.QTextStream_Delete(this.h)
}