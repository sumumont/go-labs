package rabbit

import (
	"encoding/json"
)

type Pod struct {
	Name     string `json:"name,omitempty"`
	NodeName string `json:"nodeName,omitempty"`

	// IP address of the host to which the pod is assigned. Empty if not yet scheduled.
	HostIP string `json:"hostIP,omitempty"`
	PodIP  string `json:"podIP,omitempty"`
	Phase  string `json:"phase,omitempty"`
}

type DeploymentState struct {
	Replicas            int32 `json:"replicas"`
	ReadyReplicas       int32 `json:"readyReplicas"`
	AvailableReplicas   int32 `json:"availableReplicas"`
	UnavailableReplicas int32 `json:"unavailableReplicas"`
}

type SparkAppState struct {
	AppState      ApplicationState  `json:"appState,omitempty"`
	ExecutorState map[string]string `json:"executorState"`
}

type ApplicationState struct {
	State        string `json:"state"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type JobState struct {
	Status          JobStatus        `json:"status"`
	Name            string           `json:"name"`
	Msg             string           `json:"msg"`
	Pods            []Pod            `json:"pods,omitempty"`
	DeploymentState *DeploymentState `json:"deploymentState,omitempty"`
	SparkAppState   *SparkAppState   `json:"sparkAppState,omitempty"`
}

func (j *JobState) GetStatus() JobStatus {
	return j.Status
}

func (j *JobState) GetName() string {
	return j.Name
}

func (j *JobState) GetMsg() string {
	return j.Msg
}

func (j *JobState) ToJson() string {
	data, _ := json.Marshal(j)
	return string(data)
}

func (j *JobState) SetMsg(msg string) {
	j.Msg = msg
}

func (j *JobState) SetDeploymentState(state DeploymentState) {
	j.DeploymentState = &state
}

func (j *JobState) SetSparkAppState(state SparkAppState) {
	j.SparkAppState = &state
}

func NewUnapproveJobState() *JobState {
	return &JobState{
		Status: JOB_STATUS_UNAPPROVE,
		Name:   JOB_STATUS_NAME_UNAPPROVE,
	}
}

func NewSchedulingJobState() *JobState {
	return &JobState{
		Status: JOB_STATUS_SCHEDULING,
		Name:   JOB_STATUS_NAME_SCHEDULING,
	}
}

func NewQueueJobState() *JobState {
	return &JobState{
		Status: JOB_STATUS_QUEUEING,
		Name:   JOB_STATUS_NAME_QUEUEING,
	}
}

func NewRunningJobState() *JobState {
	return &JobState{
		Status: JOB_STATUS_RUNNING,
		Name:   JOB_STATUS_NAME_RUNNING,
	}
}

func NewFinishJobState() *JobState {
	return &JobState{
		Status: JOB_STATUS_FINISH,
		Name:   JOB_STATUS_NAME_FINISHED,
	}
}

func NewErrorJobState() *JobState {
	return &JobState{
		Status: JOB_STATUS_ERROR,
		Name:   JOB_STATUS_NAME_ERROR,
	}
}

func NewTerminatingJobState() *JobState {
	return &JobState{
		Status: JOB_STATUS_TERMINATING,
		Name:   JOB_STATUS_NAME_TERMINATING,
	}
}

func NewTerminatedJobState(reason string) *JobState {
	return &JobState{
		Status: JOB_STATUS_TERMINATED,
		Name:   JOB_STATUS_NAME_TERMINATED,
		Msg:    reason,
	}
}

func NewUnkownJobState() *JobState {
	return &JobState{
		Status: JOB_STATUS_UNKOWN,
		Name:   JOB_STATUS_NAME_UNKOWN,
	}
}

func ConvertJobState(status string) *JobState {
	if status == JOB_STATUS_NAME_SCHEDULING {
		return NewSchedulingJobState()
	} else if status == JOB_STATUS_NAME_RUNNING {
		return NewRunningJobState()
	} else if status == JOB_STATUS_NAME_ERROR {
		return NewErrorJobState()
	} else if status == JOB_STATUS_NAME_FINISHED {
		return NewFinishJobState()
	} else {
		return NewUnkownJobState()
	}
}

func AbleToTransitState(resType int, src, dst string) bool {

	if src == JOB_STATUS_NAME_RUNNING {

		if dst == JOB_STATUS_NAME_QUEUEING ||
			dst == JOB_STATUS_NAME_UNAPPROVE {
			return false
		}

		// for job resource, we don't allow it to transit state
		// from running to scheduling
		if resType == RESOURCE_TYPE_JOB.Int() &&
			dst == JOB_STATUS_NAME_SCHEDULING {
			return false
		}

	} else if src == JOB_STATUS_NAME_TERMINATING {

		if dst != JOB_STATUS_NAME_TERMINATED &&
			dst != JOB_STATUS_NAME_FINISHED {
			return false
		}

	} else if src == JOB_STATUS_NAME_TERMINATED {

		if dst != JOB_STATUS_NAME_FINISHED {
			return false
		}

	} else if src == JOB_STATUS_NAME_FINISHED ||
		src == JOB_STATUS_NAME_ERROR {
		return false
	}

	return true
}
