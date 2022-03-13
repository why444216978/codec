package yaml

import (
	"bytes"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

type conf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func TestYamlCodec_Encode(t *testing.T) {
	codec := YamlCodec{}
	convey.Convey("TestYamlCodec_Encode", t, func() {
		convey.Convey("success", func() {
			c := conf{
				Host: "127.0.0.1",
				Port: 80,
			}
			r, err := codec.Encode(c)
			assert.Equal(t, err, nil)
			assert.Equal(t, r, bytes.NewBuffer([]uint8{104, 111, 115, 116, 58, 32, 49, 50, 55, 46, 48, 46, 48, 46, 49, 10, 112, 111, 114, 116, 58, 32, 56, 48, 10}))
		})
	})
}

func TestYamlCodec_Decode(t *testing.T) {
	codec := YamlCodec{}
	convey.Convey("TestYamlCodec_Decode", t, func() {
		convey.Convey("success", func() {
			c := &conf{}
			err := codec.Decode(bytes.NewBuffer([]uint8{104, 111, 115, 116, 58, 32, 49, 50, 55, 46, 48, 46, 48, 46, 49, 10, 112, 111, 114, 116, 58, 32, 56, 48, 10}), c)
			assert.Equal(t, err, nil)
			assert.Equal(t, c.Host, "127.0.0.1")
			assert.Equal(t, c.Port, 80)
		})
		convey.Convey("reader is nil", func() {
			c := &conf{}
			err := codec.Decode(nil, c)
			assert.Equal(t, err.Error(), "reader is nil")
		})
	})
}
