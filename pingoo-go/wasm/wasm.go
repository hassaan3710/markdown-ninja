package wasm

// Buffer represents a pointer to a buffer allocated in a WASM module's memory
// Because we pack the pointer with its length, it currently only supports wasm32
// [ pointer (32 bits) | length (32 bits)]
type Buffer uint64

func (buffer Buffer) Pointer() uint32 {
	return uint32(buffer >> 32)
}

func (buffer Buffer) Size() uint32 {
	return uint32(buffer)
}

func NewBuffer(pointer, length uint32) Buffer {
	return Buffer((uint64(pointer) << 32) | uint64(length))
}

type Result[T any] struct {
	Ok    *T      `json:"ok,omitempty"`
	Error *string `json:"error,omitempty"`
}

// func HandleHostFunctionCall[I, O any](ctx context.Context, module wazeroapi.Module, input I) (output O, err error) {
// 	var ret O

// 	return ret, nil
// }
