package eventcar

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/1-bi/log-api"
	"github.com/coreos/etcd/clientv3"
	"log"
	"strings"
	"time"
)

//the detail of service
type AgentInfo struct {
	LastUpdatedTime int64
}

func (myself *AgentInfo) SetLastUpdatedTime(t int64) {
	myself.LastUpdatedTime = t
}

func NewAgentInfo() *AgentInfo {
	var agentInfo = new(AgentInfo)

	return agentInfo
}

type AgentServiceRegService struct {
	nodeId  string
	stop    chan error
	leaseid clientv3.LeaseID
	client  *clientv3.Client
	_prefix string

	roles []string
}

func NewAgentRegisterService(nodeId string, cli *clientv3.Client, roles []string) *AgentServiceRegService {

	// --- create  AgentService ---
	var agentRegServ = new(AgentServiceRegService)
	agentRegServ.client = cli
	agentRegServ.nodeId = nodeId
	agentRegServ.stop = make(chan error)

	agentRegServ._prefix = "nodes"
	agentRegServ.roles = roles

	return agentRegServ

}

func (s *AgentServiceRegService) Start() error {

	repo, err := s.leaseGrant()
	if err != nil {
		structBean := logapi.NewStructBean()
		structBean.LogString("reason", err.Error())
		logapi.GetLogger("servicebus.AgentServiceWatchService.start").Fatal("Set lease time is fail.", structBean)
		return err
	}

	// --- set the key value frist ---
	ch, err := s.keepAliveFirst(repo)

	// --- connect to message
	for {
		select {
		case err := <-s.stop:
			s.revoke()
			return err
		case <-s.client.Ctx().Done():
			return errors.New("server closed")
		case ka, ok := <-ch:
			if !ok {
				logapi.GetLogger("servicebus.AgentServiceWatchService.start").Info("keep alive channel closed.", nil)
				//log.Println("keep alive channel closed")
				s.revoke()
				return nil
			} else {

				// ---  update status ---
				structBean := logapi.NewStructBean()
				structBean.LogString("nodeId", s.nodeId)
				structBean.LogInt64("ttl time ", ka.TTL)
				logapi.GetLogger("servicebus.AgentServiceWatchService.start").Debug("Recv reply from service: %fixture, ttl:%d", structBean)
				goto END
				//log.Printf("Recv reply from service: %fixture, ttl:%d", fixture.nodeId, ka.TTL)
			}
		}
	END:
		time.Sleep(3 * time.Second)
	}
}

func (myself *AgentServiceRegService) leaseGrant() (*clientv3.LeaseGrantResponse, error) {

	// create new lease
	lease := clientv3.NewLease(myself.client)

	//设置租约时间
	leaseResp, err := lease.Grant(context.TODO(), 5)
	if err != nil {
		return nil, err
	}

	return leaseResp, nil
}

func (myself *AgentServiceRegService) keepAliveFirst(resp *clientv3.LeaseGrantResponse) (<-chan *clientv3.LeaseKeepAliveResponse, error) {

	var err error
	var value []byte
	value, err = myself.getLastUpdatedAgentInfo()
	if err != nil {
		return nil, err
	}

	var key string
	for _, noderole := range myself.roles {
		key = strings.Join([]string{myself._prefix, noderole + "=" + myself.nodeId}, "/")
		_, err = myself.client.Put(context.TODO(), key, string(value), clientv3.WithLease(resp.ID))
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	myself.leaseid = resp.ID

	return myself.client.KeepAlive(context.TODO(), resp.ID)
}

func (myself *AgentServiceRegService) getLastUpdatedAgentInfo() ([]byte, error) {
	// --- get the lastupdate register service ---
	info := NewAgentInfo()
	info.SetLastUpdatedTime(time.Now().UnixNano())

	//info := &fixture.Info
	value, err := json.Marshal(info)
	return value, err
}

func (s *AgentServiceRegService) Stop() {
	s.stop <- nil
}

func (s *AgentServiceRegService) revoke() error {

	_, err := s.client.Revoke(context.TODO(), s.leaseid)
	if err != nil {
		log.Fatal(err)
	}

	structBean := logapi.NewStructBean()
	structBean.LogString("nodeId", s.nodeId)
	logapi.GetLogger("servicebus.AgentServiceWatchService.revoke").Info("servide:%fixture stop\n", structBean)

	//log.Printf("servide:%fixture stop\n", fixture.nodeId)
	return err
}
