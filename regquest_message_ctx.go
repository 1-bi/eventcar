package eventcar

import (
	"github.com/1-bi/eventcar/api"
	"github.com/1-bi/eventcar/schema"
)

type embeddedReqMsgContext struct {
	rawMsgBody []byte

	resResult *embeddedResult
}

func newEmbeddedReqMsgContext(req *schema.ReqQ) *embeddedReqMsgContext {
	var ctx = new(embeddedReqMsgContext)

	var result = new(embeddedResult)
	ctx.resResult = result
	result.req = req

	return ctx
}

func (myself *embeddedReqMsgContext) setMsgRawBody(body []byte) {
	myself.rawMsgBody = body
}

func (myself *embeddedReqMsgContext) GetMsgRawBody() []byte {
	return myself.rawMsgBody
}

func (myself *embeddedReqMsgContext) GetResResult() api.Result {
	return myself.resResult
}
