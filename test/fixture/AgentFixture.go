package fixture

import (
	"fmt"
	"github.com/1-bi/eventcar"
	"github.com/1-bi/eventcar/api"
	"github.com/1-bi/log-api"
	"github.com/1-bi/log-zap"
	"github.com/1-bi/log-zap/appender"
	zaplayout "github.com/1-bi/log-zap/layout"
	"github.com/coreos/etcd/clientv3"
	"github.com/smartystreets/gunit"
	"log"
	"time"
)

// AgentFixture Test structure framework define
type AgentFixture struct {
	*gunit.Fixture

	agent *eventcar.Agent
}

// SetupAgent
func (myself *AgentFixture) Setup() {
	// --- create agent ---
	myself.prepareLogSetting()

	conf, err := myself.prepareConfig()

	if err != nil {
		logapi.GetLogger("agent_basecase1").Fatal(err.Error(), nil)
	}

	// dstart agent to complent
	myself.agent = eventcar.NewAgent(conf)

	myself.agent.Start()
}

func (myself *AgentFixture) Teardown() {

	// --- stop server
	myself.agent.Stop()

}

func (myself *AgentFixture) prepareLogSetting() {

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

func (myself *AgentFixture) prepareConfig() (*eventcar.Config, error) {

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

func (myself *AgentFixture) Test_Subscribe_Publish() {

	myself.agent.On("agent.test.case1", func(ctx api.ReqMsgContext) {

		reqMsg := string(ctx.GetMsgRawBody())

		fmt.Println("receive message ")
		fmt.Println(reqMsg)

		ctx.GetResResult().Complete([]byte("response message "))

	})

	time.Sleep(6 * time.Second)

	myself.agent.FireByQueue("agent.test.case1", []byte("hello testabdi"))
}
