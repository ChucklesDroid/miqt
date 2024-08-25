#include "gen_qitemdelegate.h"
#include "qitemdelegate.h"

#include <QAbstractItemModel>
#include <QItemDelegate>
#include <QItemEditorFactory>
#include <QMetaObject>
#include <QModelIndex>
#include <QObject>
#include <QPainter>
#include <QSize>
#include <QString>
#include <QStyleOptionViewItem>
#include <QWidget>


extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

QItemDelegate* QItemDelegate_new() {
	return new QItemDelegate();
}

QItemDelegate* QItemDelegate_new2(QObject* parent) {
	return new QItemDelegate(parent);
}

QMetaObject* QItemDelegate_MetaObject(QItemDelegate* self) {
	return (QMetaObject*) self->metaObject();
}

void QItemDelegate_Tr(char* s, char** _out, int* _out_Strlen) {
	QString ret = QItemDelegate::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QItemDelegate_TrUtf8(char* s, char** _out, int* _out_Strlen) {
	QString ret = QItemDelegate::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

bool QItemDelegate_HasClipping(QItemDelegate* self) {
	return self->hasClipping();
}

void QItemDelegate_SetClipping(QItemDelegate* self, bool clip) {
	self->setClipping(clip);
}

void QItemDelegate_Paint(QItemDelegate* self, QPainter* painter, QStyleOptionViewItem* option, QModelIndex* index) {
	self->paint(painter, *option, *index);
}

QSize* QItemDelegate_SizeHint(QItemDelegate* self, QStyleOptionViewItem* option, QModelIndex* index) {
	QSize ret = self->sizeHint(*option, *index);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSize*>(new QSize(ret));
}

QWidget* QItemDelegate_CreateEditor(QItemDelegate* self, QWidget* parent, QStyleOptionViewItem* option, QModelIndex* index) {
	return self->createEditor(parent, *option, *index);
}

void QItemDelegate_SetEditorData(QItemDelegate* self, QWidget* editor, QModelIndex* index) {
	self->setEditorData(editor, *index);
}

void QItemDelegate_SetModelData(QItemDelegate* self, QWidget* editor, QAbstractItemModel* model, QModelIndex* index) {
	self->setModelData(editor, model, *index);
}

void QItemDelegate_UpdateEditorGeometry(QItemDelegate* self, QWidget* editor, QStyleOptionViewItem* option, QModelIndex* index) {
	self->updateEditorGeometry(editor, *option, *index);
}

QItemEditorFactory* QItemDelegate_ItemEditorFactory(QItemDelegate* self) {
	return self->itemEditorFactory();
}

void QItemDelegate_SetItemEditorFactory(QItemDelegate* self, QItemEditorFactory* factory) {
	self->setItemEditorFactory(factory);
}

void QItemDelegate_Tr2(char* s, char* c, char** _out, int* _out_Strlen) {
	QString ret = QItemDelegate::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QItemDelegate_Tr3(char* s, char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QItemDelegate::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QItemDelegate_TrUtf82(char* s, char* c, char** _out, int* _out_Strlen) {
	QString ret = QItemDelegate::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QItemDelegate_TrUtf83(char* s, char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QItemDelegate::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QItemDelegate_Delete(QItemDelegate* self) {
	delete self;
}
