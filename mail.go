// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// package ezmail implements basic mail creation to ease the net/smtp package use
package ezmail

import (
	"bytes"
	"fmt"
	"net/mail"
	"strings"
)

// A Message represents an electronic message
// Usage:
// msg := ezmail.Message{}
// msg.SetFrom("Marc Weistroff", "marc@example.org")
// msg.AddTo("Recipient 1", "recipient1@example.org")
// msg.AddTo("", "foobar@example.org")
// msg.AddCC("", "cc@example.org")
// msg.Subject = "Hey, message subject"
// msg.Body = "plain text body"
// msg.Headers["Content-Type"] = "text/plain"
// smtp.SendMail(addr, auth, msg.From.String(), msg.Recipients(), msg.Bytes())
//
type Message struct {
	From        mail.Address
	To, Cc, Bcc []mail.Address

	Subject string
	Body    string
	
	Headers map[string]string
}

func NewMessage() *Message {
	return &Message{Headers:map[string]string{}}
}

// Bytes returns the []byte representation of the message
func (msg *Message) Bytes() []byte {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintln("From:", msg.From.String()))
	if len(msg.To) > 0 {
		b.WriteString(fmt.Sprintln("To:", listToString(msg.To)))
	}
	if len(msg.Cc) > 0 {
		b.WriteString(fmt.Sprintln("Cc:", listToString(msg.Cc)))
	}
	if len(msg.Bcc) > 0 {
		b.WriteString(fmt.Sprintln("Bcc:", listToString(msg.Bcc)))
	}
	
	for k, v := range msg.Headers {
	    b.WriteString(fmt.Sprintf("%s:%s\n", k, v))
	}

	b.WriteString(fmt.Sprintln("Subject:", msg.Subject))
	b.WriteString(fmt.Sprintln())
	b.WriteString(msg.Body)

	return b.Bytes()
}

func (msg *Message) SetFrom(name, email string) {
	msg.From = mail.Address{name, email}
}

func (msg *Message) AddTo(name, email string) {
	msg.To = append(msg.To, mail.Address{name, email})
}

func (msg *Message) AddCc(name, email string) {
	msg.Cc = append(msg.Cc, mail.Address{name, email})
}

func (msg *Message) AddBcc(name, email string) {
	msg.Bcc = append(msg.Bcc, mail.Address{name, email})
}

// String returns the string representation of the message
func (msg *Message) String() string {
	return string(msg.Bytes())
}

// Returns list of recipients (Format: "Recipient Name <recipientemail@example.com>")
func (msg *Message) Recipients() []string {
	var r []string
	for _, v := range msg.To {
		r = append(r, v.String())
	}

	return r
}

// Returns list of recipient email addresses
func (msg *Message) RecipientsEmails() []string {
	var r []string
	for _, v := range msg.To {
		r = append(r, v.Address)
	}

	return r
}

func listToString(addresses []mail.Address) string {
	var list []string

	for _, a := range addresses {
		list = append(list, a.String())
	}

	return strings.Join(list, ",")
}
