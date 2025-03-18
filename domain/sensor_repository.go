package domain

type SensorRepository interface {
    Store(sensor SensorData) error
    ListAll() ([]SensorData, error)
}
