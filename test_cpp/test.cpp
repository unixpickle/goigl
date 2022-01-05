// Example copied from:
// https://github.com/libigl/libigl/blob/v2.3.0/tutorial/101_FileIO/main.cpp

#include <igl/readOFF.h>
#include <igl/writeOBJ.h>
#include <iostream>

Eigen::MatrixXd V;
Eigen::MatrixXi F;

int main(int argc, char *argv[])
{
    // Load a mesh in OFF format
    igl::readOFF("cube.off", V, F);

    // Print the vertices and faces matrices
    std::cout << "Vertices: " << std::endl
              << V << std::endl;
    std::cout << "Faces:    " << std::endl
              << F << std::endl;

    // Save the mesh in OBJ format
    igl::writeOBJ("cube.obj", V, F);
}