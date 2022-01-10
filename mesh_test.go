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

func TestMeshDecodeSTLError(t *testing.T) {
	_, err := MeshDecodeSTL(make([]byte, 10))
	if err == nil {
		t.Fatal("expected error from decoding invalid mesh")
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

func TestMeshRemoveDuplicateVertices(t *testing.T) {
	data, err := ioutil.ReadFile("test_data/cube.stl")
	if err != nil {
		t.Fatal(err)
	}
	mesh, err := MeshDecodeSTL(data)
	if err != nil {
		t.Fatal(err)
	}
	m1 := mesh.RemoveDuplicateVertices(0)
	if m1.NumVertices() != 8 {
		t.Errorf("expected %d vertices but got %d", 8, m1.NumVertices())
	}
	f := mesh.Faces()
	v := mesh.Vertices()
	f1 := m1.Faces()
	v1 := m1.Vertices()
	if len(f) != len(f1) {
		t.Fatalf("mismatching face count: %d versus %d", len(f), len(f1))
	}
	for i, faceIdx := range f {
		faceIdx1 := f1[i]
		vertex := v[faceIdx*3 : (faceIdx+1)*3]
		vertex1 := v1[faceIdx1*3 : (faceIdx1+1)*3]
		for i, x := range vertex {
			y := vertex1[i]
			if x != y {
				t.Errorf("face %d: vertex should be %v but got %v", i, vertex, vertex1)
				break
			}
		}
	}
}
