package store

type Store interface {
	Connectable
	Queryable
	Modifiable
	Closeable
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
