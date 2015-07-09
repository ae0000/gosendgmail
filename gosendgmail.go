package gosendgmail

import (
	"errors"
	"net/smtp"
)

// Gmail sends emails..... via.... GMAIL
type Gmail struct {
	Address  string
	Password string
	SMTP     string
	Port     string
}

// WeOK checks to make sure the struct has been setup correctly..
// TODO Maybe add some more validation.. but then again, the email will just
// fail so leave it as is... SISO
func (g *Gmail) WeOK() error {
	if len(g.Address) == 0 {
		return errors.New("Need an email address (Gmail.Address)")
	}
	if len(g.Password) == 0 {
		return errors.New("Need an email password (Gmail.Password)")
	}
	if len(g.SMTP) == 0 {
		return errors.New("Need a SMTP host (Gmail.SMTP)")
	}
	if len(g.Port) == 0 {
		return errors.New("Need a SMPT port (Gmail.Port)")
	}

	return nil
}

// Send an email, like this:
//     g.Send([]string("blah@blah"), "Message")
func (g *Gmail) Send(to []string, msg string) error {
	// Check we have what we need
	if err := g.WeOK(); err != nil {
		return err
	}

	auth := smtp.PlainAuth(
		"",
		g.Address,
		g.Password,
		g.SMTP,
	)

	//  sender := "fromwho@gmail.com" // change here

	// send out the email
	err := smtp.SendMail(
		g.SMTP+":"+g.Port,
		auth,
		g.Address,
		to,
		[]byte(msg),
	)

	return err
}
