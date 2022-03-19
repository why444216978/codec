package json

import (
	"bytes"
	"errors"
	"io"

	json "github.com/json-iterator/go"
	"github.com/valyala/bytebufferpool"

	"github.com/why444216978/codec"
)

type JSONCodec struct{}

var _ codec.Codec = (*JSONCodec)(nil)

func (c JSONCodec) Encode(data interface{}) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(b), nil
}

func (c JSONCodec) Decode(r io.Reader, dst interface{}) error {
	if r == nil {
		return errors.New("reader is nil")
	}

	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}

	return json.Unmarshal(buf.Bytes(), dst)
}
