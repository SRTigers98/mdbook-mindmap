package service

var (
	sectionProcessor SectionProcessor
	chapterProcessor ChapterProcessor
)

func init() {
	sectionProcessor = NewSectionProcessor()
	chapterProcessor = NewChapterProcessor()
}
