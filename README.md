# go-hash-bench

## Benchmark

```bash 
go test -run=xxx -bench=. -benchmem
```


## Benchmark Result

```bash
go version go1.19 darwin/amd64
goarch: amd64
pkg: go-hash-bench
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSum64StringMurmur3/4B-12               25956595                42.29 ns/op       94.59 MB/s           8 B/op          1 allocs/op
BenchmarkSum64StringMurmur3/16B-12              24882978                45.95 ns/op      348.21 MB/s          16 B/op          1 allocs/op
BenchmarkSum64StringMurmur3/100B-12             13237304                86.98 ns/op     1149.63 MB/s         112 B/op          1 allocs/op
BenchmarkSum64StringMurmur3/4KB-12                692950              1630 ns/op        2453.49 MB/s        4096 B/op          1 allocs/op
BenchmarkSum64StringMurmur3/10MB-12                  273           4012493 ns/op        2492.22 MB/s    10002432 B/op          1 allocs/op
BenchmarkSum64StringXxhash/4B-12                184129242                6.487 ns/op     616.64 MB/s           0 B/op          0 allocs/op
BenchmarkSum64StringXxhash/16B-12               140242632                8.534 ns/op    1874.89 MB/s           0 B/op          0 allocs/op
BenchmarkSum64StringXxhash/100B-12              52295360                22.98 ns/op     4352.22 MB/s           0 B/op          0 allocs/op
BenchmarkSum64StringXxhash/4KB-12                2945018               405.1 ns/op      9873.08 MB/s           0 B/op          0 allocs/op
BenchmarkSum64StringXxhash/10MB-12                  1066           1101171 ns/op        9081.24 MB/s           0 B/op          0 allocs/op

```
