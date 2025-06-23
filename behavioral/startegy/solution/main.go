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
	s := NotificationService{SlackNotifier{}}
	s.SendNotification("Hello World")
}

// startegy -> tăng độ phức tạp
// -> nếu logic ko thay đổi tgian thì ko cần xài
