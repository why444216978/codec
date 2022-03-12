package form

import (
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

func TestFormCodec_Encode(t *testing.T) {
	codec := FormCodec{}
	convey.Convey("TestFormCodec_Encode", t, func() {
		convey.Convey("success", func() {
			data := url.Values{
				"a": {"a"},
			}
			r, err := codec.Encode(data)
			assert.Equal(t, err, nil)

			b, err := ioutil.ReadAll(r)
			assert.Equal(t, err, nil)
			assert.Equal(t, b, []byte{97, 61, 97})
		})
		convey.Convey("data nil", func() {
			r, err := codec.Encode(nil)
			assert.Equal(t, err.Error(), "data assert to url.Values fail")
			assert.Equal(t, r, nil)
		})
	})
}

func TestFormCodec_Decode(t *testing.T) {
	codec := FormCodec{}
	convey.Convey("TestFormCodec_Decode", t, func() {
		convey.Convey("not support", func() {
			err := codec.Decode(nil, nil)
			assert.Equal(t, err.Error(), "not support")
		})
	})
}
