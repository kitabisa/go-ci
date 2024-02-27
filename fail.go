package ci

import "testing"

// FailTestIfCI is a testing helper function that fails the execution of a test
// if it is running within a CI/CD pipeline.
func FailTestIfCI(t *testing.T) {
	if IsCI() {
		t.FailNow()
	}
}

// FailTestIfNotCI is a testing helper function that fails the execution of a
// test if it is not running within a CI/CD pipeline.
func FailTestIfNotCI(t *testing.T) {
	if IsNotCI() {
		t.FailNow()
	}
}

// FailBenchmarkIfCI is a testing helper function that fails the execution of a
// benchmark if it is running within a CI/CD pipeline.
func FailBenchmarkIfCI(b *testing.B) {
	if IsCI() {
		b.FailNow()
	}
}

// FailBenchmarkIfNotCI is a testing helper function that fails the execution of
// a benchmark if it is not running within a CI/CD pipeline.
func FailBenchmarkIfNotCI(b *testing.B) {
	if IsNotCI() {
		b.FailNow()
	}
}

// FailFuzzIfCI is a testing helper function that fails the execution of a fuzz
// test if it is running within a CI/CD pipeline.
func FailFuzzIfCI(f *testing.F) {
	if IsCI() {
		f.FailNow()
	}
}

// FailFuzzIfNotCI is a testing helper function that fails the execution of a
// fuzz test if it is not running within a CI/CD pipeline.
func FailFuzzIfNotCI(f *testing.F) {
	if IsNotCI() {
		f.FailNow()
	}
}
