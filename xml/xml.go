package xml

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"

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
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return xml.Unmarshal(b, dst)
}
