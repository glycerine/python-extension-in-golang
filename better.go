// Package sum
package main

// #cgo pkg-config: python-3.7
// #define Py_LIMITED_API
// #include <Python.h>
import "C"

//export Sum
func Sum(a, b int) int {
	return (a + b)
}

//export NewDictionary
func NewDictionary() *C.PyObject {
	var dict = C.PyDict_New()
	C.PyDict_SetItem(
		dict,
		C.PyUnicode_FromString(C.CString("key")),
		C.PyUnicode_FromString(C.CString("value")))
	return dict
}

//export SumOverSlice
func SumOverSlice(sequence *C.PyObject) (total C.long) {

	sequenceLen := int(C.PySequence_Length(sequence))

	for i := 0; i < sequenceLen; i++ {
		element := C.PySequence_GetItem(sequence, C.Py_ssize_t(i))
		val := C.PyLong_AsLong(element)
		total += val
	}
	return
}

func main() {
}
