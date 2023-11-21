# Contenders:

- github.com/exaring/hoglet
- github.com/afex/hystrix-go
- github.com/sony/gobreaker
- github.com/rubyist/circuitbreaker

# Current results

```
goos: linux
goarch: amd64
pkg: benchbarm_hoglet_hystrix
cpu: 12th Gen Intel(R) Core(TM) i7-1250U
BenchmarkHoglet
BenchmarkHoglet-12       	35349103	       345.5 ns/op	     208 B/op	       5 allocs/op
BenchmarkHystrix
BenchmarkHystrix-12      	13165989	      1108 ns/op	    1217 B/op	      23 allocs/op
BenchmarkGoBreaker
BenchmarkGoBreaker-12    	34526943	       297.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCircuit
BenchmarkCircuit-12      	25142444	       664.2 ns/op	     338 B/op	       6 allocs/op
```