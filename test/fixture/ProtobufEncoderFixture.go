package fixture

import (
	"github.com/1-bi/eventcar/encoder"
	"github.com/1-bi/eventcar/encoder/protobuf"
	"github.com/1-bi/eventcar/schema"
	"github.com/smartystreets/gunit"
)

// ProtobufEncoderFixture Test structure framework define
type ProtobufEncoderFixture struct {
	*gunit.Fixture

	msgEncoder encoder.Encoder
}

func (myself *ProtobufEncoderFixture) Setup() {
	// --- create client ---
	myself.msgEncoder = new(protobuf.EncoderImpl)
}

func (myself *ProtobufEncoderFixture) Teardown() {

	// --- stop server

}

func (myself *ProtobufEncoderFixture) Test_ReqEvent() {

	msg := new(schema.ReqEvent)
	msg.ReqId = 1000
	msg.Name = "reqQTestName"
	msg.MsgBody = []byte("test body")

	var err error

	byteMsg, err := myself.msgEncoder.EncodeReqEvent(msg)

	if err != nil {
		myself.Println("encode msg message fail. ")
		return
	}

	recMsg, err := myself.msgEncoder.DecodeReqEvent(byteMsg)
	if err != nil {
		myself.Println("dencode msg message fail. ")
		return
	}

	myself.AssertEqual(recMsg.ReqId, msg.ReqId)
	myself.AssertEqual(recMsg.Name, msg.Name)
	myself.AssertEqual(string(recMsg.MsgBody), string(msg.MsgBody))

}

func (myself *ProtobufEncoderFixture) Test_ReqQ() {

	reqQ := new(schema.ReqQ)
	reqQ.ReqId = 1001
	reqQ.Name = "reqQTestName"
	reqQ.ComType = schema.ReqQ_QUE

	var err error

	byteReqQ, err := myself.msgEncoder.EncodeReqQ(reqQ)

	if err != nil {
		myself.Println("encode reqQ message fail. ")
		return
	}

	recReqQ, err := myself.msgEncoder.DecodeReqQ(byteReqQ)
	if err != nil {
		myself.Println("dencode reqQ message fail. ")
		return
	}

	myself.AssertEqual(recReqQ.ReqId, reqQ.ReqId)
	myself.AssertEqual(recReqQ.Name, reqQ.Name)
	myself.AssertEqual(recReqQ.ComType, reqQ.ComType)

}
