package dtmodels

type Converter[T any] interface {
	WithBool() (T, bool)
	// B is alias of WithBool()
	B() (T, bool)

	WithPointerBool() (*T, bool)
	// PB is alias of WithPointerBool()
	PB() (*T, bool)

	WithError() (T, error)
	// E is alias of WithError()
	E() (T, error)

	WithPointerError() (*T, error)
	// PE is alias of WithPointerError()
	PE() (*T, error)

	WithDefault(def T) T
	// D is alias of WithDefault()
	D(def T) T

	WithPointerDefault(def T) *T
	// PE is alias of WithPointerError()
	PD(def T) *T

	Force() T
	// F is alias of Force()
	F() T
}
