package common

type InnerErrCode int64

const (
	// common
	INNER_ERR_CODE_UNKNOWN_ERR                   = InnerErrCode(0)
	INNER_ERR_CODE_SUCCESS                       = InnerErrCode(1)
	INNER_ERR_CODE_NO_PERMISSION             	 = InnerErrCode(2)
	INNER_ERR_CODE_BAD_PARAM                     = InnerErrCode(3)
)
