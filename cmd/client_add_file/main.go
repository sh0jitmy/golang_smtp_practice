package main

import (
    "bytes"
    "encoding/base64"
    "fmt"
    "io/ioutil"
    //"mime/multipart"
    "net/smtp"
    "path/filepath"
)

func main() {
    smtpHost := "localhost"
    smtpPort := "2525"
    auth := smtp.PlainAuth("", "username", "password", smtpHost)

    from := "sender@example.com"
    to := []string{"recipient@example.com"}

    subject := "Subject: Test Email with Image Attachment\r\n"
    body := "This is a test email with an image attachment sent from a Go SMTP client.\r\n"

    // 画像ファイルのパス
    imagePath := "image.jpg"
    imageName := filepath.Base(imagePath)

    // 画像ファイルを読み込む
    image, err := ioutil.ReadFile(imagePath)
    if err != nil {
        fmt.Println("Failed to read image file:", err)
        return
    }

    // MIMEマルチパートの境界設定
    boundary := "my-boundary-1"
    
    // メールのヘッダー
    var msg bytes.Buffer
    msg.WriteString(subject)
    msg.WriteString("MIME-Version: 1.0\r\n")
    msg.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\r\n")
    msg.WriteString("\r\n--" + boundary + "\r\n")
    msg.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
    msg.WriteString("Content-Transfer-Encoding: quoted-printable\r\n")
    msg.WriteString("\r\n" + body + "\r\n")

    // 画像のMIMEパート
    msg.WriteString("\r\n--" + boundary + "\r\n")
    msg.WriteString("Content-Type: image/jpeg\r\n")
    msg.WriteString("Content-Transfer-Encoding: base64\r\n")
    msg.WriteString("Content-Disposition: attachment; filename=\"" + imageName + "\"\r\n")
    msg.WriteString("\r\n")
    
    // 画像の内容をBase64エンコードして追加
    base64Image := base64.StdEncoding.EncodeToString(image)
    msg.WriteString(base64Image)
    msg.WriteString("\r\n--" + boundary + "--\r\n")

    // メール送信
    err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg.Bytes())
    if err != nil {
        fmt.Println("Failed to send email:", err)
        return
    }

    fmt.Println("Email with image attachment sent successfully")
}
