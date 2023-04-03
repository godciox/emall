package user

var pathMap = map[string]struct{}{}

func init() {
	pathMap["/user/login"] = struct{}{}
	pathMap["/user/login/captcha"] = struct{}{}
	pathMap["/captcha"] = struct{}{}
	pathMap["/user/register"] = struct{}{}
}
