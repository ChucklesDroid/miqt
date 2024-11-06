#ifndef GEN_QGRAPHICSITEM_H
#define GEN_QGRAPHICSITEM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractGraphicsShapeItem;
class QBrush;
class QColor;
class QCursor;
class QFont;
class QGraphicsEffect;
class QGraphicsEllipseItem;
class QGraphicsItem;
class QGraphicsItemGroup;
class QGraphicsLineItem;
class QGraphicsObject;
class QGraphicsPathItem;
class QGraphicsPixmapItem;
class QGraphicsPolygonItem;
class QGraphicsRectItem;
class QGraphicsScene;
class QGraphicsSimpleTextItem;
class QGraphicsTextItem;
class QGraphicsTransform;
class QGraphicsWidget;
class QLineF;
class QMetaObject;
class QPainter;
class QPainterPath;
class QPen;
class QPixmap;
class QPointF;
class QRectF;
class QRegion;
class QSize;
class QStyleOptionGraphicsItem;
class QTextCursor;
class QTextDocument;
class QTransform;
class QVariant;
class QWidget;
#else
typedef struct QAbstractGraphicsShapeItem QAbstractGraphicsShapeItem;
typedef struct QBrush QBrush;
typedef struct QColor QColor;
typedef struct QCursor QCursor;
typedef struct QFont QFont;
typedef struct QGraphicsEffect QGraphicsEffect;
typedef struct QGraphicsEllipseItem QGraphicsEllipseItem;
typedef struct QGraphicsItem QGraphicsItem;
typedef struct QGraphicsItemGroup QGraphicsItemGroup;
typedef struct QGraphicsLineItem QGraphicsLineItem;
typedef struct QGraphicsObject QGraphicsObject;
typedef struct QGraphicsPathItem QGraphicsPathItem;
typedef struct QGraphicsPixmapItem QGraphicsPixmapItem;
typedef struct QGraphicsPolygonItem QGraphicsPolygonItem;
typedef struct QGraphicsRectItem QGraphicsRectItem;
typedef struct QGraphicsScene QGraphicsScene;
typedef struct QGraphicsSimpleTextItem QGraphicsSimpleTextItem;
typedef struct QGraphicsTextItem QGraphicsTextItem;
typedef struct QGraphicsTransform QGraphicsTransform;
typedef struct QGraphicsWidget QGraphicsWidget;
typedef struct QLineF QLineF;
typedef struct QMetaObject QMetaObject;
typedef struct QPainter QPainter;
typedef struct QPainterPath QPainterPath;
typedef struct QPen QPen;
typedef struct QPixmap QPixmap;
typedef struct QPointF QPointF;
typedef struct QRectF QRectF;
typedef struct QRegion QRegion;
typedef struct QSize QSize;
typedef struct QStyleOptionGraphicsItem QStyleOptionGraphicsItem;
typedef struct QTextCursor QTextCursor;
typedef struct QTextDocument QTextDocument;
typedef struct QTransform QTransform;
typedef struct QVariant QVariant;
typedef struct QWidget QWidget;
#endif

QGraphicsScene* QGraphicsItem_Scene(const QGraphicsItem* self);
QGraphicsItem* QGraphicsItem_ParentItem(const QGraphicsItem* self);
QGraphicsItem* QGraphicsItem_TopLevelItem(const QGraphicsItem* self);
QGraphicsObject* QGraphicsItem_ParentObject(const QGraphicsItem* self);
QGraphicsWidget* QGraphicsItem_ParentWidget(const QGraphicsItem* self);
QGraphicsWidget* QGraphicsItem_TopLevelWidget(const QGraphicsItem* self);
QGraphicsWidget* QGraphicsItem_Window(const QGraphicsItem* self);
QGraphicsItem* QGraphicsItem_Panel(const QGraphicsItem* self);
void QGraphicsItem_SetParentItem(QGraphicsItem* self, QGraphicsItem* parent);
struct miqt_array QGraphicsItem_ChildItems(const QGraphicsItem* self);
bool QGraphicsItem_IsWidget(const QGraphicsItem* self);
bool QGraphicsItem_IsWindow(const QGraphicsItem* self);
bool QGraphicsItem_IsPanel(const QGraphicsItem* self);
QGraphicsObject* QGraphicsItem_ToGraphicsObject(QGraphicsItem* self);
QGraphicsObject* QGraphicsItem_ToGraphicsObject2(const QGraphicsItem* self);
QGraphicsItemGroup* QGraphicsItem_Group(const QGraphicsItem* self);
void QGraphicsItem_SetGroup(QGraphicsItem* self, QGraphicsItemGroup* group);
int QGraphicsItem_Flags(const QGraphicsItem* self);
void QGraphicsItem_SetFlag(QGraphicsItem* self, int flag);
void QGraphicsItem_SetFlags(QGraphicsItem* self, int flags);
int QGraphicsItem_CacheMode(const QGraphicsItem* self);
void QGraphicsItem_SetCacheMode(QGraphicsItem* self, int mode);
int QGraphicsItem_PanelModality(const QGraphicsItem* self);
void QGraphicsItem_SetPanelModality(QGraphicsItem* self, int panelModality);
bool QGraphicsItem_IsBlockedByModalPanel(const QGraphicsItem* self);
struct miqt_string QGraphicsItem_ToolTip(const QGraphicsItem* self);
void QGraphicsItem_SetToolTip(QGraphicsItem* self, struct miqt_string toolTip);
QCursor* QGraphicsItem_Cursor(const QGraphicsItem* self);
void QGraphicsItem_SetCursor(QGraphicsItem* self, QCursor* cursor);
bool QGraphicsItem_HasCursor(const QGraphicsItem* self);
void QGraphicsItem_UnsetCursor(QGraphicsItem* self);
bool QGraphicsItem_IsVisible(const QGraphicsItem* self);
bool QGraphicsItem_IsVisibleTo(const QGraphicsItem* self, QGraphicsItem* parent);
void QGraphicsItem_SetVisible(QGraphicsItem* self, bool visible);
void QGraphicsItem_Hide(QGraphicsItem* self);
void QGraphicsItem_Show(QGraphicsItem* self);
bool QGraphicsItem_IsEnabled(const QGraphicsItem* self);
void QGraphicsItem_SetEnabled(QGraphicsItem* self, bool enabled);
bool QGraphicsItem_IsSelected(const QGraphicsItem* self);
void QGraphicsItem_SetSelected(QGraphicsItem* self, bool selected);
bool QGraphicsItem_AcceptDrops(const QGraphicsItem* self);
void QGraphicsItem_SetAcceptDrops(QGraphicsItem* self, bool on);
double QGraphicsItem_Opacity(const QGraphicsItem* self);
double QGraphicsItem_EffectiveOpacity(const QGraphicsItem* self);
void QGraphicsItem_SetOpacity(QGraphicsItem* self, double opacity);
QGraphicsEffect* QGraphicsItem_GraphicsEffect(const QGraphicsItem* self);
void QGraphicsItem_SetGraphicsEffect(QGraphicsItem* self, QGraphicsEffect* effect);
int QGraphicsItem_AcceptedMouseButtons(const QGraphicsItem* self);
void QGraphicsItem_SetAcceptedMouseButtons(QGraphicsItem* self, int buttons);
bool QGraphicsItem_AcceptHoverEvents(const QGraphicsItem* self);
void QGraphicsItem_SetAcceptHoverEvents(QGraphicsItem* self, bool enabled);
bool QGraphicsItem_AcceptTouchEvents(const QGraphicsItem* self);
void QGraphicsItem_SetAcceptTouchEvents(QGraphicsItem* self, bool enabled);
bool QGraphicsItem_FiltersChildEvents(const QGraphicsItem* self);
void QGraphicsItem_SetFiltersChildEvents(QGraphicsItem* self, bool enabled);
bool QGraphicsItem_HandlesChildEvents(const QGraphicsItem* self);
void QGraphicsItem_SetHandlesChildEvents(QGraphicsItem* self, bool enabled);
bool QGraphicsItem_IsActive(const QGraphicsItem* self);
void QGraphicsItem_SetActive(QGraphicsItem* self, bool active);
bool QGraphicsItem_HasFocus(const QGraphicsItem* self);
void QGraphicsItem_SetFocus(QGraphicsItem* self);
void QGraphicsItem_ClearFocus(QGraphicsItem* self);
QGraphicsItem* QGraphicsItem_FocusProxy(const QGraphicsItem* self);
void QGraphicsItem_SetFocusProxy(QGraphicsItem* self, QGraphicsItem* item);
QGraphicsItem* QGraphicsItem_FocusItem(const QGraphicsItem* self);
QGraphicsItem* QGraphicsItem_FocusScopeItem(const QGraphicsItem* self);
void QGraphicsItem_GrabMouse(QGraphicsItem* self);
void QGraphicsItem_UngrabMouse(QGraphicsItem* self);
void QGraphicsItem_GrabKeyboard(QGraphicsItem* self);
void QGraphicsItem_UngrabKeyboard(QGraphicsItem* self);
QPointF* QGraphicsItem_Pos(const QGraphicsItem* self);
double QGraphicsItem_X(const QGraphicsItem* self);
void QGraphicsItem_SetX(QGraphicsItem* self, double x);
double QGraphicsItem_Y(const QGraphicsItem* self);
void QGraphicsItem_SetY(QGraphicsItem* self, double y);
QPointF* QGraphicsItem_ScenePos(const QGraphicsItem* self);
void QGraphicsItem_SetPos(QGraphicsItem* self, QPointF* pos);
void QGraphicsItem_SetPos2(QGraphicsItem* self, double x, double y);
void QGraphicsItem_MoveBy(QGraphicsItem* self, double dx, double dy);
void QGraphicsItem_EnsureVisible(QGraphicsItem* self);
void QGraphicsItem_EnsureVisible2(QGraphicsItem* self, double x, double y, double w, double h);
QTransform* QGraphicsItem_Transform(const QGraphicsItem* self);
QTransform* QGraphicsItem_SceneTransform(const QGraphicsItem* self);
QTransform* QGraphicsItem_DeviceTransform(const QGraphicsItem* self, QTransform* viewportTransform);
QTransform* QGraphicsItem_ItemTransform(const QGraphicsItem* self, QGraphicsItem* other);
void QGraphicsItem_SetTransform(QGraphicsItem* self, QTransform* matrix);
void QGraphicsItem_ResetTransform(QGraphicsItem* self);
void QGraphicsItem_SetRotation(QGraphicsItem* self, double angle);
double QGraphicsItem_Rotation(const QGraphicsItem* self);
void QGraphicsItem_SetScale(QGraphicsItem* self, double scale);
double QGraphicsItem_Scale(const QGraphicsItem* self);
struct miqt_array QGraphicsItem_Transformations(const QGraphicsItem* self);
void QGraphicsItem_SetTransformations(QGraphicsItem* self, struct miqt_array /* of QGraphicsTransform* */ transformations);
QPointF* QGraphicsItem_TransformOriginPoint(const QGraphicsItem* self);
void QGraphicsItem_SetTransformOriginPoint(QGraphicsItem* self, QPointF* origin);
void QGraphicsItem_SetTransformOriginPoint2(QGraphicsItem* self, double ax, double ay);
void QGraphicsItem_Advance(QGraphicsItem* self, int phase);
double QGraphicsItem_ZValue(const QGraphicsItem* self);
void QGraphicsItem_SetZValue(QGraphicsItem* self, double z);
void QGraphicsItem_StackBefore(QGraphicsItem* self, QGraphicsItem* sibling);
QRectF* QGraphicsItem_BoundingRect(const QGraphicsItem* self);
QRectF* QGraphicsItem_ChildrenBoundingRect(const QGraphicsItem* self);
QRectF* QGraphicsItem_SceneBoundingRect(const QGraphicsItem* self);
QPainterPath* QGraphicsItem_Shape(const QGraphicsItem* self);
bool QGraphicsItem_IsClipped(const QGraphicsItem* self);
QPainterPath* QGraphicsItem_ClipPath(const QGraphicsItem* self);
bool QGraphicsItem_Contains(const QGraphicsItem* self, QPointF* point);
bool QGraphicsItem_CollidesWithItem(const QGraphicsItem* self, QGraphicsItem* other);
bool QGraphicsItem_CollidesWithPath(const QGraphicsItem* self, QPainterPath* path);
struct miqt_array QGraphicsItem_CollidingItems(const QGraphicsItem* self);
bool QGraphicsItem_IsObscured(const QGraphicsItem* self);
bool QGraphicsItem_IsObscured2(const QGraphicsItem* self, double x, double y, double w, double h);
bool QGraphicsItem_IsObscuredBy(const QGraphicsItem* self, QGraphicsItem* item);
QPainterPath* QGraphicsItem_OpaqueArea(const QGraphicsItem* self);
QRegion* QGraphicsItem_BoundingRegion(const QGraphicsItem* self, QTransform* itemToDeviceTransform);
double QGraphicsItem_BoundingRegionGranularity(const QGraphicsItem* self);
void QGraphicsItem_SetBoundingRegionGranularity(QGraphicsItem* self, double granularity);
void QGraphicsItem_Paint(QGraphicsItem* self, QPainter* painter, QStyleOptionGraphicsItem* option);
void QGraphicsItem_Update(QGraphicsItem* self);
void QGraphicsItem_Update2(QGraphicsItem* self, double x, double y, double width, double height);
void QGraphicsItem_Scroll(QGraphicsItem* self, double dx, double dy);
QPointF* QGraphicsItem_MapToItem(const QGraphicsItem* self, QGraphicsItem* item, QPointF* point);
QPointF* QGraphicsItem_MapToParent(const QGraphicsItem* self, QPointF* point);
QPointF* QGraphicsItem_MapToScene(const QGraphicsItem* self, QPointF* point);
QRectF* QGraphicsItem_MapRectToItem(const QGraphicsItem* self, QGraphicsItem* item, QRectF* rect);
QRectF* QGraphicsItem_MapRectToParent(const QGraphicsItem* self, QRectF* rect);
QRectF* QGraphicsItem_MapRectToScene(const QGraphicsItem* self, QRectF* rect);
QPainterPath* QGraphicsItem_MapToItem4(const QGraphicsItem* self, QGraphicsItem* item, QPainterPath* path);
QPainterPath* QGraphicsItem_MapToParentWithPath(const QGraphicsItem* self, QPainterPath* path);
QPainterPath* QGraphicsItem_MapToSceneWithPath(const QGraphicsItem* self, QPainterPath* path);
QPointF* QGraphicsItem_MapFromItem(const QGraphicsItem* self, QGraphicsItem* item, QPointF* point);
QPointF* QGraphicsItem_MapFromParent(const QGraphicsItem* self, QPointF* point);
QPointF* QGraphicsItem_MapFromScene(const QGraphicsItem* self, QPointF* point);
QRectF* QGraphicsItem_MapRectFromItem(const QGraphicsItem* self, QGraphicsItem* item, QRectF* rect);
QRectF* QGraphicsItem_MapRectFromParent(const QGraphicsItem* self, QRectF* rect);
QRectF* QGraphicsItem_MapRectFromScene(const QGraphicsItem* self, QRectF* rect);
QPainterPath* QGraphicsItem_MapFromItem4(const QGraphicsItem* self, QGraphicsItem* item, QPainterPath* path);
QPainterPath* QGraphicsItem_MapFromParentWithPath(const QGraphicsItem* self, QPainterPath* path);
QPainterPath* QGraphicsItem_MapFromSceneWithPath(const QGraphicsItem* self, QPainterPath* path);
QPointF* QGraphicsItem_MapToItem5(const QGraphicsItem* self, QGraphicsItem* item, double x, double y);
QPointF* QGraphicsItem_MapToParent2(const QGraphicsItem* self, double x, double y);
QPointF* QGraphicsItem_MapToScene2(const QGraphicsItem* self, double x, double y);
QRectF* QGraphicsItem_MapRectToItem2(const QGraphicsItem* self, QGraphicsItem* item, double x, double y, double w, double h);
QRectF* QGraphicsItem_MapRectToParent2(const QGraphicsItem* self, double x, double y, double w, double h);
QRectF* QGraphicsItem_MapRectToScene2(const QGraphicsItem* self, double x, double y, double w, double h);
QPointF* QGraphicsItem_MapFromItem5(const QGraphicsItem* self, QGraphicsItem* item, double x, double y);
QPointF* QGraphicsItem_MapFromParent2(const QGraphicsItem* self, double x, double y);
QPointF* QGraphicsItem_MapFromScene2(const QGraphicsItem* self, double x, double y);
QRectF* QGraphicsItem_MapRectFromItem2(const QGraphicsItem* self, QGraphicsItem* item, double x, double y, double w, double h);
QRectF* QGraphicsItem_MapRectFromParent2(const QGraphicsItem* self, double x, double y, double w, double h);
QRectF* QGraphicsItem_MapRectFromScene2(const QGraphicsItem* self, double x, double y, double w, double h);
bool QGraphicsItem_IsAncestorOf(const QGraphicsItem* self, QGraphicsItem* child);
QGraphicsItem* QGraphicsItem_CommonAncestorItem(const QGraphicsItem* self, QGraphicsItem* other);
bool QGraphicsItem_IsUnderMouse(const QGraphicsItem* self);
QVariant* QGraphicsItem_Data(const QGraphicsItem* self, int key);
void QGraphicsItem_SetData(QGraphicsItem* self, int key, QVariant* value);
int QGraphicsItem_InputMethodHints(const QGraphicsItem* self);
void QGraphicsItem_SetInputMethodHints(QGraphicsItem* self, int hints);
int QGraphicsItem_Type(const QGraphicsItem* self);
void QGraphicsItem_InstallSceneEventFilter(QGraphicsItem* self, QGraphicsItem* filterItem);
void QGraphicsItem_RemoveSceneEventFilter(QGraphicsItem* self, QGraphicsItem* filterItem);
void QGraphicsItem_SetFlag2(QGraphicsItem* self, int flag, bool enabled);
void QGraphicsItem_SetCacheMode2(QGraphicsItem* self, int mode, QSize* cacheSize);
void QGraphicsItem_SetFocus1(QGraphicsItem* self, int focusReason);
void QGraphicsItem_EnsureVisible1(QGraphicsItem* self, QRectF* rect);
void QGraphicsItem_EnsureVisible22(QGraphicsItem* self, QRectF* rect, int xmargin);
void QGraphicsItem_EnsureVisible3(QGraphicsItem* self, QRectF* rect, int xmargin, int ymargin);
void QGraphicsItem_EnsureVisible5(QGraphicsItem* self, double x, double y, double w, double h, int xmargin);
void QGraphicsItem_EnsureVisible6(QGraphicsItem* self, double x, double y, double w, double h, int xmargin, int ymargin);
QTransform* QGraphicsItem_ItemTransform2(const QGraphicsItem* self, QGraphicsItem* other, bool* ok);
void QGraphicsItem_SetTransform2(QGraphicsItem* self, QTransform* matrix, bool combine);
bool QGraphicsItem_CollidesWithItem2(const QGraphicsItem* self, QGraphicsItem* other, int mode);
bool QGraphicsItem_CollidesWithPath2(const QGraphicsItem* self, QPainterPath* path, int mode);
struct miqt_array QGraphicsItem_CollidingItems1(const QGraphicsItem* self, int mode);
bool QGraphicsItem_IsObscured1(const QGraphicsItem* self, QRectF* rect);
void QGraphicsItem_Paint3(QGraphicsItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
void QGraphicsItem_Update1(QGraphicsItem* self, QRectF* rect);
void QGraphicsItem_Scroll3(QGraphicsItem* self, double dx, double dy, QRectF* rect);
void QGraphicsItem_Delete(QGraphicsItem* self);

QMetaObject* QGraphicsObject_MetaObject(const QGraphicsObject* self);
void* QGraphicsObject_Metacast(QGraphicsObject* self, const char* param1);
struct miqt_string QGraphicsObject_Tr(const char* s);
void QGraphicsObject_GrabGesture(QGraphicsObject* self, int typeVal);
void QGraphicsObject_UngrabGesture(QGraphicsObject* self, int typeVal);
void QGraphicsObject_ParentChanged(QGraphicsObject* self);
void QGraphicsObject_connect_ParentChanged(QGraphicsObject* self, intptr_t slot);
void QGraphicsObject_OpacityChanged(QGraphicsObject* self);
void QGraphicsObject_connect_OpacityChanged(QGraphicsObject* self, intptr_t slot);
void QGraphicsObject_VisibleChanged(QGraphicsObject* self);
void QGraphicsObject_connect_VisibleChanged(QGraphicsObject* self, intptr_t slot);
void QGraphicsObject_EnabledChanged(QGraphicsObject* self);
void QGraphicsObject_connect_EnabledChanged(QGraphicsObject* self, intptr_t slot);
void QGraphicsObject_XChanged(QGraphicsObject* self);
void QGraphicsObject_connect_XChanged(QGraphicsObject* self, intptr_t slot);
void QGraphicsObject_YChanged(QGraphicsObject* self);
void QGraphicsObject_connect_YChanged(QGraphicsObject* self, intptr_t slot);
void QGraphicsObject_ZChanged(QGraphicsObject* self);
void QGraphicsObject_connect_ZChanged(QGraphicsObject* self, intptr_t slot);
void QGraphicsObject_RotationChanged(QGraphicsObject* self);
void QGraphicsObject_connect_RotationChanged(QGraphicsObject* self, intptr_t slot);
void QGraphicsObject_ScaleChanged(QGraphicsObject* self);
void QGraphicsObject_connect_ScaleChanged(QGraphicsObject* self, intptr_t slot);
void QGraphicsObject_ChildrenChanged(QGraphicsObject* self);
void QGraphicsObject_connect_ChildrenChanged(QGraphicsObject* self, intptr_t slot);
void QGraphicsObject_WidthChanged(QGraphicsObject* self);
void QGraphicsObject_connect_WidthChanged(QGraphicsObject* self, intptr_t slot);
void QGraphicsObject_HeightChanged(QGraphicsObject* self);
void QGraphicsObject_connect_HeightChanged(QGraphicsObject* self, intptr_t slot);
struct miqt_string QGraphicsObject_Tr2(const char* s, const char* c);
struct miqt_string QGraphicsObject_Tr3(const char* s, const char* c, int n);
void QGraphicsObject_GrabGesture2(QGraphicsObject* self, int typeVal, int flags);
void QGraphicsObject_Delete(QGraphicsObject* self);

QPen* QAbstractGraphicsShapeItem_Pen(const QAbstractGraphicsShapeItem* self);
void QAbstractGraphicsShapeItem_SetPen(QAbstractGraphicsShapeItem* self, QPen* pen);
QBrush* QAbstractGraphicsShapeItem_Brush(const QAbstractGraphicsShapeItem* self);
void QAbstractGraphicsShapeItem_SetBrush(QAbstractGraphicsShapeItem* self, QBrush* brush);
bool QAbstractGraphicsShapeItem_IsObscuredBy(const QAbstractGraphicsShapeItem* self, QGraphicsItem* item);
QPainterPath* QAbstractGraphicsShapeItem_OpaqueArea(const QAbstractGraphicsShapeItem* self);
void QAbstractGraphicsShapeItem_Delete(QAbstractGraphicsShapeItem* self);

QGraphicsPathItem* QGraphicsPathItem_new();
QGraphicsPathItem* QGraphicsPathItem_new2(QPainterPath* path);
QGraphicsPathItem* QGraphicsPathItem_new3(QGraphicsItem* parent);
QGraphicsPathItem* QGraphicsPathItem_new4(QPainterPath* path, QGraphicsItem* parent);
QPainterPath* QGraphicsPathItem_Path(const QGraphicsPathItem* self);
void QGraphicsPathItem_SetPath(QGraphicsPathItem* self, QPainterPath* path);
QRectF* QGraphicsPathItem_BoundingRect(const QGraphicsPathItem* self);
QPainterPath* QGraphicsPathItem_Shape(const QGraphicsPathItem* self);
bool QGraphicsPathItem_Contains(const QGraphicsPathItem* self, QPointF* point);
void QGraphicsPathItem_Paint(QGraphicsPathItem* self, QPainter* painter, QStyleOptionGraphicsItem* option);
bool QGraphicsPathItem_IsObscuredBy(const QGraphicsPathItem* self, QGraphicsItem* item);
QPainterPath* QGraphicsPathItem_OpaqueArea(const QGraphicsPathItem* self);
int QGraphicsPathItem_Type(const QGraphicsPathItem* self);
void QGraphicsPathItem_Paint3(QGraphicsPathItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
void QGraphicsPathItem_Delete(QGraphicsPathItem* self);

QGraphicsRectItem* QGraphicsRectItem_new();
QGraphicsRectItem* QGraphicsRectItem_new2(QRectF* rect);
QGraphicsRectItem* QGraphicsRectItem_new3(double x, double y, double w, double h);
QGraphicsRectItem* QGraphicsRectItem_new4(QGraphicsItem* parent);
QGraphicsRectItem* QGraphicsRectItem_new5(QRectF* rect, QGraphicsItem* parent);
QGraphicsRectItem* QGraphicsRectItem_new6(double x, double y, double w, double h, QGraphicsItem* parent);
QRectF* QGraphicsRectItem_Rect(const QGraphicsRectItem* self);
void QGraphicsRectItem_SetRect(QGraphicsRectItem* self, QRectF* rect);
void QGraphicsRectItem_SetRect2(QGraphicsRectItem* self, double x, double y, double w, double h);
QRectF* QGraphicsRectItem_BoundingRect(const QGraphicsRectItem* self);
QPainterPath* QGraphicsRectItem_Shape(const QGraphicsRectItem* self);
bool QGraphicsRectItem_Contains(const QGraphicsRectItem* self, QPointF* point);
void QGraphicsRectItem_Paint(QGraphicsRectItem* self, QPainter* painter, QStyleOptionGraphicsItem* option);
bool QGraphicsRectItem_IsObscuredBy(const QGraphicsRectItem* self, QGraphicsItem* item);
QPainterPath* QGraphicsRectItem_OpaqueArea(const QGraphicsRectItem* self);
int QGraphicsRectItem_Type(const QGraphicsRectItem* self);
void QGraphicsRectItem_Paint3(QGraphicsRectItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
void QGraphicsRectItem_Delete(QGraphicsRectItem* self);

QGraphicsEllipseItem* QGraphicsEllipseItem_new();
QGraphicsEllipseItem* QGraphicsEllipseItem_new2(QRectF* rect);
QGraphicsEllipseItem* QGraphicsEllipseItem_new3(double x, double y, double w, double h);
QGraphicsEllipseItem* QGraphicsEllipseItem_new4(QGraphicsItem* parent);
QGraphicsEllipseItem* QGraphicsEllipseItem_new5(QRectF* rect, QGraphicsItem* parent);
QGraphicsEllipseItem* QGraphicsEllipseItem_new6(double x, double y, double w, double h, QGraphicsItem* parent);
QRectF* QGraphicsEllipseItem_Rect(const QGraphicsEllipseItem* self);
void QGraphicsEllipseItem_SetRect(QGraphicsEllipseItem* self, QRectF* rect);
void QGraphicsEllipseItem_SetRect2(QGraphicsEllipseItem* self, double x, double y, double w, double h);
int QGraphicsEllipseItem_StartAngle(const QGraphicsEllipseItem* self);
void QGraphicsEllipseItem_SetStartAngle(QGraphicsEllipseItem* self, int angle);
int QGraphicsEllipseItem_SpanAngle(const QGraphicsEllipseItem* self);
void QGraphicsEllipseItem_SetSpanAngle(QGraphicsEllipseItem* self, int angle);
QRectF* QGraphicsEllipseItem_BoundingRect(const QGraphicsEllipseItem* self);
QPainterPath* QGraphicsEllipseItem_Shape(const QGraphicsEllipseItem* self);
bool QGraphicsEllipseItem_Contains(const QGraphicsEllipseItem* self, QPointF* point);
void QGraphicsEllipseItem_Paint(QGraphicsEllipseItem* self, QPainter* painter, QStyleOptionGraphicsItem* option);
bool QGraphicsEllipseItem_IsObscuredBy(const QGraphicsEllipseItem* self, QGraphicsItem* item);
QPainterPath* QGraphicsEllipseItem_OpaqueArea(const QGraphicsEllipseItem* self);
int QGraphicsEllipseItem_Type(const QGraphicsEllipseItem* self);
void QGraphicsEllipseItem_Paint3(QGraphicsEllipseItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
void QGraphicsEllipseItem_Delete(QGraphicsEllipseItem* self);

QGraphicsPolygonItem* QGraphicsPolygonItem_new();
QGraphicsPolygonItem* QGraphicsPolygonItem_new2(QGraphicsItem* parent);
int QGraphicsPolygonItem_FillRule(const QGraphicsPolygonItem* self);
void QGraphicsPolygonItem_SetFillRule(QGraphicsPolygonItem* self, int rule);
QRectF* QGraphicsPolygonItem_BoundingRect(const QGraphicsPolygonItem* self);
QPainterPath* QGraphicsPolygonItem_Shape(const QGraphicsPolygonItem* self);
bool QGraphicsPolygonItem_Contains(const QGraphicsPolygonItem* self, QPointF* point);
void QGraphicsPolygonItem_Paint(QGraphicsPolygonItem* self, QPainter* painter, QStyleOptionGraphicsItem* option);
bool QGraphicsPolygonItem_IsObscuredBy(const QGraphicsPolygonItem* self, QGraphicsItem* item);
QPainterPath* QGraphicsPolygonItem_OpaqueArea(const QGraphicsPolygonItem* self);
int QGraphicsPolygonItem_Type(const QGraphicsPolygonItem* self);
void QGraphicsPolygonItem_Paint3(QGraphicsPolygonItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
void QGraphicsPolygonItem_Delete(QGraphicsPolygonItem* self);

QGraphicsLineItem* QGraphicsLineItem_new();
QGraphicsLineItem* QGraphicsLineItem_new2(QLineF* line);
QGraphicsLineItem* QGraphicsLineItem_new3(double x1, double y1, double x2, double y2);
QGraphicsLineItem* QGraphicsLineItem_new4(QGraphicsItem* parent);
QGraphicsLineItem* QGraphicsLineItem_new5(QLineF* line, QGraphicsItem* parent);
QGraphicsLineItem* QGraphicsLineItem_new6(double x1, double y1, double x2, double y2, QGraphicsItem* parent);
QPen* QGraphicsLineItem_Pen(const QGraphicsLineItem* self);
void QGraphicsLineItem_SetPen(QGraphicsLineItem* self, QPen* pen);
QLineF* QGraphicsLineItem_Line(const QGraphicsLineItem* self);
void QGraphicsLineItem_SetLine(QGraphicsLineItem* self, QLineF* line);
void QGraphicsLineItem_SetLine2(QGraphicsLineItem* self, double x1, double y1, double x2, double y2);
QRectF* QGraphicsLineItem_BoundingRect(const QGraphicsLineItem* self);
QPainterPath* QGraphicsLineItem_Shape(const QGraphicsLineItem* self);
bool QGraphicsLineItem_Contains(const QGraphicsLineItem* self, QPointF* point);
void QGraphicsLineItem_Paint(QGraphicsLineItem* self, QPainter* painter, QStyleOptionGraphicsItem* option);
bool QGraphicsLineItem_IsObscuredBy(const QGraphicsLineItem* self, QGraphicsItem* item);
QPainterPath* QGraphicsLineItem_OpaqueArea(const QGraphicsLineItem* self);
int QGraphicsLineItem_Type(const QGraphicsLineItem* self);
void QGraphicsLineItem_Paint3(QGraphicsLineItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
void QGraphicsLineItem_Delete(QGraphicsLineItem* self);

QGraphicsPixmapItem* QGraphicsPixmapItem_new();
QGraphicsPixmapItem* QGraphicsPixmapItem_new2(QPixmap* pixmap);
QGraphicsPixmapItem* QGraphicsPixmapItem_new3(QGraphicsItem* parent);
QGraphicsPixmapItem* QGraphicsPixmapItem_new4(QPixmap* pixmap, QGraphicsItem* parent);
QPixmap* QGraphicsPixmapItem_Pixmap(const QGraphicsPixmapItem* self);
void QGraphicsPixmapItem_SetPixmap(QGraphicsPixmapItem* self, QPixmap* pixmap);
int QGraphicsPixmapItem_TransformationMode(const QGraphicsPixmapItem* self);
void QGraphicsPixmapItem_SetTransformationMode(QGraphicsPixmapItem* self, int mode);
QPointF* QGraphicsPixmapItem_Offset(const QGraphicsPixmapItem* self);
void QGraphicsPixmapItem_SetOffset(QGraphicsPixmapItem* self, QPointF* offset);
void QGraphicsPixmapItem_SetOffset2(QGraphicsPixmapItem* self, double x, double y);
QRectF* QGraphicsPixmapItem_BoundingRect(const QGraphicsPixmapItem* self);
QPainterPath* QGraphicsPixmapItem_Shape(const QGraphicsPixmapItem* self);
bool QGraphicsPixmapItem_Contains(const QGraphicsPixmapItem* self, QPointF* point);
void QGraphicsPixmapItem_Paint(QGraphicsPixmapItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
bool QGraphicsPixmapItem_IsObscuredBy(const QGraphicsPixmapItem* self, QGraphicsItem* item);
QPainterPath* QGraphicsPixmapItem_OpaqueArea(const QGraphicsPixmapItem* self);
int QGraphicsPixmapItem_Type(const QGraphicsPixmapItem* self);
int QGraphicsPixmapItem_ShapeMode(const QGraphicsPixmapItem* self);
void QGraphicsPixmapItem_SetShapeMode(QGraphicsPixmapItem* self, int mode);
void QGraphicsPixmapItem_Delete(QGraphicsPixmapItem* self);

QGraphicsTextItem* QGraphicsTextItem_new();
QGraphicsTextItem* QGraphicsTextItem_new2(struct miqt_string text);
QGraphicsTextItem* QGraphicsTextItem_new3(QGraphicsItem* parent);
QGraphicsTextItem* QGraphicsTextItem_new4(struct miqt_string text, QGraphicsItem* parent);
QMetaObject* QGraphicsTextItem_MetaObject(const QGraphicsTextItem* self);
void* QGraphicsTextItem_Metacast(QGraphicsTextItem* self, const char* param1);
struct miqt_string QGraphicsTextItem_Tr(const char* s);
struct miqt_string QGraphicsTextItem_ToHtml(const QGraphicsTextItem* self);
void QGraphicsTextItem_SetHtml(QGraphicsTextItem* self, struct miqt_string html);
struct miqt_string QGraphicsTextItem_ToPlainText(const QGraphicsTextItem* self);
void QGraphicsTextItem_SetPlainText(QGraphicsTextItem* self, struct miqt_string text);
QFont* QGraphicsTextItem_Font(const QGraphicsTextItem* self);
void QGraphicsTextItem_SetFont(QGraphicsTextItem* self, QFont* font);
void QGraphicsTextItem_SetDefaultTextColor(QGraphicsTextItem* self, QColor* c);
QColor* QGraphicsTextItem_DefaultTextColor(const QGraphicsTextItem* self);
QRectF* QGraphicsTextItem_BoundingRect(const QGraphicsTextItem* self);
QPainterPath* QGraphicsTextItem_Shape(const QGraphicsTextItem* self);
bool QGraphicsTextItem_Contains(const QGraphicsTextItem* self, QPointF* point);
void QGraphicsTextItem_Paint(QGraphicsTextItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
bool QGraphicsTextItem_IsObscuredBy(const QGraphicsTextItem* self, QGraphicsItem* item);
QPainterPath* QGraphicsTextItem_OpaqueArea(const QGraphicsTextItem* self);
int QGraphicsTextItem_Type(const QGraphicsTextItem* self);
void QGraphicsTextItem_SetTextWidth(QGraphicsTextItem* self, double width);
double QGraphicsTextItem_TextWidth(const QGraphicsTextItem* self);
void QGraphicsTextItem_AdjustSize(QGraphicsTextItem* self);
void QGraphicsTextItem_SetDocument(QGraphicsTextItem* self, QTextDocument* document);
QTextDocument* QGraphicsTextItem_Document(const QGraphicsTextItem* self);
void QGraphicsTextItem_SetTextInteractionFlags(QGraphicsTextItem* self, int flags);
int QGraphicsTextItem_TextInteractionFlags(const QGraphicsTextItem* self);
void QGraphicsTextItem_SetTabChangesFocus(QGraphicsTextItem* self, bool b);
bool QGraphicsTextItem_TabChangesFocus(const QGraphicsTextItem* self);
void QGraphicsTextItem_SetOpenExternalLinks(QGraphicsTextItem* self, bool open);
bool QGraphicsTextItem_OpenExternalLinks(const QGraphicsTextItem* self);
void QGraphicsTextItem_SetTextCursor(QGraphicsTextItem* self, QTextCursor* cursor);
QTextCursor* QGraphicsTextItem_TextCursor(const QGraphicsTextItem* self);
void QGraphicsTextItem_LinkActivated(QGraphicsTextItem* self, struct miqt_string param1);
void QGraphicsTextItem_connect_LinkActivated(QGraphicsTextItem* self, intptr_t slot);
void QGraphicsTextItem_LinkHovered(QGraphicsTextItem* self, struct miqt_string param1);
void QGraphicsTextItem_connect_LinkHovered(QGraphicsTextItem* self, intptr_t slot);
struct miqt_string QGraphicsTextItem_Tr2(const char* s, const char* c);
struct miqt_string QGraphicsTextItem_Tr3(const char* s, const char* c, int n);
void QGraphicsTextItem_Delete(QGraphicsTextItem* self);

QGraphicsSimpleTextItem* QGraphicsSimpleTextItem_new();
QGraphicsSimpleTextItem* QGraphicsSimpleTextItem_new2(struct miqt_string text);
QGraphicsSimpleTextItem* QGraphicsSimpleTextItem_new3(QGraphicsItem* parent);
QGraphicsSimpleTextItem* QGraphicsSimpleTextItem_new4(struct miqt_string text, QGraphicsItem* parent);
void QGraphicsSimpleTextItem_SetText(QGraphicsSimpleTextItem* self, struct miqt_string text);
struct miqt_string QGraphicsSimpleTextItem_Text(const QGraphicsSimpleTextItem* self);
void QGraphicsSimpleTextItem_SetFont(QGraphicsSimpleTextItem* self, QFont* font);
QFont* QGraphicsSimpleTextItem_Font(const QGraphicsSimpleTextItem* self);
QRectF* QGraphicsSimpleTextItem_BoundingRect(const QGraphicsSimpleTextItem* self);
QPainterPath* QGraphicsSimpleTextItem_Shape(const QGraphicsSimpleTextItem* self);
bool QGraphicsSimpleTextItem_Contains(const QGraphicsSimpleTextItem* self, QPointF* point);
void QGraphicsSimpleTextItem_Paint(QGraphicsSimpleTextItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
bool QGraphicsSimpleTextItem_IsObscuredBy(const QGraphicsSimpleTextItem* self, QGraphicsItem* item);
QPainterPath* QGraphicsSimpleTextItem_OpaqueArea(const QGraphicsSimpleTextItem* self);
int QGraphicsSimpleTextItem_Type(const QGraphicsSimpleTextItem* self);
void QGraphicsSimpleTextItem_Delete(QGraphicsSimpleTextItem* self);

QGraphicsItemGroup* QGraphicsItemGroup_new();
QGraphicsItemGroup* QGraphicsItemGroup_new2(QGraphicsItem* parent);
void QGraphicsItemGroup_AddToGroup(QGraphicsItemGroup* self, QGraphicsItem* item);
void QGraphicsItemGroup_RemoveFromGroup(QGraphicsItemGroup* self, QGraphicsItem* item);
QRectF* QGraphicsItemGroup_BoundingRect(const QGraphicsItemGroup* self);
void QGraphicsItemGroup_Paint(QGraphicsItemGroup* self, QPainter* painter, QStyleOptionGraphicsItem* option);
bool QGraphicsItemGroup_IsObscuredBy(const QGraphicsItemGroup* self, QGraphicsItem* item);
QPainterPath* QGraphicsItemGroup_OpaqueArea(const QGraphicsItemGroup* self);
int QGraphicsItemGroup_Type(const QGraphicsItemGroup* self);
void QGraphicsItemGroup_Paint3(QGraphicsItemGroup* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
void QGraphicsItemGroup_Delete(QGraphicsItemGroup* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif