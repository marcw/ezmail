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

## Edge Cases

When sending emails through GMail SMTP servers, if you run into an error like:

    555 5.5.2 Syntax error. z6sm27198953bkn.8 - gsmtp

Try changing the smtp.SendMail call in order to provide only the email addresses (without full names)

	smtp.SendMail(addr, auth, msg.From.Address, msg.RecipientsEmails(), msg.Bytes())

## License

The ezmail code is free to use and distribute, under the [MIT license](https://github.com/marcw/ezmail/blob/master/LICENSE).

[![Build Status](https://travis-ci.org/marcw/ezmail.png?branch=master)](https://travis-ci.org/marcw/ezmail)

