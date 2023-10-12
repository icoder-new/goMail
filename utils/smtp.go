package utils

import (
	"fmt"
	"fr33d0mz/goMail/internal"
	"net/smtp"
)

func SendMessage(email, firstname, subject, message string) error {

	msg := "From: " + AppSettings.SMTPParams.Username + "\n" +
		"To: " + email + "\n" +
		"Subject: " + subject + "\n\n" +
		"Dear, " + firstname + "!\n\n" +
		"We'll answer you as soon as possible!\n\n" +
		"Here is your message:\n\n" +
		message + "\n\n" +
		"Here is currency realtime:\n\n" +
		currencyData()

	err := smtp.SendMail(AppSettings.SMTPParams.Server+":"+AppSettings.SMTPParams.Port,
		smtp.PlainAuth("", AppSettings.SMTPParams.Username, AppSettings.SMTPParams.Password, AppSettings.SMTPParams.Server),
		AppSettings.SMTPParams.Username, []string{email}, []byte(msg))

	return err
}

func currencyData() string {
	msg := "Transfer\t\tTrade\t\tValue\t\tDate\t\tName\n\n"

	data, _ := internal.GetJSONData()

	for _, rate := range data.LocalRates {
		msg += fmt.Sprintf("%s\t\t%s\t\t%s\t\t%s\t\t%s\n\n", rate.MoneyTransferBuyValue, rate.MoneyTransferTradeValue, rate.SellValue, rate.CreatedAt, rate.Name)
	}

	return msg
}
