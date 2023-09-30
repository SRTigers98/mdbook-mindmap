package service

func ProcessBook(context map[string]any, book map[string]any) map[string]any {
	sections := book["sections"].([]any)

	var processedSections []any
	for _, section := range sections {
		processedSection := ProcessSection(section)

		processedSections = append(processedSections, processedSection)
	}

	processedBook := make(map[string]any)
	for k, v := range book {
		processedBook[k] = v
	}
	processedBook["sections"] = processedSections

	return processedBook
}
