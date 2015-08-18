// Code generated by protoc-gen-gogo.
// source: request_vote_request.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io3 "io"
import fmt12 "fmt"
import code_google_com_p_gogoprotobuf_proto6 "github.com/gogo/protobuf/proto"

import fmt13 "fmt"
import strings6 "strings"
import reflect6 "reflect"

import fmt14 "fmt"
import strings7 "strings"
import code_google_com_p_gogoprotobuf_proto7 "github.com/gogo/protobuf/proto"
import sort3 "sort"
import strconv3 "strconv"
import reflect7 "reflect"

import fmt15 "fmt"
import bytes3 "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type RequestVoteRequest struct {
	Term             *uint64 `protobuf:"varint,1,req" json:"Term,omitempty"`
	LastLogIndex     *uint64 `protobuf:"varint,2,req" json:"LastLogIndex,omitempty"`
	LastLogTerm      *uint64 `protobuf:"varint,3,req" json:"LastLogTerm,omitempty"`
	CandidateName    *string `protobuf:"bytes,4,req" json:"CandidateName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestVoteRequest) Reset()      { *m = RequestVoteRequest{} }
func (*RequestVoteRequest) ProtoMessage() {}

func (m *RequestVoteRequest) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *RequestVoteRequest) GetLastLogIndex() uint64 {
	if m != nil && m.LastLogIndex != nil {
		return *m.LastLogIndex
	}
	return 0
}

func (m *RequestVoteRequest) GetLastLogTerm() uint64 {
	if m != nil && m.LastLogTerm != nil {
		return *m.LastLogTerm
	}
	return 0
}

func (m *RequestVoteRequest) GetCandidateName() string {
	if m != nil && m.CandidateName != nil {
		return *m.CandidateName
	}
	return ""
}

func init() {
}
func (m *RequestVoteRequest) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io3.ErrUnexpectedEOF
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
			if wireType != 0 {
				return fmt12.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io3.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Term = &v
		case 2:
			if wireType != 0 {
				return fmt12.Errorf("proto: wrong wireType = %d for field LastLogIndex", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io3.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LastLogIndex = &v
		case 3:
			if wireType != 0 {
				return fmt12.Errorf("proto: wrong wireType = %d for field LastLogTerm", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io3.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LastLogTerm = &v
		case 4:
			if wireType != 2 {
				return fmt12.Errorf("proto: wrong wireType = %d for field CandidateName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io3.ErrUnexpectedEOF
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
				return io3.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.CandidateName = &s
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
			skippy, err := code_google_com_p_gogoprotobuf_proto6.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io3.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (this *RequestVoteRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings6.Join([]string{`&RequestVoteRequest{`,
		`Term:` + valueToStringRequestVoteRequest(this.Term) + `,`,
		`LastLogIndex:` + valueToStringRequestVoteRequest(this.LastLogIndex) + `,`,
		`LastLogTerm:` + valueToStringRequestVoteRequest(this.LastLogTerm) + `,`,
		`CandidateName:` + valueToStringRequestVoteRequest(this.CandidateName) + `,`,
		`XXX_unrecognized:` + fmt13.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRequestVoteRequest(v interface{}) string {
	rv := reflect6.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect6.Indirect(rv).Interface()
	return fmt13.Sprintf("*%v", pv)
}
func (m *RequestVoteRequest) Size() (n int) {
	var l int
	_ = l
	if m.Term != nil {
		n += 1 + sovRequestVoteRequest(uint64(*m.Term))
	}
	if m.LastLogIndex != nil {
		n += 1 + sovRequestVoteRequest(uint64(*m.LastLogIndex))
	}
	if m.LastLogTerm != nil {
		n += 1 + sovRequestVoteRequest(uint64(*m.LastLogTerm))
	}
	if m.CandidateName != nil {
		l = len(*m.CandidateName)
		n += 1 + l + sovRequestVoteRequest(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRequestVoteRequest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRequestVoteRequest(x uint64) (n int) {
	return sovRequestVoteRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedRequestVoteRequest(r randyRequestVoteRequest, easy bool) *RequestVoteRequest {
	this := &RequestVoteRequest{}
	v1 := uint64(r.Uint32())
	this.Term = &v1
	v2 := uint64(r.Uint32())
	this.LastLogIndex = &v2
	v3 := uint64(r.Uint32())
	this.LastLogTerm = &v3
	v4 := randStringRequestVoteRequest(r)
	this.CandidateName = &v4
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedRequestVoteRequest(r, 5)
	}
	return this
}

type randyRequestVoteRequest interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneRequestVoteRequest(r randyRequestVoteRequest) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringRequestVoteRequest(r randyRequestVoteRequest) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneRequestVoteRequest(r)
	}
	return string(tmps)
}
func randUnrecognizedRequestVoteRequest(r randyRequestVoteRequest, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldRequestVoteRequest(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldRequestVoteRequest(data []byte, r randyRequestVoteRequest, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateRequestVoteRequest(data, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		data = encodeVarintPopulateRequestVoteRequest(data, uint64(v6))
	case 1:
		data = encodeVarintPopulateRequestVoteRequest(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateRequestVoteRequest(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateRequestVoteRequest(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateRequestVoteRequest(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateRequestVoteRequest(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *RequestVoteRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RequestVoteRequest) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Term != nil {
		data[i] = 0x8
		i++
		i = encodeVarintRequestVoteRequest(data, i, uint64(*m.Term))
	}
	if m.LastLogIndex != nil {
		data[i] = 0x10
		i++
		i = encodeVarintRequestVoteRequest(data, i, uint64(*m.LastLogIndex))
	}
	if m.LastLogTerm != nil {
		data[i] = 0x18
		i++
		i = encodeVarintRequestVoteRequest(data, i, uint64(*m.LastLogTerm))
	}
	if m.CandidateName != nil {
		data[i] = 0x22
		i++
		i = encodeVarintRequestVoteRequest(data, i, uint64(len(*m.CandidateName)))
		i += copy(data[i:], *m.CandidateName)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64RequestVoteRequest(data []byte, offset int, v uint64) int {
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
func encodeFixed32RequestVoteRequest(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRequestVoteRequest(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *RequestVoteRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings7.Join([]string{`&protobuf.RequestVoteRequest{` + `Term:` + valueToGoStringRequestVoteRequest(this.Term, "uint64"), `LastLogIndex:` + valueToGoStringRequestVoteRequest(this.LastLogIndex, "uint64"), `LastLogTerm:` + valueToGoStringRequestVoteRequest(this.LastLogTerm, "uint64"), `CandidateName:` + valueToGoStringRequestVoteRequest(this.CandidateName, "string"), `XXX_unrecognized:` + fmt14.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringRequestVoteRequest(v interface{}, typ string) string {
	rv := reflect7.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect7.Indirect(rv).Interface()
	return fmt14.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringRequestVoteRequest(e map[int32]code_google_com_p_gogoprotobuf_proto7.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort3.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv3.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings7.Join(ss, ",") + "}"
	return s
}
func (this *RequestVoteRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt15.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*RequestVoteRequest)
	if !ok {
		return fmt15.Errorf("that is not of type *RequestVoteRequest")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt15.Errorf("that is type *RequestVoteRequest but is nil && this != nil")
	} else if this == nil {
		return fmt15.Errorf("that is type *RequestVoteRequestbut is not nil && this == nil")
	}
	if this.Term != nil && that1.Term != nil {
		if *this.Term != *that1.Term {
			return fmt15.Errorf("Term this(%v) Not Equal that(%v)", *this.Term, *that1.Term)
		}
	} else if this.Term != nil {
		return fmt15.Errorf("this.Term == nil && that.Term != nil")
	} else if that1.Term != nil {
		return fmt15.Errorf("Term this(%v) Not Equal that(%v)", this.Term, that1.Term)
	}
	if this.LastLogIndex != nil && that1.LastLogIndex != nil {
		if *this.LastLogIndex != *that1.LastLogIndex {
			return fmt15.Errorf("LastLogIndex this(%v) Not Equal that(%v)", *this.LastLogIndex, *that1.LastLogIndex)
		}
	} else if this.LastLogIndex != nil {
		return fmt15.Errorf("this.LastLogIndex == nil && that.LastLogIndex != nil")
	} else if that1.LastLogIndex != nil {
		return fmt15.Errorf("LastLogIndex this(%v) Not Equal that(%v)", this.LastLogIndex, that1.LastLogIndex)
	}
	if this.LastLogTerm != nil && that1.LastLogTerm != nil {
		if *this.LastLogTerm != *that1.LastLogTerm {
			return fmt15.Errorf("LastLogTerm this(%v) Not Equal that(%v)", *this.LastLogTerm, *that1.LastLogTerm)
		}
	} else if this.LastLogTerm != nil {
		return fmt15.Errorf("this.LastLogTerm == nil && that.LastLogTerm != nil")
	} else if that1.LastLogTerm != nil {
		return fmt15.Errorf("LastLogTerm this(%v) Not Equal that(%v)", this.LastLogTerm, that1.LastLogTerm)
	}
	if this.CandidateName != nil && that1.CandidateName != nil {
		if *this.CandidateName != *that1.CandidateName {
			return fmt15.Errorf("CandidateName this(%v) Not Equal that(%v)", *this.CandidateName, *that1.CandidateName)
		}
	} else if this.CandidateName != nil {
		return fmt15.Errorf("this.CandidateName == nil && that.CandidateName != nil")
	} else if that1.CandidateName != nil {
		return fmt15.Errorf("CandidateName this(%v) Not Equal that(%v)", this.CandidateName, that1.CandidateName)
	}
	if !bytes3.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt15.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *RequestVoteRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RequestVoteRequest)
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
	if this.Term != nil && that1.Term != nil {
		if *this.Term != *that1.Term {
			return false
		}
	} else if this.Term != nil {
		return false
	} else if that1.Term != nil {
		return false
	}
	if this.LastLogIndex != nil && that1.LastLogIndex != nil {
		if *this.LastLogIndex != *that1.LastLogIndex {
			return false
		}
	} else if this.LastLogIndex != nil {
		return false
	} else if that1.LastLogIndex != nil {
		return false
	}
	if this.LastLogTerm != nil && that1.LastLogTerm != nil {
		if *this.LastLogTerm != *that1.LastLogTerm {
			return false
		}
	} else if this.LastLogTerm != nil {
		return false
	} else if that1.LastLogTerm != nil {
		return false
	}
	if this.CandidateName != nil && that1.CandidateName != nil {
		if *this.CandidateName != *that1.CandidateName {
			return false
		}
	} else if this.CandidateName != nil {
		return false
	} else if that1.CandidateName != nil {
		return false
	}
	if !bytes3.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
