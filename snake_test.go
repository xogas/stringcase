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

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Empty", "", ""},
		{"OnlySpaces", "  ", ""},
		{"OnlyUnderscores", "__", "__"},
		{"SingleLower", "test", "test"},
		{"SingleUpper", "TEST", "test"},
		{"CamelCase", "camelCase", "camel_case"},
		{"PascalCase", "PascalCase", "pascal_case"},
		{"WithSpaces", "Test Case", "test_case"},
		{"WithMultiSpaces", "Test  Case", "test__case"},
		{"ExistingSnake", "test_case", "test_case"},
		{"Acronym", "JSONData", "json_data"},
		{"AcronymLower", "dataJSON", "data_json"},
		{"WithDigits", "v2Beta3", "v_2_beta_3"},
		{"LeadingUnderscore", "_privateVar", "_private_var"},
		{"TrailingUnderscore", "var_", "var_"},
		{"LeadingHyphen", "-optionFlag", "_option_flag"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToSnakeCase(tc.input); got != tc.want {
				t.Errorf("ToSnakeCase(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func BenchmarkToSnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToSnakeCase("BenchmarkToSnakeCase")
	}
}

func TestToSnakeCaseWithIgnore(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		ignore string
		want   string
	}{
		{"IgnoreUnderscore", "test_case", "_", "test_case"},
		{"IgnoreSpace", "test case", " ", "test case"},
		{"IgnoreHyphen", "test-case", "-", "test-case"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToSnakeCaseWithIgnore(tc.input, tc.ignore); got != tc.want {
				t.Errorf("ToSnakeCaseWithIgnore(%q, %q) = %q, want %q", tc.input, tc.ignore, got, tc.want)
			}
		})
	}
}

func BenchmarkToSnakeCaseWithIgnore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToSnakeCaseWithIgnore("BenchmarkToSnakeCaseWithIgnore", "_")
	}
}

func TestToScreamingSnakeCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Empty", "", ""},
		{"SingleLower", "test", "TEST"},
		{"CamelCase", "camelCase", "CAMEL_CASE"},
		{"PascalCase", "PascalCase", "PASCAL_CASE"},
		{"WithSpaces", "Test Case", "TEST_CASE"},
		{"ExistingSnake", "test_case", "TEST_CASE"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToScreamingSnakeCase(tc.input); got != tc.want {
				t.Errorf("ToScreamingSnakeCase(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func BenchmarkToScreamingSnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToScreamingSnakeCase("BenchmarkToScreamingSnakeCase")
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Empty", "", ""},
		{"SingleLower", "test", "test"},
		{"CamelCase", "camelCase", "camel-case"},
		{"PascalCase", "PascalCase", "pascal-case"},
		{"WithSpaces", "Test Case", "test-case"},
		{"ExistingKebab", "test-case", "test-case"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToKebabCase(tc.input); got != tc.want {
				t.Errorf("ToKebabCase(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func BenchmarkToKebabCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToKebabCase("BenchmarkToKebabCase")
	}
}

func TestToScreamingKebabCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Empty", "", ""},
		{"SingleLower", "test", "TEST"},
		{"CamelCase", "camelCase", "CAMEL-CASE"},
		{"PascalCase", "PascalCase", "PASCAL-CASE"},
		{"WithSpaces", "Test Case", "TEST-CASE"},
		{"ExistingKebab", "test-case", "TEST-CASE"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToScreamingKebabCase(tc.input); got != tc.want {
				t.Errorf("ToScreamingKebabCase(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func BenchmarkToScreamingKebabCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToScreamingKebabCase("BenchmarkToScreamingKebabCase")
	}
}

func TestToDelimitedCase(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		delimiter byte
		want      string
	}{
		{"Empty", "", '_', ""},
		{"SingleLower", "test", '_', "test"},
		{"CamelCase", "camelCase", '_', "camel_case"},
		{"PascalCase", "PascalCase", '_', "pascal_case"},
		{"WithSpaces", "Test Case", '_', "test_case"},
		{"ExistingSnake", "test_case", '_', "test_case"},
		{"KebabDelimiter", "test-case", '-', "test-case"},
		{"DotDelimiter", "test.case", '.', "test.case"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToDelimitedCase(tc.input, tc.delimiter); got != tc.want {
				t.Errorf("ToDelimitedCase(%q, %q) = %q, want %q", tc.input, string(tc.delimiter), got, tc.want)
			}
		})
	}
}

func BenchmarkToDelimitedCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToDelimitedCase("BenchmarkToDelimitedCase", '_')
	}
}

func TestToScreamingDelimitedCase(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		delimiter byte
		want      string
	}{
		{"Empty", "", '_', ""},
		{"SingleLower", "test", '_', "TEST"},
		{"CamelCase", "camelCase", '_', "CAMEL_CASE"},
		{"PascalCase", "PascalCase", '_', "PASCAL_CASE"},
		{"WithSpaces", "Test Case", '_', "TEST_CASE"},
		{"ExistingSnake", "test_case", '_', "TEST_CASE"},
		{"KebabDelimiter", "test-case", '-', "TEST-CASE"},
		{"DotDelimiter", "test.case", '.', "TEST.CASE"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToScreamingDelimitedCase(tc.input, tc.delimiter, "", true); got != tc.want {
				t.Errorf("ToScreamingDelimitedCase(%q, %q) = %q, want %q", tc.input, string(tc.delimiter), got, tc.want)
			}
		})
	}
}

func BenchmarkToScreamingDelimitedCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToScreamingDelimitedCase("BenchmarkToScreamingDelimitedCase", '_', "", true)
	}
}
