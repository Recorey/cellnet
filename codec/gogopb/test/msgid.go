// Generated by github.com/adamluo159/cellnet/protoc-gen-msg
// DO NOT EDIT!
// Source: pb.proto

package test

import (
	"github.com/adamluo159/cellnet"
	"reflect"
	_ "github.com/adamluo159/cellnet/codec/gogopb"
	"github.com/adamluo159/cellnet/codec"
)

func init() {

	// pb.proto
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*ContentACK)(nil)).Elem(),
		ID:    60952,
	})
}
