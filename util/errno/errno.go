package errno

const Success = 0

const ParamsError = -200000
const DBOpError = -2000001

const PhoneHasResisted = -300000

var ErrmsgMap = map[int64]string{
	Success: "success",

	DBOpError:   "数据库操作错误",
	ParamsError: "参数错误",

	PhoneHasResisted: "手机号已被注册",
}
