package toml

import (
	"bytes"
	"errors"
	"io"

	"github.com/BurntSushi/toml"
	"github.com/valyala/bytebufferpool"

	"github.com/why444216978/codec"
)

type TomlCodec struct{}

var _ codec.Codec = (*TomlCodec)(nil)

func (c TomlCodec) Encode(data interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(buf).Encode(data); err != nil {
		return nil, err
	}

	return buf, nil
}

func (c TomlCodec) Decode(r io.Reader, dst interface{}) error {
	if r == nil {
		return errors.New("reader is nil")
	}

	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}

	return toml.Unmarshal(buf.Bytes(), dst)
}
