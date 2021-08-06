package codec

import "io"

type Header struct {
	ServiceMethod string //ServiceName.MethodName
	Seq           uint64 //Sequence number
	Error         string //Error reported by server
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}
