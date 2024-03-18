package internal

func ValidateSort(s string) bool {
	sort_fields := []string{"name", "release", "rating"}
	flag := false
	for _, sort_field := range sort_fields {
		if s == sort_field {
			flag = true
		}
	}
	return flag
}
