package ci

import "testing"

// SkipTestIfCI is a testing helper function that skips the execution of a test
// if it is running within a CI/CD pipeline.
//
// Example:
//
//	func TestYourFunc(t *testing.T) {
//		defer func() {
//			println(t.Skipped())
//			// Output: true
//		}()
//
//		t.Setenv("CI", "true")
//		ci.SkipTestIfCI(t)
//
//		// do your test...
//	}
func SkipTestIfCI(t *testing.T) {
	if IsCI() {
		t.SkipNow()
	}
}

// SkipTestIfNotCI is a testing helper function that skips the execution of a
// test if it is not running within a CI/CD pipeline.
func SkipTestIfNotCI(t *testing.T) {
	if IsNotCI() {
		t.SkipNow()
	}
}

// SkipBenchmarkIfCI is a testing helper function that skips the execution of a
// benchmark if it is running within a CI/CD pipeline.
func SkipBenchmarkIfCI(b *testing.B) {
	if IsCI() {
		b.SkipNow()
	}
}

// SkipBenchmarkIfNotCI is a testing helper function that skips the execution of
// a benchmark if it is not running within a CI/CD pipeline.
func SkipBenchmarkIfNotCI(b *testing.B) {
	if IsNotCI() {
		b.SkipNow()
	}
}

// SkipFuzzIfCI is a testing helper function that skips the execution of a fuzz
// test if it is running within a CI/CD pipeline.
func SkipFuzzIfCI(f *testing.F) {
	if IsCI() {
		f.SkipNow()
	}
}

// SkipFuzzIfNotCI is a testing helper function that skips the execution of a
// fuzz test if it is not running within a CI/CD pipeline.
func SkipFuzzIfNotCI(f *testing.F) {
	if IsNotCI() {
		f.SkipNow()
	}
}
