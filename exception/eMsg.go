package ecode

var CM = map[int]map[string]string{
	0:    {"EN": "OK", "CN": "OK"},
	-304: {"EN": "NO CHANGE", "CN": "没有改动"},
	-307: {"EN": "redirect", "CN": "装车跳转"},
	-400: {"EN": "wrong request", "CN": "请求错误"},
	-401: {"EN": "not authorized", "CN": "未认证"},
	-403: {"EN": "permission denied", "CN": "访问权限不足"},
	-404: {"EN": "no data", "CN": "啥都木有"},
	-405: {"EN": "invalid method", "CN": "不支持该方法"},
	-409: {"EN": "conflict request", "CN": "冲突"},
	-498: {"EN": "client stop", "CN": "客户端取消请求"},
	-500: {"EN": "server error", "CN": "服务器错误"},
	-503: {"EN": "server overload", "CN": "过载保护,服务暂不可用"},
	-504: {"EN": "service out time", "CN": "服务调用超时"},
	-509: {"EN": "out of boundary", "CN": "超出限制"},

	10001: {"EN": "illegal param", "CN": "参数错误"},
	10101: {"EN": "login fail", "CN": "登陆失败"},
	10102: {"EN": "Token out time", "CN": "Token 过期"},
	10103: {"EN": "not login", "CN": "账号未登录"},
	10104: {"EN": "invalid identifying code", "CN": "验证码错误"},
	10105: {"EN": "fail too many times", "CN": "登录失败次数太多"},
	10106: {"EN": "invalid user", "CN": "用户不存在"},
	10107: {"EN": "invalid username or password", "CN": "用户名或密码错误"},
	10108: {"EN": "illegal Token", "CN": "Token 非法"},
	10109: {"EN": "login fail too many times,please try again in 15 minutes", "CN": "登录太频繁，请15分钟后再试"},
	10110: {"EN": "send sms too frequently,please try again in 60 seconds", "CN": "验证码发送过于频繁，请60s后重试"},
	10111: {"EN": "sms service current not support this region", "CN": "目前短信发送尚不支持此地区"},
	10130: {"EN": "invalid email sender", "CN": "不存在的发送邮箱"},

	10201: {"EN": "frozen user", "CN": "用户被封禁"},
	10202: {"EN": "duplicate user", "CN": "重复的用户"},
	10203: {"EN": "illegal username", "CN": "用户名不合法"},
	10204: {"EN": "illegal phone number", "CN": "手机号不合法"},
	10205: {"EN": "password too weak", "CN": "密码太弱"},
	10206: {"EN": "username too long", "CN": "用户名长度超过限制"},
	10207: {"EN": "request frequency too high", "CN": "触发频率限制"},
	10208: {"EN": "new password is same as old one", "CN": "新密码与旧密码相同"},
	10209: {"EN": "invalid invite code", "CN": "无效的邀请码"},
	10210: {"EN": "your ID card cannot be updated,because you already have one,if you need help,please contact the staff", "CN": "已验证过身份证，修改需要联系工作人员进行核实"},
	10211: {"EN": "illegal IdCard", "CN": "身份证号不合法"},
	10212: {"EN": "illegal role", "CN": "无效的角色"},
	10213: {"EN": "illegal role change", "CN": "无法跨平台变更角色"},
	10214: {"EN": "email verify code error", "CN": "邮箱验证码错误"},
}
