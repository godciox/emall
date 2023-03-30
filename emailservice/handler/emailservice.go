package handler

import (
	"context"
	"crypto/tls"
	"go-micro.dev/v4/logger"
	"gopkg.in/gomail.v2"

	pb "emallservice/proto"
)

type EmailService struct{}

func (es *EmailService) SendOrderConfirmation(ctx context.Context, in *pb.SendOrderConfirmationRequest, out *pb.Empty) error {
	m := gomail.NewMessage()
	host := "smtp.qq.com"
	port := 25
	userName := "2386008349@qq.com"
	password := "jygktkiocjtpebjj"
	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m.SetHeader("From", userName)
	m.SetHeader("To", in.Email)
	m.SetBody("text/plain", "订单"+in.Order.GetOrderId()+"发往"+in.Order.ShippingAddress.String()+"花费"+"")
	if err := d.DialAndSend(m); err != nil {
		logger.Errorf("email send failed of %s, email of %s, error is %s", in.Order.OrderId, in.Email, err)
		return err
	}
	logger.Infof("email send success of %s, email of %s", in.Order.OrderId, in.Email)
	return nil
}
