package store

import "time"

type Store interface {
	Connectable
	Queryable
	Modifiable
	Closeable
}

type Point struct {
	Time     time.Time
	Average  float64
	Minimum  float64
	Maximum  float64
	STD      float64
	Variance float64
}

type HistoricalData interface {
	LastNMinutes(float64) Point
	LastLocalMinimum() Point
	LastLocalMaximum() Point
}

type Connectable interface {
	Connect() error
}

type Queryable interface {
	Query(...string) (interface{}, error)
}

type Modifiable interface {
	Insert(...string) error
}
type Closeable interface {
	Close() error
}
