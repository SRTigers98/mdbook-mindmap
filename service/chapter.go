package service

import "strings"

const (
	MINDMAP_INDICATOR = "!mindmap"
)

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
	content, ok := chapter["content"]
	if !ok {
		return false
	}

	return strings.Contains(content.(string), MINDMAP_INDICATOR)
}

func (p *DefaultChapterProcessor) ProcessChapter(chapter map[string]any) map[string]any {
	return chapter
}
