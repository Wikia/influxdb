// Code generated by protoc-gen-gogo.
// source: snapshot_recovery_request.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io5 "io"
import fmt20 "fmt"
import code_google_com_p_gogoprotobuf_proto10 "github.com/gogo/protobuf/proto"

import fmt21 "fmt"
import strings10 "strings"
import reflect10 "reflect"

import fmt22 "fmt"
import strings11 "strings"
import code_google_com_p_gogoprotobuf_proto11 "github.com/gogo/protobuf/proto"
import sort5 "sort"
import strconv5 "strconv"
import reflect11 "reflect"

import fmt23 "fmt"
import bytes5 "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SnapshotRecoveryRequest struct {
	LeaderName       *string                         `protobuf:"bytes,1,req" json:"LeaderName,omitempty"`
	LastIndex        *uint64                         `protobuf:"varint,2,req" json:"LastIndex,omitempty"`
	LastTerm         *uint64                         `protobuf:"varint,3,req" json:"LastTerm,omitempty"`
	Peers            []*SnapshotRecoveryRequest_Peer `protobuf:"bytes,4,rep" json:"Peers,omitempty"`
	State            []byte                          `protobuf:"bytes,5,req" json:"State,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *SnapshotRecoveryRequest) Reset()      { *m = SnapshotRecoveryRequest{} }
func (*SnapshotRecoveryRequest) ProtoMessage() {}

func (m *SnapshotRecoveryRequest) GetLeaderName() string {
	if m != nil && m.LeaderName != nil {
		return *m.LeaderName
	}
	return ""
}

func (m *SnapshotRecoveryRequest) GetLastIndex() uint64 {
	if m != nil && m.LastIndex != nil {
		return *m.LastIndex
	}
	return 0
}

func (m *SnapshotRecoveryRequest) GetLastTerm() uint64 {
	if m != nil && m.LastTerm != nil {
		return *m.LastTerm
	}
	return 0
}

func (m *SnapshotRecoveryRequest) GetPeers() []*SnapshotRecoveryRequest_Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *SnapshotRecoveryRequest) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

type SnapshotRecoveryRequest_Peer struct {
	Name             *string `protobuf:"bytes,1,req" json:"Name,omitempty"`
	ConnectionString *string `protobuf:"bytes,2,req" json:"ConnectionString,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SnapshotRecoveryRequest_Peer) Reset()      { *m = SnapshotRecoveryRequest_Peer{} }
func (*SnapshotRecoveryRequest_Peer) ProtoMessage() {}

func (m *SnapshotRecoveryRequest_Peer) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SnapshotRecoveryRequest_Peer) GetConnectionString() string {
	if m != nil && m.ConnectionString != nil {
		return *m.ConnectionString
	}
	return ""
}

func init() {
}
func (m *SnapshotRecoveryRequest) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io5.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt20.Errorf("proto: wrong wireType = %d for field LeaderName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io5.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io5.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.LeaderName = &s
			index = postIndex
		case 2:
			if wireType != 0 {
				return fmt20.Errorf("proto: wrong wireType = %d for field LastIndex", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io5.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LastIndex = &v
		case 3:
			if wireType != 0 {
				return fmt20.Errorf("proto: wrong wireType = %d for field LastTerm", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io5.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LastTerm = &v
		case 4:
			if wireType != 2 {
				return fmt20.Errorf("proto: wrong wireType = %d for field Peers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io5.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io5.ErrUnexpectedEOF
			}
			m.Peers = append(m.Peers, &SnapshotRecoveryRequest_Peer{})
			m.Peers[len(m.Peers)-1].Unmarshal(data[index:postIndex])
			index = postIndex
		case 5:
			if wireType != 2 {
				return fmt20.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io5.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io5.ErrUnexpectedEOF
			}
			m.State = append(m.State, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto10.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io5.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *SnapshotRecoveryRequest_Peer) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io5.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt20.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io5.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io5.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.Name = &s
			index = postIndex
		case 2:
			if wireType != 2 {
				return fmt20.Errorf("proto: wrong wireType = %d for field ConnectionString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io5.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io5.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.ConnectionString = &s
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto10.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io5.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (this *SnapshotRecoveryRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings10.Join([]string{`&SnapshotRecoveryRequest{`,
		`LeaderName:` + valueToStringSnapshotRecoveryRequest(this.LeaderName) + `,`,
		`LastIndex:` + valueToStringSnapshotRecoveryRequest(this.LastIndex) + `,`,
		`LastTerm:` + valueToStringSnapshotRecoveryRequest(this.LastTerm) + `,`,
		`Peers:` + strings10.Replace(fmt21.Sprintf("%v", this.Peers), "SnapshotRecoveryRequest_Peer", "SnapshotRecoveryRequest_Peer", 1) + `,`,
		`State:` + valueToStringSnapshotRecoveryRequest(this.State) + `,`,
		`XXX_unrecognized:` + fmt21.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SnapshotRecoveryRequest_Peer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings10.Join([]string{`&SnapshotRecoveryRequest_Peer{`,
		`Name:` + valueToStringSnapshotRecoveryRequest(this.Name) + `,`,
		`ConnectionString:` + valueToStringSnapshotRecoveryRequest(this.ConnectionString) + `,`,
		`XXX_unrecognized:` + fmt21.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSnapshotRecoveryRequest(v interface{}) string {
	rv := reflect10.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect10.Indirect(rv).Interface()
	return fmt21.Sprintf("*%v", pv)
}
func (m *SnapshotRecoveryRequest) Size() (n int) {
	var l int
	_ = l
	if m.LeaderName != nil {
		l = len(*m.LeaderName)
		n += 1 + l + sovSnapshotRecoveryRequest(uint64(l))
	}
	if m.LastIndex != nil {
		n += 1 + sovSnapshotRecoveryRequest(uint64(*m.LastIndex))
	}
	if m.LastTerm != nil {
		n += 1 + sovSnapshotRecoveryRequest(uint64(*m.LastTerm))
	}
	if len(m.Peers) > 0 {
		for _, e := range m.Peers {
			l = e.Size()
			n += 1 + l + sovSnapshotRecoveryRequest(uint64(l))
		}
	}
	if m.State != nil {
		l = len(m.State)
		n += 1 + l + sovSnapshotRecoveryRequest(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *SnapshotRecoveryRequest_Peer) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovSnapshotRecoveryRequest(uint64(l))
	}
	if m.ConnectionString != nil {
		l = len(*m.ConnectionString)
		n += 1 + l + sovSnapshotRecoveryRequest(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSnapshotRecoveryRequest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSnapshotRecoveryRequest(x uint64) (n int) {
	return sovSnapshotRecoveryRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedSnapshotRecoveryRequest(r randySnapshotRecoveryRequest, easy bool) *SnapshotRecoveryRequest {
	this := &SnapshotRecoveryRequest{}
	v1 := randStringSnapshotRecoveryRequest(r)
	this.LeaderName = &v1
	v2 := uint64(r.Uint32())
	this.LastIndex = &v2
	v3 := uint64(r.Uint32())
	this.LastTerm = &v3
	if r.Intn(10) != 0 {
		v4 := r.Intn(10)
		this.Peers = make([]*SnapshotRecoveryRequest_Peer, v4)
		for i := 0; i < v4; i++ {
			this.Peers[i] = NewPopulatedSnapshotRecoveryRequest_Peer(r, easy)
		}
	}
	v5 := r.Intn(100)
	this.State = make([]byte, v5)
	for i := 0; i < v5; i++ {
		this.State[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSnapshotRecoveryRequest(r, 6)
	}
	return this
}

func NewPopulatedSnapshotRecoveryRequest_Peer(r randySnapshotRecoveryRequest, easy bool) *SnapshotRecoveryRequest_Peer {
	this := &SnapshotRecoveryRequest_Peer{}
	v6 := randStringSnapshotRecoveryRequest(r)
	this.Name = &v6
	v7 := randStringSnapshotRecoveryRequest(r)
	this.ConnectionString = &v7
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSnapshotRecoveryRequest(r, 3)
	}
	return this
}

type randySnapshotRecoveryRequest interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSnapshotRecoveryRequest(r randySnapshotRecoveryRequest) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringSnapshotRecoveryRequest(r randySnapshotRecoveryRequest) string {
	v8 := r.Intn(100)
	tmps := make([]rune, v8)
	for i := 0; i < v8; i++ {
		tmps[i] = randUTF8RuneSnapshotRecoveryRequest(r)
	}
	return string(tmps)
}
func randUnrecognizedSnapshotRecoveryRequest(r randySnapshotRecoveryRequest, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldSnapshotRecoveryRequest(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldSnapshotRecoveryRequest(data []byte, r randySnapshotRecoveryRequest, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateSnapshotRecoveryRequest(data, uint64(key))
		v9 := r.Int63()
		if r.Intn(2) == 0 {
			v9 *= -1
		}
		data = encodeVarintPopulateSnapshotRecoveryRequest(data, uint64(v9))
	case 1:
		data = encodeVarintPopulateSnapshotRecoveryRequest(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateSnapshotRecoveryRequest(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateSnapshotRecoveryRequest(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateSnapshotRecoveryRequest(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateSnapshotRecoveryRequest(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *SnapshotRecoveryRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SnapshotRecoveryRequest) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.LeaderName != nil {
		data[i] = 0xa
		i++
		i = encodeVarintSnapshotRecoveryRequest(data, i, uint64(len(*m.LeaderName)))
		i += copy(data[i:], *m.LeaderName)
	}
	if m.LastIndex != nil {
		data[i] = 0x10
		i++
		i = encodeVarintSnapshotRecoveryRequest(data, i, uint64(*m.LastIndex))
	}
	if m.LastTerm != nil {
		data[i] = 0x18
		i++
		i = encodeVarintSnapshotRecoveryRequest(data, i, uint64(*m.LastTerm))
	}
	if len(m.Peers) > 0 {
		for _, msg := range m.Peers {
			data[i] = 0x22
			i++
			i = encodeVarintSnapshotRecoveryRequest(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.State != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintSnapshotRecoveryRequest(data, i, uint64(len(m.State)))
		i += copy(data[i:], m.State)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *SnapshotRecoveryRequest_Peer) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SnapshotRecoveryRequest_Peer) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		data[i] = 0xa
		i++
		i = encodeVarintSnapshotRecoveryRequest(data, i, uint64(len(*m.Name)))
		i += copy(data[i:], *m.Name)
	}
	if m.ConnectionString != nil {
		data[i] = 0x12
		i++
		i = encodeVarintSnapshotRecoveryRequest(data, i, uint64(len(*m.ConnectionString)))
		i += copy(data[i:], *m.ConnectionString)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64SnapshotRecoveryRequest(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32SnapshotRecoveryRequest(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSnapshotRecoveryRequest(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *SnapshotRecoveryRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings11.Join([]string{`&protobuf.SnapshotRecoveryRequest{` + `LeaderName:` + valueToGoStringSnapshotRecoveryRequest(this.LeaderName, "string"), `LastIndex:` + valueToGoStringSnapshotRecoveryRequest(this.LastIndex, "uint64"), `LastTerm:` + valueToGoStringSnapshotRecoveryRequest(this.LastTerm, "uint64"), `Peers:` + fmt22.Sprintf("%#v", this.Peers), `State:` + valueToGoStringSnapshotRecoveryRequest(this.State, "byte"), `XXX_unrecognized:` + fmt22.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *SnapshotRecoveryRequest_Peer) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings11.Join([]string{`&protobuf.SnapshotRecoveryRequest_Peer{` + `Name:` + valueToGoStringSnapshotRecoveryRequest(this.Name, "string"), `ConnectionString:` + valueToGoStringSnapshotRecoveryRequest(this.ConnectionString, "string"), `XXX_unrecognized:` + fmt22.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringSnapshotRecoveryRequest(v interface{}, typ string) string {
	rv := reflect11.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect11.Indirect(rv).Interface()
	return fmt22.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringSnapshotRecoveryRequest(e map[int32]code_google_com_p_gogoprotobuf_proto11.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort5.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv5.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings11.Join(ss, ",") + "}"
	return s
}
func (this *SnapshotRecoveryRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt23.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*SnapshotRecoveryRequest)
	if !ok {
		return fmt23.Errorf("that is not of type *SnapshotRecoveryRequest")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt23.Errorf("that is type *SnapshotRecoveryRequest but is nil && this != nil")
	} else if this == nil {
		return fmt23.Errorf("that is type *SnapshotRecoveryRequestbut is not nil && this == nil")
	}
	if this.LeaderName != nil && that1.LeaderName != nil {
		if *this.LeaderName != *that1.LeaderName {
			return fmt23.Errorf("LeaderName this(%v) Not Equal that(%v)", *this.LeaderName, *that1.LeaderName)
		}
	} else if this.LeaderName != nil {
		return fmt23.Errorf("this.LeaderName == nil && that.LeaderName != nil")
	} else if that1.LeaderName != nil {
		return fmt23.Errorf("LeaderName this(%v) Not Equal that(%v)", this.LeaderName, that1.LeaderName)
	}
	if this.LastIndex != nil && that1.LastIndex != nil {
		if *this.LastIndex != *that1.LastIndex {
			return fmt23.Errorf("LastIndex this(%v) Not Equal that(%v)", *this.LastIndex, *that1.LastIndex)
		}
	} else if this.LastIndex != nil {
		return fmt23.Errorf("this.LastIndex == nil && that.LastIndex != nil")
	} else if that1.LastIndex != nil {
		return fmt23.Errorf("LastIndex this(%v) Not Equal that(%v)", this.LastIndex, that1.LastIndex)
	}
	if this.LastTerm != nil && that1.LastTerm != nil {
		if *this.LastTerm != *that1.LastTerm {
			return fmt23.Errorf("LastTerm this(%v) Not Equal that(%v)", *this.LastTerm, *that1.LastTerm)
		}
	} else if this.LastTerm != nil {
		return fmt23.Errorf("this.LastTerm == nil && that.LastTerm != nil")
	} else if that1.LastTerm != nil {
		return fmt23.Errorf("LastTerm this(%v) Not Equal that(%v)", this.LastTerm, that1.LastTerm)
	}
	if len(this.Peers) != len(that1.Peers) {
		return fmt23.Errorf("Peers this(%v) Not Equal that(%v)", len(this.Peers), len(that1.Peers))
	}
	for i := range this.Peers {
		if !this.Peers[i].Equal(that1.Peers[i]) {
			return fmt23.Errorf("Peers this[%v](%v) Not Equal that[%v](%v)", i, this.Peers[i], i, that1.Peers[i])
		}
	}
	if !bytes5.Equal(this.State, that1.State) {
		return fmt23.Errorf("State this(%v) Not Equal that(%v)", this.State, that1.State)
	}
	if !bytes5.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt23.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *SnapshotRecoveryRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SnapshotRecoveryRequest)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.LeaderName != nil && that1.LeaderName != nil {
		if *this.LeaderName != *that1.LeaderName {
			return false
		}
	} else if this.LeaderName != nil {
		return false
	} else if that1.LeaderName != nil {
		return false
	}
	if this.LastIndex != nil && that1.LastIndex != nil {
		if *this.LastIndex != *that1.LastIndex {
			return false
		}
	} else if this.LastIndex != nil {
		return false
	} else if that1.LastIndex != nil {
		return false
	}
	if this.LastTerm != nil && that1.LastTerm != nil {
		if *this.LastTerm != *that1.LastTerm {
			return false
		}
	} else if this.LastTerm != nil {
		return false
	} else if that1.LastTerm != nil {
		return false
	}
	if len(this.Peers) != len(that1.Peers) {
		return false
	}
	for i := range this.Peers {
		if !this.Peers[i].Equal(that1.Peers[i]) {
			return false
		}
	}
	if !bytes5.Equal(this.State, that1.State) {
		return false
	}
	if !bytes5.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SnapshotRecoveryRequest_Peer) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt23.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*SnapshotRecoveryRequest_Peer)
	if !ok {
		return fmt23.Errorf("that is not of type *SnapshotRecoveryRequest_Peer")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt23.Errorf("that is type *SnapshotRecoveryRequest_Peer but is nil && this != nil")
	} else if this == nil {
		return fmt23.Errorf("that is type *SnapshotRecoveryRequest_Peerbut is not nil && this == nil")
	}
	if this.Name != nil && that1.Name != nil {
		if *this.Name != *that1.Name {
			return fmt23.Errorf("Name this(%v) Not Equal that(%v)", *this.Name, *that1.Name)
		}
	} else if this.Name != nil {
		return fmt23.Errorf("this.Name == nil && that.Name != nil")
	} else if that1.Name != nil {
		return fmt23.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.ConnectionString != nil && that1.ConnectionString != nil {
		if *this.ConnectionString != *that1.ConnectionString {
			return fmt23.Errorf("ConnectionString this(%v) Not Equal that(%v)", *this.ConnectionString, *that1.ConnectionString)
		}
	} else if this.ConnectionString != nil {
		return fmt23.Errorf("this.ConnectionString == nil && that.ConnectionString != nil")
	} else if that1.ConnectionString != nil {
		return fmt23.Errorf("ConnectionString this(%v) Not Equal that(%v)", this.ConnectionString, that1.ConnectionString)
	}
	if !bytes5.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt23.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *SnapshotRecoveryRequest_Peer) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SnapshotRecoveryRequest_Peer)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != nil && that1.Name != nil {
		if *this.Name != *that1.Name {
			return false
		}
	} else if this.Name != nil {
		return false
	} else if that1.Name != nil {
		return false
	}
	if this.ConnectionString != nil && that1.ConnectionString != nil {
		if *this.ConnectionString != *that1.ConnectionString {
			return false
		}
	} else if this.ConnectionString != nil {
		return false
	} else if that1.ConnectionString != nil {
		return false
	}
	if !bytes5.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
