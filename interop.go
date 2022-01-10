package goigl

import "github.com/unixpickle/model3d/model3d"

// NewMeshModel3d converts a model3d mesh to a *Mesh.
func NewMeshModel3d(m *model3d.Mesh) *Mesh {
	var vertices []Vertex
	var faces []Face
	m.Iterate(func(t *model3d.Triangle) {
		var face Face
		for i, v := range t {
			face[i] = len(vertices)
			vertices = append(vertices, v.Array())
		}
		faces = append(faces, face)
	})
	return NewMesh(vertices, faces).RemoveDuplicateVertices(0)
}

// Model3d converts the mesh to a model3d mesh.
func (m *Mesh) Model3d() *model3d.Mesh {
	res := model3d.NewMesh()
	vertices := m.Vertices()
	for _, f := range m.Faces() {
		t := &model3d.Triangle{}
		for i, vIdx := range f {
			t[i] = model3d.NewCoord3DArray(vertices[vIdx])
		}
		res.Add(t)
	}
	return res
}
