package fixture

import (
	"github.com/1-bi/eventcar"
	"github.com/1-bi/log-api"
	"github.com/1-bi/log-zap"
	"github.com/1-bi/log-zap/appender"
	zaplayout "github.com/1-bi/log-zap/layout"
	"github.com/coreos/etcd/clientv3"
	"github.com/nats-io/go-nats-streaming"
	"github.com/smartystreets/gunit"
	"log"
	"strings"
	"time"
)

// ClientFixture Test structure framework define
type ClientFixture struct {
	*gunit.Fixture

	client *eventcar.ClientEndpoint
}

// SetupAgent
func (myself *ClientFixture) Setup() {
	// --- create client ---
	myself.prepareLogSetting()

	_, err := myself.prepareConfig()

	if err != nil {
		logapi.GetLogger("client_object ").Fatal(err.Error(), nil)
	}

	// --- create client --
	natsHost := []string{"nats://localhost:4222"}
	natsServer := strings.Join(natsHost, ",")
	natsConn, err := stan.Connect("test-cluster", "clienttest", stan.NatsURL(natsServer))
	if err != nil {
		structBean := logapi.NewStructBean()
		structBean.LogStringArray("nats.server", natsHost)
		logapi.GetLogger("serviebus.Start").Fatal("Connect nats server fail.", structBean)
		return
	}

	// dstart client to complent
	myself.client = eventcar.NewClient(1, natsConn, nil)

}

func (myself *ClientFixture) Teardown() {

	// --- stop server

}

func (myself *ClientFixture) prepareLogSetting() {

	// --- construct layout ---
	var jsonLayout = zaplayout.NewJsonLayout()
	//jsonLayout.SetTimeFormat("2006-01-02 15:04:05")
	jsonLayout.SetTimeFormat("2006-01-02 15:04:05 +0800 CST")
	//fmt.Println( time.Now().Location() )

	// --- set appender
	var consoleAppender = appender.NewConsoleAppender(jsonLayout)

	var mainOpt = logzap.NewLoggerOption()
	mainOpt.SetLevel("debug")
	mainOpt.AddAppender(consoleAppender)

	var agentOpt = logzap.NewLoggerOption()
	agentOpt.SetLoggerPattern("servicebus")
	agentOpt.SetLevel("warn")
	agentOpt.AddAppender(consoleAppender)

	var implReg = new(logzap.ZapFactoryRegister)

	_, err := logapi.RegisterLoggerFactory(implReg, mainOpt, agentOpt)

	if err != nil {
		log.Fatal(err)
	}
}

func (myself *ClientFixture) prepareConfig() (*eventcar.Config, error) {

	var conf = eventcar.NewConfig()

	conf.SetRegisterServer(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 2 * time.Second,
	})

	conf.SetNodeRoles([]string{
		"master", "minion",
	})

	conf.SetNatsHost([]string{"nats://localhost:4222"})

	var configErr = conf.CheckBeforeStart()

	return conf, configErr

}

func (myself *ClientFixture) Test_Publish_Queue() {

	myself.client.FireByQueue("client.test.case1", []byte("hello testabdi"))
}
