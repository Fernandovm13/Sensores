package webhook

import (
	"net/http"

	"webhook-sensors/domain"
	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
<<<<<<< HEAD:infraestructure/webhook/webhook.go
    notificationSender domain.NotificationSender
}

func NewWebhookHandler(sender domain.NotificationSender) *WebhookHandler {
    return &WebhookHandler{notificationSender: sender}
=======
    notificationSender ports.NotificationSender
    sensorRepo         domain.SensorRepository
}

func NewWebhookHandler(sender ports.NotificationSender, repo domain.SensorRepository) *WebhookHandler {
    return &WebhookHandler{
        notificationSender: sender,
        sensorRepo:         repo,
    }
>>>>>>> rama_Fernando:adapters/webhook/webhook.go
}

func (h *WebhookHandler) HandleSensorData(c *gin.Context) {
    var sensorData domain.SensorData
    if err := c.ShouldBindJSON(&sensorData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

<<<<<<< HEAD:infraestructure/webhook/webhook.go
=======
    // Guardar la lectura
    if err := h.sensorRepo.Store(sensorData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error guardando lectura"})
        return
    }

    // Validar el valor del sensor
>>>>>>> rama_Fernando:adapters/webhook/webhook.go
    isOutOfRange, message := domain.ValidateSensor(sensorData)
    if isOutOfRange {
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
