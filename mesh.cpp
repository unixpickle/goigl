#include "mesh.h"

#include <igl/readSTL.h>

#include <sstream>

struct mesh_t {
  Eigen::MatrixXd V;
  Eigen::MatrixXi F;
};

static void copy_error_message(const char *msg, char **output) {
  const int error_len = strlen(msg);
  *output = (char *)malloc(error_len + 1);
  memcpy(*output, msg, error_len + 1);
}

mesh_t *mesh_decode_stl(const char *data, size_t data_len, char **error_out) {
  mesh_t *mesh = new mesh_t;
  Eigen::MatrixXd N;
  std::istringstream stream(std::string(data, (size_t)data_len));
  try {
    bool success = igl::readSTL(stream, mesh->V, mesh->F, N);
    if (!success) {
      throw 0;
    }
  } catch (const std::runtime_error &re) {
    copy_error_message(re.what(), error_out);
    delete mesh;
    return NULL;
  } catch (...) {
    copy_error_message("unknown error", error_out);
    delete mesh;
    return NULL;
  }
  return mesh;
}

double *mesh_vertices(mesh_t *mesh) {
  double *result = (double *)malloc(sizeof(double) * mesh->V.size());
  size_t idx = 0;
  for (size_t i = 0; i < mesh->V.rows(); i++) {
    for (size_t j = 0; j < mesh->V.cols(); j++) {
      result[idx++] = mesh->V(i, j);
    }
  }
  return result;
}

size_t mesh_vertices_size(mesh_t *mesh) { return mesh->V.size(); }

int mesh_num_vertices(mesh_t *mesh) { return mesh->V.rows(); }

int *mesh_faces(mesh_t *mesh) {
  int *result = (int *)malloc(sizeof(int) * mesh->F.size());
  size_t idx = 0;
  for (size_t i = 0; i < mesh->F.rows(); i++) {
    for (size_t j = 0; j < mesh->F.cols(); j++) {
      result[idx] = mesh->F(i, j);
      idx++;
    }
  }
  return result;
}

size_t mesh_faces_size(mesh_t *mesh) { return mesh->F.size(); }
int mesh_num_faces(mesh_t *mesh) { return mesh->F.rows(); }
void mesh_free(mesh_t *mesh) { delete mesh; }
