package util

func AppendPath(p1 string, p2 string) string {
	if p1 == "" {
		return p2
	}
	return p1 + "/" + p2
}

func StrOr(s1 string, s2 string) string {
	if s1 == "" {
		return s2
	}
	return s1
}
