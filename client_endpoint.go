package eventcar

import (
	"github.com/1-bi/eventcar/api"
	"github.com/1-bi/eventcar/encoder"
	"github.com/1-bi/eventcar/encoder/protobuf"
	"github.com/1-bi/eventcar/etcd"
	"github.com/1-bi/eventcar/schema"
	"github.com/1-bi/eventcar/worker"
	"github.com/1-bi/log-api"
	"github.com/bwmarrin/snowflake"
	"github.com/coreos/etcd/clientv3"
	"github.com/nats-io/go-nats-streaming"
	"strconv"
	"strings"
)

type ClientEndpoint struct {
	nodeGenerater *snowflake.Node
	natsConn      stan.Conn
	regListeners  map[string]func(api.ReqMsgContext)
	etcdServOpt   *etcd.EtcdServiceOperations
	respWatcher   *worker.EtcdWatcherWorker

	msgEncoder encoder.MsgEncoder
}

// NewAgent check agent
func NewClient(nodeNum int64, natsConn stan.Conn, cli *clientv3.Client) *ClientEndpoint {

	var agentLogger = logapi.GetLogger("eventcar.Client")
	agentLogger.Info("base", nil)

	var client = new(ClientEndpoint)

	// define the client node number
	node, err := snowflake.NewNode(nodeNum)
	if err != nil {
		logapi.GetLogger("NewClient").Fatal(err.Error(), nil)
	} else {
		client.nodeGenerater = node
	}

	// set nats mq connection
	client.natsConn = natsConn

	// --- connect client ---
	client.regListeners = make(map[string]func(api.ReqMsgContext))

	if cli != nil {
		servOptsMap := make(map[string]string)
		client.etcdServOpt = etcd.NewEtcdServiceOperations(cli, servOptsMap)
	}

	client.msgEncoder = new(protobuf.ProtobufWithBase64MsgEncoder)

	// start describe queue to nats
	return client

}

// On implement event name
func (myself *ClientEndpoint) On(eventName string, fn func(api.ReqMsgContext)) error {

	myself.regListeners[eventName] = fn

	return nil
}

// FireByQueue call by event name and define callback
func (myself *ClientEndpoint) FireByQueue(eventName string, msgBody []byte, callback ...api.Callback) error {
	// --- send message to  nats ---

	// serialization runtimeArgs
	reqEvent := new(schema.ReqEvent)

	reqEvent.ReqId = myself.nodeGenerater.Generate().Int64()
	reqEvent.Name = eventName
	reqEvent.MsgBody = msgBody

	// --- sent msg body ---
	var reqMsg []byte

	reqMsg, err := myself.msgEncoder.EncodeReqEvent(reqEvent)

	if err != nil {
		return err
	}

	reqQ := new(schema.ReqQ)
	reqQ.ReqId = reqEvent.ReqId
	reqQ.Name = reqEvent.Name
	reqQ.ComType = schema.ReqQ_QUE

	var req []byte
	req, err = myself.msgEncoder.EncodeReqQ(reqQ)

	if err != nil {
		return err
	}

	// get minion runinng node

	// --- key ---
	var key = strings.Join([]string{"reqm", strconv.FormatInt(reqEvent.ReqId, 10)}, "/")

	// --- set the key value ---
	if myself.etcdServOpt != nil {

		err = myself.etcdServOpt.SetMessage(key, reqMsg)

		if err != nil {
			return err
		}

		// ---- start watcher listener ---
		// --- connect client ---

		go func() {
			myself.respWatcher.WatcherQueueResp(strings.Join([]string{"resm", strconv.FormatInt(reqEvent.ReqId, 10)}, "/"), callback...)
		}()

	}

	err = myself.natsConn.Publish("reqm", req)
	return err

}

func (myself *ClientEndpoint) FireByPublish(eventName string, msgBody []byte, callback ...api.Callback) error {
	// --- send message to  nats ---

	// serialization runtimeArgs
	reqEvent := new(schema.ReqEvent)

	reqEvent.ReqId = myself.nodeGenerater.Generate().Int64()
	reqEvent.Name = eventName
	reqEvent.MsgBody = msgBody

	// --- sent msg body ---+
	var reqMsg []byte

	reqMsg, err := myself.msgEncoder.EncodeReqEvent(reqEvent)

	if err != nil {
		return err
	}

	reqQ := new(schema.ReqQ)
	reqQ.ReqId = reqEvent.ReqId
	reqQ.Name = reqEvent.Name
	reqQ.ComType = schema.ReqQ_SUB

	var req []byte
	req, err = myself.msgEncoder.EncodeReqQ(reqQ)

	if err != nil {
		return err
	}

	// get minion runinng node
	if myself.etcdServOpt != nil {

		nodes, err := myself.etcdServOpt.GetAllNodeIds("minion")
		if err != nil {
			return err
		}

		// pub the message to content
		for _, node := range nodes {

			// --- key ---
			var key = strings.Join([]string{"reqm", strconv.FormatInt(reqEvent.ReqId, 10), "mi=" + node}, "/")

			// --- set the key value ---
			err = myself.etcdServOpt.SetMessage(key, reqMsg)

			if err != nil {
				break
			}

		}

	}

	err = myself.natsConn.Publish("reqm", req)
	if err != nil {
		return err
	}

	// --- start new watcher ---

	return nil
}
