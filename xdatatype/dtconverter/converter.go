package dtconverter

type Converter[T any] interface {
	WithBool(input any) (T, bool)
	// B is alias of WithBool()
	B(input any) (T, bool)

	WithPointerBool(input any) (*T, bool)
	// PB is alias of WithPointerBool()
	PB(input any) (*T, bool)

	WithError(input any) (T, error)
	// E is alias of WithError()
	E(input any) (T, error)

	WithPointerError(input any) (*T, error)
	// PE is alias of WithPointerError()
	PE(input any) (*T, error)

	WithDefault(input any, def T) T
	// D is alias of WithDefault()
	D(input any, def T) T

	WithPointerDefault(input any, def T) *T
	// PE is alias of WithPointerError()
	PD(input any, def T) *T

	Force(input any) T
	// F is alias of Force()
	F(input any) T
}
