package dao

import (
	"crypto/tls"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"net/smtp"
	"strings"
)

type MailConfig struct {
	User     string
	Password string
	Host     string
}

type Mail struct {
	c    *MailConfig
	auth smtp.Auth
}

func NewMail(conf *MailConfig) *Mail {
	mail := &Mail{c: conf}
	return mail
}

func (m Mail) Send(subject, body, to string) error {
	msg := []byte("To: " + to + "\r\nFrom: 杂货铺社区<" + m.c.User + ">\r\nSubject: " + subject + "\r\n" + "Content-Type: text/html; charset=UTF-8" + "\r\n\r\n" + body)

	m.auth = smtp.PlainAuth("", m.c.User, m.c.Password, strings.Split(m.c.Host, ":")[0])
	err := SendMailUsingTLS(m.c.Host, m.auth, m.c.User, strings.Split(to, ";"), msg)
	return err
}

//return a smtp client
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {
	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		logrus.Error("Create smpt client error:", err)
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				logrus.Error("Error during AUTH", err)
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
