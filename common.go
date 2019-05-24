/**
 * =========================================================================== *
 * define all interface in this file
 * =========================================================================== *
 */
package eventcar

import (
	"github.com/1-bi/log-api"
	"github.com/bwmarrin/snowflake"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/**
 * --------------------------------- *
 * constant region
 * --------------------------------- *
 */
const (
	PREFIX = "eventcar"

	STATUS_DONE    int8 = 5
	STATUS_RUNNING int8 = 1
	STATUS_CANNEL  int8 = 3

	REQ_TIMEOUT = 3 * time.Second

	// ---- value is for  interface "FutureReturnResult"
	NONE         int8 = 0
	ALL_COMPLETE int8 = 1
	ANY_ERRORS   int8 = 2
	ALL_ERRORS   int8 = 3

	// worker command  define
	IDEL      = 0
	CMD_RUN   = 1
	CMD_PAUSE = 2
	CMD_STOP  = 3
)

//  --- check the node id propertes
func createNodeIdFile(path string, nodeNum int64) string {

	// --- get the path ---
	lastFolder := strings.LastIndex(path, "/")
	parentFolder := path[0:lastFolder]

	existed, err := pathExists(parentFolder)

	if err != nil {
		logapi.GetLogger("serviebus.common").Debug(err.Error(), nil)
	}

	if !existed {
		os.MkdirAll(parentFolder, os.ModePerm)
	}

	// --- generate file and id
	f, err := os.Create(path)

	defer f.Close()

	if err != nil {
		logapi.GetLogger("serviebus.common").Debug(err.Error(), nil)
	}

	node, err := snowflake.NewNode(nodeNum)

	if err != nil {
		logapi.GetLogger("serviebus.common").Debug(err.Error(), nil)
	}

	// Generate a snowflake ID.
	id := node.Generate()
	locIndex, err := f.WriteString(id.String())

	f.Sync()

	logapi.GetLogger("serviebus.common").Debug("Location index is "+strconv.Itoa(locIndex), nil)

	var nodeId = id.String()
	return nodeId
}

// pathExists file exist
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
 * --------------------------------- *
 * function region
 * --------------------------------- *
 */
func FunctypeInObject(object interface{}) string {
	objectType := reflect.TypeOf(object)
	function := objectType.Elem().String()
	return function
}

/**

 */
type EventMapping struct {
	EventName string

	// binding reference instance
	Ref interface{}

	/**
	 * define event name
	 */
	FunctionName string
}

/**
 * define event delegate
 */
type EventsDelegate struct {
	eventholder       map[string]*EventMapping
	ctlHandlerMapping map[string]map[string]reflect.Value
}

func (this *EventsDelegate) AddEvent(event string, handler interface{}, functionName string) {

	eventMapping := new(EventMapping)
	eventMapping.EventName = event
	eventMapping.Ref = handler
	eventMapping.FunctionName = functionName

	this.eventholder[event] = eventMapping
}

func (this *EventsDelegate) AllEventsPredefined() map[string]*EventMapping {
	return this.eventholder
}

/**
 * get all method by define ---
 */
func (this *EventsDelegate) AllEventMethods() map[string]reflect.Value {

	eventMethods := make(map[string]reflect.Value)

	for serviceId, eventMapping := range this.eventholder {

		structMethods := this.getAllMethodMapByStruct(eventMapping.Ref)

		eventMethods[serviceId] = structMethods[eventMapping.FunctionName]

	}

	return eventMethods

}

/**
 * get method by method name
 */
func (this *EventsDelegate) getAllMethodMapByStruct(structType interface{}) map[string]reflect.Value {

	t := reflect.TypeOf(structType)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	e := t.String()

	mapSize := len(this.ctlHandlerMapping[e])

	if mapSize == 0 {
		// --- create new mapping ---
		this.ctlHandlerMapping[e] = make(map[string]reflect.Value, 0)

		// --- scan all method in this object ----
		objInstRef := reflect.ValueOf(structType)
		typ := objInstRef.Type()

		for i := 0; i < objInstRef.NumMethod(); i++ {

			methodInst := objInstRef.Method(i)
			methodName := typ.Method(i).Name

			this.ctlHandlerMapping[e][methodName] = methodInst
		}

	}

	return this.ctlHandlerMapping[e]

}

func NewEventsDelegate() *EventsDelegate {
	delegate := new(EventsDelegate)
	delegate.eventholder = make(map[string]*EventMapping, 0)
	delegate.ctlHandlerMapping = make(map[string]map[string]reflect.Value, 0)
	return delegate
}
