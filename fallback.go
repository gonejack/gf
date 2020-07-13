package gf

func EmptyFallback(s, def string) string {
	if s == "" {
		return def
	}
	return s
}

func ZeroFallback(i, def int) int {
	if i == 0 {
		return def
	}
	return i
}

func Max(i, max int) int {
	if i > max {
		return max
	}
	return i
}

func Min(i, min int) int {
	if i < min {
		return min
	}
	return i
}
