// Written in 2019 by Peter Li

// Package discreteprobability is to generate random values with a corresponding weights.
// Example usage:
//
//		values := []int{1, 2, 3}
//		weights := []float64{0.25, 0.5, 0.25}
//		generator, err := discreteprobability.New(values, weights)
//		if err != nil {
//			panic(err) // Error handlers
//		}
//		num := generator.RandomInt()
//
//		The num would have a 50% probability value of 2, a 25% probability value of 1 or 3.
package discreteprobability

import (
	"errors"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

// ErrType is returned when type assertion is failed
var ErrType				= errors.New("failed to assert type")
// ErrNotSlice is returned when the type of value is not a slice
var ErrNotSlice			= errors.New("value is not a slice")
// ErrLength is returned when the length of values and weights are different
var ErrLength			= errors.New("length of values and weights not match")
// ErrWeightSum is returned when the sum of weights is not 1
var ErrWeightSum		= errors.New("")

var seed = time.Now().UnixNano()

// Generator is the struct to store the sorted values and weights
// and can generate random values which based on the corresponding weight
type Generator struct {
	values 			[]reflect.Value
	weights 		[]float64
	size 			int
	source			rand.Source
}

func (g *Generator) Len() int { return len(g.values) }
func (g *Generator) Swap(i, j int) {
	g.values[i], g.values[j] = g.values[j], g.values[i]
	g.weights[i], g.weights[j] = g.weights[j], g.weights[i]
}
func (g *Generator) Less(i, j int) bool { return g.weights[i] < g.weights[j] }


// New returns a new Generator. It will return error if values and weights have different length
// or the sum of weights not equal to 1
func New(v interface{}, w []float64) (*Generator, error) {
	t := reflect.TypeOf(v).Kind()
	if t != reflect.Slice {
		return nil, ErrNotSlice
	}

	val := reflect.ValueOf(v)
	values := make([]reflect.Value, val.Len())
	for i := 0; i < val.Len(); i++ {
		values[i] = val.Index(i)
	}

	if len(values) != len(w) {
		return nil, ErrLength
	}
	s := &Generator{
		values: 		values,
		weights: 		w,
		size:			len(values),
		source:			rand.NewSource(seed),
	}

	sort.Sort(s)
	sum := float64(0)

	for i, weight := range s.weights {
		sum += weight
		s.weights[i] = sum
	}
	if sum - 1 > 1e-4 {
		return nil, ErrWeightSum
	}

	return s, nil
}


// SetSeed is to set a custom random seed other than the time stamp.
func (g *Generator) SetSeed(s int64) {
	g.source = rand.NewSource(s)
}

func (g *Generator) random() reflect.Value {
	f := float64(g.source.Int63()) / (1 << 63)
	i := sort.Search(g.size, func(i int) bool {
		return g.weights[i] >= f
	})

	return g.values[i]
}

// RandomInt returns the int value from the value set with corresponding weights without type assertion.
// Will panic if input value is not ([]int, []float64)
func (g *Generator) RandomInt() int {
	return int(g.random().Int())
}

// RandomFloat64 returns the float64 value from the value set with corresponding weights without type assertion.
// Will panic if input value is not ([]float64, []float64)
func (g *Generator) RandomFloat64() float64 {
	return g.random().Float()
}

// RandomString returns the string value from the value set with corresponding weights without type assertion.
// The input value should be ([]string, []float64)
func (g *Generator) RandomString() string {
	return g.random().String()
}

// RandomIntSafe returns the int value from the value set with corresponding weights.
func (g *Generator) RandomIntSafe() (int, error) {
	r, ok := g.random().Interface().(int)
	if !ok {
		return r, ErrType
	}
	return r, nil
}


// RandomStringSafe returns the int value from the value set with corresponding weights.
func (g *Generator) RandomStringSafe() (string, error) {
	r, ok := g.random().Interface().(string)
	if !ok {
		return r, ErrType
	}
	return r, nil
}

// RandomFloat64Safe returns the int value from the value set with corresponding weights.
func (g *Generator) RandomFloat64Safe() (float64, error) {
	r, ok := g.random().Interface().(float64)
	if !ok {
		return r, ErrType
	}
	return r, nil
}



