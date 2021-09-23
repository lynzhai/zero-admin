package common

import (
	"fmt"
	"github.com/aliyun-sdk/mail-go/smtp"
)

type AliEmail struct{
	//SenderAddr string
	//Password string
	//SmtpUrl string
	Client *smtp.Client
}

func NewAliEmail(senderAddr,password,SmtpUrl string)*AliEmail{
	aliEmail := AliEmail{}
	client := smtp.New(SmtpUrl, senderAddr, password)
	aliEmail.Client = client
	return &aliEmail
}

func(s *AliEmail) Send(destEmail ,subject,code string)error{

	emailContent:= fmt.Sprintf("<h2>您的注册验证码是 %s</h2>",code)
	err := s.Client.Send(
		smtp.From("smartjiajia"),
		smtp.Subject(subject),
		smtp.SendTo(destEmail),
		smtp.Content(smtp.Html,emailContent),
	)

	return err
}