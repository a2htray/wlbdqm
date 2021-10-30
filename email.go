package wlbdqm

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
	"text/template"
)

const (
	MailFromKey  = "WLBDQM_MAIL_FROM"
	MailPassword = "WLBDQM_MAIL_PASSWORD"
	MailHost     = "WLBDQM_MAIL_HOST"
	MailPort     = "WLBDQM_MAIL_PORT"
)

var (
	tplDiskQuotaInsufficient = `<div style="font-size: 18px;">
	Dear Boss @T_T@: <br />
    &nbsp;&nbsp;&nbsp;&nbsp;您的存储使用占比或文件数量占比已达指定比例 {{.MaxPercentage}}%, 请及时清理磁盘。<br />
	存储使用比例：<font color="red">{{.DPercentage}}%</font>，文件数量占比：<font color="red">{{.FPercentage}}%</font><br />
	<br />
	<div style="padding: 4px;">
		{{.DPOutput}}
	</div>
	<br />
	!^_^! Best Wishes.
</div>`
)

// ContentItem
type ContentItem struct {
	MaxPercentage float64
	DPercentage   float64
	FPercentage   float64
	DPOutput      string
}

type Message struct {
	Subject string
	Body    string
}

func SendEmail(toList []string, message Message) error {
	from := os.Getenv(MailFromKey)
	password := os.Getenv(MailPassword)
	host := os.Getenv(MailHost)
	port := os.Getenv(MailPort)

	portInt64, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", toList...)
	m.SetHeader("Subject", "程序自动发送，勿回 - "+message.Subject)
	m.SetBody("text/html", message.Body)

	d := gomail.NewDialer(host, int(portInt64), from, password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func PrepareDiskInsufficientContent(item ContentItem) (Message, error) {
	t, err := template.New("").Parse(tplDiskQuotaInsufficient)
	if err != nil {
		return Message{}, err
	}

	var buffer bytes.Buffer
	if err := t.Execute(&buffer, item); err != nil {
		return Message{}, err
	}
	content := buffer.String()
	//DebugPrintln(content)

	return Message{
		Subject: "存储空间不足",
		Body:    content,
	}, nil
}
