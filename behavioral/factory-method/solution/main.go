package main

import "fmt"

type Notifier interface {
	SendNotification(string)
}

type SlackNotifier struct{}

func (s SlackNotifier) SendNotification(message string) {
	fmt.Printf("Sending slack: %s\n", message)
}

type EmailNotifier struct{}

func (e EmailNotifier) SendNotification(message string) {
	fmt.Printf("Sending email: %s\n", message)
}

type NotificationService struct {
	notifier Notifier
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.SendNotification(message)
}

func CreateNotifier(t string) Notifier {
	if t == "slack" {
		return SlackNotifier{}
	}

	return EmailNotifier{}
}

func main() {
	s := NotificationService{
		notifier: CreateNotifier("slack"),
	}

	s.notifier.SendNotification("Hello World")
}
