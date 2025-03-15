package domain

// SensorData representa los datos recibidos de un sensor
type SensorData struct {
    SensorType string  `json:"sensorType"` // Tipo de sensor (sound, temperature, light)
    Value      float64 `json:"value"`      // Valor del sensor
}

// ValidateSensor verifica si el valor del sensor está fuera de rango por al menos 1 entero
func ValidateSensor(data SensorData) (bool, string) {
    switch data.SensorType {
    case "sound":
        if data.Value < 30 || data.Value > 40 {
            if data.Value <= 29 || data.Value >= 41 { // Fuera de rango por al menos 1 entero
                return true, "Nivel de sonido fuera de rango permitido (30dB - 40dB)"
            }
        }
    case "temperature":
        if data.Value < 21 || data.Value > 23 {
            if data.Value <= 20 || data.Value >= 24 { // Fuera de rango por al menos 1 entero
                return true, "Temperatura fuera de rango agradable (21°C - 23°C)"
            }
        }
    case "light":
        if data.Value < 300 || data.Value > 500 {
            if data.Value <= 299 || data.Value >= 501 { // Fuera de rango por al menos 1 entero
                return true, "Nivel de iluminación fuera de rango óptimo (300 lux - 500 lux)"
            }
        }
    default:
        return false, "Tipo de sensor no reconocido"
    }
    return false, ""
}