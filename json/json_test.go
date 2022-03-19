package json

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

func TestJSONCodec_Encode(t *testing.T) {
	codec := JSONCodec{}
	convey.Convey("TestJSONCodec_Encode", t, func() {
		convey.Convey("success", func() {
			data := map[string]string{"a": "a"}
			r, err := codec.Encode(data)
			assert.Equal(t, err, nil)

			b, err := ioutil.ReadAll(r)
			assert.Equal(t, err, nil)
			assert.Equal(t, b, []byte{123, 34, 97, 34, 58, 34, 97, 34, 125})
		})
	})
}

func TestJSONCodec_Decode(t *testing.T) {
	codec := JSONCodec{}
	convey.Convey("TestJSONCodec_Decode", t, func() {
		convey.Convey("success", func() {
			type Data struct {
				A string `json:"a"`
			}
			data := &Data{}
			err := codec.Decode(bytes.NewReader([]byte(`{"a":"a"}`)), data)
			assert.Equal(t, err, nil)
			assert.Equal(t, data.A, "a")
		})
		convey.Convey("err", func() {
			type Data struct {
				A int `json:"a"`
			}
			data := &Data{}
			err := codec.Decode(bytes.NewReader([]byte(`{"a":"a"}`)), data)
			assert.Equal(t, err != nil, true)
			assert.Equal(t, data.A, 0)
		})
		convey.Convey("reader is nil", func() {
			type Data struct {
				A int `json:"a"`
			}
			data := &Data{}
			err := codec.Decode(nil, data)
			assert.Equal(t, err.Error(), "reader is nil")
		})
	})
}
