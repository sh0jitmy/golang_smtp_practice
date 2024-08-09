package main

import (
    "fmt"
    "net/smtp"
)

func main() {
    // SMTPサーバー情報
    smtpHost := "localhost"
    smtpPort := "2525"
    auth := smtp.PlainAuth("", "username", "password", smtpHost)
    
    // 送信者と受信者のメールアドレス
    from := "sender@example.com"
    to := []string{"recipient@example.com"}

    // メールの内容
    subject := "Subject: Test Email\r\n"
    body := "This is a test email sent from a Go SMTP client.\r\n"
    msg := []byte(subject + "\r\n" + body)

    // メール送信
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
    if err != nil {
        fmt.Println("Failed to send email:", err)
        return
    }

    fmt.Println("Email sent successfully")
}
