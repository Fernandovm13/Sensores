package webhook

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "webhook-sensors/domain"
    "webhook-sensors/ports"
)

type WebhookHandler struct {
    notificationSender ports.NotificationSender
    sensorRepo         domain.SensorRepository
}

func NewWebhookHandler(sender ports.NotificationSender, repo domain.SensorRepository) *WebhookHandler {
    return &WebhookHandler{
        notificationSender: sender,
        sensorRepo:         repo,
    }
}

// HandleSensorData procesa los datos del sensor
func (h *WebhookHandler) HandleSensorData(c *gin.Context) {
    var sensorData domain.SensorData
    if err := c.ShouldBindJSON(&sensorData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    // Guardar la lectura
    if err := h.sensorRepo.Store(sensorData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error guardando lectura"})
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

// GetSensorReadings devuelve todas las lecturas almacenadas
func (h *WebhookHandler) GetSensorReadings(c *gin.Context) {
    readings, err := h.sensorRepo.ListAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo lecturas"})
        return
    }
    c.JSON(http.StatusOK, readings)
}
