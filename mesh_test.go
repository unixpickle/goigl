package goigl

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

func TestMeshVerticesFaces(t *testing.T) {
	data, err := ioutil.ReadFile("test_data/cube.stl")
	if err != nil {
		t.Fatal(err)
	}
	mesh, err := MeshDecodeSTL(data)
	if err != nil {
		t.Fatal(err)
	}
	numZeros := 0
	numOnes := 0
	for _, x := range mesh.Vertices() {
		if x == 0 {
			numZeros++
		} else if x == 1 {
			numOnes++
		} else {
			t.Fatalf("unexpected vertex: %f", x)
		}
	}
	if numZeros != 12*9/2 || numOnes != 12*9/2 {
		t.Errorf("unexpected number of ones (%d) or zeros (%d)", numOnes, numZeros)
	}
	for i, x := range mesh.Faces() {
		if x != i {
			t.Errorf("face %d is %d but should be %d", i, x, i)
		}
	}
}
