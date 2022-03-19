package proto

import (
	"bytes"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

type message struct {
	Message *string `protobuf:"bytes,1,req,name=message" json:"echo,omitempty"`
}

func (m *message) Reset() { *m = message{} }

func (m *message) String() string { return proto.CompactTextString(m) }

func (m *message) ProtoMessage() {}

var msg = &message{Message: proto.String("protobuf")}

func TestProtoCodec_Encode(t *testing.T) {
	codec := ProtoCodec{}
	convey.Convey("TestProtoCodec_Encode", t, func() {
		convey.Convey("success", func() {
			r, err := codec.Encode(msg)
			assert.Equal(t, err, nil)
			assert.Equal(t, r, bytes.NewBuffer([]byte{10, 8, 112, 114, 111, 116, 111, 98, 117, 102}))
		})
		convey.Convey("data not proto.Message", func() {
			r, err := codec.Encode(1)
			assert.Equal(t, err.Error(), "data assert to proto.Message fail")
			assert.Equal(t, r, nil)
		})
	})
}

func TestProtoCodec_Decode(t *testing.T) {
	codec := ProtoCodec{}
	convey.Convey("TestProtoCodec_Decode", t, func() {
		convey.Convey("success", func() {
			msg := &message{}
			err := codec.Decode(bytes.NewReader([]byte{10, 8, 112, 114, 111, 116, 111, 98, 117, 102}), msg)
			assert.Equal(t, err, nil)
			assert.Equal(t, *msg.Message, "protobuf")
		})
		convey.Convey("data not proto.Message", func() {
			err := codec.Decode(bytes.NewBuffer([]byte{10, 8, 112, 114, 111, 116, 111, 98, 117, 102}), nil)
			assert.Equal(t, err.Error(), "dst assert to proto.Message fail")
		})
		convey.Convey("reader is nil", func() {
			err := codec.Decode(nil, nil)
			assert.Equal(t, err.Error(), "reader is nil")
		})
	})
}
