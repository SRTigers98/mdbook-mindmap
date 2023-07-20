use mdbook::{
    book::Book,
    errors::Result,
    preprocess::{Preprocessor, PreprocessorContext},
};

pub struct MindmapPreprocessor;

impl MindmapPreprocessor {
    pub fn new() -> Self {
        Self
    }
}

impl Preprocessor for MindmapPreprocessor {
    fn name(&self) -> &str {
        "mdbook-mindmap"
    }

    fn run(&self, _: &PreprocessorContext, book: Book) -> Result<Book> {
        Result::Ok(book)
    }

    fn supports_renderer(&self, _renderer: &str) -> bool {
        _renderer == "html"
    }
}
