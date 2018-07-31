

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>


class ViewModel34b3c9: public QObject
{
Q_OBJECT
Q_PROPERTY(QString initStatus READ initStatus WRITE setInitStatus NOTIFY initStatusChanged)
Q_PROPERTY(QString activePage READ activePage WRITE setActivePage NOTIFY activePageChanged)
public:
	ViewModel34b3c9(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaType();ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaTypes();callbackViewModel34b3c9_Constructor(this);};
	QString initStatus() { return ({ Moc_PackedString tempVal = callbackViewModel34b3c9_InitStatus(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setInitStatus(QString initStatus) { QByteArray tda7f7b = initStatus.toUtf8(); Moc_PackedString initStatusPacked = { const_cast<char*>(tda7f7b.prepend("WHITESPACE").constData()+10), tda7f7b.size()-10 };callbackViewModel34b3c9_SetInitStatus(this, initStatusPacked); };
	void Signal_InitStatusChanged(QString initStatus) { QByteArray tda7f7b = initStatus.toUtf8(); Moc_PackedString initStatusPacked = { const_cast<char*>(tda7f7b.prepend("WHITESPACE").constData()+10), tda7f7b.size()-10 };callbackViewModel34b3c9_InitStatusChanged(this, initStatusPacked); };
	QString activePage() { return ({ Moc_PackedString tempVal = callbackViewModel34b3c9_ActivePage(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setActivePage(QString activePage) { QByteArray t38639e = activePage.toUtf8(); Moc_PackedString activePagePacked = { const_cast<char*>(t38639e.prepend("WHITESPACE").constData()+10), t38639e.size()-10 };callbackViewModel34b3c9_SetActivePage(this, activePagePacked); };
	void Signal_ActivePageChanged(QString activePage) { QByteArray t38639e = activePage.toUtf8(); Moc_PackedString activePagePacked = { const_cast<char*>(t38639e.prepend("WHITESPACE").constData()+10), t38639e.size()-10 };callbackViewModel34b3c9_ActivePageChanged(this, activePagePacked); };
	 ~ViewModel34b3c9() { callbackViewModel34b3c9_DestroyViewModel(this); };
	bool event(QEvent * e) { return callbackViewModel34b3c9_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackViewModel34b3c9_EventFilter(this, watched, event) != 0; };
	
	void childEvent(QChildEvent * event) { callbackViewModel34b3c9_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackViewModel34b3c9_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackViewModel34b3c9_CustomEvent(this, event); };
	void deleteLater() { callbackViewModel34b3c9_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackViewModel34b3c9_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackViewModel34b3c9_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackViewModel34b3c9_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackViewModel34b3c9_TimerEvent(this, event); };
	
	QString initStatusDefault() { return _initStatus; };
	void setInitStatusDefault(QString p) { if (p != _initStatus) { _initStatus = p; initStatusChanged(_initStatus); } };
	QString activePageDefault() { return _activePage; };
	void setActivePageDefault(QString p) { if (p != _activePage) { _activePage = p; activePageChanged(_activePage); } };
signals:
	void initStatusChanged(QString initStatus);
	void activePageChanged(QString activePage);
public slots:
private:
	QString _initStatus;
	QString _activePage;
};

Q_DECLARE_METATYPE(ViewModel34b3c9*)


void ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

struct Moc_PackedString ViewModel34b3c9_InitStatus(void* ptr)
{
	return ({ QByteArray t1cd2d8 = static_cast<ViewModel34b3c9*>(ptr)->initStatus().toUtf8(); Moc_PackedString { const_cast<char*>(t1cd2d8.prepend("WHITESPACE").constData()+10), t1cd2d8.size()-10 }; });
}

struct Moc_PackedString ViewModel34b3c9_InitStatusDefault(void* ptr)
{
	return ({ QByteArray ta1e537 = static_cast<ViewModel34b3c9*>(ptr)->initStatusDefault().toUtf8(); Moc_PackedString { const_cast<char*>(ta1e537.prepend("WHITESPACE").constData()+10), ta1e537.size()-10 }; });
}

void ViewModel34b3c9_SetInitStatus(void* ptr, struct Moc_PackedString initStatus)
{
	static_cast<ViewModel34b3c9*>(ptr)->setInitStatus(QString::fromUtf8(initStatus.data, initStatus.len));
}

void ViewModel34b3c9_SetInitStatusDefault(void* ptr, struct Moc_PackedString initStatus)
{
	static_cast<ViewModel34b3c9*>(ptr)->setInitStatusDefault(QString::fromUtf8(initStatus.data, initStatus.len));
}

void ViewModel34b3c9_ConnectInitStatusChanged(void* ptr)
{
	QObject::connect(static_cast<ViewModel34b3c9*>(ptr), static_cast<void (ViewModel34b3c9::*)(QString)>(&ViewModel34b3c9::initStatusChanged), static_cast<ViewModel34b3c9*>(ptr), static_cast<void (ViewModel34b3c9::*)(QString)>(&ViewModel34b3c9::Signal_InitStatusChanged));
}

void ViewModel34b3c9_DisconnectInitStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<ViewModel34b3c9*>(ptr), static_cast<void (ViewModel34b3c9::*)(QString)>(&ViewModel34b3c9::initStatusChanged), static_cast<ViewModel34b3c9*>(ptr), static_cast<void (ViewModel34b3c9::*)(QString)>(&ViewModel34b3c9::Signal_InitStatusChanged));
}

void ViewModel34b3c9_InitStatusChanged(void* ptr, struct Moc_PackedString initStatus)
{
	static_cast<ViewModel34b3c9*>(ptr)->initStatusChanged(QString::fromUtf8(initStatus.data, initStatus.len));
}

struct Moc_PackedString ViewModel34b3c9_ActivePage(void* ptr)
{
	return ({ QByteArray t7ca24b = static_cast<ViewModel34b3c9*>(ptr)->activePage().toUtf8(); Moc_PackedString { const_cast<char*>(t7ca24b.prepend("WHITESPACE").constData()+10), t7ca24b.size()-10 }; });
}

struct Moc_PackedString ViewModel34b3c9_ActivePageDefault(void* ptr)
{
	return ({ QByteArray t183999 = static_cast<ViewModel34b3c9*>(ptr)->activePageDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t183999.prepend("WHITESPACE").constData()+10), t183999.size()-10 }; });
}

void ViewModel34b3c9_SetActivePage(void* ptr, struct Moc_PackedString activePage)
{
	static_cast<ViewModel34b3c9*>(ptr)->setActivePage(QString::fromUtf8(activePage.data, activePage.len));
}

void ViewModel34b3c9_SetActivePageDefault(void* ptr, struct Moc_PackedString activePage)
{
	static_cast<ViewModel34b3c9*>(ptr)->setActivePageDefault(QString::fromUtf8(activePage.data, activePage.len));
}

void ViewModel34b3c9_ConnectActivePageChanged(void* ptr)
{
	QObject::connect(static_cast<ViewModel34b3c9*>(ptr), static_cast<void (ViewModel34b3c9::*)(QString)>(&ViewModel34b3c9::activePageChanged), static_cast<ViewModel34b3c9*>(ptr), static_cast<void (ViewModel34b3c9::*)(QString)>(&ViewModel34b3c9::Signal_ActivePageChanged));
}

void ViewModel34b3c9_DisconnectActivePageChanged(void* ptr)
{
	QObject::disconnect(static_cast<ViewModel34b3c9*>(ptr), static_cast<void (ViewModel34b3c9::*)(QString)>(&ViewModel34b3c9::activePageChanged), static_cast<ViewModel34b3c9*>(ptr), static_cast<void (ViewModel34b3c9::*)(QString)>(&ViewModel34b3c9::Signal_ActivePageChanged));
}

void ViewModel34b3c9_ActivePageChanged(void* ptr, struct Moc_PackedString activePage)
{
	static_cast<ViewModel34b3c9*>(ptr)->activePageChanged(QString::fromUtf8(activePage.data, activePage.len));
}

int ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaType()
{
	return qRegisterMetaType<ViewModel34b3c9*>();
}

int ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ViewModel34b3c9*>(const_cast<const char*>(typeName));
}

int ViewModel34b3c9_ViewModel34b3c9_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ViewModel34b3c9>();
#else
	return 0;
#endif
}

int ViewModel34b3c9_ViewModel34b3c9_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ViewModel34b3c9>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* ViewModel34b3c9___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ViewModel34b3c9___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ViewModel34b3c9___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ViewModel34b3c9___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ViewModel34b3c9___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ViewModel34b3c9___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ViewModel34b3c9___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ViewModel34b3c9___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ViewModel34b3c9___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ViewModel34b3c9___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ViewModel34b3c9___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ViewModel34b3c9___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ViewModel34b3c9___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ViewModel34b3c9___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ViewModel34b3c9___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ViewModel34b3c9_NewViewModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new ViewModel34b3c9(static_cast<QWindow*>(parent));
	} else {
		return new ViewModel34b3c9(static_cast<QObject*>(parent));
	}
}

void ViewModel34b3c9_DestroyViewModel(void* ptr)
{
	static_cast<ViewModel34b3c9*>(ptr)->~ViewModel34b3c9();
}

void ViewModel34b3c9_DestroyViewModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char ViewModel34b3c9_EventDefault(void* ptr, void* e)
{
	return static_cast<ViewModel34b3c9*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char ViewModel34b3c9_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ViewModel34b3c9*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void ViewModel34b3c9_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ViewModel34b3c9*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void ViewModel34b3c9_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ViewModel34b3c9*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ViewModel34b3c9_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ViewModel34b3c9*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void ViewModel34b3c9_DeleteLaterDefault(void* ptr)
{
	static_cast<ViewModel34b3c9*>(ptr)->QObject::deleteLater();
}

void ViewModel34b3c9_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ViewModel34b3c9*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void ViewModel34b3c9_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ViewModel34b3c9*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
