package eventcar

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
)

type AgentServiceWatchService struct {
	nodeId  string
	client  *clientv3.Client
	_prefix string
}

func NewAgentWatchService(nodeId string, cli *clientv3.Client) *AgentServiceWatchService {

	// --- create  AgentService ---
	var agentWatchServ = new(AgentServiceWatchService)
	agentWatchServ.client = cli
	agentWatchServ.nodeId = nodeId

	agentWatchServ._prefix = "nodes/"
	return agentWatchServ
}

func (myself *AgentServiceWatchService) Start() error {

	// go and define object

	return nil
}

func (myself *AgentServiceWatchService) Stop() error {
	return nil
}

func (myself *AgentServiceWatchService) watchNodeChange() {

	// --- watch message of node changed
	rch := myself.client.Watch(context.Background(), myself._prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				fmt.Printf("[%fixture] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)

			case clientv3.EventTypeDelete:
				fmt.Printf("[%fixture] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				fmt.Println(string(ev.Kv.Key))
			}
		}
	}

}

// GetRuntimeNodeIds
func (myself *AgentServiceWatchService) GetRuntimeNodeIds() []string {
	return nil
}
