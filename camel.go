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

// ToCamelCase converts a string to camelCase.
func ToCamelCase(s string) string {
	return camelPascalCase(s, false)
}

// ToPascalCase converts a string to PascalCase.
func ToPascalCase(s string) string {
	return camelPascalCase(s, true)
}

func camelPascalCase(s string, pascal bool) string {
	s = strings.TrimSpace(s)
	v, hasAcronym := upperAcronyms.Load(s)
	if hasAcronym {
		s = v.(string)
	}

	result := strings.Builder{}
	result.Grow(len(s))
	capNext := pascal
	prevRIsUpper := false
	runes := []rune(s)
	for k, r := range runes {
		rIsUpper := unicode.IsUpper(r)
		rIsLower := unicode.IsLower(r)
		rIsDigit := unicode.IsDigit(r)
		if capNext {
			if rIsLower {
				r = unicode.ToUpper(r)
			}
		} else if k == 0 {
			if rIsUpper {
				r = unicode.ToLower(r)
			}
		} else if rIsUpper && prevRIsUpper && !hasAcronym {
			r = unicode.ToLower(r)
		}
		prevRIsUpper = rIsUpper

		if rIsUpper || rIsLower {
			result.WriteRune(r)
			capNext = false
		} else if rIsDigit {
			result.WriteRune(r)
			capNext = true
		} else {
			capNext = r == ' ' || r == '_' || r == '-' || r == '.'
		}
	}

	return result.String()
}
