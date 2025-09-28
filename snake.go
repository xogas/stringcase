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

import (
	"strings"
	"unicode"
)

// ToSnakeCase converts the input string to snake_case.
func ToSnakeCase(s string) string {
	return ToDelimitedCase(s, '_')
}

// ToSnakeCaseWithIgnore converts the input string to snake_case, ignoring specified characters.
func ToSnakeCaseWithIgnore(s string, ignore string) string {
	return ToScreamingDelimitedCase(s, '_', ignore, false)
}

// ToScreamingSnakeCase converts the input string to SCREAMING_SNAKE_CASE.
func ToScreamingSnakeCase(s string) string {
	return ToScreamingDelimitedCase(s, '_', "", true)
}

// ToKebabCase converts the input string to kebab-case.
func ToKebabCase(s string) string {
	return ToDelimitedCase(s, '-')
}

// ToScreamingKebabCase converts the input string to SCREAMING-KEBAB-CASE.
func ToScreamingKebabCase(s string) string {
	return ToScreamingDelimitedCase(s, '-', "", true)
}

// ToDelimitedCase converts the input string to a delimited format using the specified delimiter.
func ToDelimitedCase(s string, delimiter byte) string {
	return ToScreamingDelimitedCase(s, delimiter, "", false)
}

// ToScreamingDelimitedCase converts a string to delimited format with ignore options and screaming case.
func ToScreamingDelimitedCase(s string, delimiter byte, ignore string, screaming bool) string {
	s = strings.TrimSpace(s)
	result := strings.Builder{}
	result.Grow(len(s) + len(s)/4) // Preallocate memory to reduce allocations

	runes := []rune(s)
	for k, r := range runes {
		rIsUpper := unicode.IsUpper(r)
		rIsLower := unicode.IsLower(r)
		rIsDigit := unicode.IsDigit(r)
		if rIsUpper && !screaming {
			r = unicode.ToLower(r)
		} else if rIsLower && screaming {
			r = unicode.ToUpper(r)
		}

		if k+1 < len(runes) {
			nextR := runes[k+1]
			nextIsUpper := unicode.IsUpper(nextR)
			nextIsLower := unicode.IsLower(nextR)
			nextIsDigit := unicode.IsDigit(nextR)

			// add underscore if next letter case is different
			shouldAddDelimiter := (rIsUpper && (nextIsLower || nextIsDigit)) ||
				(rIsLower && (nextIsUpper || nextIsDigit)) ||
				(rIsDigit && (nextIsUpper || nextIsLower))

			if shouldAddDelimiter {
				prevRIgnore := ignore != "" && k > 0 && strings.ContainsRune(ignore, runes[k-1])
				if !prevRIgnore {
					if rIsUpper && nextIsLower && k > 0 && unicode.IsUpper(runes[k-1]) {
						result.WriteByte(delimiter)
					}
					result.WriteRune(r)
					if rIsLower || rIsDigit || nextIsDigit {
						result.WriteByte(delimiter)
					}
					continue
				}
			}
		}

		if (r == ' ' || r == '-' || r == '_' || r == '.') && !strings.ContainsRune(ignore, r) {
			// replace space/underscore/hyphen/dot with delimiter if not ignored
			result.WriteByte(delimiter)
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}
