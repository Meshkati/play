# Sync-map

Benchmarking go standard `sync.map` with some alternatives, against different situations of read/write.

Run `go test -benchmem -bench .`

Run results:
```
goos: darwin
goarch: arm64
pkg: github.com/meshkati/go-benchmarks/sync-map
BenchmarkSyncMap-8               2971382               354.5 ns/op           125 B/op          4 allocs/op
BenchmarkRegularMap-8           11041466               144.9 ns/op            65 B/op          0 allocs/op
BenchmarkSyncMapRw-8             1743649               736.0 ns/op           182 B/op          4 allocs/op
BenchmarkRegularMapRW-8          2362556               555.9 ns/op            74 B/op          0 allocs/op
BenchmarkSyncMapR5W1-8           1926480               743.0 ns/op           189 B/op          4 allocs/op
BenchmarkRegularMapR5W1-8        2115121               664.9 ns/op            83 B/op          0 allocs/op
BenchmarkSyncMapR100W1-8         1401072               894.1 ns/op           146 B/op          4 allocs/op
BenchmarkRegularMapR100W1-8       176671              7645 ns/op              63 B/op          0 allocs/op
BenchmarkSyncMapR999W1-8         1000000              1103 ns/op             176 B/op          4 allocs/op
BenchmarkRegularMapR999W1-8       164307              7411 ns/op              66 B/op          0 allocs/op
PASS
ok      github.com/meshkati/go-benchmarks/sync-map      17.186s
```