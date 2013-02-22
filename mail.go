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

// A Message represents the most basic electronic message
type Message struct {
	From        mail.Address
	To, Cc, Bcc []mail.Address

	Subject string
	Body    string
}

func NewMessage() *Message {
	return &Message{}
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

func listToString(addresses []mail.Address) string {
	var list []string

	for _, a := range addresses {
		list = append(list, a.String())
	}

	return strings.Join(list, ",")
}
