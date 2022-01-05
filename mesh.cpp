#include "mesh.h"

#include <igl/readSTL.h>

#include <sstream>

struct mesh_t {
  Eigen::MatrixXd V;
  Eigen::MatrixXi F;
};

mesh_t *mesh_decode_stl(const char *data, size_t data_len) {
  mesh_t *mesh = new mesh_t;
  Eigen::MatrixXd N;
  std::istringstream stream(std::string(data, (size_t)data_len));
  bool success = igl::readSTL(stream, mesh->V, mesh->F, N);
  if (!success) {
    delete mesh;
    return NULL;
  }
  return mesh;
}

int mesh_num_vertices(mesh_t *mesh) { return mesh->V.rows(); }
int mesh_num_faces(mesh_t *mesh) { return mesh->F.rows(); }
void mesh_free(mesh_t *mesh) { delete mesh; }
