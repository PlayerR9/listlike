//go:generate go run cmd/stack/main.go -name=LinkedStack -type=T -g=T/any -o=stack/linked_stack_generic.go
//go:generate go run cmd/stack/main.go -name=BoolStack -type=bool -o=stack/linked_stack_bool.go
//go:generate go run cmd/stack/main.go -name=ByteStack -type=byte -o=stack/linked_stack_byte.go
//go:generate go run cmd/stack/main.go -name=Complex64Stack -type=complex64 -o=stack/linked_stack_complex64.go
//go:generate go run cmd/stack/main.go -name=Complex128Stack -type=complex128 -o=stack/linked_stack_complex128.go
//go:generate go run cmd/stack/main.go -name=ErrorStack -type=error -o=stack/linked_stack_error.go
//go:generate go run cmd/stack/main.go -name=Float32Stack -type=float32 -o=stack/linked_stack_float32.go
//go:generate go run cmd/stack/main.go -name=Float64Stack -type=float64 -o=stack/linked_stack_float64.go
//go:generate go run cmd/stack/main.go -name=IntStack -type=int -o=stack/linked_stack_int.go
//go:generate go run cmd/stack/main.go -name=Int8Stack -type=int8 -o=stack/linked_stack_int8.go
//go:generate go run cmd/stack/main.go -name=Int16Stack -type=int16 -o=stack/linked_stack_int16.go
//go:generate go run cmd/stack/main.go -name=Int32Stack -type=int32 -o=stack/linked_stack_int32.go
//go:generate go run cmd/stack/main.go -name=Int64Stack -type=int64 -o=stack/linked_stack_int64.go
//go:generate go run cmd/stack/main.go -name=RuneStack -type=rune -o=stack/linked_stack_rune.go
//go:generate go run cmd/stack/main.go -name=StringStack -type=string -o=stack/linked_stack_string.go
//go:generate go run cmd/stack/main.go -name=UintStack -type=uint -o=stack/linked_stack_uint.go
//go:generate go run cmd/stack/main.go -name=Uint8Stack -type=uint8 -o=stack/linked_stack_uint8.go
//go:generate go run cmd/stack/main.go -name=Uint16Stack -type=uint16 -o=stack/linked_stack_uint16.go
//go:generate go run cmd/stack/main.go -name=Uint32Stack -type=uint32 -o=stack/linked_stack_uint32.go
//go:generate go run cmd/stack/main.go -name=Uint64Stack -type=uint64 -o=stack/linked_stack_uint64.go
//go:generate go run cmd/stack/main.go -name=UintptrStack -type=uintptr -o=stack/linked_stack_uintptr.go

package stack
