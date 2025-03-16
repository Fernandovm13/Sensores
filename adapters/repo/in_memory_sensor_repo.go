package repo

import (
    "sync"
    "webhook-sensors/domain"
)

type InMemorySensorRepo struct {
    data []domain.SensorData
    mu   sync.Mutex
}

// NewInMemorySensorRepo crea un repositorio en memoria
func NewInMemorySensorRepo() *InMemorySensorRepo {
    return &InMemorySensorRepo{
        data: make([]domain.SensorData, 0),
    }
}

func (r *InMemorySensorRepo) Store(sensor domain.SensorData) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.data = append(r.data, sensor)
    return nil
}

func (r *InMemorySensorRepo) ListAll() ([]domain.SensorData, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    // Devuelve una copia para evitar que se modifique externamente
    readings := make([]domain.SensorData, len(r.data))
    copy(readings, r.data)
    return readings, nil
}
