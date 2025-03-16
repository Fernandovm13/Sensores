package webhook

import (
	"net/http"

	"webhook-sensors/domain"
	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
    notificationSender domain.NotificationSender
}

func NewWebhookHandler(sender domain.NotificationSender) *WebhookHandler {
    return &WebhookHandler{notificationSender: sender}
}

func (h *WebhookHandler) HandleSensorData(c *gin.Context) {
    var sensorData domain.SensorData
    if err := c.ShouldBindJSON(&sensorData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    isOutOfRange, message := domain.ValidateSensor(sensorData)
    if isOutOfRange {
        if err := h.notificationSender.SendNotification(message); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error enviando notificación"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Datos procesados correctamente"})
}