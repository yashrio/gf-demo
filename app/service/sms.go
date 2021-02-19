package service

var Sms = new(smsService)

type smsService struct {}

// 发送手机验证码
func (s *smsService) SendCode(mobile string) error {

	return nil
}

// 校验
func (s *smsService) VerifyCode(mobile string, code string) error  {
	return nil
}
