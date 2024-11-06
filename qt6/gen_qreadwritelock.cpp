#include <QReadLocker>
#include <QReadWriteLock>
#include <QWriteLocker>
#include <qreadwritelock.h>
#include "gen_qreadwritelock.h"
#include "_cgo_export.h"

QReadWriteLock* QReadWriteLock_new() {
	return new QReadWriteLock();
}

QReadWriteLock* QReadWriteLock_new2(int recursionMode) {
	return new QReadWriteLock(static_cast<QReadWriteLock::RecursionMode>(recursionMode));
}

void QReadWriteLock_LockForRead(QReadWriteLock* self) {
	self->lockForRead();
}

bool QReadWriteLock_TryLockForRead(QReadWriteLock* self) {
	return self->tryLockForRead();
}

bool QReadWriteLock_TryLockForReadWithTimeout(QReadWriteLock* self, int timeout) {
	return self->tryLockForRead(static_cast<int>(timeout));
}

void QReadWriteLock_LockForWrite(QReadWriteLock* self) {
	self->lockForWrite();
}

bool QReadWriteLock_TryLockForWrite(QReadWriteLock* self) {
	return self->tryLockForWrite();
}

bool QReadWriteLock_TryLockForWriteWithTimeout(QReadWriteLock* self, int timeout) {
	return self->tryLockForWrite(static_cast<int>(timeout));
}

void QReadWriteLock_Unlock(QReadWriteLock* self) {
	self->unlock();
}

void QReadWriteLock_Delete(QReadWriteLock* self) {
	delete self;
}

QReadLocker* QReadLocker_new(QReadWriteLock* readWriteLock) {
	return new QReadLocker(readWriteLock);
}

void QReadLocker_Unlock(QReadLocker* self) {
	self->unlock();
}

void QReadLocker_Relock(QReadLocker* self) {
	self->relock();
}

QReadWriteLock* QReadLocker_ReadWriteLock(const QReadLocker* self) {
	return self->readWriteLock();
}

void QReadLocker_Delete(QReadLocker* self) {
	delete self;
}

QWriteLocker* QWriteLocker_new(QReadWriteLock* readWriteLock) {
	return new QWriteLocker(readWriteLock);
}

void QWriteLocker_Unlock(QWriteLocker* self) {
	self->unlock();
}

void QWriteLocker_Relock(QWriteLocker* self) {
	self->relock();
}

QReadWriteLock* QWriteLocker_ReadWriteLock(const QWriteLocker* self) {
	return self->readWriteLock();
}

void QWriteLocker_Delete(QWriteLocker* self) {
	delete self;
}
