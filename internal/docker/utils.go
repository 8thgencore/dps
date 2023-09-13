package docker

import (
	"fmt"
	"math"
	"strings"
)

func parseRepoTag(repoTag string) (repo, tag string) {
	if idx := lastIndexOf(repoTag, ":"); idx != -1 {
		repo = repoTag[:idx]
		tag = repoTag[idx+1:]
	} else {
		repo = repoTag
		tag = "<none>"
	}
	return repo, tag
}

func lastIndexOf(s, substr string) int {
	idx := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == substr[0] && len(s[i:]) >= len(substr) && s[i:i+len(substr)] == substr {
			idx = i
			break
		}
	}
	return idx
}

func sliteToString(sl []string) string {
	// Создаем мапу для отслеживания уникальных элементов
	uniqueSlice := make(map[string]bool)

	// Итерируем по слайсу и добавляем уникальные элементы в мапу
	for _, el := range sl {
		if el != "" {
			uniqueSlice[el] = true
		}
	}

	// Создаем слайс для хранения уникальных элементов
	uniquePortSlice := make([]string, 0, len(uniqueSlice))

	// Итерируем по мапе и добавляем уникальные элементы в слайс
	for el := range uniqueSlice {
		uniquePortSlice = append(uniquePortSlice, el)
	}

	return strings.Join(uniquePortSlice, " ")
}

func formatBytes(bytes float64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%.0fБ", bytes)
	} else if bytes < math.Pow(1024, 2) {
		return fmt.Sprintf("%.1fКб", bytes/1024)
	} else if bytes < math.Pow(1024, 3) {
		return fmt.Sprintf("%.1fМб", bytes/math.Pow(1024, 2))
	} else {
		return fmt.Sprintf("%.1fГб", bytes/math.Pow(1024, 3))
	}
}
