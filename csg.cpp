#include "csg.h"

#include <igl/copyleft/cgal/mesh_boolean.h>

struct mesh_t {
  Eigen::MatrixXd V;
  Eigen::MatrixXi F;
};

mesh_t *mesh_boolean(mesh_t *m1, mesh_t *m2, int boolean_type) {
  igl::MeshBooleanType bt = static_cast<igl::MeshBooleanType>(boolean_type);
  mesh_t *result = new mesh_t;
  igl::copyleft::cgal::mesh_boolean(m1->V, m1->F, m2->V, m2->F, bt, result->V,
                                    result->F);
  return result;
}
