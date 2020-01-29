package errno

const Success = 0

const ParamsError = -200000
const DBOpError = -2000001
const RedisOpError = -2000002

const PhoneHasResisted = -300000
const SmsCodeNotRight = -300001
const SmsSendFailed = -300002

var ErrmsgMap = map[int64]string{
	Success: "success",

	DBOpError:    "数据库操作错误",
	RedisOpError: "Redis操作错误",
	ParamsError:  "参数错误",

	PhoneHasResisted: "手机号已被注册",
	SmsCodeNotRight:  "短信验证码错误",
	SmsSendFailed:    "发送短信失败",
}
