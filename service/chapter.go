package service

import (
	"fmt"
	"strings"
)

const (
	MINDMAP_INDICATOR = "!mindmap"
	MERMAID_IMPORT    = `<script type="module">
    import mermaid from "https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs";
    mermaid.initialize({ startOnLoad: true });
</script>`
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
	processedContent := processContent(chapter)
	processedSubItems := processSubItems(chapter)

	processedChapter := make(map[string]any)
	for k, v := range chapter {
		processedChapter[k] = v
	}
	processedChapter["content"] = processedContent
	processedChapter["sub_items"] = processedSubItems

	return processedChapter
}

func processContent(chapter map[string]any) string {
	content := chapter["content"].(string)

	withImport := fmt.Sprintf("%s\n%s", content, MERMAID_IMPORT)

	mindmap := mindmapCreator.CreateMindmap(chapter)
	processedContent := strings.ReplaceAll(withImport, MINDMAP_INDICATOR, mindmap)

	return processedContent
}

func processSubItems(chapter map[string]any) []any {
	subItems := chapter["sub_items"].([]any)

	var processedItems []any
	for _, item := range subItems {
		processedItem := sectionProcessor.ProcessSection(item)

		processedItems = append(processedItems, processedItem)
	}

	return processedItems
}
