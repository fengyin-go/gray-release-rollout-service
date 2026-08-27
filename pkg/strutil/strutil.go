// Package strutil 提供字符串处理工具。
package strutil

import (
	"strings"
	"unicode"
)

// ToSnakeCase 将驼峰或空格分隔字符串转换为下划线分隔。
func ToSnakeCase(s string) string {
	var b strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 && b.Len() > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		if r == ' ' || r == '-' || r == '.' {
			b.WriteByte('_')
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// ToCamelCase 将下划线或连字符分隔字符串转换为小驼峰。
func ToCamelCase(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})
	var b strings.Builder
	for i, p := range parts {
		if p == "" {
			continue
		}
		if i == 0 {
			b.WriteString(strings.ToLower(p))
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		b.WriteString(p[1:])
	}
	return b.String()
}

// Truncate 截断字符串到 max 长度，超出追加省略号。
func Truncate(s string, max int) string {
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	if max <= 3 {
		return string(r[:max])
	}
	return string(r[:max-3]) + "..."
}

// ContainsFold 大小写不敏感判断是否包含子串。
func ContainsFold(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// PadLeft 在左侧填充字符到指定长度。
func PadLeft(s string, pad rune, width int) string {
	r := []rune(s)
	if len(r) >= width {
		return s
	}
	var b strings.Builder
	for i := len(r); i < width; i++ {
		b.WriteRune(pad)
	}
	b.WriteString(s)
	return b.String()
}

// PadRight 在右侧填充字符到指定长度。
func PadRight(s string, pad rune, width int) string {
	r := []rune(s)
	if len(r) >= width {
		return s
	}
	var b strings.Builder
	b.WriteString(s)
	for i := len(r); i < width; i++ {
		b.WriteRune(pad)
	}
	return b.String()
}

// Reverse 反转字符串（按 rune 处理，支持中文）。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
