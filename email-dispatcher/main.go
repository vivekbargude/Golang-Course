package main

import (
	"bytes"
	"html/template"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	recipientChannel := make(chan Recipient)

	go func() {
		loadRecipient("./email.csv", recipientChannel)
	}()

	var wg sync.WaitGroup

	workerCount := 5
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}

	wg.Wait()

}


func executeTemplate(recipient Recipient) (string, error) {
	// Placeholder for template execution logic
	t, err := template.ParseFiles("email.tmpl")

	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, recipient)

	if err != nil {
		return "", err
	}
	// In a real implementation, this would parse and execute the template with the recipient data
	return tpl.String(), nil
}
