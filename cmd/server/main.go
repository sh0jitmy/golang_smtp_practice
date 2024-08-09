package main

import (
    "bytes"
    "log"
    "io"
    "net"
    "net/mail"

    "github.com/mhale/smtpd"
)

func mailHandler(origin net.Addr, from string, to []string, data []byte) error {
    msg, _ := mail.ReadMessage(bytes.NewReader(data))
    subject := msg.Header.Get("Subject")
    log.Printf("Received mail from %s for %s with subject %s", from, to[0], subject)
    bytes, _ := io.ReadAll(msg.Body)
    log.Printf("Body: %s", string(bytes))  
    return nil
}

func authHandler(remoteAddr net.Addr, mechanism string, username []byte, password []byte, shared []byte) (bool, error) {
    log.Printf("auth handle call\n")
    return string(username) == "username" && string(password) == "password", nil
}

func ListenAndServe(addr string, handler smtpd.Handler, authHandler smtpd.AuthHandler) error {
    mechs := map[string]bool{"PLAIN": true}
    srv := &smtpd.Server{
        Addr:        addr,
        Handler:     handler,
        Appname:     "MyServerApp",
        Hostname:    "",
        AuthHandler: authHandler,
        AuthRequired: true,
        AuthMechs : mechs,
    }
    return srv.ListenAndServe()
}

func main() {
    ListenAndServe("127.0.0.1:2525", mailHandler, authHandler)
}
