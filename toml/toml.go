package toml

import (
	"io"

	"github.com/pkg/errors"
	// "github.com/BurntSushi/toml"

	"github.com/why444216978/codec"
)

type TomlCodec struct{}

var _ codec.Codec = (*TomlCodec)(nil)

func (c TomlCodec) Encode(data interface{}) (io.Reader, error) {
	// TODO
	return nil, errors.New("not support")
}

func (c TomlCodec) Decode(r io.Reader, dst interface{}) error {
	// TODO
	return errors.New("not support")
}
