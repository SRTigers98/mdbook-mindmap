package service

var (
	sectionProcessor SectionProcessor
	chapterProcessor ChapterProcessor
	mindmapCreator   MindmapCreator
)

func init() {
	sectionProcessor = NewSectionProcessor()
	chapterProcessor = NewChapterProcessor()
	mindmapCreator = NewMindmapCreator()
}
