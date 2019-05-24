package api

import "github.com/1-bi/uerrors"

type Callback interface {
}

type SuccessCallback interface {
	Callback

	Succeed(content []byte)
}

type FailureCallback interface {
	Callback

	Fail(err uerrors.CodeError)
}
