package main

import (
	"log"
	"webhook-sensors/adapters/repo"
	"webhook-sensors/infraestructure/adapters/fcm"
	"webhook-sensors/infraestructure/webhook"
	"webhook-sensors/simulations"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
        fcmSender, err := fcm.NewFCMSender("serviceAccountKey.json")
        if err != nil {
                log.Fatalf("Error inicializando FCM: %v\n", err)
        }
        sensorRepo := repo.NewInMemorySensorRepo()


        webhookHandler := webhook.NewWebhookHandler(fcmSender, sensorRepo)

        r := gin.Default()

        config := cors.DefaultConfig()
        config.AllowAllOrigins = true

        r.Use(cors.New(config))

        r.POST("/webhook", webhookHandler.HandleSensorData)

        go simulations.SimulateSensors("http://localhost:8080/webhook")

        if err := r.Run(":8080"); err != nil {
                log.Fatalf("Error iniciando el servidor: %v\n", err)
        }
}
