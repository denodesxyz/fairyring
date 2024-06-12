// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package pep

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_PepPacketData        protoreflect.MessageDescriptor
	fd_PepPacketData_noData protoreflect.FieldDescriptor
)

func init() {
	file_fairyring_pep_packet_proto_init()
	md_PepPacketData = File_fairyring_pep_packet_proto.Messages().ByName("PepPacketData")
	fd_PepPacketData_noData = md_PepPacketData.Fields().ByName("noData")
}

var _ protoreflect.Message = (*fastReflection_PepPacketData)(nil)

type fastReflection_PepPacketData PepPacketData

func (x *PepPacketData) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PepPacketData)(x)
}

func (x *PepPacketData) slowProtoReflect() protoreflect.Message {
	mi := &file_fairyring_pep_packet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PepPacketData_messageType fastReflection_PepPacketData_messageType
var _ protoreflect.MessageType = fastReflection_PepPacketData_messageType{}

type fastReflection_PepPacketData_messageType struct{}

func (x fastReflection_PepPacketData_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PepPacketData)(nil)
}
func (x fastReflection_PepPacketData_messageType) New() protoreflect.Message {
	return new(fastReflection_PepPacketData)
}
func (x fastReflection_PepPacketData_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PepPacketData
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PepPacketData) Descriptor() protoreflect.MessageDescriptor {
	return md_PepPacketData
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PepPacketData) Type() protoreflect.MessageType {
	return _fastReflection_PepPacketData_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PepPacketData) New() protoreflect.Message {
	return new(fastReflection_PepPacketData)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PepPacketData) Interface() protoreflect.ProtoMessage {
	return (*PepPacketData)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PepPacketData) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Packet != nil {
		switch o := x.Packet.(type) {
		case *PepPacketData_NoData:
			v := o.NoData
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_PepPacketData_noData, value) {
				return
			}
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
func (x *fastReflection_PepPacketData) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "fairyring.pep.PepPacketData.noData":
		if x.Packet == nil {
			return false
		} else if _, ok := x.Packet.(*PepPacketData_NoData); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.PepPacketData"))
		}
		panic(fmt.Errorf("message fairyring.pep.PepPacketData does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PepPacketData) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "fairyring.pep.PepPacketData.noData":
		x.Packet = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.PepPacketData"))
		}
		panic(fmt.Errorf("message fairyring.pep.PepPacketData does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PepPacketData) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "fairyring.pep.PepPacketData.noData":
		if x.Packet == nil {
			return protoreflect.ValueOfMessage((*NoData)(nil).ProtoReflect())
		} else if v, ok := x.Packet.(*PepPacketData_NoData); ok {
			return protoreflect.ValueOfMessage(v.NoData.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*NoData)(nil).ProtoReflect())
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.PepPacketData"))
		}
		panic(fmt.Errorf("message fairyring.pep.PepPacketData does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_PepPacketData) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "fairyring.pep.PepPacketData.noData":
		cv := value.Message().Interface().(*NoData)
		x.Packet = &PepPacketData_NoData{NoData: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.PepPacketData"))
		}
		panic(fmt.Errorf("message fairyring.pep.PepPacketData does not contain field %s", fd.FullName()))
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
func (x *fastReflection_PepPacketData) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "fairyring.pep.PepPacketData.noData":
		if x.Packet == nil {
			value := &NoData{}
			oneofValue := &PepPacketData_NoData{NoData: value}
			x.Packet = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Packet.(type) {
		case *PepPacketData_NoData:
			return protoreflect.ValueOfMessage(m.NoData.ProtoReflect())
		default:
			value := &NoData{}
			oneofValue := &PepPacketData_NoData{NoData: value}
			x.Packet = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.PepPacketData"))
		}
		panic(fmt.Errorf("message fairyring.pep.PepPacketData does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PepPacketData) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "fairyring.pep.PepPacketData.noData":
		value := &NoData{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.PepPacketData"))
		}
		panic(fmt.Errorf("message fairyring.pep.PepPacketData does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PepPacketData) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "fairyring.pep.PepPacketData.packet":
		if x.Packet == nil {
			return nil
		}
		switch x.Packet.(type) {
		case *PepPacketData_NoData:
			return x.Descriptor().Fields().ByName("noData")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in fairyring.pep.PepPacketData", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PepPacketData) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PepPacketData) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_PepPacketData) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PepPacketData) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PepPacketData)
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
		switch x := x.Packet.(type) {
		case *PepPacketData_NoData:
			if x == nil {
				break
			}
			l = options.Size(x.NoData)
			n += 1 + l + runtime.Sov(uint64(l))
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
		x := input.Message.Interface().(*PepPacketData)
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
		switch x := x.Packet.(type) {
		case *PepPacketData_NoData:
			encoded, err := options.Marshal(x.NoData)
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
		x := input.Message.Interface().(*PepPacketData)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PepPacketData: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PepPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
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
				v := &NoData{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Packet = &PepPacketData_NoData{v}
				iNdEx = postIndex
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
	md_NoData protoreflect.MessageDescriptor
)

func init() {
	file_fairyring_pep_packet_proto_init()
	md_NoData = File_fairyring_pep_packet_proto.Messages().ByName("NoData")
}

var _ protoreflect.Message = (*fastReflection_NoData)(nil)

type fastReflection_NoData NoData

func (x *NoData) ProtoReflect() protoreflect.Message {
	return (*fastReflection_NoData)(x)
}

func (x *NoData) slowProtoReflect() protoreflect.Message {
	mi := &file_fairyring_pep_packet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_NoData_messageType fastReflection_NoData_messageType
var _ protoreflect.MessageType = fastReflection_NoData_messageType{}

type fastReflection_NoData_messageType struct{}

func (x fastReflection_NoData_messageType) Zero() protoreflect.Message {
	return (*fastReflection_NoData)(nil)
}
func (x fastReflection_NoData_messageType) New() protoreflect.Message {
	return new(fastReflection_NoData)
}
func (x fastReflection_NoData_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_NoData
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_NoData) Descriptor() protoreflect.MessageDescriptor {
	return md_NoData
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_NoData) Type() protoreflect.MessageType {
	return _fastReflection_NoData_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_NoData) New() protoreflect.Message {
	return new(fastReflection_NoData)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_NoData) Interface() protoreflect.ProtoMessage {
	return (*NoData)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_NoData) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_NoData) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.NoData"))
		}
		panic(fmt.Errorf("message fairyring.pep.NoData does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NoData) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.NoData"))
		}
		panic(fmt.Errorf("message fairyring.pep.NoData does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_NoData) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.NoData"))
		}
		panic(fmt.Errorf("message fairyring.pep.NoData does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_NoData) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.NoData"))
		}
		panic(fmt.Errorf("message fairyring.pep.NoData does not contain field %s", fd.FullName()))
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
func (x *fastReflection_NoData) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.NoData"))
		}
		panic(fmt.Errorf("message fairyring.pep.NoData does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_NoData) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.pep.NoData"))
		}
		panic(fmt.Errorf("message fairyring.pep.NoData does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_NoData) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in fairyring.pep.NoData", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_NoData) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NoData) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_NoData) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_NoData) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*NoData)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*NoData)
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
		x := input.Message.Interface().(*NoData)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NoData: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
// source: fairyring/pep/packet.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PepPacketData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Packet:
	//
	//	*PepPacketData_NoData
	Packet isPepPacketData_Packet `protobuf_oneof:"packet"`
}

func (x *PepPacketData) Reset() {
	*x = PepPacketData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fairyring_pep_packet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PepPacketData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PepPacketData) ProtoMessage() {}

// Deprecated: Use PepPacketData.ProtoReflect.Descriptor instead.
func (*PepPacketData) Descriptor() ([]byte, []int) {
	return file_fairyring_pep_packet_proto_rawDescGZIP(), []int{0}
}

func (x *PepPacketData) GetPacket() isPepPacketData_Packet {
	if x != nil {
		return x.Packet
	}
	return nil
}

func (x *PepPacketData) GetNoData() *NoData {
	if x, ok := x.GetPacket().(*PepPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

type isPepPacketData_Packet interface {
	isPepPacketData_Packet()
}

type PepPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof"`
}

func (*PepPacketData_NoData) isPepPacketData_Packet() {}

type NoData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoData) Reset() {
	*x = NoData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fairyring_pep_packet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoData) ProtoMessage() {}

// Deprecated: Use NoData.ProtoReflect.Descriptor instead.
func (*NoData) Descriptor() ([]byte, []int) {
	return file_fairyring_pep_packet_proto_rawDescGZIP(), []int{1}
}

var File_fairyring_pep_packet_proto protoreflect.FileDescriptor

var file_fairyring_pep_packet_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x66, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x70, 0x2f,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x61,
	0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x65, 0x70, 0x22, 0x4a, 0x0a, 0x0d, 0x50,
	0x65, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x06,
	0x6e, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66,
	0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x65, 0x70, 0x2e, 0x4e, 0x6f, 0x44,
	0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x42, 0x08, 0x0a,
	0x06, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x08, 0x0a, 0x06, 0x4e, 0x6f, 0x44, 0x61, 0x74,
	0x61, 0x42, 0x95, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x61, 0x69, 0x72, 0x79, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x65, 0x70, 0x42, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64,
	0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x70, 0x65, 0x70, 0xa2, 0x02, 0x03, 0x46, 0x50, 0x58, 0xaa, 0x02, 0x0d, 0x46,
	0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x65, 0x70, 0xca, 0x02, 0x0d, 0x46,
	0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x50, 0x65, 0x70, 0xe2, 0x02, 0x19, 0x46,
	0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x50, 0x65, 0x70, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x46, 0x61, 0x69, 0x72, 0x79,
	0x72, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x50, 0x65, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_fairyring_pep_packet_proto_rawDescOnce sync.Once
	file_fairyring_pep_packet_proto_rawDescData = file_fairyring_pep_packet_proto_rawDesc
)

func file_fairyring_pep_packet_proto_rawDescGZIP() []byte {
	file_fairyring_pep_packet_proto_rawDescOnce.Do(func() {
		file_fairyring_pep_packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_fairyring_pep_packet_proto_rawDescData)
	})
	return file_fairyring_pep_packet_proto_rawDescData
}

var file_fairyring_pep_packet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fairyring_pep_packet_proto_goTypes = []interface{}{
	(*PepPacketData)(nil), // 0: fairyring.pep.PepPacketData
	(*NoData)(nil),        // 1: fairyring.pep.NoData
}
var file_fairyring_pep_packet_proto_depIdxs = []int32{
	1, // 0: fairyring.pep.PepPacketData.noData:type_name -> fairyring.pep.NoData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fairyring_pep_packet_proto_init() }
func file_fairyring_pep_packet_proto_init() {
	if File_fairyring_pep_packet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fairyring_pep_packet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PepPacketData); i {
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
		file_fairyring_pep_packet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoData); i {
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
	file_fairyring_pep_packet_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PepPacketData_NoData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fairyring_pep_packet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fairyring_pep_packet_proto_goTypes,
		DependencyIndexes: file_fairyring_pep_packet_proto_depIdxs,
		MessageInfos:      file_fairyring_pep_packet_proto_msgTypes,
	}.Build()
	File_fairyring_pep_packet_proto = out.File
	file_fairyring_pep_packet_proto_rawDesc = nil
	file_fairyring_pep_packet_proto_goTypes = nil
	file_fairyring_pep_packet_proto_depIdxs = nil
}
