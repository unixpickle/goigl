package goigl

import "github.com/unixpickle/model3d/model3d"

// NewMeshModel3d converts a model3d mesh to a *Mesh.
func NewMeshModel3d(m *model3d.Mesh) *Mesh {
	var vertices []float64
	var faces []int
	m.Iterate(func(t *model3d.Triangle) {
		for _, v := range t {
			idx := len(vertices) / 3
			faces = append(faces, idx)
			arr := v.Array()
			vertices = append(vertices, arr[:]...)
		}
	})
	return NewMesh(vertices, faces)
}

// Model3d converts the mesh to a model3d mesh.
func (m *Mesh) Model3d() *model3d.Mesh {
	vertices := m.Vertices()
	faces := m.Faces()
	res := model3d.NewMesh()
	for i := 0; i < len(faces); i += 3 {
		t := &model3d.Triangle{}
		for j := 0; j < 3; j++ {
			idx := faces[i+j]
			t[j] = model3d.XYZ(vertices[idx*3], vertices[idx*3+1], vertices[idx*3+2])
		}
		res.Add(t)
	}
	return res
}
