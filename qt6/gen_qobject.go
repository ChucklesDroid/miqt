package qt6

/*

#include "gen_qobject.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QObjectData__ int

const (
	QObjectData__CheckForParentChildLoopsWarnDepth QObjectData__ = 4096
)

type QObjectData struct {
	h *C.QObjectData
}

func (this *QObjectData) cPointer() *C.QObjectData {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QObjectData) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

func newQObjectData(h *C.QObjectData) *QObjectData {
	if h == nil {
		return nil
	}
	return &QObjectData{h: h}
}

func UnsafeNewQObjectData(h unsafe.Pointer) *QObjectData {
	return newQObjectData((*C.QObjectData)(h))
}

func (this *QObjectData) DynamicMetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QObjectData_DynamicMetaObject(this.h)))
}

// Delete this object from C++ memory.
func (this *QObjectData) Delete() {
	C.QObjectData_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QObjectData) GoGC() {
	runtime.SetFinalizer(this, func(this *QObjectData) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QObject struct {
	h *C.QObject
}

func (this *QObject) cPointer() *C.QObject {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QObject) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

func newQObject(h *C.QObject) *QObject {
	if h == nil {
		return nil
	}
	return &QObject{h: h}
}

func UnsafeNewQObject(h unsafe.Pointer) *QObject {
	return newQObject((*C.QObject)(h))
}

// NewQObject constructs a new QObject object.
func NewQObject() *QObject {
	ret := C.QObject_new()
	return newQObject(ret)
}

// NewQObject2 constructs a new QObject object.
func NewQObject2(parent *QObject) *QObject {
	ret := C.QObject_new2(parent.cPointer())
	return newQObject(ret)
}

func (this *QObject) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QObject_MetaObject(this.h)))
}

func (this *QObject) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QObject_Metacast(this.h, param1_Cstring))
}

func QObject_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QObject_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QObject) Event(event *QEvent) bool {
	return (bool)(C.QObject_Event(this.h, event.cPointer()))
}

func (this *QObject) EventFilter(watched *QObject, event *QEvent) bool {
	return (bool)(C.QObject_EventFilter(this.h, watched.cPointer(), event.cPointer()))
}

func (this *QObject) ObjectName() string {
	var _ms C.struct_miqt_string = C.QObject_ObjectName(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QObject) SetObjectName(name QAnyStringView) {
	C.QObject_SetObjectName(this.h, name.cPointer())
}

func (this *QObject) IsWidgetType() bool {
	return (bool)(C.QObject_IsWidgetType(this.h))
}

func (this *QObject) IsWindowType() bool {
	return (bool)(C.QObject_IsWindowType(this.h))
}

func (this *QObject) IsQuickItemType() bool {
	return (bool)(C.QObject_IsQuickItemType(this.h))
}

func (this *QObject) SignalsBlocked() bool {
	return (bool)(C.QObject_SignalsBlocked(this.h))
}

func (this *QObject) BlockSignals(b bool) bool {
	return (bool)(C.QObject_BlockSignals(this.h, (C.bool)(b)))
}

func (this *QObject) Thread() *QThread {
	return UnsafeNewQThread(unsafe.Pointer(C.QObject_Thread(this.h)))
}

func (this *QObject) MoveToThread(thread *QThread) {
	C.QObject_MoveToThread(this.h, thread.cPointer())
}

func (this *QObject) StartTimer(interval int) int {
	return (int)(C.QObject_StartTimer(this.h, (C.int)(interval)))
}

func (this *QObject) KillTimer(id int) {
	C.QObject_KillTimer(this.h, (C.int)(id))
}

func (this *QObject) Children() []*QObject {
	var _ma C.struct_miqt_array = C.QObject_Children(this.h)
	_ret := make([]*QObject, int(_ma.len))
	_outCast := (*[0xffff]*C.QObject)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = UnsafeNewQObject(unsafe.Pointer(_outCast[i]))
	}
	return _ret
}

func (this *QObject) SetParent(parent *QObject) {
	C.QObject_SetParent(this.h, parent.cPointer())
}

func (this *QObject) InstallEventFilter(filterObj *QObject) {
	C.QObject_InstallEventFilter(this.h, filterObj.cPointer())
}

func (this *QObject) RemoveEventFilter(obj *QObject) {
	C.QObject_RemoveEventFilter(this.h, obj.cPointer())
}

func QObject_Connect(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	_ret := C.QObject_Connect(sender.cPointer(), signal.cPointer(), receiver.cPointer(), method.cPointer())
	_goptr := newQMetaObject__Connection(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QObject) Connect2(sender *QObject, signal string, member string) *QMetaObject__Connection {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	_ret := C.QObject_Connect2(this.h, sender.cPointer(), signal_Cstring, member_Cstring)
	_goptr := newQMetaObject__Connection(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QObject_Disconnect(sender *QObject, signal *QMetaMethod, receiver *QObject, member *QMetaMethod) bool {
	return (bool)(C.QObject_Disconnect(sender.cPointer(), signal.cPointer(), receiver.cPointer(), member.cPointer()))
}

func QObject_DisconnectWithQMetaObjectConnection(param1 *QMetaObject__Connection) bool {
	return (bool)(C.QObject_DisconnectWithQMetaObjectConnection(param1.cPointer()))
}

func (this *QObject) DumpObjectTree() {
	C.QObject_DumpObjectTree(this.h)
}

func (this *QObject) DumpObjectInfo() {
	C.QObject_DumpObjectInfo(this.h)
}

func (this *QObject) SetProperty(name string, value *QVariant) bool {
	name_Cstring := C.CString(name)
	defer C.free(unsafe.Pointer(name_Cstring))
	return (bool)(C.QObject_SetProperty(this.h, name_Cstring, value.cPointer()))
}

func (this *QObject) Property(name string) *QVariant {
	name_Cstring := C.CString(name)
	defer C.free(unsafe.Pointer(name_Cstring))
	_ret := C.QObject_Property(this.h, name_Cstring)
	_goptr := newQVariant(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QObject) DynamicPropertyNames() [][]byte {
	var _ma C.struct_miqt_array = C.QObject_DynamicPropertyNames(this.h)
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray C.struct_miqt_string = _outCast[i]
		_lv_ret := C.GoBytes(unsafe.Pointer(_lv_bytearray.data), C.int(int64(_lv_bytearray.len)))
		C.free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QObject) BindingStorage() *QBindingStorage {
	return UnsafeNewQBindingStorage(unsafe.Pointer(C.QObject_BindingStorage(this.h)))
}

func (this *QObject) BindingStorage2() *QBindingStorage {
	return UnsafeNewQBindingStorage(unsafe.Pointer(C.QObject_BindingStorage2(this.h)))
}

func (this *QObject) Destroyed() {
	C.QObject_Destroyed(this.h)
}
func (this *QObject) OnDestroyed(slot func()) {
	C.QObject_connect_Destroyed(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObject_Destroyed
func miqt_exec_callback_QObject_Destroyed(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QObject) Parent() *QObject {
	return UnsafeNewQObject(unsafe.Pointer(C.QObject_Parent(this.h)))
}

func (this *QObject) Inherits(classname string) bool {
	classname_Cstring := C.CString(classname)
	defer C.free(unsafe.Pointer(classname_Cstring))
	return (bool)(C.QObject_Inherits(this.h, classname_Cstring))
}

func (this *QObject) DeleteLater() {
	C.QObject_DeleteLater(this.h)
}

func QObject_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QObject_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QObject_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QObject_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QObject) StartTimer2(interval int, timerType TimerType) int {
	return (int)(C.QObject_StartTimer2(this.h, (C.int)(interval), (C.int)(timerType)))
}

func QObject_Connect5(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, typeVal ConnectionType) *QMetaObject__Connection {
	_ret := C.QObject_Connect5(sender.cPointer(), signal.cPointer(), receiver.cPointer(), method.cPointer(), (C.int)(typeVal))
	_goptr := newQMetaObject__Connection(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QObject) Connect4(sender *QObject, signal string, member string, typeVal ConnectionType) *QMetaObject__Connection {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	_ret := C.QObject_Connect4(this.h, sender.cPointer(), signal_Cstring, member_Cstring, (C.int)(typeVal))
	_goptr := newQMetaObject__Connection(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QObject) Destroyed1(param1 *QObject) {
	C.QObject_Destroyed1(this.h, param1.cPointer())
}
func (this *QObject) OnDestroyed1(slot func(param1 *QObject)) {
	C.QObject_connect_Destroyed1(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObject_Destroyed1
func miqt_exec_callback_QObject_Destroyed1(cb C.intptr_t, param1 *C.QObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQObject(unsafe.Pointer(param1))

	gofunc(slotval1)
}

// Delete this object from C++ memory.
func (this *QObject) Delete() {
	C.QObject_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QObject) GoGC() {
	runtime.SetFinalizer(this, func(this *QObject) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QSignalBlocker struct {
	h *C.QSignalBlocker
}

func (this *QSignalBlocker) cPointer() *C.QSignalBlocker {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QSignalBlocker) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

func newQSignalBlocker(h *C.QSignalBlocker) *QSignalBlocker {
	if h == nil {
		return nil
	}
	return &QSignalBlocker{h: h}
}

func UnsafeNewQSignalBlocker(h unsafe.Pointer) *QSignalBlocker {
	return newQSignalBlocker((*C.QSignalBlocker)(h))
}

// NewQSignalBlocker constructs a new QSignalBlocker object.
func NewQSignalBlocker(o *QObject) *QSignalBlocker {
	ret := C.QSignalBlocker_new(o.cPointer())
	return newQSignalBlocker(ret)
}

// NewQSignalBlocker2 constructs a new QSignalBlocker object.
func NewQSignalBlocker2(o *QObject) *QSignalBlocker {
	ret := C.QSignalBlocker_new2(o.cPointer())
	return newQSignalBlocker(ret)
}

func (this *QSignalBlocker) Reblock() {
	C.QSignalBlocker_Reblock(this.h)
}

func (this *QSignalBlocker) Unblock() {
	C.QSignalBlocker_Unblock(this.h)
}

// Delete this object from C++ memory.
func (this *QSignalBlocker) Delete() {
	C.QSignalBlocker_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QSignalBlocker) GoGC() {
	runtime.SetFinalizer(this, func(this *QSignalBlocker) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}