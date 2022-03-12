package json

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/why444216978/codec"
)

type JSONCodec struct{}

var _ codec.Codec = (*JSONCodec)(nil)

func (c JSONCodec) Encode(data interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return nil, err
	}
	return buf, nil
}

func (c JSONCodec) Decode(r io.Reader, dst interface{}) error {
	return json.NewDecoder(r).Decode(dst)
}
