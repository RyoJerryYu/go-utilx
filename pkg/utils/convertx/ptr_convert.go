package convertx

func BoolPtr(b bool) *bool                   { return &b }
func StrPtr(s string) *string                { return &s }
func IntPtr(i int) *int                      { return &i }
func Int8Ptr(i int8) *int8                   { return &i }
func Int16Ptr(i int16) *int16                { return &i }
func Int32Ptr(i int32) *int32                { return &i }
func Int64Ptr(i int64) *int64                { return &i }
func UIntPtr(i uint) *uint                   { return &i }
func UInt8Ptr(i uint8) *uint8                { return &i }
func UInt16Ptr(i uint16) *uint16             { return &i }
func UInt32Ptr(i uint32) *uint32             { return &i }
func UInt64Ptr(i uint64) *uint64             { return &i }
func UIntPtrPtr(i uintptr) *uintptr          { return &i }
func BytePtr(b byte) *byte                   { return &b }
func RunePtr(r rune) *rune                   { return &r }
func Float32Ptr(f float32) *float32          { return &f }
func Float64Ptr(f float64) *float64          { return &f }
func Complex64Ptr(c complex64) *complex64    { return &c }
func Complex128Ptr(c complex128) *complex128 { return &c }
func AnyPtr(a any) *any                      { return &a }
