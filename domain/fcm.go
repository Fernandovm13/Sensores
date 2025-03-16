package domain

type NotificationSender interface {
    SendNotification(message string) error
}