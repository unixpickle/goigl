package iglcgal

// #cgo CPPFLAGS: -I../cvendor/libigl/include -I../cvendor/eigen -DBOOST_BIND_GLOBAL_PLACEHOLDERS
// #cgo LDFLAGS: -lgmp -lmpfr
// #include "csg.h"
// #include "stdlib.h"
import "C"

import (
	"unsafe"

	"github.com/unixpickle/goigl"
)

type MeshBooleanType int

const (
	Union MeshBooleanType = iota
	Intersect
	Minus
	Xor
	Resolve
)

// MeshBoolean performs a robust boolean operation on the pair of meshes.
// This can be composed multiple times to perform more general operations on a
// larger set of meshes.
func MeshBoolean(m1, m2 *goigl.Mesh, t MeshBooleanType) *goigl.Mesh {
	m1.Check()
	m2.Check()
	ptr := C.mesh_boolean((*C.mesh_t)(m1.Pointer()), (*C.mesh_t)(m2.Pointer()), C.int(t))
	return goigl.NewMeshPointer(unsafe.Pointer(ptr))
}
