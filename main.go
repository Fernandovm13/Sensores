package main

import (
    "log"
    "webhook-sensors/adapters/fcm"
    "webhook-sensors/adapters/repo"
    "webhook-sensors/adapters/webhook"
    "webhook-sensors/simulations"

    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializar FCM
    fcmSender, err := fcm.NewFCMSender("serviceAccountKey.json")
    if err != nil {
        log.Fatalf("Error inicializando FCM: %v\n", err)
    }

    // Crear repositorio en memoria
    sensorRepo := repo.NewInMemorySensorRepo()

    // Configurar el webhook
    webhookHandler := webhook.NewWebhookHandler(fcmSender, sensorRepo)

    // Configurar el servidor Gin
    r := gin.Default()
    r.POST("/webhook", webhookHandler.HandleSensorData)
    r.GET("/readings", webhookHandler.GetSensorReadings)

    // Iniciar la simulación de sensores
    go simulations.SimulateSensors("http://localhost:8080/webhook")

    // Iniciar el servidor
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error iniciando el servidor: %v\n", err)
    }
}
