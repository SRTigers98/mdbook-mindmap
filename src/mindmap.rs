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
    import_mermaid_js(chapter.content.clone()).replace(MINDMAP_INDICATOR, &create_mindmap(chapter))
}

fn import_mermaid_js(content: String) -> String {
    format!("{}\n{}", content, MERMAID_IMPORT)
}

fn create_mindmap(chapter: &Chapter) -> String {
    let root_name = &chapter.name;

    let mut mindmap = Vec::<String>::new();
    mindmap.push("<pre class='mermaid'>".into());
    mindmap.push("mindmap".into());
    mindmap.push(format!("  root(({}))", root_name));
    mindmap.push("</pre>".into());

    mindmap.join("\n")
}
