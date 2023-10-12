package utils

import (
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
		"Here is currency realtime:\n\n"
		// + "wtf to write?"

	err := smtp.SendMail(AppSettings.SMTPParams.Server+":"+AppSettings.SMTPParams.Port,
		smtp.PlainAuth("", AppSettings.SMTPParams.Username, AppSettings.SMTPParams.Password, AppSettings.SMTPParams.Server),
		AppSettings.SMTPParams.Username, []string{email}, []byte(msg))

	return err
}
