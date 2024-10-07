package benchmark

import (
	"fmt"
	"testing"
)

func HelloWorld(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

func BenchmarkHelloWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ivan")
	}
}

// benchmark sub
func BenchmarkSub(b *testing.B) {
	b.Run("Adi", func(b *testing.B)  {
		for i := 0; i < b.N; i++ {
			HelloWorld("Adi")
		}
	})
	b.Run("Saputra", func(b *testing.B)  {
		for i := 0; i < b.N; i++ {
			HelloWorld("Saputra")
		}
	})
}

// benchmark table
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct{
		name string
		request string
	} {
		{
			name: "Ivan",
            request: "Ivan",
		},
		{
			name: "Adi",
            request: "Adi",
		},
		{
			name: "Saputra",
            request: "Saputra",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func (b *testing.B)  {
			for i := 0; i < b.N; i++ {
                HelloWorld(benchmark.request)
            }			
		})
	}
}