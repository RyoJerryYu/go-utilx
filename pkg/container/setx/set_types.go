package setx

import "github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"

// This file defines type aliases and constructors for common set types.
// Each type alias is a specialized version of Set[T] for a specific type.
// Each New* function creates an empty set of the corresponding type.

type (
	BoolSet       = Set[bool]
	StrSet        = Set[string]
	IntSet        = Set[int]
	Int8Set       = Set[int8]
	Int16Set      = Set[int16]
	Int32Set      = Set[int32]
	Int64Set      = Set[int64]
	UIntSet       = Set[uint]
	UInt8Set      = Set[uint8]
	UInt16Set     = Set[uint16]
	UInt32Set     = Set[uint32]
	UInt64Set     = Set[uint64]
	UIntPtrSet    = Set[uintptr]
	ByteSet       = Set[byte]
	RuneSet       = Set[rune]
	Float32Set    = Set[float32]
	Float64Set    = Set[float64]
	Complex64Set  = Set[complex64]
	Complex128Set = Set[complex128]
)

func NewBool() BoolSet             { return New[bool]() }
func NewStr() StrSet               { return New[string]() }
func NewInt() IntSet               { return New[int]() }
func NewInt8() Int8Set             { return New[int8]() }
func NewInt16() Int16Set           { return New[int16]() }
func NewInt32() Int32Set           { return New[int32]() }
func NewInt64() Int64Set           { return New[int64]() }
func NewUInt() UIntSet             { return New[uint]() }
func NewUInt8() UInt8Set           { return New[uint8]() }
func NewUInt16() UInt16Set         { return New[uint16]() }
func NewUInt32() UInt32Set         { return New[uint32]() }
func NewUInt64() UInt64Set         { return New[uint64]() }
func NewUIntPtr() UIntPtrSet       { return New[uintptr]() }
func NewByte() ByteSet             { return New[byte]() }
func NewRune() RuneSet             { return New[rune]() }
func NewFloat32() Float32Set       { return New[float32]() }
func NewFloat64() Float64Set       { return New[float64]() }
func NewComplex64() Complex64Set   { return New[complex64]() }
func NewComplex128() Complex128Set { return New[complex128]() }

var (
	_ icontainer.Container[bool]       = (BoolSet)(nil)
	_ icontainer.Container[string]     = (StrSet)(nil)
	_ icontainer.Container[int]        = (IntSet)(nil)
	_ icontainer.Container[int8]       = (Int8Set)(nil)
	_ icontainer.Container[int16]      = (Int16Set)(nil)
	_ icontainer.Container[int32]      = (Int32Set)(nil)
	_ icontainer.Container[int64]      = (Int64Set)(nil)
	_ icontainer.Container[uint]       = (UIntSet)(nil)
	_ icontainer.Container[uint8]      = (UInt8Set)(nil)
	_ icontainer.Container[uint16]     = (UInt16Set)(nil)
	_ icontainer.Container[uint32]     = (UInt32Set)(nil)
	_ icontainer.Container[uint64]     = (UInt64Set)(nil)
	_ icontainer.Container[uintptr]    = (UIntPtrSet)(nil)
	_ icontainer.Container[byte]       = (ByteSet)(nil)
	_ icontainer.Container[rune]       = (RuneSet)(nil)
	_ icontainer.Container[float32]    = (Float32Set)(nil)
	_ icontainer.Container[float64]    = (Float64Set)(nil)
	_ icontainer.Container[complex64]  = (Complex64Set)(nil)
	_ icontainer.Container[complex128] = (Complex128Set)(nil)
)
