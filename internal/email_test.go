package internal

import (
	"github.com/Glebvvss/exchange-rate-api/internal"
	"regexp"
	"testing"
)

func TestGetRateEmailBody(t *testing.T) {
	var rate float32 = 39.45
	body, err := internal.GetRateEmailBody(rate)
	if err != nil {
		t.Fatal(err)
	}

	r, _ := regexp.Compile("Actual currency rates USD/UAH: 39.45")
	if !r.MatchString(body) {
		t.Fatalf("email body intompatible")
	}
}

func TestEmailMsgToString(t *testing.T) {
	var from string = "test@gmail.com"
	var to string = "test2@gmail.com"
	var subject string = "subject"
	var body string = "content"

	msg := internal.EmailMsg{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	actualResult := msg.ToString()
	expectedResult := "From: " + from + "\n" + "To: " + to + "\n" + "Subject: " + subject + "\n\n" + body
	if actualResult != expectedResult {
		t.Fatalf("email message intompatible")
	}
}
