package multimedia

/*

#include "gen_qaudiosink.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QAudioSink struct {
	h *C.QAudioSink
	*qt6.QObject
}

func (this *QAudioSink) cPointer() *C.QAudioSink {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QAudioSink) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

func newQAudioSink(h *C.QAudioSink) *QAudioSink {
	if h == nil {
		return nil
	}
	return &QAudioSink{h: h, QObject: qt6.UnsafeNewQObject(unsafe.Pointer(h))}
}

func UnsafeNewQAudioSink(h unsafe.Pointer) *QAudioSink {
	return newQAudioSink((*C.QAudioSink)(h))
}

// NewQAudioSink constructs a new QAudioSink object.
func NewQAudioSink() *QAudioSink {
	ret := C.QAudioSink_new()
	return newQAudioSink(ret)
}

// NewQAudioSink2 constructs a new QAudioSink object.
func NewQAudioSink2(audioDeviceInfo *QAudioDevice) *QAudioSink {
	ret := C.QAudioSink_new2(audioDeviceInfo.cPointer())
	return newQAudioSink(ret)
}

// NewQAudioSink3 constructs a new QAudioSink object.
func NewQAudioSink3(format *QAudioFormat) *QAudioSink {
	ret := C.QAudioSink_new3(format.cPointer())
	return newQAudioSink(ret)
}

// NewQAudioSink4 constructs a new QAudioSink object.
func NewQAudioSink4(format *QAudioFormat, parent *qt6.QObject) *QAudioSink {
	ret := C.QAudioSink_new4(format.cPointer(), (*C.QObject)(parent.UnsafePointer()))
	return newQAudioSink(ret)
}

// NewQAudioSink5 constructs a new QAudioSink object.
func NewQAudioSink5(audioDeviceInfo *QAudioDevice, format *QAudioFormat) *QAudioSink {
	ret := C.QAudioSink_new5(audioDeviceInfo.cPointer(), format.cPointer())
	return newQAudioSink(ret)
}

// NewQAudioSink6 constructs a new QAudioSink object.
func NewQAudioSink6(audioDeviceInfo *QAudioDevice, format *QAudioFormat, parent *qt6.QObject) *QAudioSink {
	ret := C.QAudioSink_new6(audioDeviceInfo.cPointer(), format.cPointer(), (*C.QObject)(parent.UnsafePointer()))
	return newQAudioSink(ret)
}

func (this *QAudioSink) MetaObject() *qt6.QMetaObject {
	return qt6.UnsafeNewQMetaObject(unsafe.Pointer(C.QAudioSink_MetaObject(this.h)))
}

func (this *QAudioSink) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QAudioSink_Metacast(this.h, param1_Cstring))
}

func QAudioSink_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAudioSink_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioSink) IsNull() bool {
	return (bool)(C.QAudioSink_IsNull(this.h))
}

func (this *QAudioSink) Format() *QAudioFormat {
	_ret := C.QAudioSink_Format(this.h)
	_goptr := newQAudioFormat(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioSink) Start(device *qt6.QIODevice) {
	C.QAudioSink_Start(this.h, (*C.QIODevice)(device.UnsafePointer()))
}

func (this *QAudioSink) Start2() *qt6.QIODevice {
	return qt6.UnsafeNewQIODevice(unsafe.Pointer(C.QAudioSink_Start2(this.h)))
}

func (this *QAudioSink) Stop() {
	C.QAudioSink_Stop(this.h)
}

func (this *QAudioSink) Reset() {
	C.QAudioSink_Reset(this.h)
}

func (this *QAudioSink) Suspend() {
	C.QAudioSink_Suspend(this.h)
}

func (this *QAudioSink) Resume() {
	C.QAudioSink_Resume(this.h)
}

func (this *QAudioSink) SetBufferSize(bytes int64) {
	C.QAudioSink_SetBufferSize(this.h, (C.ptrdiff_t)(bytes))
}

func (this *QAudioSink) BufferSize() int64 {
	return (int64)(C.QAudioSink_BufferSize(this.h))
}

func (this *QAudioSink) BytesFree() int64 {
	return (int64)(C.QAudioSink_BytesFree(this.h))
}

func (this *QAudioSink) ProcessedUSecs() int64 {
	return (int64)(C.QAudioSink_ProcessedUSecs(this.h))
}

func (this *QAudioSink) ElapsedUSecs() int64 {
	return (int64)(C.QAudioSink_ElapsedUSecs(this.h))
}

func (this *QAudioSink) Error() QAudio__Error {
	return (QAudio__Error)(C.QAudioSink_Error(this.h))
}

func (this *QAudioSink) State() QAudio__State {
	return (QAudio__State)(C.QAudioSink_State(this.h))
}

func (this *QAudioSink) SetVolume(volume float64) {
	C.QAudioSink_SetVolume(this.h, (C.double)(volume))
}

func (this *QAudioSink) Volume() float64 {
	return (float64)(C.QAudioSink_Volume(this.h))
}

func (this *QAudioSink) StateChanged(state QAudio__State) {
	C.QAudioSink_StateChanged(this.h, (C.int)(state))
}
func (this *QAudioSink) OnStateChanged(slot func(state QAudio__State)) {
	C.QAudioSink_connect_StateChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioSink_StateChanged
func miqt_exec_callback_QAudioSink_StateChanged(cb C.intptr_t, state C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(state QAudio__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAudio__State)(state)

	gofunc(slotval1)
}

func QAudioSink_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAudioSink_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAudioSink_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAudioSink_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QAudioSink) Delete() {
	C.QAudioSink_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QAudioSink) GoGC() {
	runtime.SetFinalizer(this, func(this *QAudioSink) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}