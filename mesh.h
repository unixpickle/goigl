#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct mesh_t mesh_t;

mesh_t *mesh_decode_stl(const char *data, size_t data_len);
double *mesh_vertices(mesh_t *mesh);
size_t mesh_vertices_size(mesh_t *mesh);
int mesh_num_vertices(mesh_t *mesh);
int *mesh_faces(mesh_t *mesh);
size_t mesh_faces_size(mesh_t *mesh);
int mesh_num_faces(mesh_t *mesh);
void mesh_free(mesh_t *mesh);

#ifdef __cplusplus
}
#endif