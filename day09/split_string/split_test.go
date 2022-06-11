/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/26 8:14
 */
package split_string

import (
	"fmt"
	"reflect"
	"testing"
)

//func TestSplit(t *testing.T) {
//
//	got := Split("abcbe", "b")
//	want := []string{"a", "c", "e"}
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf("execepted:%v, got:%v\n", want, got)
//	}
//}
//func Test1Split(t *testing.T) {
//	got := Split("a:b:c:b:e", ":")
//	want := []string{"a", "b", "c", "b", "e"}
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf("execepted:%v, got:%v\n", want, got)
//	}
//}

//func TestSplit(t *testing.T) {
//	type testCase struct {
//		s    string
//		sep  string
//		want []string
//	}
//	testGroup := []testCase{
//		{"abcbe", "b", []string{"a", "c", "e"}},
//		{"a:b:c:b:e", ":", []string{"a", "b", "c", "b", "e"}},
//		{"abcefbce", "bc", []string{"a", "ef", "e"}},
//	}
//	for _, tc := range testGroup {
//		got := Split(tc.s, tc.sep)
//		if !reflect.DeepEqual(got, tc.want) {
//			t.Errorf("got:%#v want:%#v\n", got, tc.want)
//		}
//	}
//}

func TestSplit(t *testing.T) {

	type testCase struct {
		s    string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case_01": {"abcbe", "b", []string{"a", "c", "e"}},
		"case_02": {"a:b:c:b:e", ":", []string{"a", "b", "c", "b", "e"}},
		"case_03": {"abcefbce", "bc", []string{"a", "ef", "e"}},
	}
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.s, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%#v got:%#v want:%#v\n", name, got, tc.want)
			}
		})
		//if !reflect.DeepEqual(got, tc.want) {
		//	t.Errorf("name:%#v got:%#v want:%#v\n", name, got, tc.want)
		//}

	}

}

// BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:b:e", ":")
	}
}

// fib_test.go

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

//func benchmark(b *testing.B, size int) { /* ... */ }
//func Benchmark10(b *testing.B)         { benchmark(b, 10) }
//func Benchmark100(b *testing.B)        { benchmark(b, 100) }
//func Benchmark1000(b *testing.B)       { benchmark(b, 1000) }

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("沙河有沙又有河", "沙"))
	// Output:
	// [a b c]
	// [ 河有 又有河]
}

func ExampleExamples_output() {
	fmt.Println("Hello")
	// Output: Hello
}
