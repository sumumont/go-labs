package rabbit

type MsgType int

const (
	MSG_TYPE_NONE                MsgType = iota
	MSG_TYPE_JOB_UPDATE          MsgType = 1
	MSG_TYPE_SERVICE_UPDATE      MsgType = 2
	MSG_TYPE_DEPLOYMENT_UPDATE   MsgType = 3
	MSG_TYPE_DEPLOYMENT_DELETE   MsgType = 4
	MSG_TYPE_ARGOWORKFLOW_UPDATE MsgType = 5
	MSG_TYPE_ARGOWORKFLOW_DELETE MsgType = 6
)

// 3
// message queue protocol
type JobMsg struct {
	JobId     string
	Namespace string
	JobState  JobState

	MsgType MsgType `json:"msgType"`
	Body    string  `json:"body"`
}

func NewJobStateMsg(jobId string, msgType MsgType, jobState JobState) *JobMsg {
	return &JobMsg{
		JobId:    jobId,
		MsgType:  msgType,
		JobState: jobState,
	}
}

func (j *JobMsg) GetBody() string {
	return j.Body
}

func (j *JobMsg) SetBody(body string) {
	j.Body = body
}

func (j *JobMsg) GetJobId() string {
	return j.JobId
}

func (j *JobMsg) SetJobId(jobId string) {
	j.JobId = jobId
}

func (j *JobMsg) GetMsgType() MsgType {
	return j.MsgType
}

func (j *JobMsg) SetMsgType(msgType MsgType) {
	j.MsgType = msgType
}

type ServiceState struct {
	Name     string `json:"name"`
	JobId    string `json:"jobId"`
	NodePort int32  `json:"nodePort"`
}

func NewServiceState(name, jobId string, nodePort int32) *ServiceState {
	return &ServiceState{
		Name:     name,
		JobId:    jobId,
		NodePort: nodePort,
	}
}

func (s *ServiceState) GetName() string {
	return s.Name
}

func (s *ServiceState) GetJobId() string {
	return s.JobId
}

func (s *ServiceState) GetNodePort() int32 {
	return s.NodePort
}
