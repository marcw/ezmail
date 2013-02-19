// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package ezmail

import (
	"bytes"
	"io/ioutil"
	"net/mail"
	"testing"
)

func TestMessage(t *testing.T) {
	msg := &Message{}

	msg.From = mail.Address{Name: "Marc Weistroff", Address: "marc+from@weistroff.net"}
	msg.To = []mail.Address{mail.Address{Name: "Marc Weistroff", Address: "marc+to@weistroff.net"},
		mail.Address{Name: "Foo Bar", Address: "foo@example.org"}}

	msg.Cc = []mail.Address{mail.Address{Name: "Bar Foo", Address: "bar@example.org"}}
	msg.Bcc = []mail.Address{mail.Address{Name: "The Boss", Address: "boss@example.org"}}
	msg.Subject = "Lorem ipsum gollum"
	msg.Body = `Oy Mate!
Wanna drink a beer tonight?

Cheers`

	proof, _ := mail.ReadMessage(bytes.NewBuffer(msg.Bytes()))
	body, _ := ioutil.ReadAll(proof.Body)
	if string(body) != msg.Body {
		t.Log(string(body))
		t.FailNow()
	}

	from, _ := proof.Header.AddressList("From")
	if from[0].Name != "Marc Weistroff" || from[0].Address != "marc+from@weistroff.net" {
		t.Log(from[0])
		t.FailNow()
	}

	to, _ := proof.Header.AddressList("To")
	if len(to) != 2 {
		proof.Header.Get("To")
		t.Log(to)
		t.FailNow()
	}
	if to[0].Name != "Marc Weistroff" || to[0].Address != "marc+to@weistroff.net" {
		t.Log(to[0])
		t.FailNow()
	}
	if to[1].Name != "Foo Bar" || to[1].Address != "foo@example.org" {
		t.Log(to[1])
		t.FailNow()
	}

	cc, _ := proof.Header.AddressList("Cc")
	if cc[0].Name != "Bar Foo" || cc[0].Address != "bar@example.org" {
		t.Log(cc[0])
		t.FailNow()
	}

	bcc, _ := proof.Header.AddressList("Bcc")
	if bcc[0].Name != "The Boss" || bcc[0].Address != "boss@example.org" {
		t.Log(bcc[0])
		t.FailNow()
	}

	subject := proof.Header.Get("Subject")
	if subject != msg.Subject {
		t.Log(subject)
		t.FailNow()
	}
}
