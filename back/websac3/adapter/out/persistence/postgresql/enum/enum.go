package enum

import (
	"time"
	"websac3/app/port/out/persistence/db"
)

type GetPortBase[T any] interface {
	GetByName(name string, ctx db.Context) (T, error)
}

type cacheEntry[T any] struct {
	value      T
	expiryTime time.Time
}

type Enum[T any] struct {
	persistenceManager db.Manager
	getPortBase        GetPortBase[T]
	registry           map[string]cacheEntry[T]
	ttl                time.Duration
}

func New[T any](
	persistenceManager db.Manager,
	getPortBase GetPortBase[T],
	ttl time.Duration,
) *Enum[T] {
	return &Enum[T]{
		registry:           make(map[string]cacheEntry[T]),
		persistenceManager: persistenceManager,
		getPortBase:        getPortBase,
		ttl:                ttl,
	}
}

func (e *Enum[T]) GetByName(name string) (T, error) {
	var zeroValue T

	if entry, exists := e.registry[name]; exists {
		if time.Now().Before(entry.expiryTime) {
			return entry.value, nil
		}

		delete(e.registry, name)
	}

	var value T
	var err error = e.persistenceManager.ExecuteNonTransactional(
		func(tx db.Context) error {
			v, getErr := e.getPortBase.GetByName(name, tx)
			if getErr != nil {
				return getErr
			}
			value = v
			return nil
		},
	)

	if err != nil {
		return zeroValue, err
	}

	e.registry[name] = cacheEntry[T]{
		value:      value,
		expiryTime: time.Now().Add(e.ttl),
	}

	return value, nil
}
