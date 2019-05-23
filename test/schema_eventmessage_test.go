package test

import (
	"fmt"
	"github.com/1-bi/eventcar/schema"
	"github.com/gogo/protobuf/proto"
	"log"
	"testing"
)

/**
 * defined publish message
 */
func Test_Schema_eventmessage(t *testing.T) {

	reqQ := new(schema.ReqQ)
	reqQ.ReqId = 123433050033
	reqQ.Name = "req message"

	out, err := proto.Marshal(reqQ)
	if err != nil {
		log.Fatal("failed to marshal: ", err)
	}

	// 解码
	unmaReqQ := new(schema.ReqQ)
	if err := proto.Unmarshal(out, unmaReqQ); err != nil {
		log.Fatal("failed to unmarshal: ", err)
	}

	fmt.Println(reqQ)
	fmt.Println(unmaReqQ)
	fmt.Println(string(out))

}
