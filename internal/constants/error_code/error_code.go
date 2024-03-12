package error_code

const (
	Success = 0

	Failed         = 1
	InvalidRequest = 98
	SystemError    = 99

	ErrorIsNotDefineMsg = "Error unknown"

	InvalidRequestMsg = "Invalid request"
	SystemErrorMsg    = "System error"

	SucessMsg = "Success"
	FailedMsg = "Failed"

	ErrorContentType = 415
)

var MapErrorCode = map[int]string{
	InvalidRequest: InvalidRequestMsg,
	SystemError:    SystemErrorMsg,
	Success:        SucessMsg,
	Failed:         FailedMsg,
}

func GetErrorMsg(errorCode int) string {
	errMsg := MapErrorCode[errorCode]
	if errMsg == "" {
		return ErrorIsNotDefineMsg
	} else {
		return errMsg
	}
}
