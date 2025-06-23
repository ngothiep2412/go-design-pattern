package main

import "fmt"

type Notifier interface {
	SendNotification(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) SendNotification(message string) {
	fmt.Printf("Sending email: %s\n", message)
}

type SlackNotifier struct{}

func (SlackNotifier) SendNotification(message string) {
	fmt.Printf("Sending slack: %s\n", message)
}

type NotificationService struct {
	notifier Notifier
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.SendNotification(message)
}

func main() {
	s := NotificationService{
		// I don;t want my users init a new notifier like this.
		// They should call to something to produce a notifier with its specific type
		// CreateNotifier(type) Notifier
		notifier: SlackNotifier{},
	}

	s.SendNotification("Hello World")
}
