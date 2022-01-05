package igl

// #cgo CPPFLAGS: -Ivendor/libigl/include -Ivendor/eigen
// #include "mesh.h"
import "C"
import (
	"errors"
	"runtime"
)

// Mesh stores an array of faces and vertices.
// Memory should automatically be released when this object is garbage
// collected, but the memory can be freed before then by calling Clear().
type Mesh struct {
	valid bool
	ptr   *C.mesh_t
}

// MeshDecodeSTL decodes data from an STL file into a Mesh.
func MeshDecodeSTL(data []byte) (*Mesh, error) {
	ptr := C.mesh_decode_stl(C.CString(string(data)), C.size_t(len(data)))
	if ptr == nil {
		return nil, errors.New("failed to decode STL data")
	}
	res := &Mesh{valid: true, ptr: ptr}
	runtime.SetFinalizer(res, (*Mesh).Delete)
	return res, nil
}

// NumFaces returns the total number of faces in the mesh.
func (m *Mesh) NumFaces() int {
	return int(C.mesh_num_faces(m.ptr))
}

// NumVertices returns the total number of vertices in the mesh.
func (m *Mesh) NumVertices() int {
	return int(C.mesh_num_vertices(m.ptr))
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
