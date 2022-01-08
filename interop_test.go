package goigl

import (
	"testing"

	"github.com/unixpickle/model3d/model3d"
)

func TestMeshInterop(t *testing.T) {
	start := model3d.NewMeshRect(model3d.XYZ(-1, -2, -3), model3d.XYZ(3, 4, 5))
	intermediate := NewMeshModel3d(start)
	end := intermediate.Model3d()

	if !meshesEqual(start, end) {
		t.Fatal("meshes not equal")
	}
}

func meshesEqual(m1, m2 *model3d.Mesh) bool {
	seg1 := meshOrderedSegments(m1)
	seg2 := meshOrderedSegments(m2)
	if len(seg1) != len(seg2) {
		return false
	}
	for s, c := range seg1 {
		if seg2[s] != c {
			return false
		}
	}
	return true
}

func meshOrderedSegments(m *model3d.Mesh) map[[2]model3d.Coord3D]int {
	res := map[[2]model3d.Coord3D]int{}
	m.Iterate(func(t *model3d.Triangle) {
		for i := 0; i < 3; i++ {
			seg := [2]model3d.Coord3D{t[i], t[(i+1)%3]}
			res[seg]++
		}
	})
	return res
}
