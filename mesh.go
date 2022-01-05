package igl

// #cgo CPPFLAGS: -Ivendor/libigl/include -Ivendor/eigen
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

// MeshDecodeSTL decodes data from an STL file into a Mesh.
func MeshDecodeSTL(data []byte) (*Mesh, error) {
	ptr := C.mesh_decode_stl(C.CString(string(data)), C.size_t(len(data)))
	if ptr == nil {
		return nil, errors.New("failed to decode STL data")
	}
	res := &Mesh{ptr: ptr}
	runtime.SetFinalizer(res, (*Mesh).Delete)
	return res, nil
}

// Vertices creates a copy of the vertices array.
func (m *Mesh) Vertices() []float64 {
	m.check()
	data := C.mesh_vertices(m.ptr)
	carr := (*[1 << 32]C.double)(unsafe.Pointer(data))[:C.mesh_vertices_size(m.ptr)]
	result := make([]float64, len(carr))
	for i, x := range carr {
		result[i] = float64(x)
	}
	C.free(unsafe.Pointer(data))
	return result
}

// NumVertices returns the total number of vertices in the mesh.
func (m *Mesh) NumVertices() int {
	m.check()
	if m.ptr == nil {
		panic("mesh has been freed")
	}
	return int(C.mesh_num_vertices(m.ptr))
}

// Faces creates a copy of the faces array.
func (m *Mesh) Faces() []int {
	m.check()
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
	m.check()
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

func (m *Mesh) check() {
	if m.ptr == nil {
		panic("mesh has been freed")
	}
}
