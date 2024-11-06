package qt6

/*

#include "gen_qpluginloader.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QPluginLoader struct {
	h *C.QPluginLoader
	*QObject
}

func (this *QPluginLoader) cPointer() *C.QPluginLoader {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QPluginLoader) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

func newQPluginLoader(h *C.QPluginLoader) *QPluginLoader {
	if h == nil {
		return nil
	}
	return &QPluginLoader{h: h, QObject: UnsafeNewQObject(unsafe.Pointer(h))}
}

func UnsafeNewQPluginLoader(h unsafe.Pointer) *QPluginLoader {
	return newQPluginLoader((*C.QPluginLoader)(h))
}

// NewQPluginLoader constructs a new QPluginLoader object.
func NewQPluginLoader() *QPluginLoader {
	ret := C.QPluginLoader_new()
	return newQPluginLoader(ret)
}

// NewQPluginLoader2 constructs a new QPluginLoader object.
func NewQPluginLoader2(fileName string) *QPluginLoader {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	ret := C.QPluginLoader_new2(fileName_ms)
	return newQPluginLoader(ret)
}

// NewQPluginLoader3 constructs a new QPluginLoader object.
func NewQPluginLoader3(parent *QObject) *QPluginLoader {
	ret := C.QPluginLoader_new3(parent.cPointer())
	return newQPluginLoader(ret)
}

// NewQPluginLoader4 constructs a new QPluginLoader object.
func NewQPluginLoader4(fileName string, parent *QObject) *QPluginLoader {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	ret := C.QPluginLoader_new4(fileName_ms, parent.cPointer())
	return newQPluginLoader(ret)
}

func (this *QPluginLoader) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QPluginLoader_MetaObject(this.h)))
}

func (this *QPluginLoader) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QPluginLoader_Metacast(this.h, param1_Cstring))
}

func QPluginLoader_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QPluginLoader_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPluginLoader) Instance() *QObject {
	return UnsafeNewQObject(unsafe.Pointer(C.QPluginLoader_Instance(this.h)))
}

func (this *QPluginLoader) MetaData() *QJsonObject {
	_ret := C.QPluginLoader_MetaData(this.h)
	_goptr := newQJsonObject(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QPluginLoader_StaticInstances() []*QObject {
	var _ma C.struct_miqt_array = C.QPluginLoader_StaticInstances()
	_ret := make([]*QObject, int(_ma.len))
	_outCast := (*[0xffff]*C.QObject)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = UnsafeNewQObject(unsafe.Pointer(_outCast[i]))
	}
	return _ret
}

func QPluginLoader_StaticPlugins() []QStaticPlugin {
	var _ma C.struct_miqt_array = C.QPluginLoader_StaticPlugins()
	_ret := make([]QStaticPlugin, int(_ma.len))
	_outCast := (*[0xffff]*C.QStaticPlugin)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_ret := _outCast[i]
		_lv_goptr := newQStaticPlugin(_lv_ret)
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QPluginLoader) Load() bool {
	return (bool)(C.QPluginLoader_Load(this.h))
}

func (this *QPluginLoader) Unload() bool {
	return (bool)(C.QPluginLoader_Unload(this.h))
}

func (this *QPluginLoader) IsLoaded() bool {
	return (bool)(C.QPluginLoader_IsLoaded(this.h))
}

func (this *QPluginLoader) SetFileName(fileName string) {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	C.QPluginLoader_SetFileName(this.h, fileName_ms)
}

func (this *QPluginLoader) FileName() string {
	var _ms C.struct_miqt_string = C.QPluginLoader_FileName(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPluginLoader) ErrorString() string {
	var _ms C.struct_miqt_string = C.QPluginLoader_ErrorString(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPluginLoader) SetLoadHints(loadHints QLibrary__LoadHint) {
	C.QPluginLoader_SetLoadHints(this.h, (C.int)(loadHints))
}

func (this *QPluginLoader) LoadHints() QLibrary__LoadHint {
	return (QLibrary__LoadHint)(C.QPluginLoader_LoadHints(this.h))
}

func QPluginLoader_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QPluginLoader_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPluginLoader_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QPluginLoader_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QPluginLoader) Delete() {
	C.QPluginLoader_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QPluginLoader) GoGC() {
	runtime.SetFinalizer(this, func(this *QPluginLoader) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}