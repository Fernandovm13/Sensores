package domain

type SensorRepository interface {
    Store(SensorData) error
    ListAll() ([]SensorData, error)
}
