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

	processedBook := map[string]any{
		"sections":         processedSections,
		"__non_exhaustive": book["__non_exhaustive"],
	}

	return processedBook
}
