package main

import "fmt"

// Vi phạm nguyên lí S (single) và O (open/closed)
// S -> 1 lớp chỉ neên chứa 1 logic
// O -> ko thể sửa đổi, mở rộng
type NotificationService struct {
	notifierType string
}

func (s NotificationService) SendNotification(message string) {
	if s.notifierType == "email" {
		fmt.Printf("Sending email: %s\n", message)
	} else if s.notifierType == "slack" {
		fmt.Printf("Sending slack: %s\n", message)
	}
}

func main() {
	s := NotificationService{"email"}
	s.SendNotification("hello world")
}
