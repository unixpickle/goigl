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
	for _, v := range mesh.Vertices() {
		for _, x := range v {
			if x == 0 {
				numZeros++
			} else if x == 1 {
				numOnes++
			} else {
				t.Fatalf("unexpected vertex: %f", x)
			}
		}
	}
	if numZeros != 12*9/2 || numOnes != 12*9/2 {
		t.Errorf("unexpected number of ones (%d) or zeros (%d)", numOnes, numZeros)
	}
	for i, f := range mesh.Faces() {
		for j, x := range f {
			realIdx := i*3 + j
			if x != realIdx {
				t.Errorf("face %d is %d but should be %d", realIdx, x, realIdx)
			}
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
	for i, face := range f {
		face1 := f1[i]
		for j, faceIdx := range face {
			faceIdx1 := face1[j]
			vertex := v[faceIdx]
			vertex1 := v1[faceIdx1]
			if vertex != vertex1 {
				t.Errorf("face %d[%d]: vertex should be %v but got %v", i, j, vertex, vertex1)
				break
			}
		}
	}
}
