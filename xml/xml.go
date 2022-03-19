package xml

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"

	"github.com/valyala/bytebufferpool"
	"github.com/why444216978/codec"
)

type XMLCodec struct{}

var _ codec.Codec = (*XMLCodec)(nil)

func (c XMLCodec) Encode(data interface{}) (io.Reader, error) {
	b, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

func (c XMLCodec) Decode(r io.Reader, dst interface{}) error {
	if r == nil {
		return errors.New("reader is nil")
	}

	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}

	return xml.Unmarshal(buf.Bytes(), dst)
}
