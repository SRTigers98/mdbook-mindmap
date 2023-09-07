package service

type Preprocessor interface {
	ProcessBook(context map[string]any, book map[string]any) map[string]any
}

type DefaultPreprocessor struct {
}

func NewPreprocessor() Preprocessor {
	return &DefaultPreprocessor{}
}

func (p *DefaultPreprocessor) ProcessBook(context map[string]any, book map[string]any) map[string]any {
	return book
}
