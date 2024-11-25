#include <QPaintDevice>
#include <QPainter>
#include <QPalette>
#include <QPixmap>
#include <QRect>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionComplex>
#include <QStylePainter>
#include <QWidget>
#include <qstylepainter.h>
#include "gen_qstylepainter.h"
#include "_cgo_export.h"

void QStylePainter_new(QWidget* w, QStylePainter** outptr_QStylePainter, QPainter** outptr_QPainter) {
	QStylePainter* ret = new QStylePainter(w);
	*outptr_QStylePainter = ret;
	*outptr_QPainter = static_cast<QPainter*>(ret);
}

void QStylePainter_new2(QStylePainter** outptr_QStylePainter, QPainter** outptr_QPainter) {
	QStylePainter* ret = new QStylePainter();
	*outptr_QStylePainter = ret;
	*outptr_QPainter = static_cast<QPainter*>(ret);
}

void QStylePainter_new3(QPaintDevice* pd, QWidget* w, QStylePainter** outptr_QStylePainter, QPainter** outptr_QPainter) {
	QStylePainter* ret = new QStylePainter(pd, w);
	*outptr_QStylePainter = ret;
	*outptr_QPainter = static_cast<QPainter*>(ret);
}

bool QStylePainter_Begin(QStylePainter* self, QWidget* w) {
	return self->begin(w);
}

bool QStylePainter_Begin2(QStylePainter* self, QPaintDevice* pd, QWidget* w) {
	return self->begin(pd, w);
}

void QStylePainter_DrawPrimitive(QStylePainter* self, int pe, QStyleOption* opt) {
	self->drawPrimitive(static_cast<QStyle::PrimitiveElement>(pe), *opt);
}

void QStylePainter_DrawControl(QStylePainter* self, int ce, QStyleOption* opt) {
	self->drawControl(static_cast<QStyle::ControlElement>(ce), *opt);
}

void QStylePainter_DrawComplexControl(QStylePainter* self, int cc, QStyleOptionComplex* opt) {
	self->drawComplexControl(static_cast<QStyle::ComplexControl>(cc), *opt);
}

void QStylePainter_DrawItemText(QStylePainter* self, QRect* r, int flags, QPalette* pal, bool enabled, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->drawItemText(*r, static_cast<int>(flags), *pal, enabled, text_QString);
}

void QStylePainter_DrawItemPixmap(QStylePainter* self, QRect* r, int flags, QPixmap* pixmap) {
	self->drawItemPixmap(*r, static_cast<int>(flags), *pixmap);
}

QStyle* QStylePainter_Style(const QStylePainter* self) {
	return self->style();
}

void QStylePainter_DrawItemText6(QStylePainter* self, QRect* r, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->drawItemText(*r, static_cast<int>(flags), *pal, enabled, text_QString, static_cast<QPalette::ColorRole>(textRole));
}

void QStylePainter_Delete(QStylePainter* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QStylePainter*>( self );
	} else {
		delete self;
	}
}

