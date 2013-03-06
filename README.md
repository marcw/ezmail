# ezmail

A library to create mail messages and ease golang net/smtp package usage.

## Usage

Here is how to use this package to send a basic mail message.

    msg := ezmail.Message{}
    msg.SetFrom("Marc Weistroff", "marc@example.org")
    msg.AddTo("Recipient 1", "recipient1@example.org")
    msg.AddTo("", "foobar@example.org")
    msg.AddCC("", "cc@example.org")
    msg.Subject = "Hey, message subject"
    msg.Body = "plain text body"
    smtp.SendMail(addr, auth, msg.From.String(), msg.Recipients(), msg.Bytes())

## License

The ezmail code is free to use and distribute, under the [MIT license](https://github.com/marcw/ezmail/blob/master/LICENSE).

[![Build Status](https://travis-ci.org/marcw/ezmail.png?branch=master)](https://travis-ci.org/marcw/ezmail)
