package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type ViewModel_ITF interface {
	std_core.QObject_ITF
	ViewModel_PTR() *ViewModel
}

func (ptr *ViewModel) ViewModel_PTR() *ViewModel {
	return ptr
}

func (ptr *ViewModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *ViewModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromViewModel(ptr ViewModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ViewModel_PTR().Pointer()
	}
	return nil
}

func NewViewModelFromPointer(ptr unsafe.Pointer) (n *ViewModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ViewModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ViewModel:
			n = deduced

		case *std_core.QObject:
			n = &ViewModel{QObject: *deduced}

		default:
			n = new(ViewModel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackViewModel34b3c9_Constructor
func callbackViewModel34b3c9_Constructor(ptr unsafe.Pointer) {
	this := NewViewModelFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackViewModel34b3c9_InitStatus
func callbackViewModel34b3c9_InitStatus(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "initStatus"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewViewModelFromPointer(ptr).InitStatusDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ViewModel) ConnectInitStatus(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "initStatus"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "initStatus", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "initStatus", f)
		}
	}
}

func (ptr *ViewModel) DisconnectInitStatus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "initStatus")
	}
}

func (ptr *ViewModel) InitStatus() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ViewModel34b3c9_InitStatus(ptr.Pointer()))
	}
	return ""
}

func (ptr *ViewModel) InitStatusDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ViewModel34b3c9_InitStatusDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackViewModel34b3c9_SetInitStatus
func callbackViewModel34b3c9_SetInitStatus(ptr unsafe.Pointer, initStatus C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setInitStatus"); signal != nil {
		signal.(func(string))(cGoUnpackString(initStatus))
	} else {
		NewViewModelFromPointer(ptr).SetInitStatusDefault(cGoUnpackString(initStatus))
	}
}

func (ptr *ViewModel) ConnectSetInitStatus(f func(initStatus string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setInitStatus"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setInitStatus", func(initStatus string) {
				signal.(func(string))(initStatus)
				f(initStatus)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setInitStatus", f)
		}
	}
}

func (ptr *ViewModel) DisconnectSetInitStatus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setInitStatus")
	}
}

func (ptr *ViewModel) SetInitStatus(initStatus string) {
	if ptr.Pointer() != nil {
		var initStatusC *C.char
		if initStatus != "" {
			initStatusC = C.CString(initStatus)
			defer C.free(unsafe.Pointer(initStatusC))
		}
		C.ViewModel34b3c9_SetInitStatus(ptr.Pointer(), C.struct_Moc_PackedString{data: initStatusC, len: C.longlong(len(initStatus))})
	}
}

func (ptr *ViewModel) SetInitStatusDefault(initStatus string) {
	if ptr.Pointer() != nil {
		var initStatusC *C.char
		if initStatus != "" {
			initStatusC = C.CString(initStatus)
			defer C.free(unsafe.Pointer(initStatusC))
		}
		C.ViewModel34b3c9_SetInitStatusDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: initStatusC, len: C.longlong(len(initStatus))})
	}
}

//export callbackViewModel34b3c9_InitStatusChanged
func callbackViewModel34b3c9_InitStatusChanged(ptr unsafe.Pointer, initStatus C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "initStatusChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(initStatus))
	}

}

func (ptr *ViewModel) ConnectInitStatusChanged(f func(initStatus string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "initStatusChanged") {
			C.ViewModel34b3c9_ConnectInitStatusChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "initStatusChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "initStatusChanged", func(initStatus string) {
				signal.(func(string))(initStatus)
				f(initStatus)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "initStatusChanged", f)
		}
	}
}

func (ptr *ViewModel) DisconnectInitStatusChanged() {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9_DisconnectInitStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "initStatusChanged")
	}
}

func (ptr *ViewModel) InitStatusChanged(initStatus string) {
	if ptr.Pointer() != nil {
		var initStatusC *C.char
		if initStatus != "" {
			initStatusC = C.CString(initStatus)
			defer C.free(unsafe.Pointer(initStatusC))
		}
		C.ViewModel34b3c9_InitStatusChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: initStatusC, len: C.longlong(len(initStatus))})
	}
}

//export callbackViewModel34b3c9_ActivePage
func callbackViewModel34b3c9_ActivePage(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "activePage"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewViewModelFromPointer(ptr).ActivePageDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ViewModel) ConnectActivePage(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "activePage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activePage", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activePage", f)
		}
	}
}

func (ptr *ViewModel) DisconnectActivePage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "activePage")
	}
}

func (ptr *ViewModel) ActivePage() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ViewModel34b3c9_ActivePage(ptr.Pointer()))
	}
	return ""
}

func (ptr *ViewModel) ActivePageDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ViewModel34b3c9_ActivePageDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackViewModel34b3c9_SetActivePage
func callbackViewModel34b3c9_SetActivePage(ptr unsafe.Pointer, activePage C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setActivePage"); signal != nil {
		signal.(func(string))(cGoUnpackString(activePage))
	} else {
		NewViewModelFromPointer(ptr).SetActivePageDefault(cGoUnpackString(activePage))
	}
}

func (ptr *ViewModel) ConnectSetActivePage(f func(activePage string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setActivePage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setActivePage", func(activePage string) {
				signal.(func(string))(activePage)
				f(activePage)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setActivePage", f)
		}
	}
}

func (ptr *ViewModel) DisconnectSetActivePage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setActivePage")
	}
}

func (ptr *ViewModel) SetActivePage(activePage string) {
	if ptr.Pointer() != nil {
		var activePageC *C.char
		if activePage != "" {
			activePageC = C.CString(activePage)
			defer C.free(unsafe.Pointer(activePageC))
		}
		C.ViewModel34b3c9_SetActivePage(ptr.Pointer(), C.struct_Moc_PackedString{data: activePageC, len: C.longlong(len(activePage))})
	}
}

func (ptr *ViewModel) SetActivePageDefault(activePage string) {
	if ptr.Pointer() != nil {
		var activePageC *C.char
		if activePage != "" {
			activePageC = C.CString(activePage)
			defer C.free(unsafe.Pointer(activePageC))
		}
		C.ViewModel34b3c9_SetActivePageDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: activePageC, len: C.longlong(len(activePage))})
	}
}

//export callbackViewModel34b3c9_ActivePageChanged
func callbackViewModel34b3c9_ActivePageChanged(ptr unsafe.Pointer, activePage C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "activePageChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(activePage))
	}

}

func (ptr *ViewModel) ConnectActivePageChanged(f func(activePage string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activePageChanged") {
			C.ViewModel34b3c9_ConnectActivePageChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activePageChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activePageChanged", func(activePage string) {
				signal.(func(string))(activePage)
				f(activePage)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activePageChanged", f)
		}
	}
}

func (ptr *ViewModel) DisconnectActivePageChanged() {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9_DisconnectActivePageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activePageChanged")
	}
}

func (ptr *ViewModel) ActivePageChanged(activePage string) {
	if ptr.Pointer() != nil {
		var activePageC *C.char
		if activePage != "" {
			activePageC = C.CString(activePage)
			defer C.free(unsafe.Pointer(activePageC))
		}
		C.ViewModel34b3c9_ActivePageChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: activePageC, len: C.longlong(len(activePage))})
	}
}

func ViewModel_QRegisterMetaType() int {
	return int(int32(C.ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaType()))
}

func (ptr *ViewModel) QRegisterMetaType() int {
	return int(int32(C.ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaType()))
}

func ViewModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaType2(typeNameC)))
}

func (ptr *ViewModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaType2(typeNameC)))
}

func ViewModel_QmlRegisterType() int {
	return int(int32(C.ViewModel34b3c9_ViewModel34b3c9_QmlRegisterType()))
}

func (ptr *ViewModel) QmlRegisterType() int {
	return int(int32(C.ViewModel34b3c9_ViewModel34b3c9_QmlRegisterType()))
}

func ViewModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ViewModel34b3c9_ViewModel34b3c9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ViewModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ViewModel34b3c9_ViewModel34b3c9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ViewModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ViewModel34b3c9___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ViewModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ViewModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ViewModel34b3c9___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ViewModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ViewModel34b3c9___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ViewModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ViewModel) __findChildren_newList2() unsafe.Pointer {
	return C.ViewModel34b3c9___findChildren_newList2(ptr.Pointer())
}

func (ptr *ViewModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ViewModel34b3c9___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ViewModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ViewModel) __findChildren_newList3() unsafe.Pointer {
	return C.ViewModel34b3c9___findChildren_newList3(ptr.Pointer())
}

func (ptr *ViewModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ViewModel34b3c9___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ViewModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ViewModel) __findChildren_newList() unsafe.Pointer {
	return C.ViewModel34b3c9___findChildren_newList(ptr.Pointer())
}

func (ptr *ViewModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ViewModel34b3c9___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ViewModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ViewModel) __children_newList() unsafe.Pointer {
	return C.ViewModel34b3c9___children_newList(ptr.Pointer())
}

func NewViewModel(parent std_core.QObject_ITF) *ViewModel {
	tmpValue := NewViewModelFromPointer(C.ViewModel34b3c9_NewViewModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackViewModel34b3c9_DestroyViewModel
func callbackViewModel34b3c9_DestroyViewModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~ViewModel"); signal != nil {
		signal.(func())()
	} else {
		NewViewModelFromPointer(ptr).DestroyViewModelDefault()
	}
}

func (ptr *ViewModel) ConnectDestroyViewModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~ViewModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~ViewModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~ViewModel", f)
		}
	}
}

func (ptr *ViewModel) DisconnectDestroyViewModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~ViewModel")
	}
}

func (ptr *ViewModel) DestroyViewModel() {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9_DestroyViewModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ViewModel) DestroyViewModelDefault() {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9_DestroyViewModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackViewModel34b3c9_Event
func callbackViewModel34b3c9_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewViewModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ViewModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ViewModel34b3c9_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackViewModel34b3c9_EventFilter
func callbackViewModel34b3c9_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewViewModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ViewModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ViewModel34b3c9_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackViewModel34b3c9_ChildEvent
func callbackViewModel34b3c9_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewViewModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ViewModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackViewModel34b3c9_ConnectNotify
func callbackViewModel34b3c9_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewViewModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ViewModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackViewModel34b3c9_CustomEvent
func callbackViewModel34b3c9_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewViewModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ViewModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackViewModel34b3c9_DeleteLater
func callbackViewModel34b3c9_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewViewModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ViewModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackViewModel34b3c9_Destroyed
func callbackViewModel34b3c9_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackViewModel34b3c9_DisconnectNotify
func callbackViewModel34b3c9_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewViewModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ViewModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackViewModel34b3c9_ObjectNameChanged
func callbackViewModel34b3c9_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackViewModel34b3c9_TimerEvent
func callbackViewModel34b3c9_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewViewModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ViewModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ViewModel34b3c9_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
