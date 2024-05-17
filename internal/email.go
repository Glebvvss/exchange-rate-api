package internal

import (
	"bytes"
	"context"
	"html/template"
	"net/smtp"
	"os"

	"github.com/badoux/checkmail"
)

type SubscribeQueryResult struct {
	Email string
}

func GetRateEmailBody(rate float32) (string, error) {
	t := template.New("rate")

	var err error
	t, err = t.ParseFiles(os.Getenv("MAILERS_DIR") + "rate.html")
	if err != nil {
		return "", err
	}

	data := struct {
		Rate float32
	}{
		Rate: rate,
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "rate.html", data); err != nil {
		return "", nil
	}

	return tpl.String(), nil
}

func GetEmails() ([]string, error) {
	db, connErr := DbOpen()
	if connErr != nil {
		return nil, connErr
	}

	results, err := db.Query("SELECT `email` FROM `emails`")
	if err != nil {
		return nil, err
	}

	var result []string
	for results.Next() {
		var row SubscribeQueryResult
		if err = results.Scan(&row.Email); err != nil {
			return nil, err
		}

		result = append(result, row.Email)
	}

	return result, nil
}

func SubscribeEmail(email string) error {
	if err := checkmail.ValidateFormat(email); err != nil {
		return err
	}

	db, connErr := DbOpen()
	if connErr != nil {
		return connErr
	}

	_, err := db.ExecContext(
		context.Background(),
		"INSERT IGNORE INTO `emails` (`email`) VALUES (?)",
		email,
	)

	return err
}

func SendEmail(to string, subject string, body string) error {
	from := os.Getenv("EMAIL_FROM")
	pass := os.Getenv("EMAIL_FROM_PASS")
	msg := EmailMsg{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		smtp.PlainAuth("", from, pass, os.Getenv("SMTP_HOST")),
		from, []string{to}, []byte(msg.ToString()),
	)

	return err
}

type EmailMsg struct {
	From    string
	To      string
	Subject string
	Body    string
}

func (e *EmailMsg) ToString() string {
	return "From: " + e.From + "\n" +
		"To: " + e.To + "\n" +
		"Subject: " + e.Subject + "\n\n" +
		e.Body
}
