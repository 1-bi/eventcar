package main

import (
	"fmt"
	"github.com/1-bi/eventcar"
	"github.com/1-bi/log-api"
	"github.com/1-bi/log-zap"
	"github.com/1-bi/log-zap/appender"
	zaplayout "github.com/1-bi/log-zap/layout"
	"github.com/coreos/etcd/clientv3"
	"log"
	"runtime"
	"time"
)

func main() {
	prepareLogSetting()
	// --- set the global logger ---

	var conf = eventcar.NewConfig()

	conf.SetRegisterServer(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 2 * time.Second,
	})

	conf.SetNatsHost([]string{"nats://localhost:4222"})

	conf.SetNodeRoles([]string{
		"master", "minion",
	})

	var configErr = conf.CheckBeforeStart()

	if configErr != nil {
		logapi.GetLogger("agent").Fatal(configErr.Error(), nil)
		return
	}

	// detect the properties
	var agent = eventcar.NewAgent(conf)

	agent.Start()

	//defer agent.Stop()
	Client_AddListener(agent)

	// --- call time service
	go func() {
		time.Sleep(6 * time.Second)
		Client_fire(agent)
	}()

	// connect api --

	// ---- keep program running ----
	runtime.Goexit()

}

func Client_AddListener(clientApi eventcar.ClientApi) {

	clientApi.On("test.event1", func(ctx eventcar.ReqMsgContext) {

		fmt.Println(" call define event . ")
		fmt.Println(string(ctx.GetMsgRawBody()))

		var res = ctx.GetResResult()

		res.Complete([]byte("response ok "))

	})

}

func Client_fire(clientApi eventcar.ClientApi) {

	var msg = "hello test case 1"

	var cb eventcar.SuccessCallback

	cb = new(SuccessCallbackImpl)

	clientApi.FireByQueue("test.event1", []byte(msg), cb)

}

type SuccessCallbackImpl struct {
}

func (myself *SuccessCallbackImpl) Succeed(content []byte) {
	fmt.Println(" callback successfully .")
	fmt.Println(string(content))
}

func prepareLogSetting() {

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

func regServer() {
	/*
		serviceName := "fixture-test"
		serviceInfo := servicebus.AgentInfo{IP: "vicenteyou"}

		fixture, err := servicebus.NewAgentRegisterService(serviceName, serviceInfo, []string{
			"http://localhost:2379",
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("name:%fixture, ip:%fixture\n", fixture.nodeId, fixture.Info.IP)

		go func() {
			time.Sleep(time.Second * 20)
			fixture.Stop()
		}()

		fixture.Start()
	*/
}

func myFunc() {
	/*
		i := 0
		c := cron.New()
		spec := "@every 2s"
		c.AddFunc(spec, func() {
			i++
			log.Println("cron running:", i)
		})
		c.Start()

		//关闭着计划任务, 但是不能关闭已经在执行中的任务.
		defer c.Stop()

		select {}
	*/
}
