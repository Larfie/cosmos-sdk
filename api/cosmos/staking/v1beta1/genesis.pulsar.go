package stakingv1beta1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_GenesisState_3_list)(nil)

type _GenesisState_3_list struct {
	list *[]*LastValidatorPower
}

func (x *_GenesisState_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*LastValidatorPower)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*LastValidatorPower)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_3_list) AppendMutable() protoreflect.Value {
	v := new(LastValidatorPower)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_3_list) NewElement() protoreflect.Value {
	v := new(LastValidatorPower)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_4_list)(nil)

type _GenesisState_4_list struct {
	list *[]*Validator
}

func (x *_GenesisState_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_4_list) AppendMutable() protoreflect.Value {
	v := new(Validator)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_4_list) NewElement() protoreflect.Value {
	v := new(Validator)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_4_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_5_list)(nil)

type _GenesisState_5_list struct {
	list *[]*Delegation
}

func (x *_GenesisState_5_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_5_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_5_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Delegation)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_5_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Delegation)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_5_list) AppendMutable() protoreflect.Value {
	v := new(Delegation)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_5_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_5_list) NewElement() protoreflect.Value {
	v := new(Delegation)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_5_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_6_list)(nil)

type _GenesisState_6_list struct {
	list *[]*UnbondingDelegation
}

func (x *_GenesisState_6_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_6_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_6_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*UnbondingDelegation)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_6_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*UnbondingDelegation)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_6_list) AppendMutable() protoreflect.Value {
	v := new(UnbondingDelegation)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_6_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_6_list) NewElement() protoreflect.Value {
	v := new(UnbondingDelegation)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_6_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_7_list)(nil)

type _GenesisState_7_list struct {
	list *[]*Redelegation
}

func (x *_GenesisState_7_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_7_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_7_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Redelegation)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_7_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Redelegation)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_7_list) AppendMutable() protoreflect.Value {
	v := new(Redelegation)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_7_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_7_list) NewElement() protoreflect.Value {
	v := new(Redelegation)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_7_list) IsValid() bool {
	return x.list != nil
}

var (
	md_GenesisState                       protoreflect.MessageDescriptor
	fd_GenesisState_params                protoreflect.FieldDescriptor
	fd_GenesisState_last_total_power      protoreflect.FieldDescriptor
	fd_GenesisState_last_validator_powers protoreflect.FieldDescriptor
	fd_GenesisState_validators            protoreflect.FieldDescriptor
	fd_GenesisState_delegations           protoreflect.FieldDescriptor
	fd_GenesisState_unbonding_delegations protoreflect.FieldDescriptor
	fd_GenesisState_redelegations         protoreflect.FieldDescriptor
	fd_GenesisState_exported              protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_genesis_proto_init()
	md_GenesisState = File_cosmos_staking_v1beta1_genesis_proto.Messages().ByName("GenesisState")
	fd_GenesisState_params = md_GenesisState.Fields().ByName("params")
	fd_GenesisState_last_total_power = md_GenesisState.Fields().ByName("last_total_power")
	fd_GenesisState_last_validator_powers = md_GenesisState.Fields().ByName("last_validator_powers")
	fd_GenesisState_validators = md_GenesisState.Fields().ByName("validators")
	fd_GenesisState_delegations = md_GenesisState.Fields().ByName("delegations")
	fd_GenesisState_unbonding_delegations = md_GenesisState.Fields().ByName("unbonding_delegations")
	fd_GenesisState_redelegations = md_GenesisState.Fields().ByName("redelegations")
	fd_GenesisState_exported = md_GenesisState.Fields().ByName("exported")
}

var _ protoreflect.Message = (*fastReflection_GenesisState)(nil)

type fastReflection_GenesisState GenesisState

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GenesisState)(x)
}

func (x *GenesisState) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GenesisState_messageType fastReflection_GenesisState_messageType
var _ protoreflect.MessageType = fastReflection_GenesisState_messageType{}

type fastReflection_GenesisState_messageType struct{}

func (x fastReflection_GenesisState_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GenesisState)(nil)
}
func (x fastReflection_GenesisState_messageType) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}
func (x fastReflection_GenesisState_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GenesisState) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GenesisState) Type() protoreflect.MessageType {
	return _fastReflection_GenesisState_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GenesisState) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GenesisState) Interface() protoreflect.ProtoMessage {
	return (*GenesisState)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GenesisState) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Params != nil {
		value := protoreflect.ValueOfMessage(x.Params.ProtoReflect())
		if !f(fd_GenesisState_params, value) {
			return
		}
	}
	if len(x.LastTotalPower) != 0 {
		value := protoreflect.ValueOfBytes(x.LastTotalPower)
		if !f(fd_GenesisState_last_total_power, value) {
			return
		}
	}
	if len(x.LastValidatorPowers) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_3_list{list: &x.LastValidatorPowers})
		if !f(fd_GenesisState_last_validator_powers, value) {
			return
		}
	}
	if len(x.Validators) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_4_list{list: &x.Validators})
		if !f(fd_GenesisState_validators, value) {
			return
		}
	}
	if len(x.Delegations) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_5_list{list: &x.Delegations})
		if !f(fd_GenesisState_delegations, value) {
			return
		}
	}
	if len(x.UnbondingDelegations) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_6_list{list: &x.UnbondingDelegations})
		if !f(fd_GenesisState_unbonding_delegations, value) {
			return
		}
	}
	if len(x.Redelegations) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_7_list{list: &x.Redelegations})
		if !f(fd_GenesisState_redelegations, value) {
			return
		}
	}
	if x.Exported != false {
		value := protoreflect.ValueOfBool(x.Exported)
		if !f(fd_GenesisState_exported, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GenesisState) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.GenesisState.params":
		return x.Params != nil
	case "cosmos.staking.v1beta1.GenesisState.last_total_power":
		return len(x.LastTotalPower) != 0
	case "cosmos.staking.v1beta1.GenesisState.last_validator_powers":
		return len(x.LastValidatorPowers) != 0
	case "cosmos.staking.v1beta1.GenesisState.validators":
		return len(x.Validators) != 0
	case "cosmos.staking.v1beta1.GenesisState.delegations":
		return len(x.Delegations) != 0
	case "cosmos.staking.v1beta1.GenesisState.unbonding_delegations":
		return len(x.UnbondingDelegations) != 0
	case "cosmos.staking.v1beta1.GenesisState.redelegations":
		return len(x.Redelegations) != 0
	case "cosmos.staking.v1beta1.GenesisState.exported":
		return x.Exported != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.GenesisState.params":
		x.Params = nil
	case "cosmos.staking.v1beta1.GenesisState.last_total_power":
		x.LastTotalPower = nil
	case "cosmos.staking.v1beta1.GenesisState.last_validator_powers":
		x.LastValidatorPowers = nil
	case "cosmos.staking.v1beta1.GenesisState.validators":
		x.Validators = nil
	case "cosmos.staking.v1beta1.GenesisState.delegations":
		x.Delegations = nil
	case "cosmos.staking.v1beta1.GenesisState.unbonding_delegations":
		x.UnbondingDelegations = nil
	case "cosmos.staking.v1beta1.GenesisState.redelegations":
		x.Redelegations = nil
	case "cosmos.staking.v1beta1.GenesisState.exported":
		x.Exported = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GenesisState) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.GenesisState.params":
		value := x.Params
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.GenesisState.last_total_power":
		value := x.LastTotalPower
		return protoreflect.ValueOfBytes(value)
	case "cosmos.staking.v1beta1.GenesisState.last_validator_powers":
		if len(x.LastValidatorPowers) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_3_list{})
		}
		listValue := &_GenesisState_3_list{list: &x.LastValidatorPowers}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.GenesisState.validators":
		if len(x.Validators) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_4_list{})
		}
		listValue := &_GenesisState_4_list{list: &x.Validators}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.GenesisState.delegations":
		if len(x.Delegations) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_5_list{})
		}
		listValue := &_GenesisState_5_list{list: &x.Delegations}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.GenesisState.unbonding_delegations":
		if len(x.UnbondingDelegations) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_6_list{})
		}
		listValue := &_GenesisState_6_list{list: &x.UnbondingDelegations}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.GenesisState.redelegations":
		if len(x.Redelegations) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_7_list{})
		}
		listValue := &_GenesisState_7_list{list: &x.Redelegations}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.GenesisState.exported":
		value := x.Exported
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.GenesisState does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.GenesisState.params":
		x.Params = value.Message().Interface().(*Params)
	case "cosmos.staking.v1beta1.GenesisState.last_total_power":
		x.LastTotalPower = value.Bytes()
	case "cosmos.staking.v1beta1.GenesisState.last_validator_powers":
		lv := value.List()
		clv := lv.(*_GenesisState_3_list)
		x.LastValidatorPowers = *clv.list
	case "cosmos.staking.v1beta1.GenesisState.validators":
		lv := value.List()
		clv := lv.(*_GenesisState_4_list)
		x.Validators = *clv.list
	case "cosmos.staking.v1beta1.GenesisState.delegations":
		lv := value.List()
		clv := lv.(*_GenesisState_5_list)
		x.Delegations = *clv.list
	case "cosmos.staking.v1beta1.GenesisState.unbonding_delegations":
		lv := value.List()
		clv := lv.(*_GenesisState_6_list)
		x.UnbondingDelegations = *clv.list
	case "cosmos.staking.v1beta1.GenesisState.redelegations":
		lv := value.List()
		clv := lv.(*_GenesisState_7_list)
		x.Redelegations = *clv.list
	case "cosmos.staking.v1beta1.GenesisState.exported":
		x.Exported = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.GenesisState.params":
		if x.Params == nil {
			x.Params = new(Params)
		}
		return protoreflect.ValueOfMessage(x.Params.ProtoReflect())
	case "cosmos.staking.v1beta1.GenesisState.last_validator_powers":
		if x.LastValidatorPowers == nil {
			x.LastValidatorPowers = []*LastValidatorPower{}
		}
		value := &_GenesisState_3_list{list: &x.LastValidatorPowers}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.GenesisState.validators":
		if x.Validators == nil {
			x.Validators = []*Validator{}
		}
		value := &_GenesisState_4_list{list: &x.Validators}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.GenesisState.delegations":
		if x.Delegations == nil {
			x.Delegations = []*Delegation{}
		}
		value := &_GenesisState_5_list{list: &x.Delegations}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.GenesisState.unbonding_delegations":
		if x.UnbondingDelegations == nil {
			x.UnbondingDelegations = []*UnbondingDelegation{}
		}
		value := &_GenesisState_6_list{list: &x.UnbondingDelegations}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.GenesisState.redelegations":
		if x.Redelegations == nil {
			x.Redelegations = []*Redelegation{}
		}
		value := &_GenesisState_7_list{list: &x.Redelegations}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.GenesisState.last_total_power":
		panic(fmt.Errorf("field last_total_power of message cosmos.staking.v1beta1.GenesisState is not mutable"))
	case "cosmos.staking.v1beta1.GenesisState.exported":
		panic(fmt.Errorf("field exported of message cosmos.staking.v1beta1.GenesisState is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GenesisState) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.GenesisState.params":
		m := new(Params)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.GenesisState.last_total_power":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.staking.v1beta1.GenesisState.last_validator_powers":
		list := []*LastValidatorPower{}
		return protoreflect.ValueOfList(&_GenesisState_3_list{list: &list})
	case "cosmos.staking.v1beta1.GenesisState.validators":
		list := []*Validator{}
		return protoreflect.ValueOfList(&_GenesisState_4_list{list: &list})
	case "cosmos.staking.v1beta1.GenesisState.delegations":
		list := []*Delegation{}
		return protoreflect.ValueOfList(&_GenesisState_5_list{list: &list})
	case "cosmos.staking.v1beta1.GenesisState.unbonding_delegations":
		list := []*UnbondingDelegation{}
		return protoreflect.ValueOfList(&_GenesisState_6_list{list: &list})
	case "cosmos.staking.v1beta1.GenesisState.redelegations":
		list := []*Redelegation{}
		return protoreflect.ValueOfList(&_GenesisState_7_list{list: &list})
	case "cosmos.staking.v1beta1.GenesisState.exported":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GenesisState) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.GenesisState", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GenesisState) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GenesisState) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GenesisState) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Params != nil {
			l = options.Size(x.Params)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.LastTotalPower)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.LastValidatorPowers) > 0 {
			for _, e := range x.LastValidatorPowers {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.Validators) > 0 {
			for _, e := range x.Validators {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.Delegations) > 0 {
			for _, e := range x.Delegations {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.UnbondingDelegations) > 0 {
			for _, e := range x.UnbondingDelegations {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.Redelegations) > 0 {
			for _, e := range x.Redelegations {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Exported {
			n += 2
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Exported {
			i--
			if x.Exported {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x40
		}
		if len(x.Redelegations) > 0 {
			for iNdEx := len(x.Redelegations) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Redelegations[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x3a
			}
		}
		if len(x.UnbondingDelegations) > 0 {
			for iNdEx := len(x.UnbondingDelegations) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.UnbondingDelegations[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x32
			}
		}
		if len(x.Delegations) > 0 {
			for iNdEx := len(x.Delegations) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Delegations[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x2a
			}
		}
		if len(x.Validators) > 0 {
			for iNdEx := len(x.Validators) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Validators[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x22
			}
		}
		if len(x.LastValidatorPowers) > 0 {
			for iNdEx := len(x.LastValidatorPowers) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.LastValidatorPowers[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1a
			}
		}
		if len(x.LastTotalPower) > 0 {
			i -= len(x.LastTotalPower)
			copy(dAtA[i:], x.LastTotalPower)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.LastTotalPower)))
			i--
			dAtA[i] = 0x12
		}
		if x.Params != nil {
			encoded, err := options.Marshal(x.Params)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Params == nil {
					x.Params = &Params{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Params); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastTotalPower", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.LastTotalPower = append(x.LastTotalPower[:0], dAtA[iNdEx:postIndex]...)
				if x.LastTotalPower == nil {
					x.LastTotalPower = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastValidatorPowers", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.LastValidatorPowers = append(x.LastValidatorPowers, &LastValidatorPower{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.LastValidatorPowers[len(x.LastValidatorPowers)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Validators = append(x.Validators, &Validator{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Validators[len(x.Validators)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Delegations", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Delegations = append(x.Delegations, &Delegation{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Delegations[len(x.Delegations)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UnbondingDelegations", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.UnbondingDelegations = append(x.UnbondingDelegations, &UnbondingDelegation{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.UnbondingDelegations[len(x.UnbondingDelegations)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Redelegations", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Redelegations = append(x.Redelegations, &Redelegation{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Redelegations[len(x.Redelegations)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 8:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Exported", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Exported = bool(v != 0)
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_LastValidatorPower         protoreflect.MessageDescriptor
	fd_LastValidatorPower_address protoreflect.FieldDescriptor
	fd_LastValidatorPower_power   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_genesis_proto_init()
	md_LastValidatorPower = File_cosmos_staking_v1beta1_genesis_proto.Messages().ByName("LastValidatorPower")
	fd_LastValidatorPower_address = md_LastValidatorPower.Fields().ByName("address")
	fd_LastValidatorPower_power = md_LastValidatorPower.Fields().ByName("power")
}

var _ protoreflect.Message = (*fastReflection_LastValidatorPower)(nil)

type fastReflection_LastValidatorPower LastValidatorPower

func (x *LastValidatorPower) ProtoReflect() protoreflect.Message {
	return (*fastReflection_LastValidatorPower)(x)
}

func (x *LastValidatorPower) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_genesis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_LastValidatorPower_messageType fastReflection_LastValidatorPower_messageType
var _ protoreflect.MessageType = fastReflection_LastValidatorPower_messageType{}

type fastReflection_LastValidatorPower_messageType struct{}

func (x fastReflection_LastValidatorPower_messageType) Zero() protoreflect.Message {
	return (*fastReflection_LastValidatorPower)(nil)
}
func (x fastReflection_LastValidatorPower_messageType) New() protoreflect.Message {
	return new(fastReflection_LastValidatorPower)
}
func (x fastReflection_LastValidatorPower_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_LastValidatorPower
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_LastValidatorPower) Descriptor() protoreflect.MessageDescriptor {
	return md_LastValidatorPower
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_LastValidatorPower) Type() protoreflect.MessageType {
	return _fastReflection_LastValidatorPower_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_LastValidatorPower) New() protoreflect.Message {
	return new(fastReflection_LastValidatorPower)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_LastValidatorPower) Interface() protoreflect.ProtoMessage {
	return (*LastValidatorPower)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_LastValidatorPower) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_LastValidatorPower_address, value) {
			return
		}
	}
	if x.Power != int64(0) {
		value := protoreflect.ValueOfInt64(x.Power)
		if !f(fd_LastValidatorPower_power, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_LastValidatorPower) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.LastValidatorPower.address":
		return x.Address != ""
	case "cosmos.staking.v1beta1.LastValidatorPower.power":
		return x.Power != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.LastValidatorPower"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.LastValidatorPower does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LastValidatorPower) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.LastValidatorPower.address":
		x.Address = ""
	case "cosmos.staking.v1beta1.LastValidatorPower.power":
		x.Power = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.LastValidatorPower"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.LastValidatorPower does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_LastValidatorPower) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.LastValidatorPower.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.LastValidatorPower.power":
		value := x.Power
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.LastValidatorPower"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.LastValidatorPower does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LastValidatorPower) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.LastValidatorPower.address":
		x.Address = value.Interface().(string)
	case "cosmos.staking.v1beta1.LastValidatorPower.power":
		x.Power = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.LastValidatorPower"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.LastValidatorPower does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LastValidatorPower) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.LastValidatorPower.address":
		panic(fmt.Errorf("field address of message cosmos.staking.v1beta1.LastValidatorPower is not mutable"))
	case "cosmos.staking.v1beta1.LastValidatorPower.power":
		panic(fmt.Errorf("field power of message cosmos.staking.v1beta1.LastValidatorPower is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.LastValidatorPower"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.LastValidatorPower does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_LastValidatorPower) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.LastValidatorPower.address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.LastValidatorPower.power":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.LastValidatorPower"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.LastValidatorPower does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_LastValidatorPower) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.LastValidatorPower", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_LastValidatorPower) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LastValidatorPower) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_LastValidatorPower) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_LastValidatorPower) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*LastValidatorPower)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Power != 0 {
			n += 1 + runtime.Sov(uint64(x.Power))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*LastValidatorPower)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Power != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Power))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*LastValidatorPower)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LastValidatorPower: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LastValidatorPower: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
				}
				x.Power = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Power |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/staking/v1beta1/genesis.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GenesisState defines the staking module's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params defines all the paramaters of related to deposit.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	// last_total_power tracks the total amounts of bonded tokens recorded during
	// the previous end block.
	LastTotalPower []byte `protobuf:"bytes,2,opt,name=last_total_power,json=lastTotalPower,proto3" json:"last_total_power,omitempty"`
	// last_validator_powers is a special index that provides a historical list
	// of the last-block's bonded validators.
	LastValidatorPowers []*LastValidatorPower `protobuf:"bytes,3,rep,name=last_validator_powers,json=lastValidatorPowers,proto3" json:"last_validator_powers,omitempty"`
	// delegations defines the validator set at genesis.
	Validators []*Validator `protobuf:"bytes,4,rep,name=validators,proto3" json:"validators,omitempty"`
	// delegations defines the delegations active at genesis.
	Delegations []*Delegation `protobuf:"bytes,5,rep,name=delegations,proto3" json:"delegations,omitempty"`
	// unbonding_delegations defines the unbonding delegations active at genesis.
	UnbondingDelegations []*UnbondingDelegation `protobuf:"bytes,6,rep,name=unbonding_delegations,json=unbondingDelegations,proto3" json:"unbonding_delegations,omitempty"`
	// redelegations defines the redelegations active at genesis.
	Redelegations []*Redelegation `protobuf:"bytes,7,rep,name=redelegations,proto3" json:"redelegations,omitempty"`
	Exported      bool            `protobuf:"varint,8,opt,name=exported,proto3" json:"exported,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetLastTotalPower() []byte {
	if x != nil {
		return x.LastTotalPower
	}
	return nil
}

func (x *GenesisState) GetLastValidatorPowers() []*LastValidatorPower {
	if x != nil {
		return x.LastValidatorPowers
	}
	return nil
}

func (x *GenesisState) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *GenesisState) GetDelegations() []*Delegation {
	if x != nil {
		return x.Delegations
	}
	return nil
}

func (x *GenesisState) GetUnbondingDelegations() []*UnbondingDelegation {
	if x != nil {
		return x.UnbondingDelegations
	}
	return nil
}

func (x *GenesisState) GetRedelegations() []*Redelegation {
	if x != nil {
		return x.Redelegations
	}
	return nil
}

func (x *GenesisState) GetExported() bool {
	if x != nil {
		return x.Exported
	}
	return false
}

// LastValidatorPower required for validator set update logic.
type LastValidatorPower struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the address of the validator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// power defines the power of the validator.
	Power int64 `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (x *LastValidatorPower) Reset() {
	*x = LastValidatorPower{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_genesis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastValidatorPower) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastValidatorPower) ProtoMessage() {}

// Deprecated: Use LastValidatorPower.ProtoReflect.Descriptor instead.
func (*LastValidatorPower) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_genesis_proto_rawDescGZIP(), []int{1}
}

func (x *LastValidatorPower) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *LastValidatorPower) GetPower() int64 {
	if x != nil {
		return x.Power
	}
	return 0
}

var File_cosmos_staking_v1beta1_genesis_proto protoreflect.FileDescriptor

var file_cosmos_staking_v1beta1_genesis_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x04, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x58, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x2e,
	0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x0e,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x64,
	0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52,
	0x13, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f,
	0x00, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x4a, 0x0a,
	0x0b, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0b, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x66, 0x0a, 0x15, 0x75, 0x6e, 0x62,
	0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x14, 0x75, 0x6e, 0x62,
	0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x50, 0x0a, 0x0d, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04,
	0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0d, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x22,
	0x68, 0x0a, 0x12, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x3a,
	0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x42, 0xec, 0x01, 0x0a, 0x1a, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x3b, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca,
	0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x22, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_staking_v1beta1_genesis_proto_rawDescOnce sync.Once
	file_cosmos_staking_v1beta1_genesis_proto_rawDescData = file_cosmos_staking_v1beta1_genesis_proto_rawDesc
)

func file_cosmos_staking_v1beta1_genesis_proto_rawDescGZIP() []byte {
	file_cosmos_staking_v1beta1_genesis_proto_rawDescOnce.Do(func() {
		file_cosmos_staking_v1beta1_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_staking_v1beta1_genesis_proto_rawDescData)
	})
	return file_cosmos_staking_v1beta1_genesis_proto_rawDescData
}

var file_cosmos_staking_v1beta1_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_staking_v1beta1_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil),        // 0: cosmos.staking.v1beta1.GenesisState
	(*LastValidatorPower)(nil),  // 1: cosmos.staking.v1beta1.LastValidatorPower
	(*Params)(nil),              // 2: cosmos.staking.v1beta1.Params
	(*Validator)(nil),           // 3: cosmos.staking.v1beta1.Validator
	(*Delegation)(nil),          // 4: cosmos.staking.v1beta1.Delegation
	(*UnbondingDelegation)(nil), // 5: cosmos.staking.v1beta1.UnbondingDelegation
	(*Redelegation)(nil),        // 6: cosmos.staking.v1beta1.Redelegation
}
var file_cosmos_staking_v1beta1_genesis_proto_depIdxs = []int32{
	2, // 0: cosmos.staking.v1beta1.GenesisState.params:type_name -> cosmos.staking.v1beta1.Params
	1, // 1: cosmos.staking.v1beta1.GenesisState.last_validator_powers:type_name -> cosmos.staking.v1beta1.LastValidatorPower
	3, // 2: cosmos.staking.v1beta1.GenesisState.validators:type_name -> cosmos.staking.v1beta1.Validator
	4, // 3: cosmos.staking.v1beta1.GenesisState.delegations:type_name -> cosmos.staking.v1beta1.Delegation
	5, // 4: cosmos.staking.v1beta1.GenesisState.unbonding_delegations:type_name -> cosmos.staking.v1beta1.UnbondingDelegation
	6, // 5: cosmos.staking.v1beta1.GenesisState.redelegations:type_name -> cosmos.staking.v1beta1.Redelegation
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cosmos_staking_v1beta1_genesis_proto_init() }
func file_cosmos_staking_v1beta1_genesis_proto_init() {
	if File_cosmos_staking_v1beta1_genesis_proto != nil {
		return
	}
	file_cosmos_staking_v1beta1_staking_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_staking_v1beta1_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_genesis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastValidatorPower); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_staking_v1beta1_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_staking_v1beta1_genesis_proto_goTypes,
		DependencyIndexes: file_cosmos_staking_v1beta1_genesis_proto_depIdxs,
		MessageInfos:      file_cosmos_staking_v1beta1_genesis_proto_msgTypes,
	}.Build()
	File_cosmos_staking_v1beta1_genesis_proto = out.File
	file_cosmos_staking_v1beta1_genesis_proto_rawDesc = nil
	file_cosmos_staking_v1beta1_genesis_proto_goTypes = nil
	file_cosmos_staking_v1beta1_genesis_proto_depIdxs = nil
}