/**
 * =========================================================================== *
 * define all interface in this file
 * =========================================================================== *
 */
package api

import (
	"github.com/1-bi/uerrors"
	"time"
)

// ServiceEvent ServiceEvent is define full message
type ServiceEvent interface {

	/**
	 * call one service handle
	 */
	On(eventName string, fn func(ReqMsgContext)) error

	/**
	 * fire service in synchronous mode
	 */
	FireSyncService(serviceId string, runtimeArgs interface{}, timeout time.Duration, fn func(FutureReturnResult, uerrors.CodeError))

	/**
	 * fire service in asynchronous mode
	 */
	FireService(serviceId string, runtimeArgs interface{}) error
}

// --- create bus context ----
type EventbusContext interface {
	GetRequestData() interface{}

	/**
	 * Get the result inteface
	 */
	GetResult() Result

	/**
	 * support service event
	 */
	GetServiceEvent() ServiceEvent
}

/**
 * define base result
 */
type Result interface {
	/**
	 *	complement object
	 */
	Complete(refobj []byte)

	/**
	 *  add the result
	 */
	Fail(err uerrors.CodeError)
}

/**
 * contruct service event handler
 * @deplecated this object is not to use
 */
type ReqMsgContext interface {

	/**
	 * define request body object
	 */
	//ConvertRequestBody(func() (reqData interface{}))

	/**
	 * define process handler
	 */
	//Process(func(bc EventbusContext) uerrors.CodeError)

	// GetMsgRawBody get request message from body
	GetMsgRawBody() []byte

	// get the response result context
	GetResResult() Result
}

/**
 * return the service event futrue
 */
type Future interface {

	/**
	 * wait the event complete
	 */
	Await() uerrors.CodeError

	/**
	 * get the result object after await
	 */
	GetResult() (FutureReturnResult, uerrors.CodeError)
}

/**
 *  This interface will map the futurn result handle
 */
type FutureReturnResult interface {

	/**
	 * define,  ALL_COMPLETE ,  ANY_ERRORS , ALL_ERRORS
	 */
	State() int8

	/**
	 *  return all error from service event running
	 */
	Errors(procName string) uerrors.CodeError

	/**
	 * rturn all return Results
	 */
	ReturnResults(procName string, inReturn interface{}) uerrors.CodeError

	/**
	 * get the first error directly
	 */
	Error() uerrors.CodeError

	/**
	 * get the first result directly
	 */
	ReturnResult(inReturn interface{}) uerrors.CodeError
}

/**
 * define root type decoder
 */
type RootTypeDecoder interface {
	ProvideRootRef() interface{}
}
