package encoder

import "github.com/1-bi/eventcar/schema"

type Encoder interface {
	DecodeReqEvent(in []byte) (*schema.ReqEvent, error)

	EncodeReqEvent(in *schema.ReqEvent) ([]byte, error)

	DecodeReqQ(in []byte) (*schema.ReqQ, error)

	EncodeReqQ(in *schema.ReqQ) ([]byte, error)
}
