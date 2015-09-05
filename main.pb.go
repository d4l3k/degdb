// Code generated by protoc-gen-go.
// source: main.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	main.proto

It has these top-level messages:
	TripleSet
	Triple
	Node
	Query
	Filter
	QueryResp
	AddTriplesRequest
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Filter_Type int32

const (
	Filter_EXISTS    Filter_Type = 0
	Filter_EQUAL     Filter_Type = 1
	Filter_NOT_EQUAL Filter_Type = 2
)

var Filter_Type_name = map[int32]string{
	0: "EXISTS",
	1: "EQUAL",
	2: "NOT_EQUAL",
}
var Filter_Type_value = map[string]int32{
	"EXISTS":    0,
	"EQUAL":     1,
	"NOT_EQUAL": 2,
}

func (x Filter_Type) String() string {
	return proto.EnumName(Filter_Type_name, int32(x))
}

type TripleSet struct {
	Triples []*Triple `protobuf:"bytes,1,rep,name=triples" json:"triples,omitempty"`
}

func (m *TripleSet) Reset()         { *m = TripleSet{} }
func (m *TripleSet) String() string { return proto.CompactTextString(m) }
func (*TripleSet) ProtoMessage()    {}

func (m *TripleSet) GetTriples() []*Triple {
	if m != nil {
		return m.Triples
	}
	return nil
}

type Triple struct {
	Subj string `protobuf:"bytes,1,opt,name=subj" json:"subj,omitempty"`
	Pred string `protobuf:"bytes,2,opt,name=pred" json:"pred,omitempty"`
	Obj  string `protobuf:"bytes,3,opt,name=obj" json:"obj,omitempty"`
}

func (m *Triple) Reset()         { *m = Triple{} }
func (m *Triple) String() string { return proto.CompactTextString(m) }
func (*Triple) ProtoMessage()    {}

type Node struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}

type Query struct {
	Source *Node     `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Id     int64     `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Subj   string    `protobuf:"bytes,3,opt,name=subj" json:"subj,omitempty"`
	Filter []*Filter `protobuf:"bytes,4,rep,name=filter" json:"filter,omitempty"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}

func (m *Query) GetSource() *Node {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Query) GetFilter() []*Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type Filter struct {
	Pred string      `protobuf:"bytes,1,opt,name=pred" json:"pred,omitempty"`
	Obj  string      `protobuf:"bytes,2,opt,name=obj" json:"obj,omitempty"`
	Type Filter_Type `protobuf:"varint,3,opt,name=type,enum=Filter_Type" json:"type,omitempty"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}

type QueryResp struct {
	Source  *Node     `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Id      int64     `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Triples []*Triple `protobuf:"bytes,3,rep,name=triples" json:"triples,omitempty"`
}

func (m *QueryResp) Reset()         { *m = QueryResp{} }
func (m *QueryResp) String() string { return proto.CompactTextString(m) }
func (*QueryResp) ProtoMessage()    {}

func (m *QueryResp) GetSource() *Node {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *QueryResp) GetTriples() []*Triple {
	if m != nil {
		return m.Triples
	}
	return nil
}

type AddTriplesRequest struct {
	Source  *Node     `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Triples []*Triple `protobuf:"bytes,2,rep,name=triples" json:"triples,omitempty"`
}

func (m *AddTriplesRequest) Reset()         { *m = AddTriplesRequest{} }
func (m *AddTriplesRequest) String() string { return proto.CompactTextString(m) }
func (*AddTriplesRequest) ProtoMessage()    {}

func (m *AddTriplesRequest) GetSource() *Node {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AddTriplesRequest) GetTriples() []*Triple {
	if m != nil {
		return m.Triples
	}
	return nil
}

func init() {
	proto.RegisterEnum("Filter_Type", Filter_Type_name, Filter_Type_value)
}
