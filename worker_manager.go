package eventcar

import (
	"fmt"
	"github.com/1-bi/eventcar/schema"
	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/go-nats-streaming"
	"log"
	"time"
)

const (
	IDEL      = 0
	CMD_RUN   = 1
	CMD_PAUSE = 2
	CMD_STOP  = 3
)

// workerManager monitor all running services, if facade mode
type WorkerManager struct {
	natsConn  stan.Conn
	controlCh chan int

	reqHandler func(req *schema.ReqQ)
}

func NewWorkerManager(natsConn stan.Conn) *WorkerManager {
	wm := new(WorkerManager)
	wm.natsConn = natsConn
	return wm
}

func (myself *WorkerManager) startSubscribe() (stan.Subscription, error) {

	sub, err := myself.natsConn.Subscribe("reqm", func(m *stan.Msg) {

		reqQ := new(schema.ReqQ)

		if err := proto.Unmarshal(m.Data, reqQ); err != nil {
			fmt.Println(err)
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

func (myself *WorkerManager) RequestHandler(fn func(req *schema.ReqQ)) {
	myself.reqHandler = fn
}

func (myself *WorkerManager) Run() {

	myself.controlCh = make(chan int)
	go func() {

		time.Sleep(2 * time.Second)

		myself.controlCh <- CMD_RUN

	}()

	go func() {

		time.Sleep(4 * time.Second)

		myself.controlCh <- CMD_PAUSE

	}()

	go func() {

		//time.Sleep(6 * time.Second)

		//myself.controlCh <- CMD_STOP

	}()

	// --- connect to message
	var sub stan.Subscription
	var err error

	for recSign := range myself.controlCh {

		if CMD_RUN == recSign {
			go func() {
				//controlCh <- 2

				sub, err = myself.startSubscribe()
				if err != nil {
					log.Println(err)
				}
				err = nil
			}()
		} else if CMD_PAUSE == recSign {

			go func() {

				if sub != nil {
					fmt.Println(" -----555 ---- ")

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

			fmt.Println("-=--00-")

			break
		}

	}

}
