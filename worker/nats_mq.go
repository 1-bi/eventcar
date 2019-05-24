package worker

import (
	"fmt"
	"github.com/1-bi/eventcar/schema"
	"github.com/1-bi/log-api"
	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/go-nats-streaming"
	"log"
)

// workerManager monitor all running services, if facade mode
type NatsMQWorker struct {
	natsConn   stan.Conn
	controlCh  chan int
	reqHandler func(req *schema.ReqQ)
	logger     logapi.Logger
}

func NewNatsMQWorker(natsConn stan.Conn) *NatsMQWorker {
	wm := new(NatsMQWorker)
	wm.natsConn = natsConn

	wm.logger = logapi.GetLogger("eventcar.NatsMQWorker")

	return wm
}

func (myself *NatsMQWorker) startSubscribe() (stan.Subscription, error) {

	sub, err := myself.natsConn.Subscribe("reqm", func(m *stan.Msg) {

		reqQ := new(schema.ReqQ)

		if err := proto.Unmarshal(m.Data, reqQ); err != nil {
			fmt.Println(err)
		}

		if myself.logger.IsDebugEnabled() {
			sb := logapi.NewStructBean()
			sb.LogInt64("reqId", reqQ.ReqId)
			sb.LogString("eventName", reqQ.Name)
			myself.logger.Debug("Receive req from message content ", sb)
		}

		// --- stop recevie message
		myself.controlCh <- CMD_PAUSE

		// --- add logic start  ---
		if myself.reqHandler != nil {
			myself.reqHandler(reqQ)
		}

		// ----add logic end ---

		// --- restart recevie message
		myself.controlCh <- CMD_RUN

	})

	if err != nil {
		return nil, err
	}

	return sub, err

}

func (myself *NatsMQWorker) RequestHandler(fn func(req *schema.ReqQ)) {
	myself.reqHandler = fn
}

// Stop send commond stop
func (myself *NatsMQWorker) Stop() {
	go func() {
		myself.controlCh <- CMD_STOP
	}()
}

func (myself *NatsMQWorker) Run() {

	myself.controlCh = make(chan int)

	// --- start  the run
	go func() {
		myself.controlCh <- CMD_RUN
	}()

	// --- close channel --
	defer close(myself.controlCh)

	// --- connect to message
	var sub stan.Subscription
	var err error

	for recSign := range myself.controlCh {

		if CMD_RUN == recSign {

			go func() {

				sub, err = myself.startSubscribe()
				if err != nil {
					log.Println(err)
				}
				err = nil
			}()

		} else if CMD_PAUSE == recSign {

			go func() {

				if sub != nil {

					err = sub.Unsubscribe()

					if err != nil {
						log.Println(err)
					}
				}

			}()

		} else if CMD_STOP == recSign {

			if sub != nil {
				err = sub.Close()

				if err != nil {
					log.Println(err)
				}

			}

			break
		}

	}

}
