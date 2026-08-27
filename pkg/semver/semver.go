// Package semver 提供简易语义化版本比较能力。
package semver

import (
	"strconv"
	"strings"
)

// Compare 比较两个形如 x.y.z 的版本号，返回 -1（a<b）、0（相等）、1（a>b）。
func Compare(a, b string) int {
	pa := parse(a)
	pb := parse(b)
	for i := 0; i < 3; i++ {
		if pa[i] != pb[i] {
			if pa[i] < pb[i] {
				return -1
			}
			return 1
		}
	}
	return 0
}

// Valid 判断字符串是否为合法的 x.y.z 版本号。
func Valid(v string) bool {
	parts := strings.Split(v, ".")
	if len(parts) != 3 {
		return false
	}
	for _, p := range parts {
		if _, err := strconv.Atoi(strings.TrimSpace(p)); err != nil {
			return false
		}
	}
	return true
}

func parse(v string) [3]int {
	var out [3]int
	parts := strings.SplitN(v, ".", 3)
	for i, p := range parts {
		if idx := strings.IndexAny(p, "-+"); idx >= 0 {
			p = p[:idx]
		}
		n, _ := strconv.Atoi(strings.TrimSpace(p))
		out[i] = n
	}
	return out
}

// Greater 判断 a 是否大于 b。
func Greater(a, b string) bool { return Compare(a, b) > 0 }

// Less 判断 a 是否小于 b。
func Less(a, b string) bool { return Compare(a, b) < 0 }

// Equal 判断 a 是否等于 b。
func Equal(a, b string) bool { return Compare(a, b) == 0 }

// Max 返回两个版本中的较大者。
func Max(a, b string) string {
	if Compare(a, b) >= 0 {
		return a
	}
	return b
}

// Min 返回两个版本中的较小者。
func Min(a, b string) string {
	if Compare(a, b) <= 0 {
		return a
	}
	return b
}
