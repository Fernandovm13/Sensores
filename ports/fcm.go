package ports

type NotificationSender interface {
    SendNotification(message string) error
}