package igl

import (
	"io/ioutil"
	"testing"
)

func TestMeshDecodeSTL(t *testing.T) {
	data, err := ioutil.ReadFile("test_data/cube.stl")
	if err != nil {
		t.Fatal(err)
	}
	mesh, err := MeshDecodeSTL(data)
	if err != nil {
		t.Fatal(err)
	}
	if n := mesh.NumFaces(); n != 12 {
		t.Errorf("expected 12 faces but got %d", n)
	}
	if n := mesh.NumVertices(); n != 12*3 {
		t.Errorf("expected %d vertices but got %d", 12*3, n)
	}
}
