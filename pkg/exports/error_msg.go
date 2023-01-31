package exports

import (
	"fmt"
	"strings"
)

type ErrorEnum map[int]string

var msgDefault = "system error"

var errText = ErrorEnum{
	SUCCESS:                                            "success",
	APULIS_IQI_ERROR_BEGIN:                             "begin error",
	APULIS_IQI_NOT_FOUND:                               "not found",
	APULIS_IQI_PARAM_ERROR:                             "param errror: << .V >>",
	APULIS_IQI_ALREADY_EXISTS:                          "already exists",
	APULIS_IQI_NO_API:                                  "no api",
	APULIS_IQI_NOT_IMPLEMENT:                           "not implement",
	APULIS_IQI_UNKNOWN_ERROR:                           "unknown error",
	APULIS_IQI_FILE_NOT_FOUND:                          "file not found",
	APULIS_IQI_PATH_NOT_FOUND:                          "path not fount",
	APULIS_IQI_OS_IO_ERROR:                             "io error",
	APULIS_IQI_OS_READ_DIR_ERROR:                       "read dir error",
	APULIS_IQI_OS_REMOVE_FILE:                          "remove file error",
	APULIS_IQI_OS_CREATE_FILE:                          "create file error",
	APULIS_IQI_FILE_TOO_LARGE:                          "file too large",
	APULIS_IQI_FILE_TYPE_ERROR:                         "file type error",
	APULIS_IQI_COMPUTE_EXCEED_LIMIT:                    "compute exceed limit : << .V >>",
	APULIS_IQI_COMPUTE_WOULD_BLOCK:                     "compute would block",
	APULIS_IQI_ERROR_REQUEST_NEW:                       "http new request error: << .V >>",
	APULIS_IQI_ERROR_REQUEST_DO:                        "http do request error: << .V >>",
	APULIS_IQI_ERROR_RESPONSE_IO:                       "http response ioread error: << .V >>",
	APULIS_IQI_ERROR_JSON_UNMARSHAL:                    "json unmarshal error: << .V >>",
	APULIS_IQI_ERROR_JSON_MARSHAL:                      "json marshal error: << .V >>",
	APULIS_IQI_ERROR_URL_PARSE:                         "url parse error: << .V >>",
	APULIS_IQI_WOULD_BLOCK:                             "would block",
	APULIS_IQI_NO_AUTH:                                 "no auth",
	APULIS_IQI_REMOTE_NETWORK_ERROR:                    "remote network error: << .V >>",
	APULIS_IQI_REMOTE_REST_ERROR:                       "remote rest error: << .V >>",
	APULIS_IQI_GROUP_NO_ERROR:                          "group no error",
	APULIS_IQI_ERROR_DIR_MAKE:                          "make dir error: << .V >>",
	APULIS_IQI_DB_ERROR:                                "db error",
	APULIS_IQI_DB_QUERY_FAILED:                         "db query failed: << .V >>",
	APULIS_IQI_DB_EXEC_FAILED:                          "db sql exec failed: << .V >>",
	APULIS_IQI_DB_DUPLICATE:                            "db data duplication: << .V >>",
	APULIS_IQI_DB_UPDATE_UNEXPECT:                      "db update unexpect: << .V >>",
	APULIS_IQI_DB_WRONG_TYPE:                           "db wrong type",
	APULIS_IQI_DB_READ_ROWS:                            "db read rows failed: << .V >>",
	APULIS_IQI_ERROR_PARSE_INT:                         "parse int error: << .V >>",
	APULIS_IQI_ERROR_SYSTEM_NAME:                       "sysparams name: << .V >> error",
	APULIS_IQI_ERROR_TASK_TYPE_EMPTY:                   "tasktype is empty",
	APULIS_IQI_ERROR_OS_RENAME_FILE:                    "rename failed: << .V >>",
	APULIS_IQI_ERROR_PRESET_UPDATE_TYPE:                "publish modle update type error",
	APULIS_IQI_ERROR_PERMISSION_DENIED:                 "permission denied",
	APULIS_IQI_ERROR_MODELS_EXITS:                      "model name exits",
	APULIS_IQI_ERROR_READ_FILE:                         "read file failed: << .V >>",
	APULIS_IQI_PROJECT_NOT_EXIST:                       "project not exits",
	APULIS_IQI_PROJECT_ERROR:                           "project error: << .V >>",
	APULIS_IQI_PROJECT_EXIST:                           "name already exits",
	APULIS_IQI_PROJECT_MODEL_MATCH_FAILED:              "models does not match to the sdk",
	APULIS_IQI_PROJECT_INFER_CHECK_FAILED:              "check running infer service failed: << .V >>",
	APULIS_IQI_PROJECT_USED_BY_INFERS:                  "project used by some infers: << .V >>",
	APULIS_IQI_PROJECT_LAB_NOT_FOUND:                   "lab not found:<< .V >>",
	APULIS_IQI_PROJECT_SDK_NOT_FOUND:                   "sdk not found",
	APULIS_IQI_PROJECT_SDK_MODEL_TITLE_DUPLICATE:       "sdk model title cannot be duplicate",
	APULIS_IQI_DB_DELETE_FAILED:                        "delete failed",
	APULIS_IQI_PROJECT_MODEL_LENS_ERROR:                "project models num does not between MinimumNum and MaximumNum!",
	APULIS_IQI_PROJECT_STOP_JOB_ERROR:                  "stop job error: << .V >>",
	APULIS_IQI_PROJECT_MODEL_MIN:                       "need at least one model",
	APULIS_IQI_RECHECK_FAILED:                          "infer recheck failed : << .V >>",
	APULIS_IQI_ANALYSIS_FAILED:                         "analysis failed: << .V >>",
	APULIS_IQI_PREDICT_FAILED:                          "predict failed: << .V >>",
	APULIS_IQI_PREDICT_MUST_RUN:                        "infer service must running",
	APULIS_IQI_PREDICT_RUNING:                          "predict job is running",
	APULIS_IQI_PREDICT_NEED_DATASET:                    "need dataset",
	APULIS_IQI_PREDICT_NEED_FILE:                       "need file",
	APULIS_IQI_PREDICT_NEED_INFER_URL:                  "find no infer url",
	APULIS_IQI_PREDICT_TASK_DATASET_CREATE_FAILED:      "dataset create failed : << .V >>",
	APULIS_IQI_PREDICT_TASK_DATASET_PUBLISH_FAILED:     "dataset publish failed : << .V >>",
	APULIS_IQI_PREDICT_TASK_SUBMIT_FIND_NO_RECHECK:     "find no images rechecked",
	APULIS_IQI_PREDICT_TASK_SUBMIT_MQ_SEND_RECHECK:     "send MQ message to pipeline failed : << .V >>",
	APULIS_IQI_ERROR_DATASET_PUBLISHED:                 "dataset is published , can not do this",
	APULIS_IQI_ERROR_DATASET_PUBLISHING:                "dataset is publishing , can not do this",
	APULIS_IQI_ERROR_KEYMAP:                            "keymap error : << .V >>",
	APULIS_IQI_ERROR_KEYMAP_EMPTY:                      "keymap is empty",
	APULIS_IQI_ERROR_CANNOT_SUBMIT:                     "there has some files to annot",
	APULIS_IQI_ERROR_PRODUCT_TEMPLATE_NAME_EXIST:       "product template name exists",
	APULIS_IQI_ERROR_PRODUCT_TEMPLATE_ZIP_NAME_INVALID: "invalid zip name",
	APULIS_IQI_ERROR_INFER_OP_FAILED:                   "infer failed: << .V >>",
	APULIS_IQI_RECKECK_ANNOT_NOTFOUNT:                  "annot not found",
	APULIS_IQI_ERROR_INFER_TEMPLATE_NOT_FOUND:          "template not found: << .V >>",
	APULIS_IQI_PREDICT_TASK_IMAGES_NOT_FOUND:           "task images not found",
	APULIS_IQI_ERROR_INFER_COORDINATE_NOT_FOUND:        "coordinate not found:<< .V >>",
}

func ErrMsg(code int, args ...interface{}) string {
	if msgs, ok := errText[code]; ok {
		// todo 所有带<<>>的 里面不管什么字段都替换为%v 兼容local的msg和国际化的msg
		msg := strings.ReplaceAll(msgs, "<< .V >>", "%v")
		return fmt.Sprintf(msg, args...)
	}
	return msgDefault
}
