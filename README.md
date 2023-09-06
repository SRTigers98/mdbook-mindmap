# mdbook-mindmap

[![Build](https://github.com/SRTigers98/mdbook-mindmap/actions/workflows/build.yml/badge.svg)](https://github.com/SRTigers98/mdbook-mindmap/actions/workflows/build.yml)

A mdbook preprocessor to create mindmaps from the book structure.

## Installation

Pre-built binaries will be available via GitHub Releases as soon the first version is released.

## Usage

A mindmap can be generated with the following command:

> !mindmap

This will generate a mindmap containing the current chapter name as root and all sub chapters with their chapter names.
For this the library [mermaid.js](https://mermaid.js.org/) will be included on every page the command is used.
