use mdbook::{
    book::{Book, Chapter},
    errors::Result,
    preprocess::{Preprocessor, PreprocessorContext},
    BookItem,
};

use crate::mindmap::{check_mindmap_indicator, process_chapter_content};

pub struct MindmapPreprocessor;

impl MindmapPreprocessor {
    pub fn new() -> Self {
        Self
    }
}

impl Default for MindmapPreprocessor {
    fn default() -> Self {
        Self::new()
    }
}

impl Preprocessor for MindmapPreprocessor {
    fn name(&self) -> &str {
        "mdbook-mindmap"
    }

    fn run(&self, _: &PreprocessorContext, book: Book) -> Result<Book> {
        let mut processed_book = book.clone();

        processed_book.for_each_mut(process_book_item);

        Result::Ok(processed_book)
    }

    fn supports_renderer(&self, _renderer: &str) -> bool {
        _renderer == "html"
    }
}

fn process_book_item(item: &mut BookItem) {
    if let BookItem::Chapter(chapter) = item {
        process_chapter(chapter);
    }
}

fn process_chapter(chapter: &mut Chapter) {
    if check_mindmap_indicator(chapter) {
        chapter.content = process_chapter_content(chapter);
    }

    chapter.sub_items.iter_mut().for_each(process_book_item);
}
