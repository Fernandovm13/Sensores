package fcm

import (
    "context"
    "log"

    "firebase.google.com/go"
    "firebase.google.com/go/messaging"
    "google.golang.org/api/option"
)

// FCMSender implementa NotificationSender para FCM
type FCMSender struct {
    client *messaging.Client
}

// NewFCMSender inicializa un nuevo cliente FCM
func NewFCMSender(credentialsFile string) (*FCMSender, error) {
    app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(credentialsFile))
    if err != nil {
        return nil, err
    }

    client, err := app.Messaging(context.Background())
    if err != nil {
        return nil, err
    }

    return &FCMSender{client: client}, nil
}

// SendNotification envía una notificación a través de FCM
func (f *FCMSender) SendNotification(message string) error {
    notification := &messaging.Message{
        Notification: &messaging.Notification{
            Title: "Alerta de Sensor",
            Body:  message,
        },
        Topic: "sensor_alerts",
    }

    _, err := f.client.Send(context.Background(), notification)
    if err != nil {
        log.Printf("Error enviando notificación: %v\n", err)
        return err
    }

    log.Println("Notificación enviada con éxito:", message)
    return nil
}