/*
 * MIT License
 *
 * Copyright (c) 2025 xogas <57179186+xogas@users.noreply.github.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package stringcase

import "testing"

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple", "test case", "testCase"},
		{"camel", "testCase", "testCase"},
		{"pascal", "TestCase", "testCase"},
		{"kebab", "test-case", "testCase"},
		{"snake", "test_case", "testCase"},
		{"mixed", "TEST_CASE", "testCase"},
		{"digits", "v_2_bate_3", "v2Bate3"},
		{"dot", "test.case", "testCase"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToCamelCase(tt.input)
			if got != tt.want {
				t.Errorf("ToCamelCase(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkToCamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToCamelCase("Benchmark to camel case.")
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple", "test case", "TestCase"},
		{"camel", "testCase", "TestCase"},
		{"pascal", "TestCase", "TestCase"},
		{"kebab", "test-case", "TestCase"},
		{"snake", "test_case", "TestCase"},
		{"mixed", "TEST_CASE", "TestCase"},
		{"digits", "v_2_bate_3", "V2Bate3"},
		{"dot", "test.case", "TestCase"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToPascalCase(tt.input)
			if got != tt.want {
				t.Errorf("ToPascalCase(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkToPascalCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToPascalCase("Benchmark to pascal case.")
	}
}
