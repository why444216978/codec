package proto

import (
	"io"

	"github.com/pkg/errors"
	// "github.com/golang/protobuf"

	"github.com/why444216978/codec"
)

type ProtoCodec struct{}

var _ codec.Codec = (*ProtoCodec)(nil)

func (c ProtoCodec) Encode(data interface{}) (io.Reader, error) {
	// TODO
	return nil, errors.New("not support")
}

func (c ProtoCodec) Decode(r io.Reader, dst interface{}) error {
	// TODO
	return errors.New("not support")
}
