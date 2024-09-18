#include "duckdb_extension.h"
#include "xxhash.h"

DUCKDB_EXTENSION_EXTERN

DUCKDB_EXTENSION_ENTRYPOINT(duckdb_connection connection, duckdb_extension_info info, duckdb_extension_access *access) {
	Register(connection, info);
}
