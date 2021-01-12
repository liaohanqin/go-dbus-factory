package networkmanager

import "errors"
import "fmt"
import "github.com/linuxdeepin/go-dbus-factory/object_manager"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type ObjectManager struct {
	object_manager.ObjectManager // interface org.freedesktop.DBus.ObjectManager
	proxy.Object
}

func NewObjectManager(conn *dbus.Conn) *ObjectManager {
	obj := new(ObjectManager)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", "/org/freedesktop")
	return obj
}

type AccessPoint struct {
	accessPoint // interface org.freedesktop.NetworkManager.AccessPoint
	proxy.Object
}

func NewAccessPoint(conn *dbus.Conn, path dbus.ObjectPath) (*AccessPoint, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(AccessPoint)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

func (obj *AccessPoint) AccessPoint() *accessPoint {
	return &obj.accessPoint
}

type accessPoint struct{}

func (v *accessPoint) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*accessPoint) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.AccessPoint"
}

// signal PropertiesChanged

func (v *accessPoint) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Flags u

func (v *accessPoint) Flags() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Flags",
	}
}

// property WpaFlags u

func (v *accessPoint) WpaFlags() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "WpaFlags",
	}
}

// property RsnFlags u

func (v *accessPoint) RsnFlags() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "RsnFlags",
	}
}

// property Ssid ay

func (v *accessPoint) Ssid() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "Ssid",
	}
}

// property Frequency u

func (v *accessPoint) Frequency() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Frequency",
	}
}

// property HwAddress s

func (v *accessPoint) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property Mode u

func (v *accessPoint) Mode() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Mode",
	}
}

// property MaxBitrate u

func (v *accessPoint) MaxBitrate() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "MaxBitrate",
	}
}

// property Strength y

func (v *accessPoint) Strength() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "Strength",
	}
}

// property LastSeen i

func (v *accessPoint) LastSeen() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "LastSeen",
	}
}

type AgentManager struct {
	agentManager // interface org.freedesktop.NetworkManager.AgentManager
	proxy.Object
}

func NewAgentManager(conn *dbus.Conn) *AgentManager {
	obj := new(AgentManager)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/AgentManager")
	return obj
}

type agentManager struct{}

func (v *agentManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*agentManager) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.AgentManager"
}

// method Register

func (v *agentManager) GoRegister(flags dbus.Flags, ch chan *dbus.Call, identifier string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Register", flags, ch, identifier)
}

func (v *agentManager) Register(flags dbus.Flags, identifier string) error {
	return (<-v.GoRegister(flags, make(chan *dbus.Call, 1), identifier).Done).Err
}

// method RegisterWithCapabilities

func (v *agentManager) GoRegisterWithCapabilities(flags dbus.Flags, ch chan *dbus.Call, identifier string, capabilities uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterWithCapabilities", flags, ch, identifier, capabilities)
}

func (v *agentManager) RegisterWithCapabilities(flags dbus.Flags, identifier string, capabilities uint32) error {
	return (<-v.GoRegisterWithCapabilities(flags, make(chan *dbus.Call, 1), identifier, capabilities).Done).Err
}

// method Unregister

func (v *agentManager) GoUnregister(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Unregister", flags, ch)
}

func (v *agentManager) Unregister(flags dbus.Flags) error {
	return (<-v.GoUnregister(flags, make(chan *dbus.Call, 1)).Done).Err
}

type Checkpoint struct {
	checkpoint // interface org.freedesktop.NetworkManager.Checkpoint
	proxy.Object
}

func NewCheckpoint(conn *dbus.Conn, path dbus.ObjectPath) (*Checkpoint, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Checkpoint)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type checkpoint struct{}

func (v *checkpoint) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*checkpoint) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Checkpoint"
}

// signal PropertiesChanged

func (v *checkpoint) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Devices ao

func (v *checkpoint) Devices() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Devices",
	}
}

// property Created x

func (v *checkpoint) Created() proxy.PropInt64 {
	return proxy.PropInt64{
		Impl: v,
		Name: "Created",
	}
}

// property RollbackTimeout u

func (v *checkpoint) RollbackTimeout() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "RollbackTimeout",
	}
}

type ActiveConnection struct {
	activeConnection // interface org.freedesktop.NetworkManager.Connection.Active
	proxy.Object
}

func NewActiveConnection(conn *dbus.Conn, path dbus.ObjectPath) (*ActiveConnection, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(ActiveConnection)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type activeConnection struct{}

func (v *activeConnection) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*activeConnection) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Connection.Active"
}

// signal StateChanged

func (v *activeConnection) ConnectStateChanged(cb func(state uint32, reason uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StateChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StateChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var state uint32
		var reason uint32
		err := dbus.Store(sig.Body, &state, &reason)
		if err == nil {
			cb(state, reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PropertiesChanged

func (v *activeConnection) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Connection o

func (v *activeConnection) Connection() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Connection",
	}
}

// property SpecificObject o

func (v *activeConnection) SpecificObject() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "SpecificObject",
	}
}

// property Id s

func (v *activeConnection) Id() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Id",
	}
}

// property Uuid s

func (v *activeConnection) Uuid() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Uuid",
	}
}

// property Type s

func (v *activeConnection) Type() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Type",
	}
}

// property Devices ao

func (v *activeConnection) Devices() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Devices",
	}
}

// property State u

func (v *activeConnection) State() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "State",
	}
}

// property Default b

func (v *activeConnection) Default() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Default",
	}
}

// property Ip4Config o

func (v *activeConnection) Ip4Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Ip4Config",
	}
}

// property Dhcp4Config o

func (v *activeConnection) Dhcp4Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Dhcp4Config",
	}
}

// property Default6 b

func (v *activeConnection) Default6() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Default6",
	}
}

// property Ip6Config o

func (v *activeConnection) Ip6Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Ip6Config",
	}
}

// property Dhcp6Config o

func (v *activeConnection) Dhcp6Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Dhcp6Config",
	}
}

// property Vpn b

func (v *activeConnection) Vpn() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Vpn",
	}
}

// property Master o

func (v *activeConnection) Master() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Master",
	}
}

type Dhcp4Config struct {
	dhcp4Config // interface org.freedesktop.NetworkManager.DHCP4Config
	proxy.Object
}

func NewDhcp4Config(conn *dbus.Conn, path dbus.ObjectPath) (*Dhcp4Config, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Dhcp4Config)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type dhcp4Config struct{}

func (v *dhcp4Config) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*dhcp4Config) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.DHCP4Config"
}

// signal PropertiesChanged

func (v *dhcp4Config) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Options a{sv}

func (v *dhcp4Config) Options() PropMapStringVariant {
	return PropMapStringVariant{
		Impl: v,
		Name: "Options",
	}
}

type Dhcp6Config struct {
	dhcp6Config // interface org.freedesktop.NetworkManager.DHCP6Config
	proxy.Object
}

func NewDhcp6Config(conn *dbus.Conn, path dbus.ObjectPath) (*Dhcp6Config, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Dhcp6Config)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type dhcp6Config struct{}

func (v *dhcp6Config) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*dhcp6Config) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.DHCP6Config"
}

// signal PropertiesChanged

func (v *dhcp6Config) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Options a{sv}

func (v *dhcp6Config) Options() PropMapStringVariant {
	return PropMapStringVariant{
		Impl: v,
		Name: "Options",
	}
}

type DnsManager struct {
	dnsManager // interface org.freedesktop.NetworkManager.DnsManager
	proxy.Object
}

func NewDnsManager(conn *dbus.Conn) *DnsManager {
	obj := new(DnsManager)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/DnsManager")
	return obj
}

type dnsManager struct{}

func (v *dnsManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*dnsManager) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.DnsManager"
}

// property Mode s

func (v *dnsManager) Mode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Mode",
	}
}

// property RcManager s

func (v *dnsManager) RcManager() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "RcManager",
	}
}

// property Configuration aa{sv}

func (v *dnsManager) Configuration() PropMapSVSlice {
	return PropMapSVSlice{
		Impl: v,
		Name: "Configuration",
	}
}

type IP4Config struct {
	ip4Config // interface org.freedesktop.NetworkManager.IP4Config
	proxy.Object
}

func NewIP4Config(conn *dbus.Conn, path dbus.ObjectPath) (*IP4Config, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(IP4Config)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type ip4Config struct{}

func (v *ip4Config) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*ip4Config) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.IP4Config"
}

// signal PropertiesChanged

func (v *ip4Config) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Addresses aau

func (v *ip4Config) Addresses() PropUint32SliceSlice {
	return PropUint32SliceSlice{
		Impl: v,
		Name: "Addresses",
	}
}

// property AddressData aa{sv}

func (v *ip4Config) AddressData() PropMapSVSlice {
	return PropMapSVSlice{
		Impl: v,
		Name: "AddressData",
	}
}

// property Gateway s

func (v *ip4Config) Gateway() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Gateway",
	}
}

// property Routes aau

func (v *ip4Config) Routes() PropUint32SliceSlice {
	return PropUint32SliceSlice{
		Impl: v,
		Name: "Routes",
	}
}

// property RouteData aa{sv}

func (v *ip4Config) RouteData() PropMapSVSlice {
	return PropMapSVSlice{
		Impl: v,
		Name: "RouteData",
	}
}

// property Nameservers au

func (v *ip4Config) Nameservers() proxy.PropUint32Array {
	return proxy.PropUint32Array{
		Impl: v,
		Name: "Nameservers",
	}
}

// property Domains as

func (v *ip4Config) Domains() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Domains",
	}
}

// property Searches as

func (v *ip4Config) Searches() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Searches",
	}
}

// property DnsOptions as

func (v *ip4Config) DnsOptions() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "DnsOptions",
	}
}

// property DnsPriority i

func (v *ip4Config) DnsPriority() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "DnsPriority",
	}
}

// property WinsServers au

func (v *ip4Config) WinsServers() proxy.PropUint32Array {
	return proxy.PropUint32Array{
		Impl: v,
		Name: "WinsServers",
	}
}

type IP6Config struct {
	ip6Config // interface org.freedesktop.NetworkManager.IP6Config
	proxy.Object
}

func NewIP6Config(conn *dbus.Conn, path dbus.ObjectPath) (*IP6Config, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(IP6Config)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type ip6Config struct{}

func (v *ip6Config) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*ip6Config) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.IP6Config"
}

// signal PropertiesChanged

func (v *ip6Config) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Addresses a(ayuay)

func (v *ip6Config) Addresses() PropIP6ConfigAddresses {
	return PropIP6ConfigAddresses{
		Impl: v,
	}
}

type PropIP6ConfigAddresses struct {
	Impl proxy.Implementer
}

func (p PropIP6ConfigAddresses) Get(flags dbus.Flags) (value []IP6Address, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Addresses", &value)
	return
}

func (p PropIP6ConfigAddresses) ConnectChanged(cb func(hasValue bool, value []IP6Address)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []IP6Address
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
		"Addresses", cb0)
}

// property AddressData aa{sv}

func (v *ip6Config) AddressData() PropMapSVSlice {
	return PropMapSVSlice{
		Impl: v,
		Name: "AddressData",
	}
}

// property Gateway s

func (v *ip6Config) Gateway() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Gateway",
	}
}

// property Routes a(ayuayu)

func (v *ip6Config) Routes() PropIP6ConfigRoutes {
	return PropIP6ConfigRoutes{
		Impl: v,
	}
}

type PropIP6ConfigRoutes struct {
	Impl proxy.Implementer
}

func (p PropIP6ConfigRoutes) Get(flags dbus.Flags) (value []IP6Route, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Routes", &value)
	return
}

func (p PropIP6ConfigRoutes) ConnectChanged(cb func(hasValue bool, value []IP6Route)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []IP6Route
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
		"Routes", cb0)
}

// property RouteData aa{sv}

func (v *ip6Config) RouteData() PropMapSVSlice {
	return PropMapSVSlice{
		Impl: v,
		Name: "RouteData",
	}
}

// property Nameservers aay

func (v *ip6Config) Nameservers() PropIP6NameServers {
	return PropIP6NameServers{
		Impl: v,
	}
}

type PropIP6NameServers struct {
	Impl proxy.Implementer
}

func (p PropIP6NameServers) Get(flags dbus.Flags) (value [][]byte, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Nameservers", &value)
	return
}

func (p PropIP6NameServers) ConnectChanged(cb func(hasValue bool, value [][]byte)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v [][]byte
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
		"Nameservers", cb0)
}

// property Domains as

func (v *ip6Config) Domains() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Domains",
	}
}

// property Searches as

func (v *ip6Config) Searches() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Searches",
	}
}

// property DnsOptions as

func (v *ip6Config) DnsOptions() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "DnsOptions",
	}
}

// property DnsPriority i

func (v *ip6Config) DnsPriority() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "DnsPriority",
	}
}

type Manager struct {
	manager // interface org.freedesktop.NetworkManager
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")
	return obj
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager"
}

// method Reload

func (v *manager) GoReload(flags dbus.Flags, ch chan *dbus.Call, flags0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reload", flags, ch, flags0)
}

func (v *manager) Reload(flags dbus.Flags, flags0 uint32) error {
	return (<-v.GoReload(flags, make(chan *dbus.Call, 1), flags0).Done).Err
}

// method GetDevices

func (v *manager) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevices", flags, ch)
}

func (*manager) StoreGetDevices(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *manager) GetDevices(flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetDevices(
		<-v.GoGetDevices(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAllDevices

func (v *manager) GoGetAllDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllDevices", flags, ch)
}

func (*manager) StoreGetAllDevices(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *manager) GetAllDevices(flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetAllDevices(
		<-v.GoGetAllDevices(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetDeviceByIpIface

func (v *manager) GoGetDeviceByIpIface(flags dbus.Flags, ch chan *dbus.Call, iface string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDeviceByIpIface", flags, ch, iface)
}

func (*manager) StoreGetDeviceByIpIface(call *dbus.Call) (device dbus.ObjectPath, err error) {
	err = call.Store(&device)
	return
}

func (v *manager) GetDeviceByIpIface(flags dbus.Flags, iface string) (device dbus.ObjectPath, err error) {
	return v.StoreGetDeviceByIpIface(
		<-v.GoGetDeviceByIpIface(flags, make(chan *dbus.Call, 1), iface).Done)
}

// method ActivateConnection

func (v *manager) GoActivateConnection(flags dbus.Flags, ch chan *dbus.Call, connection dbus.ObjectPath, device dbus.ObjectPath, specific_object dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateConnection", flags, ch, connection, device, specific_object)
}

func (*manager) StoreActivateConnection(call *dbus.Call) (active_connection dbus.ObjectPath, err error) {
	err = call.Store(&active_connection)
	return
}

func (v *manager) ActivateConnection(flags dbus.Flags, connection dbus.ObjectPath, device dbus.ObjectPath, specific_object dbus.ObjectPath) (active_connection dbus.ObjectPath, err error) {
	return v.StoreActivateConnection(
		<-v.GoActivateConnection(flags, make(chan *dbus.Call, 1), connection, device, specific_object).Done)
}

// method AddAndActivateConnection

func (v *manager) GoAddAndActivateConnection(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant, device dbus.ObjectPath, specific_object dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddAndActivateConnection", flags, ch, connection, device, specific_object)
}

func (*manager) StoreAddAndActivateConnection(call *dbus.Call) (path dbus.ObjectPath, active_connection dbus.ObjectPath, err error) {
	err = call.Store(&path, &active_connection)
	return
}

func (v *manager) AddAndActivateConnection(flags dbus.Flags, connection map[string]map[string]dbus.Variant, device dbus.ObjectPath, specific_object dbus.ObjectPath) (path dbus.ObjectPath, active_connection dbus.ObjectPath, err error) {
	return v.StoreAddAndActivateConnection(
		<-v.GoAddAndActivateConnection(flags, make(chan *dbus.Call, 1), connection, device, specific_object).Done)
}

// method DeactivateConnection

func (v *manager) GoDeactivateConnection(flags dbus.Flags, ch chan *dbus.Call, active_connection dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeactivateConnection", flags, ch, active_connection)
}

func (v *manager) DeactivateConnection(flags dbus.Flags, active_connection dbus.ObjectPath) error {
	return (<-v.GoDeactivateConnection(flags, make(chan *dbus.Call, 1), active_connection).Done).Err
}

// method Sleep

func (v *manager) GoSleep(flags dbus.Flags, ch chan *dbus.Call, sleep bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Sleep", flags, ch, sleep)
}

func (v *manager) Sleep(flags dbus.Flags, sleep bool) error {
	return (<-v.GoSleep(flags, make(chan *dbus.Call, 1), sleep).Done).Err
}

// method Enable

func (v *manager) GoEnable(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, enable)
}

func (v *manager) Enable(flags dbus.Flags, enable bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

// method GetPermissions

func (v *manager) GoGetPermissions(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetPermissions", flags, ch)
}

func (*manager) StoreGetPermissions(call *dbus.Call) (permissions map[string]string, err error) {
	err = call.Store(&permissions)
	return
}

func (v *manager) GetPermissions(flags dbus.Flags) (permissions map[string]string, err error) {
	return v.StoreGetPermissions(
		<-v.GoGetPermissions(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetLogging

func (v *manager) GoSetLogging(flags dbus.Flags, ch chan *dbus.Call, level string, domains string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLogging", flags, ch, level, domains)
}

func (v *manager) SetLogging(flags dbus.Flags, level string, domains string) error {
	return (<-v.GoSetLogging(flags, make(chan *dbus.Call, 1), level, domains).Done).Err
}

// method GetLogging

func (v *manager) GoGetLogging(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetLogging", flags, ch)
}

func (*manager) StoreGetLogging(call *dbus.Call) (level string, domains string, err error) {
	err = call.Store(&level, &domains)
	return
}

func (v *manager) GetLogging(flags dbus.Flags) (level string, domains string, err error) {
	return v.StoreGetLogging(
		<-v.GoGetLogging(flags, make(chan *dbus.Call, 1)).Done)
}

// method CheckConnectivity

func (v *manager) GoCheckConnectivity(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CheckConnectivity", flags, ch)
}

func (*manager) StoreCheckConnectivity(call *dbus.Call) (connectivity uint32, err error) {
	err = call.Store(&connectivity)
	return
}

func (v *manager) CheckConnectivity(flags dbus.Flags) (connectivity uint32, err error) {
	return v.StoreCheckConnectivity(
		<-v.GoCheckConnectivity(flags, make(chan *dbus.Call, 1)).Done)
}

// method state

//func (v *manager) GoState(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
//	return v.GetObject_().Go_(v.GetInterfaceName_()+".state", flags, ch)
//}
//
//func (*manager) StoreState(call *dbus.Call) (state uint32, err error) {
//	err = call.Store(&state)
//	return
//}
//
//func (v *manager) State(flags dbus.Flags) (state uint32, err error) {
//	return v.StoreState(
//		<-v.GoState(flags, make(chan *dbus.Call, 1)).Done)
//}

// method CheckpointCreate

func (v *manager) GoCheckpointCreate(flags dbus.Flags, ch chan *dbus.Call, devices []dbus.ObjectPath, rollback_timeout uint32, flags0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CheckpointCreate", flags, ch, devices, rollback_timeout, flags0)
}

func (*manager) StoreCheckpointCreate(call *dbus.Call) (checkpoint dbus.ObjectPath, err error) {
	err = call.Store(&checkpoint)
	return
}

func (v *manager) CheckpointCreate(flags dbus.Flags, devices []dbus.ObjectPath, rollback_timeout uint32, flags0 uint32) (checkpoint dbus.ObjectPath, err error) {
	return v.StoreCheckpointCreate(
		<-v.GoCheckpointCreate(flags, make(chan *dbus.Call, 1), devices, rollback_timeout, flags0).Done)
}

// method CheckpointDestroy

func (v *manager) GoCheckpointDestroy(flags dbus.Flags, ch chan *dbus.Call, checkpoint dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CheckpointDestroy", flags, ch, checkpoint)
}

func (v *manager) CheckpointDestroy(flags dbus.Flags, checkpoint dbus.ObjectPath) error {
	return (<-v.GoCheckpointDestroy(flags, make(chan *dbus.Call, 1), checkpoint).Done).Err
}

// method CheckpointRollback

func (v *manager) GoCheckpointRollback(flags dbus.Flags, ch chan *dbus.Call, checkpoint dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CheckpointRollback", flags, ch, checkpoint)
}

func (*manager) StoreCheckpointRollback(call *dbus.Call) (result map[string]uint32, err error) {
	err = call.Store(&result)
	return
}

func (v *manager) CheckpointRollback(flags dbus.Flags, checkpoint dbus.ObjectPath) (result map[string]uint32, err error) {
	return v.StoreCheckpointRollback(
		<-v.GoCheckpointRollback(flags, make(chan *dbus.Call, 1), checkpoint).Done)
}

// signal CheckPermissions

func (v *manager) ConnectCheckPermissions(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CheckPermissions", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CheckPermissions",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal StateChanged

func (v *manager) ConnectStateChanged(cb func(state uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StateChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StateChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var state uint32
		err := dbus.Store(sig.Body, &state)
		if err == nil {
			cb(state)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PropertiesChanged

func (v *manager) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceAdded

func (v *manager) ConnectDeviceAdded(cb func(device_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var device_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &device_path)
		if err == nil {
			cb(device_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceRemoved

func (v *manager) ConnectDeviceRemoved(cb func(device_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var device_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &device_path)
		if err == nil {
			cb(device_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Devices ao

func (v *manager) Devices() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Devices",
	}
}

// property AllDevices ao

func (v *manager) AllDevices() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "AllDevices",
	}
}

// property NetworkingEnabled b

func (v *manager) NetworkingEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NetworkingEnabled",
	}
}

// property WirelessEnabled b

func (v *manager) WirelessEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WirelessEnabled",
	}
}

// property WirelessHardwareEnabled b

func (v *manager) WirelessHardwareEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WirelessHardwareEnabled",
	}
}

// property WwanEnabled b

func (v *manager) WwanEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WwanEnabled",
	}
}

// property WwanHardwareEnabled b

func (v *manager) WwanHardwareEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WwanHardwareEnabled",
	}
}

// property WimaxEnabled b

func (v *manager) WimaxEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WimaxEnabled",
	}
}

// property WimaxHardwareEnabled b

func (v *manager) WimaxHardwareEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WimaxHardwareEnabled",
	}
}

// property ActiveConnections ao

func (v *manager) ActiveConnections() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "ActiveConnections",
	}
}

// property PrimaryConnection o

func (v *manager) PrimaryConnection() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "PrimaryConnection",
	}
}

// property PrimaryConnectionType s

func (v *manager) PrimaryConnectionType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PrimaryConnectionType",
	}
}

// property Metered u

func (v *manager) Metered() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Metered",
	}
}

// property ActivatingConnection o

func (v *manager) ActivatingConnection() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "ActivatingConnection",
	}
}

// property Startup b

func (v *manager) Startup() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Startup",
	}
}

// property Version s

func (v *manager) Version() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Version",
	}
}

// property Capabilities au

func (v *manager) Capabilities() proxy.PropUint32Array {
	return proxy.PropUint32Array{
		Impl: v,
		Name: "Capabilities",
	}
}

// property State u

func (v *manager) State() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "State",
	}
}

// property Connectivity u

func (v *manager) Connectivity() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Connectivity",
	}
}

// property GlobalDnsConfiguration a{sv}

func (v *manager) GlobalDnsConfiguration() PropMapStringVariant {
	return PropMapStringVariant{
		Impl: v,
		Name: "GlobalDnsConfiguration",
	}
}

type PPP struct {
	ppp // interface org.freedesktop.NetworkManager.PPP
	proxy.Object
}

func NewPPP(conn *dbus.Conn, path dbus.ObjectPath) (*PPP, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(PPP)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type ppp struct{}

func (v *ppp) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*ppp) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.PPP"
}

// method NeedSecrets

func (v *ppp) GoNeedSecrets(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NeedSecrets", flags, ch)
}

func (*ppp) StoreNeedSecrets(call *dbus.Call) (username string, password string, err error) {
	err = call.Store(&username, &password)
	return
}

func (v *ppp) NeedSecrets(flags dbus.Flags) (username string, password string, err error) {
	return v.StoreNeedSecrets(
		<-v.GoNeedSecrets(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetIp4Config

func (v *ppp) GoSetIp4Config(flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIp4Config", flags, ch, config)
}

func (v *ppp) SetIp4Config(flags dbus.Flags, config map[string]dbus.Variant) error {
	return (<-v.GoSetIp4Config(flags, make(chan *dbus.Call, 1), config).Done).Err
}

// method SetIp6Config

func (v *ppp) GoSetIp6Config(flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIp6Config", flags, ch, config)
}

func (v *ppp) SetIp6Config(flags dbus.Flags, config map[string]dbus.Variant) error {
	return (<-v.GoSetIp6Config(flags, make(chan *dbus.Call, 1), config).Done).Err
}

// method SetState

func (v *ppp) GoSetState(flags dbus.Flags, ch chan *dbus.Call, state uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetState", flags, ch, state)
}

func (v *ppp) SetState(flags dbus.Flags, state uint32) error {
	return (<-v.GoSetState(flags, make(chan *dbus.Call, 1), state).Done).Err
}

type SecretAgent struct {
	secretAgent // interface org.freedesktop.NetworkManager.SecretAgent
	proxy.Object
}

func NewSecretAgent(conn *dbus.Conn, path dbus.ObjectPath) (*SecretAgent, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(SecretAgent)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type secretAgent struct{}

func (v *secretAgent) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*secretAgent) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.SecretAgent"
}

// method GetSecrets

func (v *secretAgent) GoGetSecrets(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath, setting_name string, hints []string, flags0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSecrets", flags, ch, connection, connection_path, setting_name, hints, flags0)
}

func (*secretAgent) StoreGetSecrets(call *dbus.Call) (secrets map[string]map[string]dbus.Variant, err error) {
	err = call.Store(&secrets)
	return
}

func (v *secretAgent) GetSecrets(flags dbus.Flags, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath, setting_name string, hints []string, flags0 uint32) (secrets map[string]map[string]dbus.Variant, err error) {
	return v.StoreGetSecrets(
		<-v.GoGetSecrets(flags, make(chan *dbus.Call, 1), connection, connection_path, setting_name, hints, flags0).Done)
}

// method CancelGetSecrets

func (v *secretAgent) GoCancelGetSecrets(flags dbus.Flags, ch chan *dbus.Call, connection_path dbus.ObjectPath, setting_name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelGetSecrets", flags, ch, connection_path, setting_name)
}

func (v *secretAgent) CancelGetSecrets(flags dbus.Flags, connection_path dbus.ObjectPath, setting_name string) error {
	return (<-v.GoCancelGetSecrets(flags, make(chan *dbus.Call, 1), connection_path, setting_name).Done).Err
}

// method SaveSecrets

func (v *secretAgent) GoSaveSecrets(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SaveSecrets", flags, ch, connection, connection_path)
}

func (v *secretAgent) SaveSecrets(flags dbus.Flags, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath) error {
	return (<-v.GoSaveSecrets(flags, make(chan *dbus.Call, 1), connection, connection_path).Done).Err
}

// method DeleteSecrets

func (v *secretAgent) GoDeleteSecrets(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteSecrets", flags, ch, connection, connection_path)
}

func (v *secretAgent) DeleteSecrets(flags dbus.Flags, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath) error {
	return (<-v.GoDeleteSecrets(flags, make(chan *dbus.Call, 1), connection, connection_path).Done).Err
}

type ConnectionSettings struct {
	connectionSettings // interface org.freedesktop.NetworkManager.Settings.Connection
	proxy.Object
}

func NewConnectionSettings(conn *dbus.Conn, path dbus.ObjectPath) (*ConnectionSettings, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(ConnectionSettings)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type connectionSettings struct{}

func (v *connectionSettings) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*connectionSettings) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Settings.Connection"
}

// method Update

func (v *connectionSettings) GoUpdate(flags dbus.Flags, ch chan *dbus.Call, properties map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Update", flags, ch, properties)
}

func (v *connectionSettings) Update(flags dbus.Flags, properties map[string]map[string]dbus.Variant) error {
	return (<-v.GoUpdate(flags, make(chan *dbus.Call, 1), properties).Done).Err
}

// method UpdateUnsaved

func (v *connectionSettings) GoUpdateUnsaved(flags dbus.Flags, ch chan *dbus.Call, properties map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateUnsaved", flags, ch, properties)
}

func (v *connectionSettings) UpdateUnsaved(flags dbus.Flags, properties map[string]map[string]dbus.Variant) error {
	return (<-v.GoUpdateUnsaved(flags, make(chan *dbus.Call, 1), properties).Done).Err
}

// method Delete

func (v *connectionSettings) GoDelete(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch)
}

func (v *connectionSettings) Delete(flags dbus.Flags) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method GetSettings

func (v *connectionSettings) GoGetSettings(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSettings", flags, ch)
}

func (*connectionSettings) StoreGetSettings(call *dbus.Call) (settings map[string]map[string]dbus.Variant, err error) {
	err = call.Store(&settings)
	return
}

func (v *connectionSettings) GetSettings(flags dbus.Flags) (settings map[string]map[string]dbus.Variant, err error) {
	return v.StoreGetSettings(
		<-v.GoGetSettings(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetSecrets

func (v *connectionSettings) GoGetSecrets(flags dbus.Flags, ch chan *dbus.Call, setting_name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSecrets", flags, ch, setting_name)
}

func (*connectionSettings) StoreGetSecrets(call *dbus.Call) (secrets map[string]map[string]dbus.Variant, err error) {
	err = call.Store(&secrets)
	return
}

func (v *connectionSettings) GetSecrets(flags dbus.Flags, setting_name string) (secrets map[string]map[string]dbus.Variant, err error) {
	return v.StoreGetSecrets(
		<-v.GoGetSecrets(flags, make(chan *dbus.Call, 1), setting_name).Done)
}

// method ClearSecrets

func (v *connectionSettings) GoClearSecrets(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearSecrets", flags, ch)
}

func (v *connectionSettings) ClearSecrets(flags dbus.Flags) error {
	return (<-v.GoClearSecrets(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Save

func (v *connectionSettings) GoSave(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Save", flags, ch)
}

func (v *connectionSettings) Save(flags dbus.Flags) error {
	return (<-v.GoSave(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Updated

func (v *connectionSettings) ConnectUpdated(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Updated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Updated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Removed

func (v *connectionSettings) ConnectRemoved(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Removed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Removed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PropertiesChanged

func (v *connectionSettings) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Unsaved b

func (v *connectionSettings) Unsaved() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Unsaved",
	}
}

type Settings struct {
	settings // interface org.freedesktop.NetworkManager.Settings
	proxy.Object
}

func NewSettings(conn *dbus.Conn) *Settings {
	obj := new(Settings)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/Settings")
	return obj
}

type settings struct{}

func (v *settings) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*settings) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Settings"
}

// method ListConnections

func (v *settings) GoListConnections(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListConnections", flags, ch)
}

func (*settings) StoreListConnections(call *dbus.Call) (connections []dbus.ObjectPath, err error) {
	err = call.Store(&connections)
	return
}

func (v *settings) ListConnections(flags dbus.Flags) (connections []dbus.ObjectPath, err error) {
	return v.StoreListConnections(
		<-v.GoListConnections(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetConnectionByUuid

func (v *settings) GoGetConnectionByUuid(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionByUuid", flags, ch, uuid)
}

func (*settings) StoreGetConnectionByUuid(call *dbus.Call) (connection dbus.ObjectPath, err error) {
	err = call.Store(&connection)
	return
}

func (v *settings) GetConnectionByUuid(flags dbus.Flags, uuid string) (connection dbus.ObjectPath, err error) {
	return v.StoreGetConnectionByUuid(
		<-v.GoGetConnectionByUuid(flags, make(chan *dbus.Call, 1), uuid).Done)
}

// method AddConnection

func (v *settings) GoAddConnection(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddConnection", flags, ch, connection)
}

func (*settings) StoreAddConnection(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *settings) AddConnection(flags dbus.Flags, connection map[string]map[string]dbus.Variant) (path dbus.ObjectPath, err error) {
	return v.StoreAddConnection(
		<-v.GoAddConnection(flags, make(chan *dbus.Call, 1), connection).Done)
}

// method AddConnectionUnsaved

func (v *settings) GoAddConnectionUnsaved(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddConnectionUnsaved", flags, ch, connection)
}

func (*settings) StoreAddConnectionUnsaved(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *settings) AddConnectionUnsaved(flags dbus.Flags, connection map[string]map[string]dbus.Variant) (path dbus.ObjectPath, err error) {
	return v.StoreAddConnectionUnsaved(
		<-v.GoAddConnectionUnsaved(flags, make(chan *dbus.Call, 1), connection).Done)
}

// method LoadConnections

func (v *settings) GoLoadConnections(flags dbus.Flags, ch chan *dbus.Call, filenames []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LoadConnections", flags, ch, filenames)
}

func (*settings) StoreLoadConnections(call *dbus.Call) (status bool, failures []string, err error) {
	err = call.Store(&status, &failures)
	return
}

func (v *settings) LoadConnections(flags dbus.Flags, filenames []string) (status bool, failures []string, err error) {
	return v.StoreLoadConnections(
		<-v.GoLoadConnections(flags, make(chan *dbus.Call, 1), filenames).Done)
}

// method ReloadConnections

func (v *settings) GoReloadConnections(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReloadConnections", flags, ch)
}

func (*settings) StoreReloadConnections(call *dbus.Call) (status bool, err error) {
	err = call.Store(&status)
	return
}

func (v *settings) ReloadConnections(flags dbus.Flags) (status bool, err error) {
	return v.StoreReloadConnections(
		<-v.GoReloadConnections(flags, make(chan *dbus.Call, 1)).Done)
}

// method SaveHostname

func (v *settings) GoSaveHostname(flags dbus.Flags, ch chan *dbus.Call, hostname string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SaveHostname", flags, ch, hostname)
}

func (v *settings) SaveHostname(flags dbus.Flags, hostname string) error {
	return (<-v.GoSaveHostname(flags, make(chan *dbus.Call, 1), hostname).Done).Err
}

// signal PropertiesChanged

func (v *settings) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NewConnection

func (v *settings) ConnectNewConnection(cb func(connection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NewConnection", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NewConnection",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var connection dbus.ObjectPath
		err := dbus.Store(sig.Body, &connection)
		if err == nil {
			cb(connection)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ConnectionRemoved

func (v *settings) ConnectConnectionRemoved(cb func(connection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ConnectionRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ConnectionRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var connection dbus.ObjectPath
		err := dbus.Store(sig.Body, &connection)
		if err == nil {
			cb(connection)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Connections ao

func (v *settings) Connections() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Connections",
	}
}

// property Hostname s

func (v *settings) Hostname() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Hostname",
	}
}

// property CanModify b

func (v *settings) CanModify() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanModify",
	}
}

type VpnConnection struct {
	vpnConnection // interface org.freedesktop.NetworkManager.VPN.Connection
	proxy.Object
}

func NewVpnConnection(conn *dbus.Conn, path dbus.ObjectPath) (*VpnConnection, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(VpnConnection)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type vpnConnection struct{}

func (v *vpnConnection) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*vpnConnection) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.VPN.Connection"
}

// signal PropertiesChanged

func (v *vpnConnection) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VpnStateChanged

func (v *vpnConnection) ConnectVpnStateChanged(cb func(state uint32, reason uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VpnStateChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VpnStateChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var state uint32
		var reason uint32
		err := dbus.Store(sig.Body, &state, &reason)
		if err == nil {
			cb(state, reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property VpnState u

func (v *vpnConnection) VpnState() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "VpnState",
	}
}

// property Banner s

func (v *vpnConnection) Banner() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Banner",
	}
}

type VpnPlugin struct {
	vpnPlugin // interface org.freedesktop.NetworkManager.VPN.Plugin
	proxy.Object
}

func NewVpnPlugin(conn *dbus.Conn, path dbus.ObjectPath) (*VpnPlugin, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(VpnPlugin)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type vpnPlugin struct{}

func (v *vpnPlugin) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*vpnPlugin) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.VPN.Plugin"
}

// method Connect

func (v *vpnPlugin) GoConnect(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Connect", flags, ch, connection)
}

func (v *vpnPlugin) Connect(flags dbus.Flags, connection map[string]map[string]dbus.Variant) error {
	return (<-v.GoConnect(flags, make(chan *dbus.Call, 1), connection).Done).Err
}

// method ConnectInteractive

func (v *vpnPlugin) GoConnectInteractive(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant, details map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ConnectInteractive", flags, ch, connection, details)
}

func (v *vpnPlugin) ConnectInteractive(flags dbus.Flags, connection map[string]map[string]dbus.Variant, details map[string]dbus.Variant) error {
	return (<-v.GoConnectInteractive(flags, make(chan *dbus.Call, 1), connection, details).Done).Err
}

// method NeedSecrets

func (v *vpnPlugin) GoNeedSecrets(flags dbus.Flags, ch chan *dbus.Call, settings map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NeedSecrets", flags, ch, settings)
}

func (*vpnPlugin) StoreNeedSecrets(call *dbus.Call) (setting_name string, err error) {
	err = call.Store(&setting_name)
	return
}

func (v *vpnPlugin) NeedSecrets(flags dbus.Flags, settings map[string]map[string]dbus.Variant) (setting_name string, err error) {
	return v.StoreNeedSecrets(
		<-v.GoNeedSecrets(flags, make(chan *dbus.Call, 1), settings).Done)
}

// method Disconnect

func (v *vpnPlugin) GoDisconnect(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Disconnect", flags, ch)
}

func (v *vpnPlugin) Disconnect(flags dbus.Flags) error {
	return (<-v.GoDisconnect(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetConfig

func (v *vpnPlugin) GoSetConfig(flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetConfig", flags, ch, config)
}

func (v *vpnPlugin) SetConfig(flags dbus.Flags, config map[string]dbus.Variant) error {
	return (<-v.GoSetConfig(flags, make(chan *dbus.Call, 1), config).Done).Err
}

// method SetIp4Config

func (v *vpnPlugin) GoSetIp4Config(flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIp4Config", flags, ch, config)
}

func (v *vpnPlugin) SetIp4Config(flags dbus.Flags, config map[string]dbus.Variant) error {
	return (<-v.GoSetIp4Config(flags, make(chan *dbus.Call, 1), config).Done).Err
}

// method SetIp6Config

func (v *vpnPlugin) GoSetIp6Config(flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIp6Config", flags, ch, config)
}

func (v *vpnPlugin) SetIp6Config(flags dbus.Flags, config map[string]dbus.Variant) error {
	return (<-v.GoSetIp6Config(flags, make(chan *dbus.Call, 1), config).Done).Err
}

// method SetFailure

func (v *vpnPlugin) GoSetFailure(flags dbus.Flags, ch chan *dbus.Call, reason string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetFailure", flags, ch, reason)
}

func (v *vpnPlugin) SetFailure(flags dbus.Flags, reason string) error {
	return (<-v.GoSetFailure(flags, make(chan *dbus.Call, 1), reason).Done).Err
}

// method NewSecrets

func (v *vpnPlugin) GoNewSecrets(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NewSecrets", flags, ch, connection)
}

func (v *vpnPlugin) NewSecrets(flags dbus.Flags, connection map[string]map[string]dbus.Variant) error {
	return (<-v.GoNewSecrets(flags, make(chan *dbus.Call, 1), connection).Done).Err
}

// signal StateChanged

func (v *vpnPlugin) ConnectStateChanged(cb func(state uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StateChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StateChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var state uint32
		err := dbus.Store(sig.Body, &state)
		if err == nil {
			cb(state)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SecretsRequired

func (v *vpnPlugin) ConnectSecretsRequired(cb func(message string, secrets []string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SecretsRequired", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SecretsRequired",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var message string
		var secrets []string
		err := dbus.Store(sig.Body, &message, &secrets)
		if err == nil {
			cb(message, secrets)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Config

func (v *vpnPlugin) ConnectConfig(cb func(config map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Config", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Config",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var config map[string]dbus.Variant
		err := dbus.Store(sig.Body, &config)
		if err == nil {
			cb(config)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Ip4Config

func (v *vpnPlugin) ConnectIp4Config(cb func(ip4config map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Ip4Config", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Ip4Config",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ip4config map[string]dbus.Variant
		err := dbus.Store(sig.Body, &ip4config)
		if err == nil {
			cb(ip4config)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Ip6Config

func (v *vpnPlugin) ConnectIp6Config(cb func(ip6config map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Ip6Config", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Ip6Config",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ip6config map[string]dbus.Variant
		err := dbus.Store(sig.Body, &ip6config)
		if err == nil {
			cb(ip6config)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal LoginBanner

func (v *vpnPlugin) ConnectLoginBanner(cb func(banner string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "LoginBanner", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".LoginBanner",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var banner string
		err := dbus.Store(sig.Body, &banner)
		if err == nil {
			cb(banner)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Failure

func (v *vpnPlugin) ConnectFailure(cb func(reason uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Failure", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Failure",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var reason uint32
		err := dbus.Store(sig.Body, &reason)
		if err == nil {
			cb(reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property State u

func (v *vpnPlugin) State() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "State",
	}
}

type Device struct {
	device           // interface org.freedesktop.NetworkManager.Device
	deviceAdsl       // interface org.freedesktop.NetworkManager.Device.Adsl
	deviceBluetooth  // interface org.freedesktop.NetworkManager.Device.Bluetooth
	deviceBond       // interface org.freedesktop.NetworkManager.Device.Bond
	deviceBridge     // interface org.freedesktop.NetworkManager.Device.Bridge
	deviceGeneric    // interface org.freedesktop.NetworkManager.Device.Generic
	deviceInfiniband // interface org.freedesktop.NetworkManager.Device.Infiniband
	deviceIPTunnel   // interface org.freedesktop.NetworkManager.Device.IPTunnel
	deviceMacsec     // interface org.freedesktop.NetworkManager.Device.Macsec
	deviceMacvlan    // interface org.freedesktop.NetworkManager.Device.Macvlan
	deviceModem      // interface org.freedesktop.NetworkManager.Device.Modem
	deviceOlpcMesh   // interface org.freedesktop.NetworkManager.Device.OlpcMesh
	deviceStatistics // interface org.freedesktop.NetworkManager.Device.Statistics
	deviceTeam       // interface org.freedesktop.NetworkManager.Device.Team
	deviceTun        // interface org.freedesktop.NetworkManager.Device.Tun
	deviceVeth       // interface org.freedesktop.NetworkManager.Device.Veth
	deviceVlan       // interface org.freedesktop.NetworkManager.Device.Vlan
	deviceVxlan      // interface org.freedesktop.NetworkManager.Device.Vxlan
	deviceWiMax      // interface org.freedesktop.NetworkManager.Device.WiMax
	deviceWired      // interface org.freedesktop.NetworkManager.Device.Wired
	deviceWireless   // interface org.freedesktop.NetworkManager.Device.Wireless
	proxy.Object
}

func NewDevice(conn *dbus.Conn, path dbus.ObjectPath) (*Device, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Device)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

func (obj *Device) Device() *device {
	return &obj.device
}

type device struct{}

func (v *device) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*device) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device"
}

// method Reapply

func (v *device) GoReapply(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant, version_id uint64, flags0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reapply", flags, ch, connection, version_id, flags0)
}

func (v *device) Reapply(flags dbus.Flags, connection map[string]map[string]dbus.Variant, version_id uint64, flags0 uint32) error {
	return (<-v.GoReapply(flags, make(chan *dbus.Call, 1), connection, version_id, flags0).Done).Err
}

// method GetAppliedConnection

func (v *device) GoGetAppliedConnection(flags dbus.Flags, ch chan *dbus.Call, flags0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAppliedConnection", flags, ch, flags0)
}

func (*device) StoreGetAppliedConnection(call *dbus.Call) (connection map[string]map[string]dbus.Variant, version_id uint64, err error) {
	err = call.Store(&connection, &version_id)
	return
}

func (v *device) GetAppliedConnection(flags dbus.Flags, flags0 uint32) (connection map[string]map[string]dbus.Variant, version_id uint64, err error) {
	return v.StoreGetAppliedConnection(
		<-v.GoGetAppliedConnection(flags, make(chan *dbus.Call, 1), flags0).Done)
}

// method Disconnect

func (v *device) GoDisconnect(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Disconnect", flags, ch)
}

func (v *device) Disconnect(flags dbus.Flags) error {
	return (<-v.GoDisconnect(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Delete

func (v *device) GoDelete(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch)
}

func (v *device) Delete(flags dbus.Flags) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal StateChanged

func (v *device) ConnectStateChanged(cb func(new_state uint32, old_state uint32, reason uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StateChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StateChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var new_state uint32
		var old_state uint32
		var reason uint32
		err := dbus.Store(sig.Body, &new_state, &old_state, &reason)
		if err == nil {
			cb(new_state, old_state, reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Udi s

func (v *device) Udi() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Udi",
	}
}

// property Path s

func (v *device) Path() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Path",
	}
}

// property Interface s

func (v *device) Interface() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Interface",
	}
}

// property IpInterface s

func (v *device) IpInterface() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IpInterface",
	}
}

// property Driver s

func (v *device) Driver() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Driver",
	}
}

// property DriverVersion s

func (v *device) DriverVersion() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DriverVersion",
	}
}

// property FirmwareVersion s

func (v *device) FirmwareVersion() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FirmwareVersion",
	}
}

// property Capabilities u

func (v *device) Capabilities() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Capabilities",
	}
}

// property Ip4Address u

func (v *device) Ip4Address() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Ip4Address",
	}
}

// property State u

func (v *device) State() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "State",
	}
}

// property StateReason (uu)

func (v *device) StateReason() PropDeviceStateReason {
	return PropDeviceStateReason{
		Impl: v,
	}
}

type PropDeviceStateReason struct {
	Impl proxy.Implementer
}

func (p PropDeviceStateReason) Get(flags dbus.Flags) (value DeviceStateReason, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"StateReason", &value)
	return
}

func (p PropDeviceStateReason) ConnectChanged(cb func(hasValue bool, value DeviceStateReason)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v DeviceStateReason
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, DeviceStateReason{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"StateReason", cb0)
}

// property ActiveConnection o

func (v *device) ActiveConnection() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "ActiveConnection",
	}
}

// property Ip4Config o

func (v *device) Ip4Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Ip4Config",
	}
}

// property Dhcp4Config o

func (v *device) Dhcp4Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Dhcp4Config",
	}
}

// property Ip6Config o

func (v *device) Ip6Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Ip6Config",
	}
}

// property Dhcp6Config o

func (v *device) Dhcp6Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Dhcp6Config",
	}
}

// property Managed b

func (v *device) Managed() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Managed",
	}
}

// property Autoconnect b

func (v *device) Autoconnect() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Autoconnect",
	}
}

// property FirmwareMissing b

func (v *device) FirmwareMissing() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "FirmwareMissing",
	}
}

// property NmPluginMissing b

func (v *device) NmPluginMissing() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NmPluginMissing",
	}
}

// property DeviceType u

func (v *device) DeviceType() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "DeviceType",
	}
}

// property AvailableConnections ao

func (v *device) AvailableConnections() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "AvailableConnections",
	}
}

// property PhysicalPortId s

func (v *device) PhysicalPortId() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PhysicalPortId",
	}
}

// property Mtu u

func (v *device) Mtu() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Mtu",
	}
}

// property Metered u

func (v *device) Metered() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Metered",
	}
}

// property LldpNeighbors aa{sv}

func (v *device) LldpNeighbors() PropMapSVSlice {
	return PropMapSVSlice{
		Impl: v,
		Name: "LldpNeighbors",
	}
}

// property Real b

func (v *device) Real() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Real",
	}
}

// property Ip4Connectivity u

func (v *device) Ip4Connectivity() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Ip4Connectivity",
	}
}

// property Ip6Connectivity u

func (v *device) Ip6Connectivity() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Ip6Connectivity",
	}
}

// property InterfaceFlags u

func (v *device) InterfaceFlags() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "InterfaceFlags",
	}
}

// property HwAddress s

func (v *device) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

func (obj *Device) Adsl() *deviceAdsl {
	return &obj.deviceAdsl
}

type deviceAdsl struct{}

func (v *deviceAdsl) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceAdsl) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Adsl"
}

// signal PropertiesChanged

func (v *deviceAdsl) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Carrier b

func (v *deviceAdsl) Carrier() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Carrier",
	}
}

func (obj *Device) Bluetooth() *deviceBluetooth {
	return &obj.deviceBluetooth
}

type deviceBluetooth struct{}

func (v *deviceBluetooth) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceBluetooth) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Bluetooth"
}

// signal PropertiesChanged

func (v *deviceBluetooth) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property HwAddress s

func (v *deviceBluetooth) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property Name s

func (v *deviceBluetooth) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property BtCapabilities u

func (v *deviceBluetooth) BtCapabilities() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "BtCapabilities",
	}
}

func (obj *Device) Bond() *deviceBond {
	return &obj.deviceBond
}

type deviceBond struct{}

func (v *deviceBond) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceBond) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Bond"
}

// signal PropertiesChanged

func (v *deviceBond) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property HwAddress s

func (v *deviceBond) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property Carrier b

func (v *deviceBond) Carrier() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Carrier",
	}
}

// property Slaves ao

func (v *deviceBond) Slaves() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Slaves",
	}
}

func (obj *Device) Bridge() *deviceBridge {
	return &obj.deviceBridge
}

type deviceBridge struct{}

func (v *deviceBridge) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceBridge) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Bridge"
}

// signal PropertiesChanged

func (v *deviceBridge) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property HwAddress s

func (v *deviceBridge) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property Carrier b

func (v *deviceBridge) Carrier() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Carrier",
	}
}

// property Slaves ao

func (v *deviceBridge) Slaves() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Slaves",
	}
}

func (obj *Device) Generic() *deviceGeneric {
	return &obj.deviceGeneric
}

type deviceGeneric struct{}

func (v *deviceGeneric) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceGeneric) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Generic"
}

// signal PropertiesChanged

func (v *deviceGeneric) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property HwAddress s

func (v *deviceGeneric) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property TypeDescription s

func (v *deviceGeneric) TypeDescription() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "TypeDescription",
	}
}

func (obj *Device) Infiniband() *deviceInfiniband {
	return &obj.deviceInfiniband
}

type deviceInfiniband struct{}

func (v *deviceInfiniband) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceInfiniband) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Infiniband"
}

// signal PropertiesChanged

func (v *deviceInfiniband) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property HwAddress s

func (v *deviceInfiniband) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property Carrier b

func (v *deviceInfiniband) Carrier() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Carrier",
	}
}

func (obj *Device) IPTunnel() *deviceIPTunnel {
	return &obj.deviceIPTunnel
}

type deviceIPTunnel struct{}

func (v *deviceIPTunnel) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceIPTunnel) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.IPTunnel"
}

// signal PropertiesChanged

func (v *deviceIPTunnel) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Mode u

func (v *deviceIPTunnel) Mode() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Mode",
	}
}

// property Parent o

func (v *deviceIPTunnel) Parent() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Parent",
	}
}

// property Local s

func (v *deviceIPTunnel) Local() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Local",
	}
}

// property Remote s

func (v *deviceIPTunnel) Remote() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Remote",
	}
}

// property Ttl y

func (v *deviceIPTunnel) Ttl() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "Ttl",
	}
}

// property Tos y

func (v *deviceIPTunnel) Tos() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "Tos",
	}
}

// property PathMtuDiscovery b

func (v *deviceIPTunnel) PathMtuDiscovery() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PathMtuDiscovery",
	}
}

// property InputKey s

func (v *deviceIPTunnel) InputKey() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "InputKey",
	}
}

// property OutputKey s

func (v *deviceIPTunnel) OutputKey() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "OutputKey",
	}
}

// property EncapsulationLimit y

func (v *deviceIPTunnel) EncapsulationLimit() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "EncapsulationLimit",
	}
}

// property FlowLabel u

func (v *deviceIPTunnel) FlowLabel() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "FlowLabel",
	}
}

func (obj *Device) Macsec() *deviceMacsec {
	return &obj.deviceMacsec
}

type deviceMacsec struct{}

func (v *deviceMacsec) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceMacsec) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Macsec"
}

// signal PropertiesChanged

func (v *deviceMacsec) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Parent o

func (v *deviceMacsec) Parent() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Parent",
	}
}

// property Sci t

func (v *deviceMacsec) Sci() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Sci",
	}
}

// property IcvLength y

func (v *deviceMacsec) IcvLength() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "IcvLength",
	}
}

// property CipherSuite t

func (v *deviceMacsec) CipherSuite() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "CipherSuite",
	}
}

// property Window u

func (v *deviceMacsec) Window() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Window",
	}
}

// property EncodingSa y

func (v *deviceMacsec) EncodingSa() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "EncodingSa",
	}
}

// property Validation s

func (v *deviceMacsec) Validation() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Validation",
	}
}

// property Encrypt b

func (v *deviceMacsec) Encrypt() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Encrypt",
	}
}

// property Protect b

func (v *deviceMacsec) Protect() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Protect",
	}
}

// property IncludeSci b

func (v *deviceMacsec) IncludeSci() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IncludeSci",
	}
}

// property Es b

func (v *deviceMacsec) Es() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Es",
	}
}

// property Scb b

func (v *deviceMacsec) Scb() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Scb",
	}
}

// property ReplayProtect b

func (v *deviceMacsec) ReplayProtect() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ReplayProtect",
	}
}

func (obj *Device) Macvlan() *deviceMacvlan {
	return &obj.deviceMacvlan
}

type deviceMacvlan struct{}

func (v *deviceMacvlan) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceMacvlan) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Macvlan"
}

// signal PropertiesChanged

func (v *deviceMacvlan) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Parent o

func (v *deviceMacvlan) Parent() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Parent",
	}
}

// property Mode s

func (v *deviceMacvlan) Mode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Mode",
	}
}

// property NoPromisc b

func (v *deviceMacvlan) NoPromisc() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NoPromisc",
	}
}

// property Tap b

func (v *deviceMacvlan) Tap() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Tap",
	}
}

func (obj *Device) Modem() *deviceModem {
	return &obj.deviceModem
}

type deviceModem struct{}

func (v *deviceModem) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceModem) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Modem"
}

// signal PropertiesChanged

func (v *deviceModem) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property ModemCapabilities u

func (v *deviceModem) ModemCapabilities() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "ModemCapabilities",
	}
}

// property CurrentCapabilities u

func (v *deviceModem) CurrentCapabilities() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "CurrentCapabilities",
	}
}

func (obj *Device) OlpcMesh() *deviceOlpcMesh {
	return &obj.deviceOlpcMesh
}

type deviceOlpcMesh struct{}

func (v *deviceOlpcMesh) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceOlpcMesh) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.OlpcMesh"
}

// signal PropertiesChanged

func (v *deviceOlpcMesh) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property HwAddress s

func (v *deviceOlpcMesh) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property Companion o

func (v *deviceOlpcMesh) Companion() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Companion",
	}
}

// property ActiveChannel u

func (v *deviceOlpcMesh) ActiveChannel() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "ActiveChannel",
	}
}

func (obj *Device) Statistics() *deviceStatistics {
	return &obj.deviceStatistics
}

type deviceStatistics struct{}

func (v *deviceStatistics) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceStatistics) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Statistics"
}

// signal PropertiesChanged

func (v *deviceStatistics) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property RefreshRateMs u

func (v *deviceStatistics) RefreshRateMs() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "RefreshRateMs",
	}
}

// property TxBytes t

func (v *deviceStatistics) TxBytes() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TxBytes",
	}
}

// property RxBytes t

func (v *deviceStatistics) RxBytes() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "RxBytes",
	}
}

func (obj *Device) Team() *deviceTeam {
	return &obj.deviceTeam
}

type deviceTeam struct{}

func (v *deviceTeam) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceTeam) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Team"
}

// signal PropertiesChanged

func (v *deviceTeam) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property HwAddress s

func (v *deviceTeam) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property Carrier b

func (v *deviceTeam) Carrier() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Carrier",
	}
}

// property Slaves ao

func (v *deviceTeam) Slaves() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Slaves",
	}
}

// property Config s

func (v *deviceTeam) Config() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Config",
	}
}

func (obj *Device) Tun() *deviceTun {
	return &obj.deviceTun
}

type deviceTun struct{}

func (v *deviceTun) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceTun) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Tun"
}

// signal PropertiesChanged

func (v *deviceTun) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Owner x

func (v *deviceTun) Owner() proxy.PropInt64 {
	return proxy.PropInt64{
		Impl: v,
		Name: "Owner",
	}
}

// property Group x

func (v *deviceTun) Group() proxy.PropInt64 {
	return proxy.PropInt64{
		Impl: v,
		Name: "Group",
	}
}

// property Mode s

func (v *deviceTun) Mode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Mode",
	}
}

// property NoPi b

func (v *deviceTun) NoPi() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NoPi",
	}
}

// property VnetHdr b

func (v *deviceTun) VnetHdr() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "VnetHdr",
	}
}

// property MultiQueue b

func (v *deviceTun) MultiQueue() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "MultiQueue",
	}
}

// property HwAddress s

func (v *deviceTun) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

func (obj *Device) Veth() *deviceVeth {
	return &obj.deviceVeth
}

type deviceVeth struct{}

func (v *deviceVeth) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceVeth) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Veth"
}

// signal PropertiesChanged

func (v *deviceVeth) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Peer o

func (v *deviceVeth) Peer() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Peer",
	}
}

func (obj *Device) Vlan() *deviceVlan {
	return &obj.deviceVlan
}

type deviceVlan struct{}

func (v *deviceVlan) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceVlan) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Vlan"
}

// signal PropertiesChanged

func (v *deviceVlan) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property HwAddress s

func (v *deviceVlan) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property Carrier b

func (v *deviceVlan) Carrier() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Carrier",
	}
}

// property Parent o

func (v *deviceVlan) Parent() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Parent",
	}
}

// property VlanId u

func (v *deviceVlan) VlanId() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "VlanId",
	}
}

func (obj *Device) Vxlan() *deviceVxlan {
	return &obj.deviceVxlan
}

type deviceVxlan struct{}

func (v *deviceVxlan) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceVxlan) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Vxlan"
}

// signal PropertiesChanged

func (v *deviceVxlan) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Parent o

func (v *deviceVxlan) Parent() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Parent",
	}
}

// property HwAddress s

func (v *deviceVxlan) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property Id u

func (v *deviceVxlan) Id() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Id",
	}
}

// property Group s

func (v *deviceVxlan) Group() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Group",
	}
}

// property Local s

func (v *deviceVxlan) Local() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Local",
	}
}

// property Tos y

func (v *deviceVxlan) Tos() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "Tos",
	}
}

// property Ttl y

func (v *deviceVxlan) Ttl() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "Ttl",
	}
}

// property Learning b

func (v *deviceVxlan) Learning() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Learning",
	}
}

// property Ageing u

func (v *deviceVxlan) Ageing() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Ageing",
	}
}

// property Limit u

func (v *deviceVxlan) Limit() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Limit",
	}
}

// property DstPort q

func (v *deviceVxlan) DstPort() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "DstPort",
	}
}

// property SrcPortMin q

func (v *deviceVxlan) SrcPortMin() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "SrcPortMin",
	}
}

// property SrcPortMax q

func (v *deviceVxlan) SrcPortMax() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "SrcPortMax",
	}
}

// property Proxy b

func (v *deviceVxlan) Proxy() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Proxy",
	}
}

// property Rsc b

func (v *deviceVxlan) Rsc() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Rsc",
	}
}

// property L2miss b

func (v *deviceVxlan) L2miss() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "L2miss",
	}
}

// property L3miss b

func (v *deviceVxlan) L3miss() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "L3miss",
	}
}

func (obj *Device) WiMax() *deviceWiMax {
	return &obj.deviceWiMax
}

type deviceWiMax struct{}

func (v *deviceWiMax) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceWiMax) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.WiMax"
}

// method GetNspList

func (v *deviceWiMax) GoGetNspList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetNspList", flags, ch)
}

func (*deviceWiMax) StoreGetNspList(call *dbus.Call) (nsps []dbus.ObjectPath, err error) {
	err = call.Store(&nsps)
	return
}

func (v *deviceWiMax) GetNspList(flags dbus.Flags) (nsps []dbus.ObjectPath, err error) {
	return v.StoreGetNspList(
		<-v.GoGetNspList(flags, make(chan *dbus.Call, 1)).Done)
}

// signal PropertiesChanged

func (v *deviceWiMax) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NspAdded

func (v *deviceWiMax) ConnectNspAdded(cb func(nsp dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NspAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NspAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var nsp dbus.ObjectPath
		err := dbus.Store(sig.Body, &nsp)
		if err == nil {
			cb(nsp)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NspRemoved

func (v *deviceWiMax) ConnectNspRemoved(cb func(nsp dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NspRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NspRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var nsp dbus.ObjectPath
		err := dbus.Store(sig.Body, &nsp)
		if err == nil {
			cb(nsp)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Nsps ao

func (v *deviceWiMax) Nsps() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Nsps",
	}
}

// property HwAddress s

func (v *deviceWiMax) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property CenterFrequency u

func (v *deviceWiMax) CenterFrequency() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "CenterFrequency",
	}
}

// property Rssi i

func (v *deviceWiMax) Rssi() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "Rssi",
	}
}

// property Cinr i

func (v *deviceWiMax) Cinr() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "Cinr",
	}
}

// property TxPower i

func (v *deviceWiMax) TxPower() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "TxPower",
	}
}

// property Bsid s

func (v *deviceWiMax) Bsid() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Bsid",
	}
}

// property ActiveNsp o

func (v *deviceWiMax) ActiveNsp() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "ActiveNsp",
	}
}

func (obj *Device) Wired() *deviceWired {
	return &obj.deviceWired
}

type deviceWired struct{}

func (v *deviceWired) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceWired) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Wired"
}

// signal PropertiesChanged

func (v *deviceWired) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property HwAddress s

func (v *deviceWired) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property PermHwAddress s

func (v *deviceWired) PermHwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PermHwAddress",
	}
}

// property Speed u

func (v *deviceWired) Speed() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Speed",
	}
}

// property S390Subchannels as

func (v *deviceWired) S390Subchannels() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "S390Subchannels",
	}
}

// property Carrier b

func (v *deviceWired) Carrier() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Carrier",
	}
}

func (obj *Device) Wireless() *deviceWireless {
	return &obj.deviceWireless
}

type deviceWireless struct{}

func (v *deviceWireless) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*deviceWireless) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Device.Wireless"
}

// method GetAccessPoints

func (v *deviceWireless) GoGetAccessPoints(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAccessPoints", flags, ch)
}

func (*deviceWireless) StoreGetAccessPoints(call *dbus.Call) (access_points []dbus.ObjectPath, err error) {
	err = call.Store(&access_points)
	return
}

func (v *deviceWireless) GetAccessPoints(flags dbus.Flags) (access_points []dbus.ObjectPath, err error) {
	return v.StoreGetAccessPoints(
		<-v.GoGetAccessPoints(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAllAccessPoints

func (v *deviceWireless) GoGetAllAccessPoints(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllAccessPoints", flags, ch)
}

func (*deviceWireless) StoreGetAllAccessPoints(call *dbus.Call) (access_points []dbus.ObjectPath, err error) {
	err = call.Store(&access_points)
	return
}

func (v *deviceWireless) GetAllAccessPoints(flags dbus.Flags) (access_points []dbus.ObjectPath, err error) {
	return v.StoreGetAllAccessPoints(
		<-v.GoGetAllAccessPoints(flags, make(chan *dbus.Call, 1)).Done)
}

// method RequestScan

func (v *deviceWireless) GoRequestScan(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestScan", flags, ch, options)
}

func (v *deviceWireless) RequestScan(flags dbus.Flags, options map[string]dbus.Variant) error {
	return (<-v.GoRequestScan(flags, make(chan *dbus.Call, 1), options).Done).Err
}

// signal PropertiesChanged

func (v *deviceWireless) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AccessPointAdded

func (v *deviceWireless) ConnectAccessPointAdded(cb func(access_point dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AccessPointAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AccessPointAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var access_point dbus.ObjectPath
		err := dbus.Store(sig.Body, &access_point)
		if err == nil {
			cb(access_point)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AccessPointRemoved

func (v *deviceWireless) ConnectAccessPointRemoved(cb func(access_point dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AccessPointRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AccessPointRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var access_point dbus.ObjectPath
		err := dbus.Store(sig.Body, &access_point)
		if err == nil {
			cb(access_point)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property HwAddress s

func (v *deviceWireless) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property PermHwAddress s

func (v *deviceWireless) PermHwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PermHwAddress",
	}
}

// property Mode u

func (v *deviceWireless) Mode() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Mode",
	}
}

// property Bitrate u

func (v *deviceWireless) Bitrate() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Bitrate",
	}
}

// property AccessPoints ao

func (v *deviceWireless) AccessPoints() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "AccessPoints",
	}
}

// property ActiveAccessPoint o

func (v *deviceWireless) ActiveAccessPoint() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "ActiveAccessPoint",
	}
}

// property WirelessCapabilities u

func (v *deviceWireless) WirelessCapabilities() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "WirelessCapabilities",
	}
}

type WiMaxNsp struct {
	wiMaxNsp // interface org.freedesktop.NetworkManager.WiMax.Nsp
	proxy.Object
}

func NewWiMaxNsp(conn *dbus.Conn, path dbus.ObjectPath) (*WiMaxNsp, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(WiMaxNsp)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type wiMaxNsp struct{}

func (v *wiMaxNsp) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*wiMaxNsp) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.WiMax.Nsp"
}

// signal PropertiesChanged

func (v *wiMaxNsp) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Name s

func (v *wiMaxNsp) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property SignalQuality u

func (v *wiMaxNsp) SignalQuality() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "SignalQuality",
	}
}

// property NetworkType u

func (v *wiMaxNsp) NetworkType() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "NetworkType",
	}
}

type PropMapStringVariant struct {
	Impl proxy.Implementer
	Name string
}

func (p PropMapStringVariant) Get(flags dbus.Flags) (value map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropMapStringVariant) Set(flags dbus.Flags, value map[string]dbus.Variant) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropMapStringVariant) ConnectChanged(cb func(hasValue bool, value map[string]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]dbus.Variant
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

type PropMapSVSlice struct {
	Impl proxy.Implementer
	Name string
}

func (p PropMapSVSlice) Get(flags dbus.Flags) (value []map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropMapSVSlice) Set(flags dbus.Flags, value []map[string]dbus.Variant) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropMapSVSlice) ConnectChanged(cb func(hasValue bool, value []map[string]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []map[string]dbus.Variant
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

type PropUint32SliceSlice struct {
	Impl proxy.Implementer
	Name string
}

func (p PropUint32SliceSlice) Get(flags dbus.Flags) (value [][]uint32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropUint32SliceSlice) Set(flags dbus.Flags, value [][]uint32) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropUint32SliceSlice) ConnectChanged(cb func(hasValue bool, value [][]uint32)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v [][]uint32
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
