package service

type Preprocessor interface {
	ProcessBook(context map[string]any, book map[string]any) map[string]any
}

type DefaultPreprocessor struct {
}

func NewPreprocessor() Preprocessor {
	return &DefaultPreprocessor{}
}

func (p *DefaultPreprocessor) ProcessBook(context map[string]any, book map[string]any) map[string]any {
	sections := book["sections"].([]any)

	var processedSections []any
	for _, section := range sections {
		processedSection := sectionProcessor.ProcessSection(section)

		processedSections = append(processedSections, processedSection)
	}

	processedBook := make(map[string]any)
	for k, v := range book {
		processedBook[k] = v
	}
	processedBook["sections"] = processedSections

	return processedBook
}
