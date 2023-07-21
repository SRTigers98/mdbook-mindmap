use mdbook::book::Chapter;

const MINDMAP_INDICATOR: &str = "!mindmap";
const MERMAID_IMPORT: &str = "<script type='module'>
    import mermaid from 'https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs';
    mermaid.initialize({ startOnLoad: true });
</script>";

pub fn check_mindmap_indicator(chapter: &Chapter) -> bool {
    chapter.content.contains(MINDMAP_INDICATOR)
}

pub fn process_chapter_content(chapter: &Chapter) -> String {
    let content = chapter.content.clone();

    import_mermaid_js(content)
}

fn import_mermaid_js(content: String) -> String {
    format!("{}\n{}", content, MERMAID_IMPORT)
}
