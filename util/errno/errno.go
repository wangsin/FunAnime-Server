package errno

const (
	Success = iota * -1
	TokenInvalid
	TokenExpired
	UnknownError
)

const ParamsError = -200000
const DBOpError = -2000001
const RedisOpError = -2000002
const NotFound = -2000003

const PhoneHasResisted = -300000
const SmsCodeNotRight = -300001
const SmsSendFailed = -300002
const PhoneNotExistence = -300003
const SmsCodeNotSend = -300004
const LoginInfoFailed = -300005
const Uncertified = -300006

var ErrmsgMap = map[int64]string{
	Success: "success",

	DBOpError:    "数据库操作错误",
	RedisOpError: "Redis操作错误",
	ParamsError:  "参数错误",

	PhoneHasResisted:  "手机号已被注册",
	PhoneNotExistence: "手机号不存在",
	SmsCodeNotRight:   "短信验证码错误",
	SmsCodeNotSend:    "短信验证码尚未发送",
	SmsSendFailed:     "您发送短信过于频繁，请稍后再试~",
	LoginInfoFailed:   "密码或验证码错误，登陆失败",
	Uncertified:       "非法用户",

	NotFound: "没有找到",
}
