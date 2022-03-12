package codec

import (
	"io"
)

type Codec interface {
	Encode(data interface{}) (io.Reader, error)
	Decode(r io.Reader, dst interface{}) error
}
