package cdata

// BaseData is a basic data interface for all data on caches
type BaseData[T any] interface {
	// Key return data key on cache
	Key() string
	// Get return data pointer (and error if exist)
	Get() (*T, error)
	// Force return raw data
	// can panic if data is nil
	Force() T
	// Update data from modifier (if and only if data is expired),
	// will return true if data got updated successfully
	Update() bool
	// Del data will mark data as deleted (w/o actual removed)
	// this will also cause Update() to do nothing
	Del() bool
	// IsDel return true if data soft deleted
	IsDel() bool
	// IsErr return true if current data contains error
	IsErr() bool
	// IsExp return true if data is expired
	IsExp() bool
	// HasExp return true if data can be expired
	HasExp() bool
}
