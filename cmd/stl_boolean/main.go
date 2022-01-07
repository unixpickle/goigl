// Command stl_boolean performs boolean operations on STL meshes.
//
// For example, to subtract mesh b.stl from mesh a.stl, do:
//
//     $ stl_boolean -op minus a.stl b.stl out.stl
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/goigl"
	"github.com/unixpickle/goigl/iglcgal"
)

func main() {
	var opName string
	flag.StringVar(&opName, "op", "union", "union, intersect, minus, xor, resolve")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: stl_boolean [flags] a.stl b.stl out.stl")
		fmt.Fprintln(os.Stderr)
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()
	if len(flag.Args()) != 3 {
		flag.Usage()
	}
	var op iglcgal.MeshBooleanType
	switch opName {
	case "union":
		op = iglcgal.Union
	case "intersect":
		op = iglcgal.Intersect
	case "minus":
		op = iglcgal.Minus
	case "xor":
		op = iglcgal.Xor
	case "resolve":
		op = iglcgal.Resolve
	default:
		fmt.Fprintln(os.Stderr, "invalid value for -op argument:", opName)
		flag.Usage()
	}

	path1, path2, pathOut := flag.Args()[0], flag.Args()[1], flag.Args()[2]

	data, err := ioutil.ReadFile(path1)
	essentials.Must(err)
	mesh1, err := goigl.MeshDecodeSTL(data)
	essentials.Must(err)
	data, err = ioutil.ReadFile(path2)
	essentials.Must(err)
	mesh2, err := goigl.MeshDecodeSTL(data)
	essentials.Must(err)

	finalMesh := iglcgal.MeshBoolean(mesh1, mesh2, op)
	essentials.Must(finalMesh.WriteSTL(pathOut))
}
