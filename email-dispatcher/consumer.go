package main

import (
	"fmt"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {

	defer wg.Done()

	for recipient := range ch {

		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\nHello %s,\r\nThis is a test email.\r\n", recipient.Email, recipient.Name)
		// msg := []byte(formattedMsg)

		msg, err := executeTemplate(recipient)

		if err != nil {
			fmt.Printf("Worker %d: Failed to execute template for %s: %v\n", id, recipient.Email, err)
			continue
		}	

		// Simulate sending email
		fmt.Printf("Worker %d: Sending email to %s at %s\n", id, recipient.Name, recipient.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "vivekbargude@gmail.com", []string{recipient.Email}, []byte(msg))

		if err != nil {
			fmt.Printf("Worker %d: Failed to send email to %s: %v\n", id, recipient.Email, err)
		}

		time.Sleep(1 * time.Millisecond) // Simulate some delay

		// Simulate sending email
		fmt.Printf("Worker %d: Sent email to %s at %s\n", id, recipient.Name, recipient.Email)
		
	}

}
