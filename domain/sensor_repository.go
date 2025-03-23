package domain

type SensorRepository interface {
    Store(sensor SensorData) error
    ListAll() ([]SensorData, error)
    StoreAggregate(aggregate SensorAggregateData) error
}
