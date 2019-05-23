package fixture

import (
	"fmt"
	"github.com/1-bi/eventcar/etcd"
	"github.com/1-bi/eventcar/schema"
	"github.com/1-bi/log-api"
	"github.com/bwmarrin/snowflake"
	"github.com/coreos/etcd/clientv3"
	"github.com/gogo/protobuf/proto"
	"github.com/smartystreets/gunit"
	"log"
	"strconv"
	"strings"
	"time"
)

// EtcdServiceOperationsFixture Test structure framework define
type EtcdServiceOperationsFixture struct {
	*gunit.Fixture

	servOper *etcd.EtcdServiceOperations
}

// SetupAgent
func (myself *EtcdServiceOperationsFixture) Setup() {
	// --- config

	var cli *clientv3.Client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		log.Println(err)
	}

	// --- create agent ---
	myself.servOper = etcd.NewEtcdServiceOperations(cli, nil)
}

func (myself *EtcdServiceOperationsFixture) Teardown() {

}

func (myself *EtcdServiceOperationsFixture) Test_GetMessage() {

	req, err := myself.servOper.GetMesssage("reqm/1129979698965647360")

	if err != nil {
		fmt.Println(err)
	}

	// 解码
	unmaReqEvent := new(schema.ReqEvent)
	if err := proto.Unmarshal(req, unmaReqEvent); err != nil {
		log.Fatal("failed to unmarshal: ", err)
	}

	fmt.Println(unmaReqEvent.ReqId)
	fmt.Println(unmaReqEvent.Name)
	fmt.Println(string(unmaReqEvent.MsgBody))

}

func (myself *EtcdServiceOperationsFixture) node(nodeNum int64) *snowflake.Node {

	node, err := snowflake.NewNode(1)
	if err != nil {
		logapi.GetLogger("start").Fatal(err.Error(), nil)
	}

	return node

}

func (myself *EtcdServiceOperationsFixture) Test_Message() {

	// --- set message ---
	// serialization runtimeArgs
	reqEvent := new(schema.ReqEvent)

	reqEvent.ReqId = myself.node(1).Generate().Int64()
	reqEvent.Name = "test.event.case1"
	reqEvent.MsgBody = []byte("hello message887")

	// --- sent msg body ---
	var reqMsg []byte
	reqMsg, err := proto.Marshal(reqEvent)

	if err != nil {
		fmt.Println(err)
	}

	// get minion runinng node

	// --- key ---
	var key = strings.Join([]string{"reqm", strconv.FormatInt(reqEvent.ReqId, 10)}, "/")

	// --- set the key value ---
	err = myself.servOper.SetMessage(key, reqMsg)

	// --- get error -------

	msgBody, err := myself.servOper.GetMesssage(key)

	recReqMsg := new(schema.ReqEvent)
	err = proto.Unmarshal(msgBody, recReqMsg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("-----------------")
	fmt.Println(recReqMsg)
	fmt.Println("================")

}
