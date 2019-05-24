package eventcar

import (
	"fmt"
	"github.com/1-bi/eventcar/schema"
	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/go-nats-streaming"
	"time"
)

const (
	IDEL    = 0
	RUNNING = 1
	PAUSE   = 2
	STOP    = 3
)

// workerManager monitor all running services, if facade mode
type WorkerManager struct {
	natsConn stan.Conn
}

func NewWorkerManager(natsConn stan.Conn) *WorkerManager {
	wm := new(WorkerManager)
	wm.natsConn = natsConn
	return wm
}

/*

func (myself *WorkerManager) sign() <- chan int{
	return <- 1
}
*/

func (myself *WorkerManager) startSubscribe() (stan.Subscription, error) {

	sub, err := myself.natsConn.Subscribe("reqm", func(m *stan.Msg) {

		reqQ := new(schema.ReqQ)

		if err := proto.Unmarshal(m.Data, reqQ); err != nil {
			fmt.Println(err)
		}

	})

	if err != nil {
		return nil, err
	}

	return sub, err

}

func (myself *WorkerManager) Run() {

	controlCh := make(chan int)
	go func() {

		time.Sleep(2 * time.Second)

		controlCh <- RUNNING

	}()

	go func() {

		time.Sleep(4 * time.Second)

		controlCh <- PAUSE

	}()

	go func() {

		time.Sleep(6 * time.Second)

		controlCh <- STOP

	}()

	// --- connect to message

	for recSign := range controlCh {

		if RUNNING == recSign {
			go func() {
				//controlCh <- 2
				fmt.Println("---- call running  ----")

			}()
		} else if PAUSE == recSign {

			go func() {
				//controlCh <- 3
				fmt.Println("---- call pause  ----")
			}()

		} else if STOP == recSign {

			go func() {
				//controlCh <- 3
				fmt.Println("---- call stop   ----")
			}()

			break
		}

	}

}
