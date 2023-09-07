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
	return section
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
