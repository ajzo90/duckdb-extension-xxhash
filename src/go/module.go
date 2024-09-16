package main

/*
#include <stdlib.h>
#include <duckdb.h>
*/
import "C"
import (
	"github.com/cespare/xxhash"
	"github.com/marcboeker/go-duckdb"
	"unsafe"
)

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}

//export Register
func Register(connection C.duckdb_connection, info C.duckdb_extension_info) {

	conn := duckdb.UpgradeConn((duckdb.Connection)(unsafe.Pointer(connection)))

	Check(duckdb.RegisterScalarUDFConn(conn, "xxhash64", hasher(xxhash.Sum64)))
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

type hasher func([]byte) uint64

func (xxh hasher) Config() duckdb.ScalarFunctionConfig {
	return duckdb.ScalarFunctionConfig{
		InputTypes: []string{duckdb.VARCHAR},
		ResultType: duckdb.UBIGINT,
	}
}
func (xxh hasher) Exec(ctx *duckdb.ExecContext) error {
	var chunkSize = ctx.ChunkSize()

	var a duckdb.Vec[duckdb.Varchar]
	_ = a.LoadCtx(ctx, 0, chunkSize)

	var out = duckdb.UDFScalarVectorResult[uint64](ctx)[:chunkSize]
	var in = a.Data[:chunkSize]

	for i, v := range in {
		out[i] = xxh(v.Bytes())
	}
	return nil
}
