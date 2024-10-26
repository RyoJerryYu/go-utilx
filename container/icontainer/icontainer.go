package icontainer

type Container[T any] interface {
	Len() int
	IsEmpty() bool
	Clear()
	Has(v T) bool
	Add(vs ...T) // Add vs to container
	Del(vs ...T) // Remove vs from container
}

type (
	BoolContainer       = Container[bool]
	StrContainer        = Container[string]
	IntContainer        = Container[int]
	Int8Container       = Container[int8]
	Int16Container      = Container[int16]
	Int32Container      = Container[int32]
	Int64Container      = Container[int64]
	UIntContainer       = Container[uint]
	UInt8Container      = Container[uint8]
	UInt16Container     = Container[uint16]
	UInt32Container     = Container[uint32]
	UInt64Container     = Container[uint64]
	UIntPtrContainer    = Container[uintptr]
	ByteContainer       = Container[byte]
	RuneContainer       = Container[rune]
	Float32Container    = Container[float32]
	Float64Container    = Container[float64]
	Complex64Container  = Container[complex64]
	Complex128Container = Container[complex128]
)
