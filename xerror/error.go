package xerror

// Error is custom interface for advance error struct on golang
// Normal error can be convert or cast to this special Error
// via `xerror.New(key, name, message)` or `xerror.Cast(<error>)`.
type Error interface {
	Key() int
	Name() string
	Error() string
}
