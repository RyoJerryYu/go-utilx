package slicex

import "github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"

type (
	BoolSlice       = Slice[bool]
	StrSlice        = Slice[string]
	IntSlice        = Slice[int]
	Int8Slice       = Slice[int8]
	Int16Slice      = Slice[int16]
	Int32Slice      = Slice[int32]
	Int64Slice      = Slice[int64]
	UIntSlice       = Slice[uint]
	UInt8Slice      = Slice[uint8]
	UInt16Slice     = Slice[uint16]
	UInt32Slice     = Slice[uint32]
	UInt64Slice     = Slice[uint64]
	UIntPtrSlice    = Slice[uintptr]
	ByteSlice       = Slice[byte]
	RuneSlice       = Slice[rune]
	Float32Slice    = Slice[float32]
	Float64Slice    = Slice[float64]
	Complex64Slice  = Slice[complex64]
	Complex128Slice = Slice[complex128]
)

func NewBool() BoolSlice             { return New[bool]() }
func NewStr() StrSlice               { return New[string]() }
func NewInt() IntSlice               { return New[int]() }
func NewInt8() Int8Slice             { return New[int8]() }
func NewInt16() Int16Slice           { return New[int16]() }
func NewInt32() Int32Slice           { return New[int32]() }
func NewInt64() Int64Slice           { return New[int64]() }
func NewUInt() UIntSlice             { return New[uint]() }
func NewUInt8() UInt8Slice           { return New[uint8]() }
func NewUInt16() UInt16Slice         { return New[uint16]() }
func NewUInt32() UInt32Slice         { return New[uint32]() }
func NewUInt64() UInt64Slice         { return New[uint64]() }
func NewUIntPtr() UIntPtrSlice       { return New[uintptr]() }
func NewByte() ByteSlice             { return New[byte]() }
func NewRune() RuneSlice             { return New[rune]() }
func NewFloat32() Float32Slice       { return New[float32]() }
func NewFloat64() Float64Slice       { return New[float64]() }
func NewComplex64() Complex64Slice   { return New[complex64]() }
func NewComplex128() Complex128Slice { return New[complex128]() }

var (
	_ icontainer.Container[bool]       = (*BoolSlice)(nil)
	_ icontainer.Container[string]     = (*StrSlice)(nil)
	_ icontainer.Container[int]        = (*IntSlice)(nil)
	_ icontainer.Container[int8]       = (*Int8Slice)(nil)
	_ icontainer.Container[int16]      = (*Int16Slice)(nil)
	_ icontainer.Container[int32]      = (*Int32Slice)(nil)
	_ icontainer.Container[int64]      = (*Int64Slice)(nil)
	_ icontainer.Container[uint]       = (*UIntSlice)(nil)
	_ icontainer.Container[uint8]      = (*UInt8Slice)(nil)
	_ icontainer.Container[uint16]     = (*UInt16Slice)(nil)
	_ icontainer.Container[uint32]     = (*UInt32Slice)(nil)
	_ icontainer.Container[uint64]     = (*UInt64Slice)(nil)
	_ icontainer.Container[uintptr]    = (*UIntPtrSlice)(nil)
	_ icontainer.Container[byte]       = (*ByteSlice)(nil)
	_ icontainer.Container[rune]       = (*RuneSlice)(nil)
	_ icontainer.Container[float32]    = (*Float32Slice)(nil)
	_ icontainer.Container[float64]    = (*Float64Slice)(nil)
	_ icontainer.Container[complex64]  = (*Complex64Slice)(nil)
	_ icontainer.Container[complex128] = (*Complex128Slice)(nil)
)
