package protobuf

import (
	"github.com/1-bi/eventcar/schema"
	"github.com/gogo/protobuf/proto"
)

type EncoderImpl struct {
}

func (myself *EncoderImpl) DecodeReqEvent(in []byte) (*schema.ReqEvent, error) {
	msg := new(schema.ReqEvent)
	if err := proto.Unmarshal(in, msg); err != nil {
		return nil, err
	}
	return msg, nil

}

func (myself *EncoderImpl) EncodeReqEvent(in *schema.ReqEvent) ([]byte, error) {
	var msg []byte
	msg, err := proto.Marshal(in)

	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (myself *EncoderImpl) DecodeReqQ(in []byte) (*schema.ReqQ, error) {
	msg := new(schema.ReqQ)
	if err := proto.Unmarshal(in, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (myself *EncoderImpl) EncodeReqQ(in *schema.ReqQ) ([]byte, error) {
	var req []byte
	req, err := proto.Marshal(in)

	if err != nil {
		return nil, err
	}
	return req, nil
}
