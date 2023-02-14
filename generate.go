//go:generate go run -mod=mod github.com/golang/mock/mockgen -package codec -source ./codec.go -destination ./mock.go  Codec
package codec
