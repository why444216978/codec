package yaml

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/why444216978/codec"
)

type YamlCodec struct{}

var _ codec.Codec = (*YamlCodec)(nil)

func (c YamlCodec) Encode(data interface{}) (io.Reader, error) {
	b, err := yaml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}

func (c YamlCodec) Decode(r io.Reader, dst interface{}) error {
	if r == nil {
		return errors.New("reader is nil")
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, dst)
}
