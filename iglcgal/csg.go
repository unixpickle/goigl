package iglcgal

// #cgo CPPFLAGS: -I../vendor/libigl/include -I../vendor/eigen -DBOOST_BIND_GLOBAL_PLACEHOLDERS
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

func MeshBoolean(m1, m2 *goigl.Mesh, t MeshBooleanType) *goigl.Mesh {
	m1.Check()
	m2.Check()
	ptr := C.mesh_boolean((*C.mesh_t)(m1.Pointer()), (*C.mesh_t)(m2.Pointer()), C.int(t))
	return goigl.NewMeshPointer(unsafe.Pointer(ptr))
}
