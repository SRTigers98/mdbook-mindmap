package service

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"strings"
	"text/template"
)

const (
	MINDMAP_TEMPLATE = `<pre class="mermaid">
mindmap
  root){{ .name }}(
{{- range .items }}
    {{ . }}
{{- end }}
</pre>`
)

type MindmapCreator interface {
	CreateMindmap(chapter map[string]any) string
}

type DefaultMindmapCreator struct {
	tmpl template.Template
}

func NewMindmapCreator() MindmapCreator {
	return &DefaultMindmapCreator{
		tmpl: *template.Must(template.New("mindmap").Parse(MINDMAP_TEMPLATE)),
	}
}

func (c *DefaultMindmapCreator) CreateMindmap(chapter map[string]any) string {
	rootName := chapter["name"].(string)

	items := createMindmapItems(chapter["sub_items"].([]any), 0)

	var mindmap bytes.Buffer
	c.tmpl.Execute(&mindmap, map[string]any{
		"name":  rootName,
		"items": items,
	})

	return mindmap.String()
}

func createMindmapItems(items []any, depth int) []string {
	spacing := strings.Repeat(" ", 2*depth)

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
			mindmapItems = append(mindmapItems, subItems...)
		}
	}

	return mindmapItems
}
