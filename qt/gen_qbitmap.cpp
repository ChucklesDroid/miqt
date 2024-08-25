#include "gen_qbitmap.h"
#include "qbitmap.h"

#include <QBitmap>
#include <QMatrix>
#include <QPixmap>
#include <QSize>
#include <QString>
#include <QTransform>


extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

QBitmap* QBitmap_new() {
	return new QBitmap();
}

QBitmap* QBitmap_new2(QPixmap* param1) {
	return new QBitmap(*param1);
}

QBitmap* QBitmap_new3(int w, int h) {
	return new QBitmap(static_cast<int>(w), static_cast<int>(h));
}

QBitmap* QBitmap_new4(QSize* param1) {
	return new QBitmap(*param1);
}

QBitmap* QBitmap_new5(const char* fileName, size_t fileName_Strlen) {
	QString fileName_QString = QString::fromUtf8(fileName, fileName_Strlen);
	return new QBitmap(fileName_QString);
}

QBitmap* QBitmap_new6(QBitmap* other) {
	return new QBitmap(*other);
}

QBitmap* QBitmap_new7(const char* fileName, size_t fileName_Strlen, char* format) {
	QString fileName_QString = QString::fromUtf8(fileName, fileName_Strlen);
	return new QBitmap(fileName_QString, format);
}

void QBitmap_OperatorAssign(QBitmap* self, QBitmap* other) {
	self->operator=(*other);
}

void QBitmap_OperatorAssignWithQPixmap(QBitmap* self, QPixmap* param1) {
	self->operator=(*param1);
}

void QBitmap_Swap(QBitmap* self, QBitmap* other) {
	self->swap(*other);
}

void QBitmap_Clear(QBitmap* self) {
	self->clear();
}

QBitmap* QBitmap_Transformed(QBitmap* self, QMatrix* param1) {
	QBitmap ret = self->transformed(*param1);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QBitmap*>(new QBitmap(ret));
}

QBitmap* QBitmap_TransformedWithMatrix(QBitmap* self, QTransform* matrix) {
	QBitmap ret = self->transformed(*matrix);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QBitmap*>(new QBitmap(ret));
}

void QBitmap_Delete(QBitmap* self) {
	delete self;
}
