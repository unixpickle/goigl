package igl

// #cgo CPPFLAGS: -Ivendor/libigl/include -Ivendor/eigen
// #cgo LDFLAGS: -lgmp -lmpfr
// #include "csg.h"
// #include "stdlib.h"
import "C"
import "runtime"

type MeshBooleanType int

const (
	Union MeshBooleanType = iota
	Intersect
	Minus
	Xor
	Resolve
)

func MeshBoolean(m1, m2 *Mesh, t MeshBooleanType) *Mesh {
	m1.check()
	m2.check()
	ptr := C.mesh_boolean(m1.ptr, m2.ptr, C.int(t))
	res := &Mesh{ptr: ptr}
	runtime.SetFinalizer(res, (*Mesh).Delete)
	return res
}
