package toml

import (
	"bytes"
	"io"

	"github.com/BurntSushi/toml"

	"github.com/why444216978/codec"
)

type TomlCodec struct{}

var _ codec.Codec = (*TomlCodec)(nil)

func (c TomlCodec) Encode(data interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := toml.NewEncoder(buf).Encode(data); err != nil {
		return nil, err
	}

	return buf, nil
}

func (c TomlCodec) Decode(r io.Reader, dst interface{}) error {
	if _, err := toml.NewDecoder(r).Decode(dst); err != nil {
		return err
	}

	return nil
}
