package ffi

// #include <plugin.h>
import "C"

func DoHTTPRequest(method int32, urlStr string, body []byte, headers map[string]string) ([]byte, error) {
	urlPtr, urlSize := unsafeSlicePointer([]byte(urlStr))
	bodyPtr, bodySize := unsafeSlicePointer(body)

	size := C.fetch_url(method, urlPtr, urlSize, bodyPtr, bodySize, Ident())

	return result(size)
}
