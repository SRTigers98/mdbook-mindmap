package service

type ChapterProcessor interface {
	HasMindmapIndicator(chapter map[string]any) bool
	ProcessChapter(chapter map[string]any) map[string]any
}

type DefaultChapterProcessor struct {
}

func NewChapterProcessor() ChapterProcessor {
	return &DefaultChapterProcessor{}
}

func (p *DefaultChapterProcessor) HasMindmapIndicator(chapter map[string]any) bool {
	return false
}

func (p *DefaultChapterProcessor) ProcessChapter(chapter map[string]any) map[string]any {
	return chapter
}
