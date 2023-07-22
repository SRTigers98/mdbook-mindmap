use mdbook::{book::Chapter, BookItem};

const MINDMAP_INDICATOR: &str = "!mindmap";
const MERMAID_IMPORT: &str = "<script type=\"module\">
    import mermaid from \"https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs\";
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

    mindmap.push("<pre class=\"mermaid\">".into());
    mindmap.push("mindmap".into());
    mindmap.push(format!("  root){}(", root_name));
    mindmap.push(create_mindmap_items(&chapter.sub_items, 0));
    mindmap.push("</pre>".into());

    mindmap.join("\n")
}

fn create_mindmap_items(items: &Vec<BookItem>, depth: usize) -> String {
    let spacing = " ".repeat(4 + (depth * 2));
    let mut mindmap_items = Vec::<String>::new();

    for item in items.iter() {
        match item {
            BookItem::Chapter(chapter) => {
                let mindmap_item = format!(
                    "{}(\"`{}`\")",
                    chapter
                        .name
                        .replace(" ", "")
                        .replace("-", "")
                        .to_lowercase(),
                    chapter.name
                );
                mindmap_items.push(format!("{}{}", spacing, mindmap_item));

                let mindmap_sub_items = create_mindmap_items(&chapter.sub_items, depth + 1);
                if mindmap_sub_items.len() > 0 {
                    mindmap_items.push(mindmap_sub_items);
                }
            }
            _ => {}
        }
    }

    mindmap_items.join("\n")
}
