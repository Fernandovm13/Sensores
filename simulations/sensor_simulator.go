package simulations

import (
    "bytes"
    "encoding/json"
    "math/rand"
    "net/http"
    "time"
)

func SimulateSensors(webhookURL string) {
    sensors := []string{"sound", "temperature", "light"}
    rand.Seed(time.Now().UnixNano())

    for {
        for _, sensor := range sensors {
            var value float64
            switch sensor {
            case "sound":
                value = 25 + rand.Float64()*20
            case "temperature":
                value = 18 + rand.Float64()*10
            case "light":
                value = 200 + rand.Float64()*600
            }

            data := map[string]interface{}{
                "sensorType": sensor,
                "value":      value,
            }
            jsonData, _ := json.Marshal(data)

            resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
            if err != nil {
                println("Error enviando datos al webhook:", err.Error())
            } else {
                println("Datos enviados:", string(jsonData))
                resp.Body.Close()
            }
        }
        time.Sleep(70 * time.Second) 
    }
}