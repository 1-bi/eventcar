package eventcar

import (
	"context"
	"fmt"
	"github.com/1-bi/eventcar/schema"
	"github.com/1-bi/uerrors"
	"github.com/coreos/etcd/clientv3"
	"github.com/gogo/protobuf/proto"
	"log"
)

type QueueWatcher struct {
	client   *clientv3.Client
	eventKey string
	cbs      []Callback

	sCallback []SuccessCallback
	fCallback []FailureCallback
}

func NewQueueWatcher(client *clientv3.Client) *QueueWatcher {
	watcher := new(QueueWatcher)
	watcher.client = client
	return watcher
}

func (myself *QueueWatcher) SetEventKey(eventKey string) {
	myself.eventKey = eventKey
}

func (myself *QueueWatcher) SetCallbacks(newCbs ...Callback) {
	myself.cbs = newCbs

	myself.sCallback = make([]SuccessCallback, 0)
	myself.fCallback = make([]FailureCallback, 0)

	for _, cb := range newCbs {

		scb, ok := cb.(SuccessCallback)

		if ok {
			myself.sCallback = append(myself.sCallback, scb)
			continue
		}

		fsb, ok := cb.(FailureCallback)
		if ok {
			myself.fCallback = append(myself.fCallback, fsb)
			continue
		}

	}
}

func (myself *QueueWatcher) run() error {
	myself.watchNodeChange()

	return nil
}

func (myself *QueueWatcher) decode(msg []byte) (*schema.ResEvent, error) {
	resEvent := new(schema.ResEvent)

	err := proto.Unmarshal(msg, resEvent)

	if err != nil {
		return nil, err
	}

	return resEvent, nil

}

func (myself *QueueWatcher) doCallback(resEvent *schema.ResEvent) {

	fmt.Println(myself.sCallback)

	if schema.ResEvent_SUCCESS == resEvent.ResultType {
		for _, cb := range myself.sCallback {
			cb.Succeed(resEvent.GetResultBody())
		}
	} else if schema.ResEvent_FAILURE == resEvent.ResultType {
		resErr := resEvent.GetError()

		uErr := uerrors.NewCodeErrorWithPrefix(resErr.GetPrefix(), resErr.GetCode(), resErr.GetMsg())

		for _, cb := range myself.fCallback {

			cb.Fail(uErr)
		}
	}

}

func (myself *QueueWatcher) watchNodeChange() {

	ctx, cancel := context.WithCancel(context.Background())

	// --- watch message of node changed
	rch := myself.client.Watch(ctx, myself.eventKey)

	// --- block thread
	for wresp := range rch {

		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:

				// --- get the message ---

				// decode value
				resEvent, err := myself.decode(ev.Kv.Value)
				if err != nil {
					log.Println(err)
				}

				myself.doCallback(resEvent)
				cancel()

			case clientv3.EventTypeDelete:
				fmt.Printf("[%fixture] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				fmt.Println(string(ev.Kv.Key))
			}
		}
	}

	// ---- delete response content by key ----
	_, err := myself.client.Delete(context.Background(), myself.eventKey)
	if err != nil {
		log.Println(err)
	}

}
