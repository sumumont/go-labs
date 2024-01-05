package rabbit

const (
	JOB_STATUS_NAME_UNAPPROVE   = "unapprove"
	JOB_STATUS_NAME_QUEUEING    = "queueing"
	JOB_STATUS_NAME_SCHEDULING  = "scheduling"
	JOB_STATUS_NAME_RUNNING     = "running"
	JOB_STATUS_NAME_FINISHED    = "finish"
	JOB_STATUS_NAME_ERROR       = "error"
	JOB_STATUS_NAME_TERMINATING = "terminating"
	JOB_STATUS_NAME_TERMINATED  = "terminated"
	JOB_STATUS_NAME_UNKOWN      = "unkown"
)

const (
	MODID_JOB_SCHEUDULER int = 108
)

type JobStatus int

const (
	// APP Status
	JOB_STATUS_UNAPPROVE JobStatus = iota
	JOB_STATUS_QUEUEING
	JOB_STATUS_SCHEDULING
	JOB_STATUS_RUNNING
	JOB_STATUS_FINISH

	// based on return code of main process within target pod
	JOB_STATUS_ERROR
	JOB_STATUS_TERMINATING
	JOB_STATUS_TERMINATED
	JOB_STATUS_UNKOWN
)

type ResourceType int

const (
	RESOURCE_TYPE_None ResourceType = iota
	RESOURCE_TYPE_POD
	RESOURCE_TYPE_JOB
	RESOURCE_TYPE_DEPLOYMENT
	RESOURCE_TYPE_CRD
	RESOURCE_TYPE_YAML
	RESOURCE_TYPE_INFERENCE_SERVICE
	RESOURCE_TYPE_SERVICE
	RESOURCE_TYPE_SPARK_APP
	RESOURCE_TYPE_DISTRIBUTED_JOB
	RESOURCE_TYPE_UNKONWN
	Resource_TYPE_ARGO_WORKFLOW
)

func (r ResourceType) Int() int {
	return int(r)
}

func (r ResourceType) IsGeneralResType() bool {

	if r.Int() >= int(RESOURCE_TYPE_POD) && r.Int() <= int(RESOURCE_TYPE_DEPLOYMENT) {
		return true
	} else {
		return false
	}
}

func (r ResourceType) IsGeneralJob() bool {

	if r == RESOURCE_TYPE_JOB {
		return true
	} else {
		return false
	}
}

func (r ResourceType) IsGeneralDeployment() bool {

	if r == RESOURCE_TYPE_DEPLOYMENT {
		return true
	} else {
		return false
	}
}

func (r ResourceType) IsDistributedJob() bool {

	if r == RESOURCE_TYPE_DISTRIBUTED_JOB {
		return true
	} else {
		return false
	}
}

func (r ResourceType) IsSparkApp() bool {
	if r == RESOURCE_TYPE_SPARK_APP {
		return true
	} else {
		return false
	}
}

func (r ResourceType) IsArgoWorkflow() bool {
	if r == Resource_TYPE_ARGO_WORKFLOW {
		return true
	} else {
		return false
	}
}

// apply for Job.Ext
type RESOURCE_FIELD_NAME string

const (

	// deployment
	FIELD_REPLICA       RESOURCE_FIELD_NAME = "replicas"
	FIELD_PORTS         RESOURCE_FIELD_NAME = "ports"
	FIELD_YAML          RESOURCE_FIELD_NAME = "yaml"
	FIELD_AFFINITY      RESOURCE_FIELD_NAME = "affinity"
	FIELD_TOLERATE      RESOURCE_FIELD_NAME = "tolerate"
	FIELD_NODE_SELECTOR RESOURCE_FIELD_NAME = "node_selector"
	FIELD_ANNOTATION    RESOURCE_FIELD_NAME = "annotation"
)

// kubernetes labels
const (
	KUBE_LABEL_JOB_ID   string = "job-id"
	KUBE_LABEL_JOB_NAME string = "job-name"
)

// service name
const (
	SERVICE_NAME string = "job-scheduler.default"
	SERVICE_PORT int    = 10100
)

// mount point
const (
	PVC_PREFIX            string = "pvc://"
	HOSTPATH_PREFIX       string = "FILE://"
	HOSTPATH_LOWER_PREFIX string = "file://"

	CM_PREFIX       string = "CM://"
	CM_LOWER_PREFIX string = "cm://"

	EMPTY_DIR_PREFIX string = "emptydir://"
)

// volume mount type

// environment
const (
	ENV_KEY_RUN_AS_ROOT   string = "RUN_AS_ROOT"
	ENV_VALUE_RUN_AS_ROOT string = "true"
)

// compute resource type
type ComputeResType int

const (
	COMPUTE_RES_TYPE_UNDEFINE ComputeResType = iota

	COMPUTE_RES_TYPE_RAM
	COMPUTE_RES_TYPE_CPU
	COMPUTE_RES_TYPE_GPU
	COMPUTE_RES_TYPE_NPU

	COMPUTE_RES_TYPE_ALL
)

// k8s service type
const (
	ServiceTypeClusterIP    string = "ClusterIP"
	ServiceTypeNodePort     string = "NodePort"
	ServiceTypeLoadBalancer string = "LoadBalancer"
	ServiceTypeExternalName string = "ExternalName"
)

// org id
const DefaultOrgId int64 = 1
const DefaultOrgName string = "systemadmin-organization"
const DefaultPullImageSecretName string = "regcred"

// node selector
const NodeSelectorKeyForPlatformSvc string = "aistudio"
const NodeSelectorValueForPlatformSvc string = "active"

const NodeSelectorKeyForOrgSvc string = "org-name"

const NodeSelectorEdgeNode string = "apedgenode"

// service scope
type ResourceScope int

const (
	// 默认选项
	ResourceScopeOrg      ResourceScope = 0
	ResourceScopePlatform ResourceScope = 1
)

const (
	DefaultCPUQuota = "1.0"
	DefaultMemQuota = "100Mi"
)

var (
	DefaultTTLSecondsAfterFinished int32 = 0
)

const (
	JobTypeDaemon = "daemon"
)

const (
	ShmMountPath = "/dev/shm"
)
