package ci

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func delCache() {
	cache.m.Lock()
	delete(cache.val, cacheKey)
	cache.m.Unlock()
}

func TestIsCI(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		t.Setenv("CI", "true")
		assert.True(t, IsCI())
		t.Cleanup(delCache)
	})

	t.Run("false", func(t *testing.T) {
		t.Setenv("CI", "false")
		assert.False(t, IsCI())
		t.Cleanup(delCache)
	})
}

func TestIsNotCI(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		t.Setenv("CI", "false")
		assert.True(t, IsNotCI())
		t.Cleanup(delCache)
	})

	t.Run("false", func(t *testing.T) {
		t.Setenv("CI", "true")
		assert.False(t, IsNotCI())
		t.Cleanup(delCache)
	})
}

func BenchmarkIsCI(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		b.ReportAllocs()
		b.Setenv("CI", "true")

		for i := 0; i < b.N; i++ {
			_ = IsCI()
		}

		b.Cleanup(delCache)
	})

	b.Run("false", func(b *testing.B) {
		b.ReportAllocs()
		b.Setenv("CI", "false")

		for i := 0; i < b.N; i++ {
			_ = IsCI()
		}

		b.Cleanup(delCache)
	})
}

func ExampleIsCI() {
	os.Setenv("CI", "false")
	fmt.Println(IsCI())
	// Output: false
}
