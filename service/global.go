package service

var (
	sectionProcessor SectionProcessor
)

func init() {
	sectionProcessor = NewSectionProcessor()
}
