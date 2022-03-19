package toml

import (
	"bytes"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

func TestTomlCodec_Encode(t *testing.T) {
	codec := TomlCodec{}
	convey.Convey("TestTomlCodec_Encode", t, func() {
		convey.Convey("success", func() {
			data := map[string]string{"k": "v"}
			r, err := codec.Encode(data)

			assert.Equal(t, err, nil)
			assert.Equal(t, r, bytes.NewBufferString(`k = "v"`+"\n"))
		})
		convey.Convey("err", func() {
			data := 1
			r, err := codec.Encode(data)

			assert.Equal(t, err.Error(), "toml: top-level values must be Go maps or structs")
			assert.Equal(t, r, nil)
		})
	})
}

func TestTomlCodec_Decode(t *testing.T) {
	codec := TomlCodec{}
	convey.Convey("TestTomlCodec_Decode", t, func() {
		convey.Convey("success", func() {
			data := map[string]string{}
			err := codec.Decode(bytes.NewBufferString(`k = "v"`+"\n"), &data)

			assert.Equal(t, err, nil)
			assert.Equal(t, data, map[string]string{"k": "v"})
		})
		convey.Convey("err", func() {
			data := 0
			err := codec.Decode(bytes.NewBufferString(`k = "v"`+"\n"), &data)

			assert.Equal(t, err.Error(), "toml: cannot decode to type int")
			assert.Equal(t, data, 0)
		})
		convey.Convey("reader is nil", func() {
			data := map[string]string{}
			err := codec.Decode(nil, &data)

			assert.Equal(t, err.Error(), "reader is nil")
		})
	})
}
