package form

import (
	"bytes"
	"io"
	"net/url"

	"github.com/pkg/errors"

	"github.com/why444216978/codec"
)

type FormCodec struct{}

var _ codec.Codec = (*FormCodec)(nil)

func (c FormCodec) Encode(data interface{}) (io.Reader, error) {
	val, ok := data.(url.Values)
	if !ok {
		return nil, errors.New("data assert to url.Values fail")
	}

	buf := bytes.NewBuffer([]byte{})
	_, err := buf.WriteString(val.Encode())

	return buf, err
}

func (c FormCodec) Decode(r io.Reader, dst interface{}) error {
	return errors.New("not support")
}
