package ffi

// #include <plugin.h>
import "C"

func result(size int32) ([]byte, HostErr) {
	allocSize := size

	if size < 0 {
		if size == -1 {
			return nil, NewHostError("unknown error returned from host")
		}

		allocSize = -size
	}

	result := make([]byte, allocSize)
	resultPtr, _ := unsafeSlicePointer(result)

	if code := C.get_ffi_result(resultPtr, Ident()); code != 0 {
		return nil, NewHostError("unknown error returned from host")
	}

	if size < 0 {
		return nil, NewHostError(string(result))
	}

	return result, nil
}

func addVar(name, value string) {
	namePtr, nameSize := unsafeSlicePointer([]byte(name))
	valuePtr, valueSize := unsafeSlicePointer([]byte(value))

	C.add_ffi_var(namePtr, nameSize, valuePtr, valueSize, Ident())
}
