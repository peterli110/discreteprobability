[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)
[![Build Status](https://travis-ci.org/peterli110/discreteprobability.svg?branch=master)](https://travis-ci.org/peterli110/discreteprobability)
[![Documentation](https://img.shields.io/badge/Documentation-GoDoc-green.svg)](https://godoc.org/github.com/peterli110/discreteprobability)

Discrete Probability Generator in Golang
=================================
Weighted random is a golang implementation of discrete probability distribution, 
which means in a set of elements, the occurrence of each element will have a
fixed probability.

For example, if we have a slice s := []int{1, 2, 3} and weight w := []float64{0.2, 0.5, 0.3},
the result should have 20% probability of 1, 50% probability of 2 and 30% probability of 3.

Documentation and examples
========================

Here is a simple example:

```
intValues := []int{1, 2, 3}
float64Values := []float64{1.1, 2.2, 3.3}
stringValues := []string{"a", "b", "c"}
weights := []float64{0.2, 0.5, 0.3}

intRNG, err := discreteprobability.New(intValues, wegihts)
if err != nil {
    // Error handlers
}

float64RNG, err := discreteprobability.New(float64Values, wegihts)
if err != nil {
    // Error handlers
}

stringRNG, err := discreteprobability.New(stringValues, wegihts)
if err != nil {
    // Error handlers
}

// If you use these methods, please make sure the input type is correct
// This will cause panic:
// intVal := intRNG.RandomFloat64()
intVal := intRNG.RandomInt()
float64Val := float64RNG.RandomFloat64()
stringVal := stringRNG.RandomString()

// Also a type safe method is provided
// but the performance is slower because of the type assertion
intVal, err = intRNG.RandomIntSafe()
if err != nil {
    // Error handlers
}
```

Testing and benchmarking
========================

To run all tests, `cd` into the directory and use:

    go test -v
    
To run benchmarks:

	go test -bench=.
	
Here is a result of benchmarks:
```
goos: darwin
goarch: amd64
BenchmarkInt/RandomInt_size_4-6                 50000000                28.9 ns/op
BenchmarkInt/RandomInt_size_8-6                 50000000                36.7 ns/op
BenchmarkInt/RandomInt_size_16-6                30000000                45.6 ns/op
BenchmarkInt/RandomInt_size_32-6                30000000                53.7 ns/op
BenchmarkFloat64/RandomFloat64_size_4-6         50000000                29.0 ns/op
BenchmarkFloat64/RandomFloat64_size_8-6         50000000                36.9 ns/op
BenchmarkFloat64/RandomFloat64_size_16-6                30000000                45.2 ns/op
BenchmarkFloat64/RandomFloat64_size_32-6                30000000                53.8 ns/op
BenchmarkString/RandomString_size_4-6                   50000000                28.6 ns/op
BenchmarkString/RandomString_size_8-6                   50000000                36.6 ns/op
BenchmarkString/RandomString_size_16-6                  30000000                45.3 ns/op
BenchmarkString/RandomString_size_32-6                  30000000                53.8 ns/op
BenchmarkIntSafe/RandomIntSafe_size_4-6                 30000000                52.6 ns/op
BenchmarkIntSafe/RandomIntSafe_size_8-6                 20000000                63.1 ns/op
BenchmarkIntSafe/RandomIntSafe_size_16-6                20000000                73.2 ns/op
BenchmarkIntSafe/RandomIntSafe_size_32-6                20000000                83.1 ns/op
BenchmarkFloat64Safe/RandomFloat64Safe_size_4-6         30000000                52.2 ns/op
BenchmarkFloat64Safe/RandomFloat64Safe_size_8-6         20000000                62.1 ns/op
BenchmarkFloat64Safe/RandomFloat64Safe_size_16-6        20000000                72.9 ns/op
BenchmarkFloat64Safe/RandomFloat64Safe_size_32-6        20000000                83.3 ns/op
BenchmarkStringSafe/RandomStringSafe_size_4-6           20000000                67.5 ns/op
BenchmarkStringSafe/RandomStringSafe_size_8-6           20000000                76.8 ns/op
BenchmarkStringSafe/RandomStringSafe_size_16-6          20000000                86.2 ns/op
BenchmarkStringSafe/RandomStringSafe_size_32-6          20000000                94.4 ns/op
PASS
```


