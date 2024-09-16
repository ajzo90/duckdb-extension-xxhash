# Example duckdb-plugin-templates for creating plugin for duckdb in Go.

## Example
```
 GEN=ninja make
 
 ./build/release/duckdb -unsigned 
 
```

```
D load './build/release/extension/xxhash/xxhash.duckdb_extension';

D select xxhash64('hello world!');

D select xxhash3('hello world!');

```
