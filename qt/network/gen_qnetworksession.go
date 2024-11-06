package network

/*

#include "gen_qnetworksession.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QNetworkSession__State int

const (
	QNetworkSession__Invalid      QNetworkSession__State = 0
	QNetworkSession__NotAvailable QNetworkSession__State = 1
	QNetworkSession__Connecting   QNetworkSession__State = 2
	QNetworkSession__Connected    QNetworkSession__State = 3
	QNetworkSession__Closing      QNetworkSession__State = 4
	QNetworkSession__Disconnected QNetworkSession__State = 5
	QNetworkSession__Roaming      QNetworkSession__State = 6
)

type QNetworkSession__SessionError int

const (
	QNetworkSession__UnknownSessionError        QNetworkSession__SessionError = 0
	QNetworkSession__SessionAbortedError        QNetworkSession__SessionError = 1
	QNetworkSession__RoamingError               QNetworkSession__SessionError = 2
	QNetworkSession__OperationNotSupportedError QNetworkSession__SessionError = 3
	QNetworkSession__InvalidConfigurationError  QNetworkSession__SessionError = 4
)

type QNetworkSession__UsagePolicy int

const (
	QNetworkSession__NoPolicy                  QNetworkSession__UsagePolicy = 0
	QNetworkSession__NoBackgroundTrafficPolicy QNetworkSession__UsagePolicy = 1
)

type QNetworkSession struct {
	h *C.QNetworkSession
	*qt.QObject
}

func (this *QNetworkSession) cPointer() *C.QNetworkSession {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QNetworkSession) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

func newQNetworkSession(h *C.QNetworkSession) *QNetworkSession {
	if h == nil {
		return nil
	}
	return &QNetworkSession{h: h, QObject: qt.UnsafeNewQObject(unsafe.Pointer(h))}
}

func UnsafeNewQNetworkSession(h unsafe.Pointer) *QNetworkSession {
	return newQNetworkSession((*C.QNetworkSession)(h))
}

// NewQNetworkSession constructs a new QNetworkSession object.
func NewQNetworkSession(connConfig *QNetworkConfiguration) *QNetworkSession {
	ret := C.QNetworkSession_new(connConfig.cPointer())
	return newQNetworkSession(ret)
}

// NewQNetworkSession2 constructs a new QNetworkSession object.
func NewQNetworkSession2(connConfig *QNetworkConfiguration, parent *qt.QObject) *QNetworkSession {
	ret := C.QNetworkSession_new2(connConfig.cPointer(), (*C.QObject)(parent.UnsafePointer()))
	return newQNetworkSession(ret)
}

func (this *QNetworkSession) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QNetworkSession_MetaObject(this.h)))
}

func (this *QNetworkSession) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QNetworkSession_Metacast(this.h, param1_Cstring))
}

func QNetworkSession_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkSession_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_TrUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkSession) IsOpen() bool {
	return (bool)(C.QNetworkSession_IsOpen(this.h))
}

func (this *QNetworkSession) Configuration() *QNetworkConfiguration {
	_ret := C.QNetworkSession_Configuration(this.h)
	_goptr := newQNetworkConfiguration(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkSession) Interface() *QNetworkInterface {
	_ret := C.QNetworkSession_Interface(this.h)
	_goptr := newQNetworkInterface(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkSession) State() QNetworkSession__State {
	return (QNetworkSession__State)(C.QNetworkSession_State(this.h))
}

func (this *QNetworkSession) Error() QNetworkSession__SessionError {
	return (QNetworkSession__SessionError)(C.QNetworkSession_Error(this.h))
}

func (this *QNetworkSession) ErrorString() string {
	var _ms C.struct_miqt_string = C.QNetworkSession_ErrorString(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkSession) SessionProperty(key string) *qt.QVariant {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	_ret := C.QNetworkSession_SessionProperty(this.h, key_ms)
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(_ret))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkSession) SetSessionProperty(key string, value *qt.QVariant) {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	C.QNetworkSession_SetSessionProperty(this.h, key_ms, (*C.QVariant)(value.UnsafePointer()))
}

func (this *QNetworkSession) BytesWritten() uint64 {
	return (uint64)(C.QNetworkSession_BytesWritten(this.h))
}

func (this *QNetworkSession) BytesReceived() uint64 {
	return (uint64)(C.QNetworkSession_BytesReceived(this.h))
}

func (this *QNetworkSession) ActiveTime() uint64 {
	return (uint64)(C.QNetworkSession_ActiveTime(this.h))
}

func (this *QNetworkSession) UsagePolicies() QNetworkSession__UsagePolicy {
	return (QNetworkSession__UsagePolicy)(C.QNetworkSession_UsagePolicies(this.h))
}

func (this *QNetworkSession) WaitForOpened() bool {
	return (bool)(C.QNetworkSession_WaitForOpened(this.h))
}

func (this *QNetworkSession) Open() {
	C.QNetworkSession_Open(this.h)
}

func (this *QNetworkSession) Close() {
	C.QNetworkSession_Close(this.h)
}

func (this *QNetworkSession) Stop() {
	C.QNetworkSession_Stop(this.h)
}

func (this *QNetworkSession) Migrate() {
	C.QNetworkSession_Migrate(this.h)
}

func (this *QNetworkSession) Ignore() {
	C.QNetworkSession_Ignore(this.h)
}

func (this *QNetworkSession) Accept() {
	C.QNetworkSession_Accept(this.h)
}

func (this *QNetworkSession) Reject() {
	C.QNetworkSession_Reject(this.h)
}

func (this *QNetworkSession) StateChanged(param1 QNetworkSession__State) {
	C.QNetworkSession_StateChanged(this.h, (C.int)(param1))
}
func (this *QNetworkSession) OnStateChanged(slot func(param1 QNetworkSession__State)) {
	C.QNetworkSession_connect_StateChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_StateChanged
func miqt_exec_callback_QNetworkSession_StateChanged(cb C.intptr_t, param1 C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 QNetworkSession__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QNetworkSession__State)(param1)

	gofunc(slotval1)
}

func (this *QNetworkSession) Opened() {
	C.QNetworkSession_Opened(this.h)
}
func (this *QNetworkSession) OnOpened(slot func()) {
	C.QNetworkSession_connect_Opened(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_Opened
func miqt_exec_callback_QNetworkSession_Opened(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkSession) Closed() {
	C.QNetworkSession_Closed(this.h)
}
func (this *QNetworkSession) OnClosed(slot func()) {
	C.QNetworkSession_connect_Closed(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_Closed
func miqt_exec_callback_QNetworkSession_Closed(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkSession) ErrorWithQNetworkSessionSessionError(param1 QNetworkSession__SessionError) {
	C.QNetworkSession_ErrorWithQNetworkSessionSessionError(this.h, (C.int)(param1))
}
func (this *QNetworkSession) OnErrorWithQNetworkSessionSessionError(slot func(param1 QNetworkSession__SessionError)) {
	C.QNetworkSession_connect_ErrorWithQNetworkSessionSessionError(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_ErrorWithQNetworkSessionSessionError
func miqt_exec_callback_QNetworkSession_ErrorWithQNetworkSessionSessionError(cb C.intptr_t, param1 C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 QNetworkSession__SessionError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QNetworkSession__SessionError)(param1)

	gofunc(slotval1)
}

func (this *QNetworkSession) PreferredConfigurationChanged(config *QNetworkConfiguration, isSeamless bool) {
	C.QNetworkSession_PreferredConfigurationChanged(this.h, config.cPointer(), (C.bool)(isSeamless))
}
func (this *QNetworkSession) OnPreferredConfigurationChanged(slot func(config *QNetworkConfiguration, isSeamless bool)) {
	C.QNetworkSession_connect_PreferredConfigurationChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_PreferredConfigurationChanged
func miqt_exec_callback_QNetworkSession_PreferredConfigurationChanged(cb C.intptr_t, config *C.QNetworkConfiguration, isSeamless C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(config *QNetworkConfiguration, isSeamless bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQNetworkConfiguration(unsafe.Pointer(config))
	slotval2 := (bool)(isSeamless)

	gofunc(slotval1, slotval2)
}

func (this *QNetworkSession) NewConfigurationActivated() {
	C.QNetworkSession_NewConfigurationActivated(this.h)
}
func (this *QNetworkSession) OnNewConfigurationActivated(slot func()) {
	C.QNetworkSession_connect_NewConfigurationActivated(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_NewConfigurationActivated
func miqt_exec_callback_QNetworkSession_NewConfigurationActivated(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkSession) UsagePoliciesChanged(usagePolicies QNetworkSession__UsagePolicy) {
	C.QNetworkSession_UsagePoliciesChanged(this.h, (C.int)(usagePolicies))
}
func (this *QNetworkSession) OnUsagePoliciesChanged(slot func(usagePolicies QNetworkSession__UsagePolicy)) {
	C.QNetworkSession_connect_UsagePoliciesChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_UsagePoliciesChanged
func miqt_exec_callback_QNetworkSession_UsagePoliciesChanged(cb C.intptr_t, usagePolicies C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(usagePolicies QNetworkSession__UsagePolicy))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QNetworkSession__UsagePolicy)(usagePolicies)

	gofunc(slotval1)
}

func QNetworkSession_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkSession_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkSession_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkSession_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkSession) WaitForOpened1(msecs int) bool {
	return (bool)(C.QNetworkSession_WaitForOpened1(this.h, (C.int)(msecs)))
}

// Delete this object from C++ memory.
func (this *QNetworkSession) Delete() {
	C.QNetworkSession_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QNetworkSession) GoGC() {
	runtime.SetFinalizer(this, func(this *QNetworkSession) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}