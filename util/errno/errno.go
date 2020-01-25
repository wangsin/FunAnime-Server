package errno

const Success = 0

const ParamsError = -200000

var ErrmsgMap = map[int64]string{
	ParamsError: "参数错误",
}
