# mdbook-mindmap

[![Build](https://github.com/SRTigers98/mdbook-mindmap/actions/workflows/build.yml/badge.svg)](https://github.com/SRTigers98/mdbook-mindmap/actions/workflows/build.yml)
[![mdBook](https://img.shields.io/badge/mdBook-0.4.35-blue.svg)](https://github.com/rust-lang/mdBook/releases/tag/v0.4.35)
[![Mermaid](https://img.shields.io/badge/Mermaid-10-pink.svg)](https://github.com/mermaid-js/mermaid)

A mdBook preprocessor to create mindmaps from the book structure.

## Installation

To use the preprocessor download the binary for your system and put it on your path.
Binaries for Linux, macOS and Windows are available from the [latest release](https://github.com/SRTigers98/mdbook-mindmap/releases/latest).
You should be able to execute the preprocessor from the command line as _mdbook-mindmap_.

### mdBook Configuration

To integrate the preprocessor in your mdBook, add the following lines to your **book.toml**:

```toml
[preprocessor.mindmap]
renderers = ["html"]
```

**Note:** The preprocessor only supports the HTML renderer!

## Usage

A mindmap can be generated with the following command:

> !mindmap

This will generate a mindmap containing the current chapter name as root and all sub chapters with their chapter names.
For this the library [mermaid.js](https://mermaid.js.org/) will be included on every page the command is used.

The mindmap will be generated with the [Mermaid mindmap syntax](https://mermaid.js.org/syntax/mindmap.html) and will render when the page is loaded in the browser.

### Example Output

```html
<pre class="mermaid">
mindmap
  root)Tools(
    408513086("`Gradle`")
    932387233("`NPM`")
    2596981504("`Kubernetes`")
      994518411("`Pod`")
      773890894("`Service`")
</pre>
<script type="module">
  import mermaid from "https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs";
  mermaid.initialize({ startOnLoad: true });
</script>
```
