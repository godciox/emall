package redisInit

import (
	"fmt"
	"time"
)

func SetCaptcha(mobile string, captcha string) error {
	cmd := RDB.SetNX(mobile, captcha, time.Minute*300)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func IsValidCaptcha(mobile, captcha string) (bool, error) {
	rs := RDB.Get(mobile)
	if rs.Err() != nil && rs.Err().Error() == "redis: nil" {
		return false, fmt.Errorf("不存在该手机号的验证码")
	}
	if rs.Val() != captcha {
		return false, fmt.Errorf("验证码不正确")
	}
	return true, nil
}
