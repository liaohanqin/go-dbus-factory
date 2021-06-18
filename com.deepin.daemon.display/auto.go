// Code generated by "./generator ./com.deepin.daemon.display"; DO NOT EDIT.

package display

import (
	"errors"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type Display interface {
	display // interface com.deepin.daemon.Display
	proxy.Object
}

type objectDisplay struct {
	interfaceDisplay // interface com.deepin.daemon.Display
	proxy.ImplObject
}

func NewDisplay(conn *dbus.Conn) Display {
	obj := new(objectDisplay)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.Display", "/com/deepin/daemon/Display")
	return obj
}

type display interface {
	GoApplyChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ApplyChanges(flags dbus.Flags) error
	GoAssociateTouch(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call
	AssociateTouch(flags dbus.Flags, arg0 string, arg1 string) error
	GoChangeBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call
	ChangeBrightness(flags dbus.Flags, arg0 bool) error
	GoDeleteCustomMode(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	DeleteCustomMode(flags dbus.Flags, arg0 string) error
	GoGetBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetBrightness(flags dbus.Flags) (map[string]float64, error)
	GoListOutputNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ListOutputNames(flags dbus.Flags) ([]string, error)
	GoListOutputsCommonModes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ListOutputsCommonModes(flags dbus.Flags) ([]ModeInfo, error)
	GoModifyConfigName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call
	ModifyConfigName(flags dbus.Flags, arg0 string, arg1 string) error
	GoRefreshBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RefreshBrightness(flags dbus.Flags) error
	GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Reset(flags dbus.Flags) error
	GoResetChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ResetChanges(flags dbus.Flags) error
	GoSave(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Save(flags dbus.Flags) error
	GoSetAndSaveBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 float64) *dbus.Call
	SetAndSaveBrightness(flags dbus.Flags, arg0 string, arg1 float64) error
	GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 float64) *dbus.Call
	SetBrightness(flags dbus.Flags, arg0 string, arg1 float64) error
	GoSetPrimary(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	SetPrimary(flags dbus.Flags, arg0 string) error
	GoSwitchMode(flags dbus.Flags, ch chan *dbus.Call, arg0 uint8, arg1 string) *dbus.Call
	SwitchMode(flags dbus.Flags, arg0 uint8, arg1 string) error
	HasChanged() proxy.PropBool
	DisplayMode() proxy.PropByte
	ScreenWidth() proxy.PropUint16
	ScreenHeight() proxy.PropUint16
	Primary() proxy.PropString
	CurrentCustomId() proxy.PropString
	CustomIdList() proxy.PropStringArray
	PrimaryRect() PropDisplayPrimaryRect
	Monitors() proxy.PropObjectPathArray
	Brightness() PropDisplayBrightness
	TouchMap() PropDisplayTouchMap
	Touchscreens() PropTouchscreens
	MaxBacklightBrightness() proxy.PropUint32
	ColorTemperatureMode() proxy.PropInt32
	ColorTemperatureManual() proxy.PropInt32
}

type interfaceDisplay struct{}

func (v *interfaceDisplay) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceDisplay) GetInterfaceName_() string {
	return "com.deepin.daemon.Display"
}

// method ApplyChanges

func (v *interfaceDisplay) GoApplyChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplyChanges", flags, ch)
}

func (v *interfaceDisplay) ApplyChanges(flags dbus.Flags) error {
	return (<-v.GoApplyChanges(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method AssociateTouch

func (v *interfaceDisplay) GoAssociateTouch(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AssociateTouch", flags, ch, arg0, arg1)
}

func (v *interfaceDisplay) AssociateTouch(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoAssociateTouch(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method ChangeBrightness

func (v *interfaceDisplay) GoChangeBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ChangeBrightness", flags, ch, arg0)
}

func (v *interfaceDisplay) ChangeBrightness(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoChangeBrightness(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method DeleteCustomMode

func (v *interfaceDisplay) GoDeleteCustomMode(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteCustomMode", flags, ch, arg0)
}

func (v *interfaceDisplay) DeleteCustomMode(flags dbus.Flags, arg0 string) error {
	return (<-v.GoDeleteCustomMode(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method GetBrightness

func (v *interfaceDisplay) GoGetBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetBrightness", flags, ch)
}

func (*interfaceDisplay) StoreGetBrightness(call *dbus.Call) (arg0 map[string]float64, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceDisplay) GetBrightness(flags dbus.Flags) (map[string]float64, error) {
	return v.StoreGetBrightness(
		<-v.GoGetBrightness(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListOutputNames

func (v *interfaceDisplay) GoListOutputNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListOutputNames", flags, ch)
}

func (*interfaceDisplay) StoreListOutputNames(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceDisplay) ListOutputNames(flags dbus.Flags) ([]string, error) {
	return v.StoreListOutputNames(
		<-v.GoListOutputNames(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListOutputsCommonModes

func (v *interfaceDisplay) GoListOutputsCommonModes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListOutputsCommonModes", flags, ch)
}

func (*interfaceDisplay) StoreListOutputsCommonModes(call *dbus.Call) (arg0 []ModeInfo, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceDisplay) ListOutputsCommonModes(flags dbus.Flags) ([]ModeInfo, error) {
	return v.StoreListOutputsCommonModes(
		<-v.GoListOutputsCommonModes(flags, make(chan *dbus.Call, 1)).Done)
}

// method ModifyConfigName

func (v *interfaceDisplay) GoModifyConfigName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ModifyConfigName", flags, ch, arg0, arg1)
}

func (v *interfaceDisplay) ModifyConfigName(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoModifyConfigName(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method RefreshBrightness

func (v *interfaceDisplay) GoRefreshBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RefreshBrightness", flags, ch)
}

func (v *interfaceDisplay) RefreshBrightness(flags dbus.Flags) error {
	return (<-v.GoRefreshBrightness(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Reset

func (v *interfaceDisplay) GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reset", flags, ch)
}

func (v *interfaceDisplay) Reset(flags dbus.Flags) error {
	return (<-v.GoReset(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ResetChanges

func (v *interfaceDisplay) GoResetChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ResetChanges", flags, ch)
}

func (v *interfaceDisplay) ResetChanges(flags dbus.Flags) error {
	return (<-v.GoResetChanges(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Save

func (v *interfaceDisplay) GoSave(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Save", flags, ch)
}

func (v *interfaceDisplay) Save(flags dbus.Flags) error {
	return (<-v.GoSave(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetAndSaveBrightness

func (v *interfaceDisplay) GoSetAndSaveBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAndSaveBrightness", flags, ch, arg0, arg1)
}

func (v *interfaceDisplay) SetAndSaveBrightness(flags dbus.Flags, arg0 string, arg1 float64) error {
	return (<-v.GoSetAndSaveBrightness(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetBrightness

func (v *interfaceDisplay) GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBrightness", flags, ch, arg0, arg1)
}

func (v *interfaceDisplay) SetBrightness(flags dbus.Flags, arg0 string, arg1 float64) error {
	return (<-v.GoSetBrightness(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetPrimary

func (v *interfaceDisplay) GoSetPrimary(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPrimary", flags, ch, arg0)
}

func (v *interfaceDisplay) SetPrimary(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetPrimary(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SwitchMode

func (v *interfaceDisplay) GoSwitchMode(flags dbus.Flags, ch chan *dbus.Call, arg0 uint8, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchMode", flags, ch, arg0, arg1)
}

func (v *interfaceDisplay) SwitchMode(flags dbus.Flags, arg0 uint8, arg1 string) error {
	return (<-v.GoSwitchMode(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// property HasChanged b

func (v *interfaceDisplay) HasChanged() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "HasChanged",
	}
}

// property DisplayMode y

func (v *interfaceDisplay) DisplayMode() proxy.PropByte {
	return &proxy.ImplPropByte{
		Impl: v,
		Name: "DisplayMode",
	}
}

// property ScreenWidth q

func (v *interfaceDisplay) ScreenWidth() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "ScreenWidth",
	}
}

// property ScreenHeight q

func (v *interfaceDisplay) ScreenHeight() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "ScreenHeight",
	}
}

// property Primary s

func (v *interfaceDisplay) Primary() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Primary",
	}
}

// property CurrentCustomId s

func (v *interfaceDisplay) CurrentCustomId() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "CurrentCustomId",
	}
}

// property CustomIdList as

func (v *interfaceDisplay) CustomIdList() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "CustomIdList",
	}
}

type PropDisplayPrimaryRect interface {
	Get(flags dbus.Flags) (value Rectangle, err error)
	Set(flags dbus.Flags, value Rectangle) error
	ConnectChanged(cb func(hasValue bool, value Rectangle)) error
}

type implPropDisplayPrimaryRect struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropDisplayPrimaryRect) Get(flags dbus.Flags) (value Rectangle, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropDisplayPrimaryRect) Set(flags dbus.Flags, value Rectangle) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropDisplayPrimaryRect) ConnectChanged(cb func(hasValue bool, value Rectangle)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v Rectangle
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, Rectangle{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property PrimaryRect (nnqq)

func (v *interfaceDisplay) PrimaryRect() PropDisplayPrimaryRect {
	return &implPropDisplayPrimaryRect{
		Impl: v,
		Name: "PrimaryRect",
	}
}

// property Monitors ao

func (v *interfaceDisplay) Monitors() proxy.PropObjectPathArray {
	return &proxy.ImplPropObjectPathArray{
		Impl: v,
		Name: "Monitors",
	}
}

type PropDisplayBrightness interface {
	Get(flags dbus.Flags) (value map[string]float64, err error)
	Set(flags dbus.Flags, value map[string]float64) error
	ConnectChanged(cb func(hasValue bool, value map[string]float64)) error
}

type implPropDisplayBrightness struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropDisplayBrightness) Get(flags dbus.Flags) (value map[string]float64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropDisplayBrightness) Set(flags dbus.Flags, value map[string]float64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropDisplayBrightness) ConnectChanged(cb func(hasValue bool, value map[string]float64)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]float64
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property Brightness a{sd}

func (v *interfaceDisplay) Brightness() PropDisplayBrightness {
	return &implPropDisplayBrightness{
		Impl: v,
		Name: "Brightness",
	}
}

type PropDisplayTouchMap interface {
	Get(flags dbus.Flags) (value map[string]string, err error)
	Set(flags dbus.Flags, value map[string]string) error
	ConnectChanged(cb func(hasValue bool, value map[string]string)) error
}

type implPropDisplayTouchMap struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropDisplayTouchMap) Get(flags dbus.Flags) (value map[string]string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropDisplayTouchMap) Set(flags dbus.Flags, value map[string]string) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropDisplayTouchMap) ConnectChanged(cb func(hasValue bool, value map[string]string)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]string
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property TouchMap a{ss}

func (v *interfaceDisplay) TouchMap() PropDisplayTouchMap {
	return &implPropDisplayTouchMap{
		Impl: v,
		Name: "TouchMap",
	}
}

type PropTouchscreens interface {
	Get(flags dbus.Flags) (value []Touchscreen, err error)
	Set(flags dbus.Flags, value []Touchscreen) error
	ConnectChanged(cb func(hasValue bool, value []Touchscreen)) error
}

type implPropTouchscreens struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropTouchscreens) Get(flags dbus.Flags) (value []Touchscreen, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropTouchscreens) Set(flags dbus.Flags, value []Touchscreen) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropTouchscreens) ConnectChanged(cb func(hasValue bool, value []Touchscreen)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []Touchscreen
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property Touchscreens a(isss)

func (v *interfaceDisplay) Touchscreens() PropTouchscreens {
	return &implPropTouchscreens{
		Impl: v,
		Name: "Touchscreens",
	}
}

// property MaxBacklightBrightness u

func (v *interfaceDisplay) MaxBacklightBrightness() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "MaxBacklightBrightness",
	}
}

// property ColorTemperatureMode i

func (v *interfaceDisplay) ColorTemperatureMode() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "ColorTemperatureMode",
	}
}

// property ColorTemperatureManual i

func (v *interfaceDisplay) ColorTemperatureManual() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "ColorTemperatureManual",
	}
}

type Monitor interface {
	monitor // interface com.deepin.daemon.Display.Monitor
	proxy.Object
}

type objectMonitor struct {
	interfaceMonitor // interface com.deepin.daemon.Display.Monitor
	proxy.ImplObject
}

func NewMonitor(conn *dbus.Conn, path dbus.ObjectPath) (Monitor, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectMonitor)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.Display", path)
	return obj, nil
}

type monitor interface {
	GoEnable(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call
	Enable(flags dbus.Flags, arg0 bool) error
	GoSetMode(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call
	SetMode(flags dbus.Flags, arg0 uint32) error
	GoSetModeBySize(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16, arg1 uint16) *dbus.Call
	SetModeBySize(flags dbus.Flags, arg0 uint16, arg1 uint16) error
	GoSetPosition(flags dbus.Flags, ch chan *dbus.Call, arg0 int16, arg1 int16) *dbus.Call
	SetPosition(flags dbus.Flags, arg0 int16, arg1 int16) error
	GoSetReflect(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16) *dbus.Call
	SetReflect(flags dbus.Flags, arg0 uint16) error
	GoSetRefreshRate(flags dbus.Flags, ch chan *dbus.Call, arg0 float64) *dbus.Call
	SetRefreshRate(flags dbus.Flags, arg0 float64) error
	GoSetRotation(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16) *dbus.Call
	SetRotation(flags dbus.Flags, arg0 uint16) error
	Name() proxy.PropString
	Enabled() proxy.PropBool
	Connected() proxy.PropBool
	X() proxy.PropInt16
	Y() proxy.PropInt16
	Width() proxy.PropUint16
	Height() proxy.PropUint16
	Rotation() proxy.PropUint16
	Reflect() proxy.PropUint16
	RefreshRate() proxy.PropDouble
	Rotations() proxy.PropUint16Array
	Reflects() proxy.PropUint16Array
	BestMode() PropModeInfo
	CurrentMode() PropModeInfo
	Modes() PropModeInfoSlice
	PreferredModes() PropModeInfoSlice
}

type interfaceMonitor struct{}

func (v *interfaceMonitor) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceMonitor) GetInterfaceName_() string {
	return "com.deepin.daemon.Display.Monitor"
}

// method Enable

func (v *interfaceMonitor) GoEnable(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, arg0)
}

func (v *interfaceMonitor) Enable(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetMode

func (v *interfaceMonitor) GoSetMode(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMode", flags, ch, arg0)
}

func (v *interfaceMonitor) SetMode(flags dbus.Flags, arg0 uint32) error {
	return (<-v.GoSetMode(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetModeBySize

func (v *interfaceMonitor) GoSetModeBySize(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16, arg1 uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetModeBySize", flags, ch, arg0, arg1)
}

func (v *interfaceMonitor) SetModeBySize(flags dbus.Flags, arg0 uint16, arg1 uint16) error {
	return (<-v.GoSetModeBySize(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetPosition

func (v *interfaceMonitor) GoSetPosition(flags dbus.Flags, ch chan *dbus.Call, arg0 int16, arg1 int16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPosition", flags, ch, arg0, arg1)
}

func (v *interfaceMonitor) SetPosition(flags dbus.Flags, arg0 int16, arg1 int16) error {
	return (<-v.GoSetPosition(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetReflect

func (v *interfaceMonitor) GoSetReflect(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetReflect", flags, ch, arg0)
}

func (v *interfaceMonitor) SetReflect(flags dbus.Flags, arg0 uint16) error {
	return (<-v.GoSetReflect(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetRefreshRate

func (v *interfaceMonitor) GoSetRefreshRate(flags dbus.Flags, ch chan *dbus.Call, arg0 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRefreshRate", flags, ch, arg0)
}

func (v *interfaceMonitor) SetRefreshRate(flags dbus.Flags, arg0 float64) error {
	return (<-v.GoSetRefreshRate(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetRotation

func (v *interfaceMonitor) GoSetRotation(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRotation", flags, ch, arg0)
}

func (v *interfaceMonitor) SetRotation(flags dbus.Flags, arg0 uint16) error {
	return (<-v.GoSetRotation(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// property Name s

func (v *interfaceMonitor) Name() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Name",
	}
}

// property Enabled b

func (v *interfaceMonitor) Enabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Enabled",
	}
}

// property Connected b

func (v *interfaceMonitor) Connected() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Connected",
	}
}

// property X n

func (v *interfaceMonitor) X() proxy.PropInt16 {
	return &proxy.ImplPropInt16{
		Impl: v,
		Name: "X",
	}
}

// property Y n

func (v *interfaceMonitor) Y() proxy.PropInt16 {
	return &proxy.ImplPropInt16{
		Impl: v,
		Name: "Y",
	}
}

// property Width q

func (v *interfaceMonitor) Width() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "Width",
	}
}

// property Height q

func (v *interfaceMonitor) Height() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "Height",
	}
}

// property Rotation q

func (v *interfaceMonitor) Rotation() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "Rotation",
	}
}

// property Reflect q

func (v *interfaceMonitor) Reflect() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "Reflect",
	}
}

// property RefreshRate d

func (v *interfaceMonitor) RefreshRate() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "RefreshRate",
	}
}

// property Rotations aq

func (v *interfaceMonitor) Rotations() proxy.PropUint16Array {
	return &proxy.ImplPropUint16Array{
		Impl: v,
		Name: "Rotations",
	}
}

// property Reflects aq

func (v *interfaceMonitor) Reflects() proxy.PropUint16Array {
	return &proxy.ImplPropUint16Array{
		Impl: v,
		Name: "Reflects",
	}
}

// property BestMode (uqqd)

func (v *interfaceMonitor) BestMode() PropModeInfo {
	return &implPropModeInfo{
		Impl: v,
		Name: "BestMode",
	}
}

// property CurrentMode (uqqd)

func (v *interfaceMonitor) CurrentMode() PropModeInfo {
	return &implPropModeInfo{
		Impl: v,
		Name: "CurrentMode",
	}
}

// property Modes a(uqqd)

func (v *interfaceMonitor) Modes() PropModeInfoSlice {
	return &implPropModeInfoSlice{
		Impl: v,
		Name: "Modes",
	}
}

// property PreferredModes a(uqqd)

func (v *interfaceMonitor) PreferredModes() PropModeInfoSlice {
	return &implPropModeInfoSlice{
		Impl: v,
		Name: "PreferredModes",
	}
}

type PropModeInfo interface {
	Get(flags dbus.Flags) (value ModeInfo, err error)
	Set(flags dbus.Flags, value ModeInfo) error
	ConnectChanged(cb func(hasValue bool, value ModeInfo)) error
}

type implPropModeInfo struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropModeInfo) Get(flags dbus.Flags) (value ModeInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropModeInfo) Set(flags dbus.Flags, value ModeInfo) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropModeInfo) ConnectChanged(cb func(hasValue bool, value ModeInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v ModeInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, ModeInfo{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

type PropModeInfoSlice interface {
	Get(flags dbus.Flags) (value []ModeInfo, err error)
	Set(flags dbus.Flags, value []ModeInfo) error
	ConnectChanged(cb func(hasValue bool, value []ModeInfo)) error
}

type implPropModeInfoSlice struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropModeInfoSlice) Get(flags dbus.Flags) (value []ModeInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropModeInfoSlice) Set(flags dbus.Flags, value []ModeInfo) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropModeInfoSlice) ConnectChanged(cb func(hasValue bool, value []ModeInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []ModeInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}
