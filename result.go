package eventcar

import (
	"github.com/1-bi/eventcar/schema"
	"github.com/1-bi/uerrors"
)

type embeddedResult struct {
	respResultType schema.ResEvent_ResultType

	respContent []byte

	respErr uerrors.CodeError

	req *schema.ReqQ
}

func (myself *embeddedResult) Complete(refobj []byte) {

	myself.respResultType = schema.ResEvent_SUCCESS
	myself.respContent = refobj

}

func (myself *embeddedResult) Fail(err uerrors.CodeError) {

	myself.respResultType = schema.ResEvent_FAILURE
	myself.respErr = err

}

func (myself *embeddedResult) ConvertRepResult() *schema.ResEvent {
	resEvent := new(schema.ResEvent)
	resEvent.ReqId = myself.req.ReqId
	resEvent.Name = myself.req.Name

	resEvent.ResultType = myself.respResultType

	if schema.ResEvent_SUCCESS == resEvent.ResultType {
		resEvent.ResultBody = myself.respContent
	} else if schema.ResEvent_FAILURE == resEvent.ResultType {

		codeErr := new(schema.CodeError)
		codeErr.Code = myself.respErr.Code()
		codeErr.Prefix = myself.respErr.Prefix()
		codeErr.Msg = myself.respErr.MsgBody()
		resEvent.Error = codeErr

	}

	return resEvent
}
