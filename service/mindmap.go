package service

type MindmapCreator interface {
	CreateMindmap(chapter map[string]any) string
}

type DefaultMindmapCreator struct {
}

func NewMindmapCreator() MindmapCreator {
	return &DefaultMindmapCreator{}
}

func (c *DefaultMindmapCreator) CreateMindmap(chapter map[string]any) string {
	return ""
}
