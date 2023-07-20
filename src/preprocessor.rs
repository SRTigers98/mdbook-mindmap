use mdbook::{
    book::{Book, Chapter},
    errors::Result,
    preprocess::{Preprocessor, PreprocessorContext},
    BookItem,
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
        let mut processed_book = book.clone();

        processed_book.for_each_mut(|item| match item {
            BookItem::Chapter(chapter) => process_chapter(chapter),
            _ => {}
        });

        Result::Ok(processed_book)
    }

    fn supports_renderer(&self, _renderer: &str) -> bool {
        _renderer == "html"
    }
}

fn process_chapter(_: &mut Chapter) {}
