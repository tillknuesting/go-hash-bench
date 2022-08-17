# go-hash-bench

## Run Benchmark

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
BenchmarkSum64StringXxhash/10MB-12                   1066           1101171 ns/op        9081.24 MB/s           0 B/op          0 allocs/op
```

## Benchmark Result when changing from Murmur3 to Xxhash

```bash
                                   Murmur3       Xxhash
benchmark                          old ns/op     new ns/op     delta
BenchmarkSum64ToString/4B-12       42.6          6.51          -84.70%
BenchmarkSum64ToString/16B-12      48.2          8.56          -82.23%
BenchmarkSum64ToString/100B-12     95.9          23.0          -76.07%
BenchmarkSum64ToString/4KB-12      2177          406           -81.36%
BenchmarkSum64ToString/10MB-12     5291419       1095498       -79.30%

benchmark                          old MB/s     new MB/s     speedup
BenchmarkSum64ToString/4B-12       93.99        614.12       6.53x
BenchmarkSum64ToString/16B-12      332.21       1869.31      5.63x
BenchmarkSum64ToString/100B-12     1042.45      4356.07      4.18x
BenchmarkSum64ToString/4KB-12      1837.21      9858.02      5.37x
BenchmarkSum64ToString/10MB-12     1889.85      9128.27      4.83x

benchmark                          old allocs     new allocs     delta
BenchmarkSum64ToString/4B-12       1              0              -100.00%
BenchmarkSum64ToString/16B-12      1              0              -100.00%
BenchmarkSum64ToString/100B-12     1              0              -100.00%
BenchmarkSum64ToString/4KB-12      1              0              -100.00%
BenchmarkSum64ToString/10MB-12     1              0              -100.00%

benchmark                          old bytes     new bytes     delta
BenchmarkSum64ToString/4B-12       8             0             -100.00%
BenchmarkSum64ToString/16B-12      16            0             -100.00%
BenchmarkSum64ToString/100B-12     112           0             -100.00%
BenchmarkSum64ToString/4KB-12      4096          0             -100.00%
BenchmarkSum64ToString/10MB-12     10002433      0             -100.00%
```


To compare using benchcmp, run:

```bash
go test -run=xxx -bench=BenchmarkSum64String -benchmem > a.txt
```

change the benchmark to the function to compare to and run

```bash

go test -run=xxx -bench=BenchmarkSum64String -benchmem > b.txt
```

then run:

```bash
benchcmp a.txt b.txt
```

or

```bash
benchstat a.txt b.txt
```

```bash
github.com/cespare/xxhash/v2 vs github.com/zeebo/xxh3
benchmark                        old ns/op     new ns/op     delta
BenchmarkSum64String/4B-12       6.60          5.46          -17.38%
BenchmarkSum64String/16B-12      8.54          5.42          -36.55%
BenchmarkSum64String/100B-12     22.9          17.4          -24.05%
BenchmarkSum64String/4KB-12      405           138           -65.95%
BenchmarkSum64String/10MB-12     1130583       428535        -62.10%

benchmark                        old MB/s     new MB/s     speedup
BenchmarkSum64String/4B-12       605.66       733.13       1.21x
BenchmarkSum64String/16B-12      1874.33      2954.30      1.58x
BenchmarkSum64String/100B-12     4357.17      5737.16      1.32x
BenchmarkSum64String/4KB-12      9869.16      28994.26     2.94x
BenchmarkSum64String/10MB-12     8845.00      23335.32     2.64x

benchmark                        old allocs     new allocs     delta
BenchmarkSum64String/4B-12       0              0              +0.00%
BenchmarkSum64String/16B-12      0              0              +0.00%
BenchmarkSum64String/100B-12     0              0              +0.00%
BenchmarkSum64String/4KB-12      0              0              +0.00%
BenchmarkSum64String/10MB-12     0              0              +0.00%

benchmark                        old bytes     new bytes     delta
BenchmarkSum64String/4B-12       0             0             +0.00%
BenchmarkSum64String/16B-12      0             0             +0.00%
BenchmarkSum64String/100B-12     0             0             +0.00%
BenchmarkSum64String/4KB-12      0             0             +0.00%
BenchmarkSum64String/10MB-12     0             0             +0.00%
```

    


