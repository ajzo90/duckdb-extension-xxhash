# DuckDB Extension that adds xxhash functions to DuckDB

Calculates xxHash checksum from a string, from https://github.com/cespare/xxhash

```
:) ./build/release/duckdb -unsigned

D load './build/release/extension/quack/quack.duckdb_extension';

D select xxhash64('hello');
┌─────────────────────┐
│  xxhash64('hello')  │
│       uint64        │
├─────────────────────┤
│ 2794345569481354659 │
└─────────────────────┘

```
