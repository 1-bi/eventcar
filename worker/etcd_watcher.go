package worker

import (
	"github.com/1-bi/eventcar/api"
	"github.com/coreos/etcd/clientv3"
	"log"
)

type EtcdWatcherWorker struct {
	nodeId  string
	client  *clientv3.Client
	_prefix string

	controlCh chan int

	queueWatcher *QueueWatcher
}

func NewEtcdWatcherWorker(nodeId string, cli *clientv3.Client) *EtcdWatcherWorker {

	// --- create  AgentService ---
	var agentWatchServ = new(EtcdWatcherWorker)
	agentWatchServ.client = cli
	agentWatchServ.nodeId = nodeId

	agentWatchServ._prefix = "nodes/"
	return agentWatchServ
}

func (myself *EtcdWatcherWorker) Start() error {

	// go and define object
	myself.controlCh = make(chan int)

	// --- close channel --
	defer close(myself.controlCh)

	// --- connect to message
	var err error

	for recSign := range myself.controlCh {

		if CMD_RUN == recSign {

			go func() {

				// --- stop accept any watcher ---

				err = myself.queueWatcher.run()

				if err != nil {
					log.Println(err)
				}

				myself.queueWatcher = nil

				err = nil
			}()

		} else if CMD_PAUSE == recSign {

			go func() {

			}()

		} else if CMD_STOP == recSign {

			break
		}

	}
	return nil
}

func (myself *EtcdWatcherWorker) Stop() error {
	myself.controlCh <- CMD_STOP

	return nil
}

func (myself *EtcdWatcherWorker) WatcherQueueResp(key string, newCbs ...api.Callback) {

	myself.queueWatcher = NewQueueWatcher(myself.client)
	myself.queueWatcher.SetEventKey(key)
	myself.queueWatcher.SetCallbacks(newCbs...)

	myself.controlCh <- CMD_RUN
}

// GetRuntimeNodeIds
func (myself *EtcdWatcherWorker) GetRuntimeNodeIds() []string {
	return nil
}
