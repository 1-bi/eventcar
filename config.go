package eventcar

import (
	"fmt"
	"github.com/1-bi/log-api"
	"github.com/coreos/etcd/clientv3"
	"io/ioutil"
)

// Config
type Config struct {
	_etcdConfig clientv3.Config

	// nats server connect
	_natsHost []string

	_natsClusterId string

	_natsClientId string

	// --- define default value --
	nodeIdPath string

	nodeRoles []string

	nodeNum int64

	_agentNodeId string
}

func NewConfig() *Config {
	var conf = new(Config)

	// --- set default value
	conf.nodeIdPath = "/etc/servicebus/agent/node_id"
	conf._natsHost = []string{"nats://localhost:4222"}
	conf._natsClusterId = "test-cluster"
	conf._natsClientId = "test-client"
	conf.nodeNum = 1
	return conf
}

func (myself *Config) SetNatsHost(hosts []string) {
	myself._natsHost = hosts
}

func (myself *Config) SetNatsClusterId(newClusterId string) {
	myself._natsClusterId = newClusterId
}

func (myself *Config) SetNatsClientId(newClientId string) {
	myself._natsClientId = newClientId
}

func (myself *Config) SetNodeRoles(newRoles []string) {
	myself.nodeRoles = newRoles
}

func (myself *Config) CheckBeforeStart() error {

	var confErr = myself.checkNodeId()
	if confErr != nil {
		return confErr
	}

	return nil
}

func (myself *Config) checkNodeId() error {

	// --- check file exist
	var existed, err = pathExists(myself.nodeIdPath)

	if err != nil {
		fmt.Println(err)
	}

	var nodeId = ""
	if !existed {
		nodeId = createNodeIdFile(myself.nodeIdPath, myself.nodeNum)
	} else {

		// --- read node id to file ---
		b, err := ioutil.ReadFile(myself.nodeIdPath)
		if err != nil {
			structBean := logapi.NewStructBean()
			structBean.LogString("fileName", myself.nodeIdPath)
			logapi.GetLogger("serviebus.Agent.checkNodeId").Debug("Get node id from file. ", structBean)
		}
		nodeId = string(b)
	}

	//  create the agentNodeId
	structBean := logapi.NewStructBean()
	structBean.LogString("nodeId", nodeId)
	logapi.GetLogger("serviebus.Agent.checkNodeId").Debug("Display the current node id. ", structBean)

	myself._agentNodeId = nodeId

	return nil

}

// SetRegisterServer define restier servers handle
func (myself *Config) SetRegisterServer(config clientv3.Config) {
	myself._etcdConfig = config
}

func (myself *Config) SetNodeIdFilePath(nodeIdFilePath string) {
	myself.nodeIdPath = nodeIdFilePath
}

func (myself *Config) SetSettingDir(paths ...string) {

}
