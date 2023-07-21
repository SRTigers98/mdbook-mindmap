use mdbook::book::Chapter;

const MINDMAP_INDICATOR: &str = "!mindmap";

pub fn check_mindmap_indicator(chapter: &Chapter) -> bool {
    chapter.content.contains(MINDMAP_INDICATOR)
}

pub fn process_mindmap_indicator(chapter: &Chapter) -> String {
    chapter.content.clone()
}
