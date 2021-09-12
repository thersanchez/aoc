package main

func setSubtract(minuend, subtraend map[string]struct{}) map[string]struct{} {
	difference := map[string]struct{}{}

	for key, _ := range minuend {
		if _, ok := subtraend[key]; ok {
			continue
		}
		difference[key] = struct{}{}
	}

	return difference
}
