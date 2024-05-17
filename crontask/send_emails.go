package crontask

import (
	"github.com/Glebvvss/exchange-rate-api/internal"
	"github.com/Glebvvss/exchange-rate-api/internal/log"
)

func SendEmails() {
	emails, err := internal.GetEmails()
	if err != nil {
		log.Error(err)
		return
	}

	rate, rateErr := internal.GetUsdRate()
	if rateErr != nil {
		log.Error(err)
		return
	}

	body, tplErr := internal.GetRateEmailBody(rate)
	if tplErr != nil {
		log.Error(err)
		return
	}

	for _, email := range emails {
		internal.SendEmail(
			email,
			"Currency rate info",
			body,
		)
	}
}
