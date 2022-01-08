package goigl

// #cgo CPPFLAGS: -Icvendor/libigl/include -Icvendor/eigen
// #include "mesh.h"
// #include "stdlib.h"
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

// Mesh stores an array of faces and vertices.
// Memory should automatically be released when this object is garbage
// collected, but the memory can be freed before then by calling Clear().
type Mesh struct {
	ptr *C.mesh_t
}

// NewMesh creates a mesh from faces and vertices.
func NewMesh(vertices []float64, faces []int) *Mesh {
	vsPtr := C.malloc(C.size_t(len(vertices)) * C.size_t(unsafe.Sizeof(C.double(0))))
	if vsPtr == nil {
		panic("allocation failed")
	}
	a := (*[1<<30 - 1]C.double)(vsPtr)
	for i, x := range vertices {
		a[i] = C.double(x)
	}

	fsPtr := C.malloc(C.size_t(len(faces)) * C.size_t(unsafe.Sizeof(C.int(0))))
	if fsPtr == nil {
		C.free(vsPtr)
		panic("allocation failed")
	}
	b := (*[1<<30 - 1]C.int)(fsPtr)
	for i, x := range faces {
		b[i] = C.int(x)
	}

	res := &Mesh{ptr: C.mesh_new(
		(*C.double)(vsPtr),
		C.size_t(len(vertices)),
		(*C.int)(fsPtr),
		C.size_t(len(faces)),
	)}
	C.free(vsPtr)
	C.free(fsPtr)
	runtime.SetFinalizer(res, (*Mesh).Delete)
	return res
}

// NewMeshPointer creates a Mesh from a backing C object.
//
// The backing C object will be owned by the result and freed by a finalizer.
func NewMeshPointer(ptr unsafe.Pointer) *Mesh {
	res := &Mesh{ptr: (*C.mesh_t)(ptr)}
	runtime.SetFinalizer(res, (*Mesh).Delete)
	return res
}

// MeshDecodeSTL decodes data from an STL file into a Mesh.
func MeshDecodeSTL(data []byte) (*Mesh, error) {
	cstring := C.CString(string(data))
	var errorOut *C.char
	ptr := C.mesh_decode_stl(cstring, C.size_t(len(data)), &errorOut)
	C.free(unsafe.Pointer(cstring))
	if ptr == nil {
		errorMsg := "decode STL: " + C.GoString(errorOut)
		C.free(unsafe.Pointer(errorOut))
		return nil, errors.New(errorMsg)
	}
	res := &Mesh{ptr: ptr}
	runtime.SetFinalizer(res, (*Mesh).Delete)
	return res, nil
}

// Vertices creates a copy of the vertices array.
func (m *Mesh) Vertices() []float64 {
	m.Check()
	data := C.mesh_vertices(m.ptr)
	carr := (*[1<<30 - 1]C.double)(unsafe.Pointer(data))[:C.mesh_vertices_size(m.ptr)]
	result := make([]float64, len(carr))
	for i, x := range carr {
		result[i] = float64(x)
	}
	C.free(unsafe.Pointer(data))
	return result
}

// NumVertices returns the total number of vertices in the mesh.
func (m *Mesh) NumVertices() int {
	m.Check()
	if m.ptr == nil {
		panic("mesh has been freed")
	}
	return int(C.mesh_num_vertices(m.ptr))
}

// Faces creates a copy of the faces array.
func (m *Mesh) Faces() []int {
	m.Check()
	data := C.mesh_faces(m.ptr)
	carr := (*[1 << 32]C.int)(unsafe.Pointer(data))[:C.mesh_faces_size(m.ptr)]
	result := make([]int, len(carr))
	for i, x := range carr {
		result[i] = int(x)
	}
	C.free(unsafe.Pointer(data))
	return result
}

// NumFaces returns the total number of faces in the mesh.
func (m *Mesh) NumFaces() int {
	m.Check()
	return int(C.mesh_num_faces(m.ptr))
}

// Delete frees the memory associated with m. Future operations on m will
// result in a panic().
//
// Calling this explicitly isn't necessary, since the mesh is deallocated by
// the garbage collector automatically.
func (m *Mesh) Delete() {
	if m.ptr != nil {
		C.mesh_free(m.ptr)
		m.ptr = nil
	}
}

// Pointer gets a pointer to the C data structure backing this object.
func (m *Mesh) Pointer() unsafe.Pointer {
	return unsafe.Pointer(m.ptr)
}

// Check panic()s if the mesh has already been Delete()ed, or is a no-op
// otherwise.
func (m *Mesh) Check() {
	if m.ptr == nil {
		panic("mesh has been freed")
	}
}

// WriteSTL writes the mesh to an STL file.
func (m *Mesh) WriteSTL(path string) error {
	cstr := C.CString(path)
	msg := C.mesh_write_stl(m.ptr, cstr)
	C.free(unsafe.Pointer(cstr))
	var err error
	if msg != nil {
		err = errors.New("write STL: " + C.GoString(msg))
		C.free(unsafe.Pointer(msg))
	}
	return err
}
