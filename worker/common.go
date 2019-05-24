package worker

const (
	CMD_RUN   = 1
	CMD_PAUSE = 2
	CMD_STOP  = 3
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
