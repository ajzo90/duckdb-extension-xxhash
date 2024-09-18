package main

/*
#include <stdlib.h>
#include <duckdb.h>

*/
import "C"
import (
	"fmt"
	"unsafe"
)

//import (
//"github.com/cespare/xxhash"
//"github.com/marcboeker/go-duckdb"
//"unsafe"
//)

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}

//export Register
func Register(connection C.duckdb_connection, info C.duckdb_extension_info) {
	if err := registerType(connection, "go_defined_type"); err != nil {
		panic(fmt.Sprintf("failed to register extension: %s, need better error handling", err.Error()))
	}
	//conn := duckdb.UpgradeConn((duckdb.Connection)(unsafe.Pointer(connection)))
	//
	//Check(duckdb.RegisterScalarUDFConn(conn, "xxhash64", hasher(xxhash.Sum64)))
}

func registerType(connection C.duckdb_connection, name string) error {

	typeName := C.CString(name)
	defer C.free(unsafe.Pointer(typeName))
	//

	x := C.DUCKDB_TYPE_INTEGER
	fmt.Println(x)

	logicalType := C.duckdb_create_logical_type(C.DUCKDB_TYPE_INTEGER)
	fmt.Println(logicalType)
	defer C.duckdb_destroy_logical_type(&logicalType)
	//
	C.duckdb_logical_type_set_alias(logicalType, typeName)
	//
	status := C.duckdb_register_logical_type(connection, logicalType, nil)
	if status != C.DuckDBSuccess {
		return fmt.Errorf("failed to register type %s", name)
	}
	//
	return nil

}

//func Check(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
//
//type hasher func([]byte) uint64
//
//func (xxh hasher) Config() duckdb.ScalarFunctionConfig {
//	return duckdb.ScalarFunctionConfig{
//		InputTypes: []string{duckdb.VARCHAR},
//		ResultType: duckdb.UBIGINT,
//	}
//}
//func (xxh hasher) Exec(ctx *duckdb.ExecContext) error {
//	var chunkSize = ctx.ChunkSize()
//
//	var a duckdb.Vec[duckdb.Varchar]
//	_ = a.LoadCtx(ctx, 0, chunkSize)
//
//	var out = duckdb.UDFScalarVectorResult[uint64](ctx)[:chunkSize]
//	var in = a.Data[:chunkSize]
//
//	for i, v := range in {
//		out[i] = xxh(v.Bytes())
//	}
//	return nil
//}
