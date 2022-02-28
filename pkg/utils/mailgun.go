package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	mailgun "github.com/mailgun/mailgun-go/v4"
)

func SendHTML(to string, subject string, html string) error {

	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_SECRET"))
	sender := "no-reply@myponyasia.com"
	message := mg.NewMessage(sender, subject, "", to)
	message.SetHtml(html)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)

	return nil
}
