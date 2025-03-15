package webhook

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "webhook-sensors/domain"
    "webhook-sensors/ports"
)

// WebhookHandler maneja las solicitudes del webhook
type WebhookHandler struct {
    notificationSender ports.NotificationSender
}

// NewWebhookHandler crea un nuevo manejador de webhook
func NewWebhookHandler(sender ports.NotificationSender) *WebhookHandler {
    return &WebhookHandler{notificationSender: sender}
}

// HandleSensorData procesa los datos del sensor
func (h *WebhookHandler) HandleSensorData(c *gin.Context) {
    var sensorData domain.SensorData
    if err := c.ShouldBindJSON(&sensorData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    // Validar el valor del sensor
    isOutOfRange, message := domain.ValidateSensor(sensorData)
    if isOutOfRange {
        // Enviar notificación si el valor está fuera de rango
        if err := h.notificationSender.SendNotification(message); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error enviando notificación"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Datos procesados correctamente"})
}