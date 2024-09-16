#include <QAbstractItemModel>
#include <QAbstractProxyModel>
#include <QList>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include "qabstractproxymodel.h"
#include "gen_qabstractproxymodel.h"
#include "_cgo_export.h"

QMetaObject* QAbstractProxyModel_MetaObject(const QAbstractProxyModel* self) {
	return (QMetaObject*) self->metaObject();
}

struct miqt_string* QAbstractProxyModel_Tr(const char* s) {
	QString _ret = QAbstractProxyModel::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QAbstractProxyModel_TrUtf8(const char* s) {
	QString _ret = QAbstractProxyModel::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QAbstractProxyModel_SetSourceModel(QAbstractProxyModel* self, QAbstractItemModel* sourceModel) {
	self->setSourceModel(sourceModel);
}

QAbstractItemModel* QAbstractProxyModel_SourceModel(const QAbstractProxyModel* self) {
	return self->sourceModel();
}

QModelIndex* QAbstractProxyModel_MapToSource(const QAbstractProxyModel* self, QModelIndex* proxyIndex) {
	QModelIndex _ret = self->mapToSource(*proxyIndex);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QModelIndex*>(new QModelIndex(_ret));
}

QModelIndex* QAbstractProxyModel_MapFromSource(const QAbstractProxyModel* self, QModelIndex* sourceIndex) {
	QModelIndex _ret = self->mapFromSource(*sourceIndex);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QModelIndex*>(new QModelIndex(_ret));
}

bool QAbstractProxyModel_Submit(QAbstractProxyModel* self) {
	return self->submit();
}

void QAbstractProxyModel_Revert(QAbstractProxyModel* self) {
	self->revert();
}

QVariant* QAbstractProxyModel_Data(const QAbstractProxyModel* self, QModelIndex* proxyIndex) {
	QVariant _ret = self->data(*proxyIndex);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QVariant*>(new QVariant(_ret));
}

QVariant* QAbstractProxyModel_HeaderData(const QAbstractProxyModel* self, int section, uintptr_t orientation) {
	QVariant _ret = self->headerData(static_cast<int>(section), static_cast<Qt::Orientation>(orientation));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QVariant*>(new QVariant(_ret));
}

int QAbstractProxyModel_Flags(const QAbstractProxyModel* self, QModelIndex* index) {
	Qt::ItemFlags _ret = self->flags(*index);
	return static_cast<int>(_ret);
}

bool QAbstractProxyModel_SetData(QAbstractProxyModel* self, QModelIndex* index, QVariant* value) {
	return self->setData(*index, *value);
}

bool QAbstractProxyModel_SetHeaderData(QAbstractProxyModel* self, int section, uintptr_t orientation, QVariant* value) {
	return self->setHeaderData(static_cast<int>(section), static_cast<Qt::Orientation>(orientation), *value);
}

QModelIndex* QAbstractProxyModel_Buddy(const QAbstractProxyModel* self, QModelIndex* index) {
	QModelIndex _ret = self->buddy(*index);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QModelIndex*>(new QModelIndex(_ret));
}

bool QAbstractProxyModel_CanFetchMore(const QAbstractProxyModel* self, QModelIndex* parent) {
	return self->canFetchMore(*parent);
}

void QAbstractProxyModel_FetchMore(QAbstractProxyModel* self, QModelIndex* parent) {
	self->fetchMore(*parent);
}

void QAbstractProxyModel_Sort(QAbstractProxyModel* self, int column) {
	self->sort(static_cast<int>(column));
}

QSize* QAbstractProxyModel_Span(const QAbstractProxyModel* self, QModelIndex* index) {
	QSize _ret = self->span(*index);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSize*>(new QSize(_ret));
}

bool QAbstractProxyModel_HasChildren(const QAbstractProxyModel* self) {
	return self->hasChildren();
}

QModelIndex* QAbstractProxyModel_Sibling(const QAbstractProxyModel* self, int row, int column, QModelIndex* idx) {
	QModelIndex _ret = self->sibling(static_cast<int>(row), static_cast<int>(column), *idx);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QModelIndex*>(new QModelIndex(_ret));
}

QMimeData* QAbstractProxyModel_MimeData(const QAbstractProxyModel* self, struct miqt_array* /* of QModelIndex* */ indexes) {
	QList<QModelIndex> indexes_QList;
	indexes_QList.reserve(indexes->len);
	QModelIndex** indexes_arr = static_cast<QModelIndex**>(indexes->data);
	for(size_t i = 0; i < indexes->len; ++i) {
		indexes_QList.push_back(*(indexes_arr[i]));
	}
	return self->mimeData(indexes_QList);
}

bool QAbstractProxyModel_CanDropMimeData(const QAbstractProxyModel* self, QMimeData* data, uintptr_t action, int row, int column, QModelIndex* parent) {
	return self->canDropMimeData(data, static_cast<Qt::DropAction>(action), static_cast<int>(row), static_cast<int>(column), *parent);
}

bool QAbstractProxyModel_DropMimeData(QAbstractProxyModel* self, QMimeData* data, uintptr_t action, int row, int column, QModelIndex* parent) {
	return self->dropMimeData(data, static_cast<Qt::DropAction>(action), static_cast<int>(row), static_cast<int>(column), *parent);
}

struct miqt_array* QAbstractProxyModel_MimeTypes(const QAbstractProxyModel* self) {
	QStringList _ret = self->mimeTypes();
	// Convert QStringList from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

int QAbstractProxyModel_SupportedDragActions(const QAbstractProxyModel* self) {
	Qt::DropActions _ret = self->supportedDragActions();
	return static_cast<int>(_ret);
}

int QAbstractProxyModel_SupportedDropActions(const QAbstractProxyModel* self) {
	Qt::DropActions _ret = self->supportedDropActions();
	return static_cast<int>(_ret);
}

struct miqt_string* QAbstractProxyModel_Tr2(const char* s, const char* c) {
	QString _ret = QAbstractProxyModel::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QAbstractProxyModel_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAbstractProxyModel::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QAbstractProxyModel_TrUtf82(const char* s, const char* c) {
	QString _ret = QAbstractProxyModel::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QAbstractProxyModel_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QAbstractProxyModel::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

QVariant* QAbstractProxyModel_Data2(const QAbstractProxyModel* self, QModelIndex* proxyIndex, int role) {
	QVariant _ret = self->data(*proxyIndex, static_cast<int>(role));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QVariant*>(new QVariant(_ret));
}

QVariant* QAbstractProxyModel_HeaderData3(const QAbstractProxyModel* self, int section, uintptr_t orientation, int role) {
	QVariant _ret = self->headerData(static_cast<int>(section), static_cast<Qt::Orientation>(orientation), static_cast<int>(role));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QVariant*>(new QVariant(_ret));
}

bool QAbstractProxyModel_SetData3(QAbstractProxyModel* self, QModelIndex* index, QVariant* value, int role) {
	return self->setData(*index, *value, static_cast<int>(role));
}

bool QAbstractProxyModel_SetHeaderData4(QAbstractProxyModel* self, int section, uintptr_t orientation, QVariant* value, int role) {
	return self->setHeaderData(static_cast<int>(section), static_cast<Qt::Orientation>(orientation), *value, static_cast<int>(role));
}

void QAbstractProxyModel_Sort2(QAbstractProxyModel* self, int column, uintptr_t order) {
	self->sort(static_cast<int>(column), static_cast<Qt::SortOrder>(order));
}

bool QAbstractProxyModel_HasChildren1(const QAbstractProxyModel* self, QModelIndex* parent) {
	return self->hasChildren(*parent);
}

void QAbstractProxyModel_Delete(QAbstractProxyModel* self) {
	delete self;
}

