#include <QAudioRoleControl>
#include <QList>
#include <QMetaObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qaudiorolecontrol.h>
#include "gen_qaudiorolecontrol.h"
#include "_cgo_export.h"

QMetaObject* QAudioRoleControl_MetaObject(const QAudioRoleControl* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAudioRoleControl_Metacast(QAudioRoleControl* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAudioRoleControl_Tr(const char* s) {
	QString _ret = QAudioRoleControl::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAudioRoleControl_TrUtf8(const char* s) {
	QString _ret = QAudioRoleControl::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QAudioRoleControl_AudioRole(const QAudioRoleControl* self) {
	QAudio::Role _ret = self->audioRole();
	return static_cast<int>(_ret);
}

void QAudioRoleControl_SetAudioRole(QAudioRoleControl* self, int role) {
	self->setAudioRole(static_cast<QAudio::Role>(role));
}

struct miqt_array /* of int */  QAudioRoleControl_SupportedAudioRoles(const QAudioRoleControl* self) {
	QList<QAudio::Role> _ret = self->supportedAudioRoles();
	// Convert QList<> from C++ memory to manually-managed C memory
	int* _arr = static_cast<int*>(malloc(sizeof(int) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QAudio::Role _lv_ret = _ret[i];
		_arr[i] = static_cast<int>(_lv_ret);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QAudioRoleControl_AudioRoleChanged(QAudioRoleControl* self, int role) {
	self->audioRoleChanged(static_cast<QAudio::Role>(role));
}

void QAudioRoleControl_connect_AudioRoleChanged(QAudioRoleControl* self, intptr_t slot) {
	QAudioRoleControl::connect(self, static_cast<void (QAudioRoleControl::*)(QAudio::Role)>(&QAudioRoleControl::audioRoleChanged), self, [=](QAudio::Role role) {
		QAudio::Role role_ret = role;
		int sigval1 = static_cast<int>(role_ret);
		miqt_exec_callback_QAudioRoleControl_AudioRoleChanged(slot, sigval1);
	});
}

struct miqt_string QAudioRoleControl_Tr2(const char* s, const char* c) {
	QString _ret = QAudioRoleControl::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAudioRoleControl_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAudioRoleControl::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAudioRoleControl_TrUtf82(const char* s, const char* c) {
	QString _ret = QAudioRoleControl::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAudioRoleControl_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QAudioRoleControl::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAudioRoleControl_Delete(QAudioRoleControl* self) {
	delete self;
}

