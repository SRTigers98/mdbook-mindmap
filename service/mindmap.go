package service

import (
	"fmt"
	"hash/crc32"
	"strings"
)

const (
	MINDMAP_TEMPLATE = `<pre class="mermaid">
mindmap
  root)%s(
%s
</pre>`
)

type MindmapCreator interface {
	CreateMindmap(chapter map[string]any) string
}

type DefaultMindmapCreator struct {
}

func NewMindmapCreator() MindmapCreator {
	return &DefaultMindmapCreator{}
}

func (c *DefaultMindmapCreator) CreateMindmap(chapter map[string]any) string {
	rootName := chapter["name"].(string)

	items := createMindmapItems(chapter["sub_items"].([]any), 0)

	return fmt.Sprintf(MINDMAP_TEMPLATE, rootName, items)
}

func createMindmapItems(items []any, depth int) string {
	spacing := strings.Repeat(" ", 4+(2*depth))

	var mindmapItems []string
	for _, item := range items {
		if !sectionProcessor.IsChapter(item) {
			continue
		}

		chapter := item.(map[string]any)["Chapter"].(map[string]any)
		chapterName := chapter["name"].(string)
		chapterId := crc32.ChecksumIEEE([]byte(chapterName))
		mindmapItem := fmt.Sprintf("%s%d(\"`%s`\")", spacing, chapterId, chapterName)

		mindmapItems = append(mindmapItems, mindmapItem)

		subItems := createMindmapItems(chapter["sub_items"].([]any), depth+1)

		if len(subItems) > 0 {
			mindmapItems = append(mindmapItems, subItems)
		}
	}

	return strings.Join(mindmapItems, "\n")
}
