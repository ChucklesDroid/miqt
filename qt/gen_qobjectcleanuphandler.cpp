#include "gen_qobjectcleanuphandler.h"
#include "qobjectcleanuphandler.h"

#include <QMetaObject>
#include <QObject>
#include <QObjectCleanupHandler>
#include <QString>


extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

QObjectCleanupHandler* QObjectCleanupHandler_new() {
	return new QObjectCleanupHandler();
}

QMetaObject* QObjectCleanupHandler_MetaObject(QObjectCleanupHandler* self) {
	return (QMetaObject*) self->metaObject();
}

void QObjectCleanupHandler_Tr(char* s, char** _out, int* _out_Strlen) {
	QString ret = QObjectCleanupHandler::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QObjectCleanupHandler_TrUtf8(char* s, char** _out, int* _out_Strlen) {
	QString ret = QObjectCleanupHandler::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

QObject* QObjectCleanupHandler_Add(QObjectCleanupHandler* self, QObject* object) {
	return self->add(object);
}

void QObjectCleanupHandler_Remove(QObjectCleanupHandler* self, QObject* object) {
	self->remove(object);
}

bool QObjectCleanupHandler_IsEmpty(QObjectCleanupHandler* self) {
	return self->isEmpty();
}

void QObjectCleanupHandler_Clear(QObjectCleanupHandler* self) {
	self->clear();
}

void QObjectCleanupHandler_Tr2(char* s, char* c, char** _out, int* _out_Strlen) {
	QString ret = QObjectCleanupHandler::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QObjectCleanupHandler_Tr3(char* s, char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QObjectCleanupHandler::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QObjectCleanupHandler_TrUtf82(char* s, char* c, char** _out, int* _out_Strlen) {
	QString ret = QObjectCleanupHandler::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QObjectCleanupHandler_TrUtf83(char* s, char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QObjectCleanupHandler::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QObjectCleanupHandler_Delete(QObjectCleanupHandler* self) {
	delete self;
}
