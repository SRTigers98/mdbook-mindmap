package service

func ProcessSection(section any) any {
	if !IsChapter(section) {
		return section
	}

	chapter := section.(map[string]any)["Chapter"].(map[string]any)

	if !HasMindmapIndicator(chapter) {
		return section
	}

	return map[string]any{
		"Chapter": ProcessChapter(chapter),
	}
}

func IsChapter(section any) bool {
	switch s := section.(type) {
	case map[string]any:
		_, ok := s["Chapter"]
		return ok
	default:
		return false
	}
}
