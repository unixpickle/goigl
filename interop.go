package goigl

import "github.com/unixpickle/model3d/model3d"

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
