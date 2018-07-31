

#pragma once

#ifndef GO_MOC_H
#define GO_MOC_H

#include <stdint.h>

#ifdef __cplusplus
class ViewModel34b3c9;
void ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
struct Moc_PackedString ViewModel34b3c9_InitStatus(void* ptr);
struct Moc_PackedString ViewModel34b3c9_InitStatusDefault(void* ptr);
void ViewModel34b3c9_SetInitStatus(void* ptr, struct Moc_PackedString initStatus);
void ViewModel34b3c9_SetInitStatusDefault(void* ptr, struct Moc_PackedString initStatus);
void ViewModel34b3c9_ConnectInitStatusChanged(void* ptr);
void ViewModel34b3c9_DisconnectInitStatusChanged(void* ptr);
void ViewModel34b3c9_InitStatusChanged(void* ptr, struct Moc_PackedString initStatus);
struct Moc_PackedString ViewModel34b3c9_ActivePage(void* ptr);
struct Moc_PackedString ViewModel34b3c9_ActivePageDefault(void* ptr);
void ViewModel34b3c9_SetActivePage(void* ptr, struct Moc_PackedString activePage);
void ViewModel34b3c9_SetActivePageDefault(void* ptr, struct Moc_PackedString activePage);
void ViewModel34b3c9_ConnectActivePageChanged(void* ptr);
void ViewModel34b3c9_DisconnectActivePageChanged(void* ptr);
void ViewModel34b3c9_ActivePageChanged(void* ptr, struct Moc_PackedString activePage);
int ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaType();
int ViewModel34b3c9_ViewModel34b3c9_QRegisterMetaType2(char* typeName);
int ViewModel34b3c9_ViewModel34b3c9_QmlRegisterType();
int ViewModel34b3c9_ViewModel34b3c9_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* ViewModel34b3c9___dynamicPropertyNames_atList(void* ptr, int i);
void ViewModel34b3c9___dynamicPropertyNames_setList(void* ptr, void* i);
void* ViewModel34b3c9___dynamicPropertyNames_newList(void* ptr);
void* ViewModel34b3c9___findChildren_atList2(void* ptr, int i);
void ViewModel34b3c9___findChildren_setList2(void* ptr, void* i);
void* ViewModel34b3c9___findChildren_newList2(void* ptr);
void* ViewModel34b3c9___findChildren_atList3(void* ptr, int i);
void ViewModel34b3c9___findChildren_setList3(void* ptr, void* i);
void* ViewModel34b3c9___findChildren_newList3(void* ptr);
void* ViewModel34b3c9___findChildren_atList(void* ptr, int i);
void ViewModel34b3c9___findChildren_setList(void* ptr, void* i);
void* ViewModel34b3c9___findChildren_newList(void* ptr);
void* ViewModel34b3c9___children_atList(void* ptr, int i);
void ViewModel34b3c9___children_setList(void* ptr, void* i);
void* ViewModel34b3c9___children_newList(void* ptr);
void* ViewModel34b3c9_NewViewModel(void* parent);
void ViewModel34b3c9_DestroyViewModel(void* ptr);
void ViewModel34b3c9_DestroyViewModelDefault(void* ptr);
char ViewModel34b3c9_EventDefault(void* ptr, void* e);
char ViewModel34b3c9_EventFilterDefault(void* ptr, void* watched, void* event);
void ViewModel34b3c9_ChildEventDefault(void* ptr, void* event);
void ViewModel34b3c9_ConnectNotifyDefault(void* ptr, void* sign);
void ViewModel34b3c9_CustomEventDefault(void* ptr, void* event);
void ViewModel34b3c9_DeleteLaterDefault(void* ptr);
void ViewModel34b3c9_DisconnectNotifyDefault(void* ptr, void* sign);
void ViewModel34b3c9_TimerEventDefault(void* ptr, void* event);
;

#ifdef __cplusplus
}
#endif

#endif