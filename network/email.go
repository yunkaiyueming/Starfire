package util

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"

	"github.com/astaxie/beego/config"
)

type BiSmtpEmail struct {
	Host string
	Port int64
	User string
	Pwd  string
	From string
	Tos  string
	Ccs  string
}

func getSmtpConfig() BiSmtpEmail {
	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println("get smtp config error:", err.Error())
		return BiSmtpEmail{}
	}

	port, _ := iniconf.Int64("bi_smtp_config::smtp_port")
	return BiSmtpEmail{
		Host: iniconf.String("bi_smtp_config::smtp_host"),
		Port: port,
		User: iniconf.String("bi_smtp_config::smtp_user"),
		Pwd:  iniconf.String("bi_smtp_config::smtp_pass"),
		From: iniconf.String("bi_smtp_config::from_name"),
		Tos:  iniconf.String("init_app_emails::to"),
		Ccs:  iniconf.String("init_app_emails::cc"),
	}
}

func SendToMail(subject, Content, tos, ccs, fromName string) error {
	BiSmtpEmail := getSmtpConfig()
	toEmails := BiSmtpEmail.Tos
	ccEmails := BiSmtpEmail.Ccs
	fromNameTitle := BiSmtpEmail.From

	if len(tos) > 0 {
		toEmails = tos
	}
	if len(ccs) > 0 {
		ccEmails = ccs
	}
	if len(fromName) > 0 {
		fromNameTitle = fromName
	}

	header := make(map[string]string)
	header["From"] = fromNameTitle + "<" + BiSmtpEmail.User + ">"
	header["To"] = toEmails
	header["Cc"] = ccEmails
	header["Subject"] = subject
	header["Content-Type"] = "text/html; charset=UTF-8"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + Content

	err := SendMailUsingTLS(
		fmt.Sprintf("%s:%d", BiSmtpEmail.Host, BiSmtpEmail.Port),
		smtp.PlainAuth("", BiSmtpEmail.User, BiSmtpEmail.Pwd, BiSmtpEmail.Host),
		BiSmtpEmail.User,
		strings.Split(toEmails+";"+ccEmails, ";"),
		[]byte(message),
	)

	return err
}

func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		fmt.Println("Dialing Error:", err)
		return nil, err
	}

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {
	c, err := Dial(addr)
	if err != nil {
		fmt.Println("Email Create smpt client error:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				fmt.Println("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
