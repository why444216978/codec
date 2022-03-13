package proto

import (
	"bytes"
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/why444216978/codec"
)

type ProtoCodec struct{}

var _ codec.Codec = (*ProtoCodec)(nil)

func (c ProtoCodec) Encode(data interface{}) (io.Reader, error) {
	message, ok := data.(proto.Message)
	if !ok {
		return nil, errors.New("data assert to proto.Message fail")
	}
	res, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(res), nil
}

func (c ProtoCodec) Decode(r io.Reader, dst interface{}) error {
	if _, ok := dst.(proto.Message); !ok {
		return errors.New("dst assert to proto.Message fail")
	}

	if buf, ok := r.(*bytes.Buffer); ok {
		return proto.Unmarshal(buf.Bytes(), dst.(proto.Message))
	}

	buf := bytes.NewBuffer([]byte{})
	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	return proto.Unmarshal(buf.Bytes(), dst.(proto.Message))
}
