#include <QCommandLineOption>
#include <QCommandLineParser>
#include <QCoreApplication>
#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include "qcommandlineparser.h"
#include "gen_qcommandlineparser.h"
#include "_cgo_export.h"

QCommandLineParser* QCommandLineParser_new() {
	return new QCommandLineParser();
}

struct miqt_string* QCommandLineParser_Tr(const char* sourceText) {
	QString _ret = QCommandLineParser::tr(sourceText);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QCommandLineParser_TrUtf8(const char* sourceText) {
	QString _ret = QCommandLineParser::trUtf8(sourceText);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QCommandLineParser_SetSingleDashWordOptionMode(QCommandLineParser* self, uintptr_t parsingMode) {
	self->setSingleDashWordOptionMode(static_cast<QCommandLineParser::SingleDashWordOptionMode>(parsingMode));
}

void QCommandLineParser_SetOptionsAfterPositionalArgumentsMode(QCommandLineParser* self, uintptr_t mode) {
	self->setOptionsAfterPositionalArgumentsMode(static_cast<QCommandLineParser::OptionsAfterPositionalArgumentsMode>(mode));
}

bool QCommandLineParser_AddOption(QCommandLineParser* self, QCommandLineOption* commandLineOption) {
	return self->addOption(*commandLineOption);
}

bool QCommandLineParser_AddOptions(QCommandLineParser* self, struct miqt_array* /* of QCommandLineOption* */ options) {
	QList<QCommandLineOption> options_QList;
	options_QList.reserve(options->len);
	QCommandLineOption** options_arr = static_cast<QCommandLineOption**>(options->data);
	for(size_t i = 0; i < options->len; ++i) {
		options_QList.push_back(*(options_arr[i]));
	}
	return self->addOptions(options_QList);
}

QCommandLineOption* QCommandLineParser_AddVersionOption(QCommandLineParser* self) {
	QCommandLineOption _ret = self->addVersionOption();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QCommandLineOption*>(new QCommandLineOption(_ret));
}

QCommandLineOption* QCommandLineParser_AddHelpOption(QCommandLineParser* self) {
	QCommandLineOption _ret = self->addHelpOption();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QCommandLineOption*>(new QCommandLineOption(_ret));
}

void QCommandLineParser_SetApplicationDescription(QCommandLineParser* self, struct miqt_string* description) {
	QString description_QString = QString::fromUtf8(&description->data, description->len);
	self->setApplicationDescription(description_QString);
}

struct miqt_string* QCommandLineParser_ApplicationDescription(const QCommandLineParser* self) {
	QString _ret = self->applicationDescription();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QCommandLineParser_AddPositionalArgument(QCommandLineParser* self, struct miqt_string* name, struct miqt_string* description) {
	QString name_QString = QString::fromUtf8(&name->data, name->len);
	QString description_QString = QString::fromUtf8(&description->data, description->len);
	self->addPositionalArgument(name_QString, description_QString);
}

void QCommandLineParser_ClearPositionalArguments(QCommandLineParser* self) {
	self->clearPositionalArguments();
}

void QCommandLineParser_Process(QCommandLineParser* self, struct miqt_array* /* of struct miqt_string* */ arguments) {
	QList<QString> arguments_QList;
	arguments_QList.reserve(arguments->len);
	miqt_string** arguments_arr = static_cast<miqt_string**>(arguments->data);
	for(size_t i = 0; i < arguments->len; ++i) {
		arguments_QList.push_back(QString::fromUtf8(& arguments_arr[i]->data, arguments_arr[i]->len));
	}
	self->process(arguments_QList);
}

void QCommandLineParser_ProcessWithApp(QCommandLineParser* self, QCoreApplication* app) {
	self->process(*app);
}

bool QCommandLineParser_Parse(QCommandLineParser* self, struct miqt_array* /* of struct miqt_string* */ arguments) {
	QList<QString> arguments_QList;
	arguments_QList.reserve(arguments->len);
	miqt_string** arguments_arr = static_cast<miqt_string**>(arguments->data);
	for(size_t i = 0; i < arguments->len; ++i) {
		arguments_QList.push_back(QString::fromUtf8(& arguments_arr[i]->data, arguments_arr[i]->len));
	}
	return self->parse(arguments_QList);
}

struct miqt_string* QCommandLineParser_ErrorText(const QCommandLineParser* self) {
	QString _ret = self->errorText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

bool QCommandLineParser_IsSet(const QCommandLineParser* self, struct miqt_string* name) {
	QString name_QString = QString::fromUtf8(&name->data, name->len);
	return self->isSet(name_QString);
}

struct miqt_string* QCommandLineParser_Value(const QCommandLineParser* self, struct miqt_string* name) {
	QString name_QString = QString::fromUtf8(&name->data, name->len);
	QString _ret = self->value(name_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_array* QCommandLineParser_Values(const QCommandLineParser* self, struct miqt_string* name) {
	QString name_QString = QString::fromUtf8(&name->data, name->len);
	QStringList _ret = self->values(name_QString);
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

bool QCommandLineParser_IsSetWithOption(const QCommandLineParser* self, QCommandLineOption* option) {
	return self->isSet(*option);
}

struct miqt_string* QCommandLineParser_ValueWithOption(const QCommandLineParser* self, QCommandLineOption* option) {
	QString _ret = self->value(*option);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_array* QCommandLineParser_ValuesWithOption(const QCommandLineParser* self, QCommandLineOption* option) {
	QStringList _ret = self->values(*option);
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

struct miqt_array* QCommandLineParser_PositionalArguments(const QCommandLineParser* self) {
	QStringList _ret = self->positionalArguments();
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

struct miqt_array* QCommandLineParser_OptionNames(const QCommandLineParser* self) {
	QStringList _ret = self->optionNames();
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

struct miqt_array* QCommandLineParser_UnknownOptionNames(const QCommandLineParser* self) {
	QStringList _ret = self->unknownOptionNames();
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

struct miqt_string* QCommandLineParser_HelpText(const QCommandLineParser* self) {
	QString _ret = self->helpText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QCommandLineParser_Tr2(const char* sourceText, const char* disambiguation) {
	QString _ret = QCommandLineParser::tr(sourceText, disambiguation);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QCommandLineParser_Tr3(const char* sourceText, const char* disambiguation, int n) {
	QString _ret = QCommandLineParser::tr(sourceText, disambiguation, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QCommandLineParser_TrUtf82(const char* sourceText, const char* disambiguation) {
	QString _ret = QCommandLineParser::trUtf8(sourceText, disambiguation);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QCommandLineParser_TrUtf83(const char* sourceText, const char* disambiguation, int n) {
	QString _ret = QCommandLineParser::trUtf8(sourceText, disambiguation, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QCommandLineParser_AddPositionalArgument3(QCommandLineParser* self, struct miqt_string* name, struct miqt_string* description, struct miqt_string* syntax) {
	QString name_QString = QString::fromUtf8(&name->data, name->len);
	QString description_QString = QString::fromUtf8(&description->data, description->len);
	QString syntax_QString = QString::fromUtf8(&syntax->data, syntax->len);
	self->addPositionalArgument(name_QString, description_QString, syntax_QString);
}

void QCommandLineParser_Delete(QCommandLineParser* self) {
	delete self;
}

