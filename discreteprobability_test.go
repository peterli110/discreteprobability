package discreteprobability

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

const (
	repeats = 100000
	sliceLen      = 10
)

func TestDistribution(t *testing.T) {
	g := generateInt(t, time.Now().Unix(), sliceLen)
	occurrence := map[int]float64{}

	for i := 0; i < repeats; i++ {
		r := g.random()
		if _, ok := occurrence[r.Interface().(int)]; !ok {
			occurrence[r.Interface().(int)] = 0
		}
		occurrence[r.Interface().(int)]++
	}
	last := float64(0)
	for index, value := range g.values {
		val := value.Interface().(int)
		// target value
		v := occurrence[val]
		// correct probability
		p := (g.weights[index] - last) * repeats
		// deviation as 3%
		d := p * 3 / 100
		if v > p + d || v < p - d {
			t.Errorf("incorrect distribution value %v, extected %f, got %f", val, p, v)
			t.FailNow()
		}
		last = g.weights[index]
	}
}

func TestGeneratorSeeding(t *testing.T) {
	seed := int64(0)
	firstRun := resultInt(t, seed, repeats)
	secondRun := resultInt(t, seed, repeats)

	for i := 0; i < repeats; i++ {
		if firstRun[i] != secondRun[i] {
			t.Errorf("position %v got different result %v and %v", i, firstRun[i], secondRun[i])
			t.FailNow()
		}
	}
}

func TestRandomIntSafe(t *testing.T) {
	g := generateInt(t, time.Now().Unix(), sliceLen)
	if _, err := g.RandomIntSafe(); err != nil {
		t.Errorf("RandomIntSafe error %v", err)
		t.FailNow()
	}
}

func TestRandomFloat64Safe(t *testing.T) {
	g := generateFloat64(t, time.Now().Unix(), sliceLen)
	if _, err := g.RandomFloat64Safe(); err != nil {
		t.Errorf("RandomFloat64Safe error %v", err)
		t.FailNow()
	}
}

func TestRandomStringSafe(t *testing.T) {
	g := generateString(t, time.Now().Unix(), sliceLen)
	if _, err := g.RandomStringSafe(); err != nil {
		t.Errorf("RandomStringSafe error %v", err)
		t.FailNow()
	}
}

func resultInt(t *testing.T, seed int64, size int) []int {
	g := generateInt(t, seed, sliceLen)
	v := make([]int, 0, size)
	for i := 0; i < size; i++ {
		s := g.RandomInt()
		v = append(v, s)
	}
	return v
}

func generateInt(t *testing.T, seed int64, size int) *Generator {
	values := make([]int, 0, size)
	weight := make([]float64, 0, size)

	p := float64(1) / float64(size)
	for i := 0; i < size; i++ {
		values = append(values, i)
		weight = append(weight, p)
	}
	g, err := New(values, weight)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	g.SetSeed(seed)
	return g
}

func generateFloat64(t *testing.T, seed int64, size int) *Generator {
	values := make([]float64, 0, size)
	weight := make([]float64, 0, size)

	p := float64(1) / float64(size)
	for i := 0; i < size; i++ {
		values = append(values, float64(i))
		weight = append(weight, p)
	}
	g, err := New(values, weight)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	g.SetSeed(seed)
	return g
}

func generateString(t *testing.T, seed int64, size int) *Generator {
	values := make([]string, 0, size)
	weight := make([]float64, 0, size)

	p := float64(1) / float64(size)
	for i := 0; i < size; i++ {
		values = append(values, strconv.Itoa(i))
		weight = append(weight, p)
	}
	g, err := New(values, weight)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	g.SetSeed(seed)
	return g
}

func BenchmarkInt(b *testing.B) {
	for size := 4; size <= 32; size = size * 2 {
		name := fmt.Sprintf("RandomInt_size_%d", size)
		b.Run(name, func(b *testing.B) {
			g := generateInt(nil, 1, size)
			b.ResetTimer()

			for n := 0; n < b.N; n++ {
				g.RandomInt()
			}
		})
	}
}

func BenchmarkFloat64(b *testing.B) {
	for size := 4; size <= 32; size = size * 2 {
		name := fmt.Sprintf("RandomFloat64_size_%d", size)
		b.Run(name, func(b *testing.B) {
			g := generateInt(nil, 1, size)
			b.ResetTimer()

			for n := 0; n < b.N; n++ {
				g.RandomInt()
			}
		})
	}
}

func BenchmarkString(b *testing.B) {
	for size := 4; size <= 32; size = size * 2 {
		name := fmt.Sprintf("RandomString_size_%d", size)
		b.Run(name, func(b *testing.B) {
			g := generateInt(nil, 1, size)
			b.ResetTimer()

			for n := 0; n < b.N; n++ {
				g.RandomInt()
			}
		})
	}
}

func BenchmarkIntSafe(b *testing.B) {
	for size := 4; size <= 32; size = size * 2 {
		name := fmt.Sprintf("RandomIntSafe_size_%d", size)
		b.Run(name, func(b *testing.B) {
			g := generateInt(nil, 1, size)
			b.ResetTimer()

			for n := 0; n < b.N; n++ {
				_, err := g.RandomIntSafe()
				if err != nil {
					b.Errorf("RandomIntSafe() error: %v", err)
					b.FailNow()
				}
			}
		})
	}
}


func BenchmarkFloat64Safe(b *testing.B) {
	for size := 4; size <= 32; size = size * 2 {
		name := fmt.Sprintf("RandomFloat64Safe_size_%d", size)
		b.Run(name, func(b *testing.B) {
			g := generateFloat64(nil, 1, size)
			b.ResetTimer()

			for n := 0; n < b.N; n++ {
				_, err := g.RandomFloat64Safe()
				if err != nil {
					b.Errorf("RandomFloat64Safe() error: %v", err)
					b.FailNow()
				}
			}
		})
	}
}

func BenchmarkStringSafe(b *testing.B) {
	for size := 4; size <= 32; size = size * 2 {
		name := fmt.Sprintf("RandomStringSafe_size_%d", size)
		b.Run(name, func(b *testing.B) {
			g := generateString(nil, 1, size)
			b.ResetTimer()

			for n := 0; n < b.N; n++ {
				_, err := g.RandomStringSafe()
				if err != nil {
					b.Errorf("RandomStringSafe() error: %v", err)
					b.FailNow()
				}
			}
		})
	}
}