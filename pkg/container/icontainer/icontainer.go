// Container is a generic interface that defines basic container operations.
// Type parameter T represents the type of elements stored in the container.
package icontainer

// Container defines the basic operations that all containers should implement.
type Container[T any] interface {
	// Len returns the number of elements in the container.
	Len() int
	// IsEmpty returns true if the container contains no elements.
	IsEmpty() bool
	// Clear removes all elements from the container.
	Clear()
	// ForEach executes the given function for each element in the container.
	ForEach(fn func(T))
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
