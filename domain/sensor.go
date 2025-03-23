package domain

//  representa una lectura individual de un sensor.
type SensorData struct {
    SensorType string  `json:"sensorType"`
    Value      float64 `json:"value"`
}

// representa un conjunto de lecturas de sensores enviado por el ESP32.
type SensorAggregateData struct {
    Temperature float64 `json:"temperature"`
    Humidity    float64 `json:"humidity"`
    Light       float64 `json:"light"`
    Sound       float64 `json:"sound"`
    AirQuality  float64 `json:"airQuality"`
}

func ValidateSensor(data SensorData) (bool, string) {
    switch data.SensorType {
    case "sound":
        if data.Value < 30 || data.Value > 40 {
            return true, "⚠️ Nivel de sonido fuera de rango permitido (30dB - 40dB)"
        }
    case "temperature":
        if data.Value < 21 || data.Value > 23 {
            return true, "🌡️ Temperatura fuera de rango óptimo (21°C - 23°C)"
        }
    case "light":
        if data.Value < 300 || data.Value > 500 {
            return true, "💡 Nivel de iluminación fuera de rango (300 - 500 lux)"
        }
    case "humidity":
        if data.Value < 30 || data.Value > 60 {
            return true, "💧 Humedad fuera de rango óptimo (30% - 60%)"
        }
    case "airQuality":
        if data.Value < 50 || data.Value > 100 {
            return true, "🌫️ Calidad de aire fuera de rango (50% - 100%)"
        }
    }
    return false, ""
}
 