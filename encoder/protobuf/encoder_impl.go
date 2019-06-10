package protobuf

import (
	"encoding/base64"
	"github.com/1-bi/eventcar/schema"
	"github.com/gogo/protobuf/proto"
)

type ProtobufWithBase64MsgEncoder struct {
}

func (myself *ProtobufWithBase64MsgEncoder) DecodeReqEvent(in []byte) (*schema.ReqEvent, error) {
	msg := new(schema.ReqEvent)

	var bmsg []byte
	// use base64 decode
	bmsg, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		return nil, err
	}

	if err = proto.Unmarshal(bmsg, msg); err != nil {
		return nil, err
	}
	return msg, nil

}

func (myself *ProtobufWithBase64MsgEncoder) DecodeReqQ(in []byte) (*schema.ReqQ, error) {
	msg := new(schema.ReqQ)

	var bmsg []byte
	// use base64 decode
	bmsg, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(bmsg, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (myself *ProtobufWithBase64MsgEncoder) DecodeResEvent(in []byte) (*schema.ResEvent, error) {
	msg := new(schema.ResEvent)

	var bmsg []byte
	// use base64 decode
	bmsg, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(bmsg, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (myself *ProtobufWithBase64MsgEncoder) EncodeReqEvent(in *schema.ReqEvent) ([]byte, error) {
	var msg []byte
	msg, err := proto.Marshal(in)

	if err != nil {
		return nil, err
	}

	// use base64 encode
	out := []byte(base64.StdEncoding.EncodeToString(msg))

	return out, nil
}

func (myself *ProtobufWithBase64MsgEncoder) EncodeReqQ(in *schema.ReqQ) ([]byte, error) {
	var bmsg []byte
	bmsg, err := proto.Marshal(in)

	if err != nil {
		return nil, err
	}

	// use base64 encode
	out := []byte(base64.StdEncoding.EncodeToString(bmsg))

	return out, nil
}

func (myself *ProtobufWithBase64MsgEncoder) EncodeResEvent(in *schema.ResEvent) ([]byte, error) {
	var bmsg []byte
	bmsg, err := proto.Marshal(in)

	if err != nil {
		return nil, err
	}

	// use base64 encode
	out := []byte(base64.StdEncoding.EncodeToString(bmsg))

	return out, nil
}
