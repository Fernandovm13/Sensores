package repo

import (
    "sync"
    "webhook-sensors/domain"
)

type InMemorySensorRepo struct {
    data []domain.SensorData
    mu   sync.Mutex
}

// crea un repositorio en memoria.
func NewInMemorySensorRepo() *InMemorySensorRepo {
    return &InMemorySensorRepo{
        data: make([]domain.SensorData, 0),
    }
}

// almacena una lectura individual.
func (r *InMemorySensorRepo) Store(sensor domain.SensorData) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.data = append(r.data, sensor)
    return nil
}

func (r *InMemorySensorRepo) StoreAggregate(aggregate domain.SensorAggregateData) error {
    r.Store(domain.SensorData{SensorType: "temperature", Value: aggregate.Temperature})
    r.Store(domain.SensorData{SensorType: "humidity", Value: aggregate.Humidity})
    r.Store(domain.SensorData{SensorType: "light", Value: aggregate.Light})
    r.Store(domain.SensorData{SensorType: "sound", Value: aggregate.Sound})
    r.Store(domain.SensorData{SensorType: "airQuality", Value: aggregate.AirQuality})
    return nil
}

// devuelve todas las lecturas almacenadas.
func (r *InMemorySensorRepo) ListAll() ([]domain.SensorData, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    readings := make([]domain.SensorData, len(r.data))
    copy(readings, r.data)
    return readings, nil
}
