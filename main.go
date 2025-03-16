package main

import (
<<<<<<< HEAD
	"log"
	"webhook-sensors/infraestructure/adapters/fcm"
	"webhook-sensors/infraestructure/webhook"
	"webhook-sensors/simulations"
=======
    "log"
    "webhook-sensors/adapters/fcm"
    "webhook-sensors/adapters/repo"
    "webhook-sensors/adapters/webhook"
    "webhook-sensors/simulations"
>>>>>>> rama_Fernando

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
        fcmSender, err := fcm.NewFCMSender("serviceAccountKey.json")
        if err != nil {
                log.Fatalf("Error inicializando FCM: %v\n", err)
        }

<<<<<<< HEAD
        webhookHandler := webhook.NewWebhookHandler(fcmSender)

        r := gin.Default()
=======
    // Crear repositorio en memoria
    sensorRepo := repo.NewInMemorySensorRepo()

    // Configurar el webhook
    webhookHandler := webhook.NewWebhookHandler(fcmSender, sensorRepo)

    // Configurar el servidor Gin
    r := gin.Default()
    r.POST("/webhook", webhookHandler.HandleSensorData)
    r.GET("/readings", webhookHandler.GetSensorReadings)
>>>>>>> rama_Fernando

        config := cors.DefaultConfig()
        config.AllowAllOrigins = true

<<<<<<< HEAD
        r.Use(cors.New(config))

        r.POST("/webhook", webhookHandler.HandleSensorData)

        go simulations.SimulateSensors("http://localhost:8080/webhook")

        if err := r.Run(":8080"); err != nil {
                log.Fatalf("Error iniciando el servidor: %v\n", err)
        }
}
=======
    // Iniciar el servidor
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error iniciando el servidor: %v\n", err)
    }
}
>>>>>>> rama_Fernando
