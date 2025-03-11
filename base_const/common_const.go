package baseConst

// 处理状态: 20.清洗待处理 21.清洗中 22.清洗失败 23.清洗成功 30.打标待处理 31.打标中 32.打标失败 33.打标成功 40.训练待处理 41.训练中 42.训练失败 43.训练成功 50.发布待处理 51.发布中 52.发布失败 53.发布成功
var (
	ModelHandleCleanPending   = &ModelHandleStatus{Code: 20, Msg: "清洗待处理"}
	ModelHandleCleanProcess   = &ModelHandleStatus{Code: 21, Msg: "清洗中"}
	ModelHandleCleanFailed    = &ModelHandleStatus{Code: 22, Msg: "清洗失败"}
	ModelHandleCleanSuccess   = &ModelHandleStatus{Code: 23, Msg: "清洗成功"}
	ModelHandleTagPending     = &ModelHandleStatus{Code: 30, Msg: "打标待处理"}
	ModelHandleTagProcess     = &ModelHandleStatus{Code: 31, Msg: "打标中"}
	ModelHandleTagFailed      = &ModelHandleStatus{Code: 32, Msg: "打标失败"}
	ModelHandleTagSuccess     = &ModelHandleStatus{Code: 33, Msg: "打标成功"}
	ModelHandleTrainPending   = &ModelHandleStatus{Code: 40, Msg: "训练待处理"}
	ModelHandleTrainProcess   = &ModelHandleStatus{Code: 41, Msg: "训练中"}
	ModelHandleTrainFailed    = &ModelHandleStatus{Code: 42, Msg: "训练失败"}
	ModelHandleTrainSuccess   = &ModelHandleStatus{Code: 43, Msg: "训练成功"}
	ModelHandlePublishPending = &ModelHandleStatus{Code: 50, Msg: "发布待处理"}
	ModelHandlePublishProcess = &ModelHandleStatus{Code: 51, Msg: "发布中"}
	ModelHandlePublishFailed  = &ModelHandleStatus{Code: 52, Msg: "发布失败"}
	ModelHandlePublishSuccess = &ModelHandleStatus{Code: 53, Msg: "发布成功"}
)

type ModelHandleStatus struct {
	Code int
	Msg  string
}

func GetModelHandleStatus(handleStatus int) *ModelHandleStatus {
	switch handleStatus {
	case ModelHandleCleanPending.Code:
		return ModelHandleCleanPending
	case ModelHandleCleanProcess.Code:
		return ModelHandleCleanProcess
	case ModelHandleCleanFailed.Code:
		return ModelHandleCleanFailed
	case ModelHandleCleanSuccess.Code:
		return ModelHandleCleanSuccess
	case ModelHandleTagPending.Code:
		return ModelHandleTagPending
	case ModelHandleTagProcess.Code:
		return ModelHandleTagProcess
	case ModelHandleTagFailed.Code:
		return ModelHandleTagFailed
	case ModelHandleTagSuccess.Code:
		return ModelHandleTagSuccess
	case ModelHandleTrainPending.Code:
		return ModelHandleTrainPending
	case ModelHandleTrainProcess.Code:
		return ModelHandleTrainProcess
	case ModelHandleTrainFailed.Code:
		return ModelHandleTrainFailed
	case ModelHandleTrainSuccess.Code:
		return ModelHandleTrainSuccess
	case ModelHandlePublishPending.Code:
		return ModelHandlePublishPending
	case ModelHandlePublishProcess.Code:
		return ModelHandlePublishProcess
	case ModelHandlePublishFailed.Code:
		return ModelHandlePublishFailed
	case ModelHandlePublishSuccess.Code:
		return ModelHandlePublishSuccess
	default:
		return nil
	}
}
