package yaml

import (
	"io"

	"github.com/pkg/errors"
	// "gopkg.in/yaml.v2"

	"github.com/why444216978/codec"
)

type YamlCodec struct{}

var _ codec.Codec = (*YamlCodec)(nil)

func (c YamlCodec) Encode(data interface{}) (io.Reader, error) {
	// TODO
	return nil, errors.New("not support")
}

func (c YamlCodec) Decode(r io.Reader, dst interface{}) error {
	// TODO
	return errors.New("not support")
}
