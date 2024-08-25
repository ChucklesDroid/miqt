package qt

/*

#cgo CFLAGS: -fPIC
#cgo pkg-config: Qt5Widgets
#include "gen_qtextlayout.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QTextInlineObject struct {
	h *C.QTextInlineObject
}

func (this *QTextInlineObject) cPointer() *C.QTextInlineObject {
	if this == nil {
		return nil
	}
	return this.h
}

func newQTextInlineObject(h *C.QTextInlineObject) *QTextInlineObject {
	return &QTextInlineObject{h: h}
}

func newQTextInlineObject_U(h unsafe.Pointer) *QTextInlineObject {
	return newQTextInlineObject((*C.QTextInlineObject)(h))
}

// NewQTextInlineObject constructs a new QTextInlineObject object.
func NewQTextInlineObject() *QTextInlineObject {
	ret := C.QTextInlineObject_new()
	return newQTextInlineObject(ret)
}

func (this *QTextInlineObject) IsValid() bool {
	ret := C.QTextInlineObject_IsValid(this.h)
	return (bool)(ret)
}

func (this *QTextInlineObject) Rect() *QRectF {
	ret := C.QTextInlineObject_Rect(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQRectF(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QRectF) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextInlineObject) Width() float64 {
	ret := C.QTextInlineObject_Width(this.h)
	return (float64)(ret)
}

func (this *QTextInlineObject) Ascent() float64 {
	ret := C.QTextInlineObject_Ascent(this.h)
	return (float64)(ret)
}

func (this *QTextInlineObject) Descent() float64 {
	ret := C.QTextInlineObject_Descent(this.h)
	return (float64)(ret)
}

func (this *QTextInlineObject) Height() float64 {
	ret := C.QTextInlineObject_Height(this.h)
	return (float64)(ret)
}

func (this *QTextInlineObject) SetWidth(w float64) {
	C.QTextInlineObject_SetWidth(this.h, (C.double)(w))
}

func (this *QTextInlineObject) SetAscent(a float64) {
	C.QTextInlineObject_SetAscent(this.h, (C.double)(a))
}

func (this *QTextInlineObject) SetDescent(d float64) {
	C.QTextInlineObject_SetDescent(this.h, (C.double)(d))
}

func (this *QTextInlineObject) TextPosition() int {
	ret := C.QTextInlineObject_TextPosition(this.h)
	return (int)(ret)
}

func (this *QTextInlineObject) FormatIndex() int {
	ret := C.QTextInlineObject_FormatIndex(this.h)
	return (int)(ret)
}

func (this *QTextInlineObject) Format() *QTextFormat {
	ret := C.QTextInlineObject_Format(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQTextFormat(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QTextFormat) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextInlineObject) Delete() {
	C.QTextInlineObject_Delete(this.h)
}

type QTextLayout struct {
	h *C.QTextLayout
}

func (this *QTextLayout) cPointer() *C.QTextLayout {
	if this == nil {
		return nil
	}
	return this.h
}

func newQTextLayout(h *C.QTextLayout) *QTextLayout {
	return &QTextLayout{h: h}
}

func newQTextLayout_U(h unsafe.Pointer) *QTextLayout {
	return newQTextLayout((*C.QTextLayout)(h))
}

// NewQTextLayout constructs a new QTextLayout object.
func NewQTextLayout() *QTextLayout {
	ret := C.QTextLayout_new()
	return newQTextLayout(ret)
}

// NewQTextLayout2 constructs a new QTextLayout object.
func NewQTextLayout2(text string) *QTextLayout {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	ret := C.QTextLayout_new2(text_Cstring, C.ulong(len(text)))
	return newQTextLayout(ret)
}

// NewQTextLayout3 constructs a new QTextLayout object.
func NewQTextLayout3(text string, font *QFont) *QTextLayout {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	ret := C.QTextLayout_new3(text_Cstring, C.ulong(len(text)), font.cPointer())
	return newQTextLayout(ret)
}

// NewQTextLayout4 constructs a new QTextLayout object.
func NewQTextLayout4(b *QTextBlock) *QTextLayout {
	ret := C.QTextLayout_new4(b.cPointer())
	return newQTextLayout(ret)
}

// NewQTextLayout5 constructs a new QTextLayout object.
func NewQTextLayout5(text string, font *QFont, paintdevice *QPaintDevice) *QTextLayout {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	ret := C.QTextLayout_new5(text_Cstring, C.ulong(len(text)), font.cPointer(), paintdevice.cPointer())
	return newQTextLayout(ret)
}

func (this *QTextLayout) SetFont(f *QFont) {
	C.QTextLayout_SetFont(this.h, f.cPointer())
}

func (this *QTextLayout) Font() *QFont {
	ret := C.QTextLayout_Font(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQFont(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QFont) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextLayout) SetRawFont(rawFont *QRawFont) {
	C.QTextLayout_SetRawFont(this.h, rawFont.cPointer())
}

func (this *QTextLayout) SetText(stringVal string) {
	stringVal_Cstring := C.CString(stringVal)
	defer C.free(unsafe.Pointer(stringVal_Cstring))
	C.QTextLayout_SetText(this.h, stringVal_Cstring, C.ulong(len(stringVal)))
}

func (this *QTextLayout) Text() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QTextLayout_Text(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextLayout) SetTextOption(option *QTextOption) {
	C.QTextLayout_SetTextOption(this.h, option.cPointer())
}

func (this *QTextLayout) TextOption() *QTextOption {
	ret := C.QTextLayout_TextOption(this.h)
	return newQTextOption_U(unsafe.Pointer(ret))
}

func (this *QTextLayout) SetPreeditArea(position int, text string) {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	C.QTextLayout_SetPreeditArea(this.h, (C.int)(position), text_Cstring, C.ulong(len(text)))
}

func (this *QTextLayout) PreeditAreaPosition() int {
	ret := C.QTextLayout_PreeditAreaPosition(this.h)
	return (int)(ret)
}

func (this *QTextLayout) PreeditAreaText() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QTextLayout_PreeditAreaText(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextLayout) ClearAdditionalFormats() {
	C.QTextLayout_ClearAdditionalFormats(this.h)
}

func (this *QTextLayout) ClearFormats() {
	C.QTextLayout_ClearFormats(this.h)
}

func (this *QTextLayout) SetCacheEnabled(enable bool) {
	C.QTextLayout_SetCacheEnabled(this.h, (C.bool)(enable))
}

func (this *QTextLayout) CacheEnabled() bool {
	ret := C.QTextLayout_CacheEnabled(this.h)
	return (bool)(ret)
}

func (this *QTextLayout) BeginLayout() {
	C.QTextLayout_BeginLayout(this.h)
}

func (this *QTextLayout) EndLayout() {
	C.QTextLayout_EndLayout(this.h)
}

func (this *QTextLayout) ClearLayout() {
	C.QTextLayout_ClearLayout(this.h)
}

func (this *QTextLayout) CreateLine() *QTextLine {
	ret := C.QTextLayout_CreateLine(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQTextLine(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QTextLine) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextLayout) LineCount() int {
	ret := C.QTextLayout_LineCount(this.h)
	return (int)(ret)
}

func (this *QTextLayout) LineAt(i int) *QTextLine {
	ret := C.QTextLayout_LineAt(this.h, (C.int)(i))
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQTextLine(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QTextLine) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextLayout) LineForTextPosition(pos int) *QTextLine {
	ret := C.QTextLayout_LineForTextPosition(this.h, (C.int)(pos))
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQTextLine(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QTextLine) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextLayout) IsValidCursorPosition(pos int) bool {
	ret := C.QTextLayout_IsValidCursorPosition(this.h, (C.int)(pos))
	return (bool)(ret)
}

func (this *QTextLayout) LeftCursorPosition(oldPos int) int {
	ret := C.QTextLayout_LeftCursorPosition(this.h, (C.int)(oldPos))
	return (int)(ret)
}

func (this *QTextLayout) RightCursorPosition(oldPos int) int {
	ret := C.QTextLayout_RightCursorPosition(this.h, (C.int)(oldPos))
	return (int)(ret)
}

func (this *QTextLayout) DrawCursor(p *QPainter, pos *QPointF, cursorPosition int) {
	C.QTextLayout_DrawCursor(this.h, p.cPointer(), pos.cPointer(), (C.int)(cursorPosition))
}

func (this *QTextLayout) DrawCursor2(p *QPainter, pos *QPointF, cursorPosition int, width int) {
	C.QTextLayout_DrawCursor2(this.h, p.cPointer(), pos.cPointer(), (C.int)(cursorPosition), (C.int)(width))
}

func (this *QTextLayout) Position() *QPointF {
	ret := C.QTextLayout_Position(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPointF(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPointF) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextLayout) SetPosition(p *QPointF) {
	C.QTextLayout_SetPosition(this.h, p.cPointer())
}

func (this *QTextLayout) BoundingRect() *QRectF {
	ret := C.QTextLayout_BoundingRect(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQRectF(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QRectF) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextLayout) MinimumWidth() float64 {
	ret := C.QTextLayout_MinimumWidth(this.h)
	return (float64)(ret)
}

func (this *QTextLayout) MaximumWidth() float64 {
	ret := C.QTextLayout_MaximumWidth(this.h)
	return (float64)(ret)
}

func (this *QTextLayout) GlyphRuns() []QGlyphRun {
	var _out **C.QGlyphRun = nil
	var _out_len C.size_t = 0
	C.QTextLayout_GlyphRuns(this.h, &_out, &_out_len)
	ret := make([]QGlyphRun, int(_out_len))
	_outCast := (*[0xffff]*C.QGlyphRun)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQGlyphRun(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextLayout) SetFlags(flags int) {
	C.QTextLayout_SetFlags(this.h, (C.int)(flags))
}

func (this *QTextLayout) GlyphRuns1(from int) []QGlyphRun {
	var _out **C.QGlyphRun = nil
	var _out_len C.size_t = 0
	C.QTextLayout_GlyphRuns1(this.h, (C.int)(from), &_out, &_out_len)
	ret := make([]QGlyphRun, int(_out_len))
	_outCast := (*[0xffff]*C.QGlyphRun)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQGlyphRun(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextLayout) GlyphRuns2(from int, length int) []QGlyphRun {
	var _out **C.QGlyphRun = nil
	var _out_len C.size_t = 0
	C.QTextLayout_GlyphRuns2(this.h, (C.int)(from), (C.int)(length), &_out, &_out_len)
	ret := make([]QGlyphRun, int(_out_len))
	_outCast := (*[0xffff]*C.QGlyphRun)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQGlyphRun(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextLayout) Delete() {
	C.QTextLayout_Delete(this.h)
}

type QTextLine struct {
	h *C.QTextLine
}

func (this *QTextLine) cPointer() *C.QTextLine {
	if this == nil {
		return nil
	}
	return this.h
}

func newQTextLine(h *C.QTextLine) *QTextLine {
	return &QTextLine{h: h}
}

func newQTextLine_U(h unsafe.Pointer) *QTextLine {
	return newQTextLine((*C.QTextLine)(h))
}

// NewQTextLine constructs a new QTextLine object.
func NewQTextLine() *QTextLine {
	ret := C.QTextLine_new()
	return newQTextLine(ret)
}

func (this *QTextLine) IsValid() bool {
	ret := C.QTextLine_IsValid(this.h)
	return (bool)(ret)
}

func (this *QTextLine) Rect() *QRectF {
	ret := C.QTextLine_Rect(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQRectF(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QRectF) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextLine) X() float64 {
	ret := C.QTextLine_X(this.h)
	return (float64)(ret)
}

func (this *QTextLine) Y() float64 {
	ret := C.QTextLine_Y(this.h)
	return (float64)(ret)
}

func (this *QTextLine) Width() float64 {
	ret := C.QTextLine_Width(this.h)
	return (float64)(ret)
}

func (this *QTextLine) Ascent() float64 {
	ret := C.QTextLine_Ascent(this.h)
	return (float64)(ret)
}

func (this *QTextLine) Descent() float64 {
	ret := C.QTextLine_Descent(this.h)
	return (float64)(ret)
}

func (this *QTextLine) Height() float64 {
	ret := C.QTextLine_Height(this.h)
	return (float64)(ret)
}

func (this *QTextLine) Leading() float64 {
	ret := C.QTextLine_Leading(this.h)
	return (float64)(ret)
}

func (this *QTextLine) SetLeadingIncluded(included bool) {
	C.QTextLine_SetLeadingIncluded(this.h, (C.bool)(included))
}

func (this *QTextLine) LeadingIncluded() bool {
	ret := C.QTextLine_LeadingIncluded(this.h)
	return (bool)(ret)
}

func (this *QTextLine) NaturalTextWidth() float64 {
	ret := C.QTextLine_NaturalTextWidth(this.h)
	return (float64)(ret)
}

func (this *QTextLine) HorizontalAdvance() float64 {
	ret := C.QTextLine_HorizontalAdvance(this.h)
	return (float64)(ret)
}

func (this *QTextLine) NaturalTextRect() *QRectF {
	ret := C.QTextLine_NaturalTextRect(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQRectF(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QRectF) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextLine) SetLineWidth(width float64) {
	C.QTextLine_SetLineWidth(this.h, (C.double)(width))
}

func (this *QTextLine) SetNumColumns(columns int) {
	C.QTextLine_SetNumColumns(this.h, (C.int)(columns))
}

func (this *QTextLine) SetNumColumns2(columns int, alignmentWidth float64) {
	C.QTextLine_SetNumColumns2(this.h, (C.int)(columns), (C.double)(alignmentWidth))
}

func (this *QTextLine) SetPosition(pos *QPointF) {
	C.QTextLine_SetPosition(this.h, pos.cPointer())
}

func (this *QTextLine) Position() *QPointF {
	ret := C.QTextLine_Position(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPointF(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPointF) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextLine) TextStart() int {
	ret := C.QTextLine_TextStart(this.h)
	return (int)(ret)
}

func (this *QTextLine) TextLength() int {
	ret := C.QTextLine_TextLength(this.h)
	return (int)(ret)
}

func (this *QTextLine) LineNumber() int {
	ret := C.QTextLine_LineNumber(this.h)
	return (int)(ret)
}

func (this *QTextLine) GlyphRuns() []QGlyphRun {
	var _out **C.QGlyphRun = nil
	var _out_len C.size_t = 0
	C.QTextLine_GlyphRuns(this.h, &_out, &_out_len)
	ret := make([]QGlyphRun, int(_out_len))
	_outCast := (*[0xffff]*C.QGlyphRun)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQGlyphRun(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextLine) GlyphRuns1(from int) []QGlyphRun {
	var _out **C.QGlyphRun = nil
	var _out_len C.size_t = 0
	C.QTextLine_GlyphRuns1(this.h, (C.int)(from), &_out, &_out_len)
	ret := make([]QGlyphRun, int(_out_len))
	_outCast := (*[0xffff]*C.QGlyphRun)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQGlyphRun(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextLine) GlyphRuns2(from int, length int) []QGlyphRun {
	var _out **C.QGlyphRun = nil
	var _out_len C.size_t = 0
	C.QTextLine_GlyphRuns2(this.h, (C.int)(from), (C.int)(length), &_out, &_out_len)
	ret := make([]QGlyphRun, int(_out_len))
	_outCast := (*[0xffff]*C.QGlyphRun)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQGlyphRun(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextLine) Delete() {
	C.QTextLine_Delete(this.h)
}