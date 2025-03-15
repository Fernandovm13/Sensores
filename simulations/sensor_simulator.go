package simulations

import (
    "bytes"
    "encoding/json"
    "math/rand"
    "net/http"
    "time"
)

// SimulateSensors envía datos de sensores simulados al webhook
func SimulateSensors(webhookURL string) {
    sensors := []string{"sound", "temperature", "light"}
    rand.Seed(time.Now().UnixNano())

    for {
        for _, sensor := range sensors {
            var value float64
            switch sensor {
            case "sound":
                // Simula valores entre 25 y 45 dB (rango permitido: 30-40 dB)
                value = 25 + rand.Float64()*20
            case "temperature":
                // Simula valores entre 18 y 28 °C (rango permitido: 21-23 °C)
                value = 18 + rand.Float64()*10
            case "light":
                // Simula valores entre 200 y 800 lux (rango permitido: 300-500 lux)
                value = 200 + rand.Float64()*600
            }

            data := map[string]interface{}{
                "sensorType": sensor,
                "value":      value,
            }
            jsonData, _ := json.Marshal(data)

            // Enviar datos al webhook
            resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
            if err != nil {
                println("Error enviando datos al webhook:", err.Error())
            } else {
                println("Datos enviados:", string(jsonData))
                resp.Body.Close()
            }
        }
        time.Sleep(10 * time.Second) // Enviar datos cada 10 segundos
    }
}