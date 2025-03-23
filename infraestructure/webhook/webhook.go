package webhook

import (
    "net/http"
    "webhook-sensors/domain"

    "github.com/gin-gonic/gin"
)

type WebhookHandler struct {
    notificationSender domain.NotificationSender
    sensorRepo         domain.SensorRepository
}

func NewWebhookHandler(sender domain.NotificationSender, repo domain.SensorRepository) *WebhookHandler {
    return &WebhookHandler{
        notificationSender: sender,
        sensorRepo:         repo,
    }
}

func (h *WebhookHandler) HandleSensorData(c *gin.Context) {
    var aggregateData domain.SensorAggregateData
    if err := c.ShouldBindJSON(&aggregateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "❌ Datos inválidos"})
        return
    }

    sensors := []domain.SensorData{
        {SensorType: "temperature", Value: aggregateData.Temperature},
        {SensorType: "humidity", Value: aggregateData.Humidity},
        {SensorType: "light", Value: aggregateData.Light},
        {SensorType: "sound", Value: aggregateData.Sound},
        {SensorType: "airQuality", Value: aggregateData.AirQuality},
    }

    for _, sensor := range sensors {
        if isOutOfRange, msg := domain.ValidateSensor(sensor); isOutOfRange {
            h.notificationSender.SendNotification(msg)
        }
    }

    if err := h.sensorRepo.StoreAggregate(aggregateData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error almacenando datos"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "✅ Datos procesados correctamente"})
}

// devuelve todas las lecturas almacenadas.
func (h *WebhookHandler) GetSensorReadings(c *gin.Context) {
    readings, err := h.sensorRepo.ListAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo lecturas"})
        return
    }
    c.JSON(http.StatusOK, readings)
}
