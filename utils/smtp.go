package utils

import "net/smtp"

func SendMessage(email, firstname, subject, message string) error {

	msg := "From: " + AppSettings.STMPParams.Username + "\n" +
		"To: " + email + "\n" +
		"Subject: " + subject + "\n\n" +
		"Dear, " + firstname + "!\n\n" +
		"We'll answer you as soon as possible!\n\n" +
		"Here is your message:\n\n" +
		message

	err := smtp.SendMail(AppSettings.STMPParams.Server+":"+AppSettings.STMPParams.Port,
		smtp.PlainAuth("", AppSettings.STMPParams.Username, AppSettings.STMPParams.Password, AppSettings.STMPParams.Server),
		AppSettings.STMPParams.Username, []string{email}, []byte(msg))

	return err
}
