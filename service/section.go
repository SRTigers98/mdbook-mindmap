package service

type SectionProcessor interface {
	ProcessSection(section any) any
	IsChapter(section any) bool
}

type DefaultSectionProcessor struct {
}

func NewSectionProcessor() SectionProcessor {
	return &DefaultSectionProcessor{}
}

func (p *DefaultSectionProcessor) ProcessSection(section any) any {
	if !p.IsChapter(section) {
		return section
	}

	chapter := section.(map[string]any)["Chapter"].(map[string]any)

	if !chapterProcessor.HasMindmapIndicator(chapter) {
		return section
	}

	return map[string]any{
		"Chapter": chapterProcessor.ProcessChapter(chapter),
	}
}

func (p *DefaultSectionProcessor) IsChapter(section any) bool {
	switch s := section.(type) {
	case map[string]any:
		_, ok := s["Chapter"]
		return ok
	default:
		return false
	}
}
