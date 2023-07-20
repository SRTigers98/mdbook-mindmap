use std::{io, process};

use clap::{Arg, ArgMatches, Command};
use mdbook::{
    errors::Error,
    preprocess::{CmdPreprocessor, Preprocessor},
};
use mdbook_mindmap::preprocessor::MindmapPreprocessor;

fn make_app() -> Command {
    Command::new("mdbook-mindmap")
        .about("A mdbook preprocessor to create mindmaps from the file structure")
        .subcommand(
            Command::new("supports")
                .arg(Arg::new("renderer").required(true))
                .about("Checks if the renderer is supported"),
        )
}

fn main() {
    let matches = make_app().get_matches();

    let preprocessor = MindmapPreprocessor::new();

    if let Some(args) = matches.subcommand_matches("supports") {
        handle_supports(&preprocessor, args);
    } else if let Err(e) = handle_preprocessing(&preprocessor) {
        eprintln!("{}", e);
        process::exit(1);
    }
}

fn handle_supports(pre: &dyn Preprocessor, args: &ArgMatches) -> ! {
    let renderer = args
        .get_one::<String>("renderer")
        .expect("required argument");

    if pre.supports_renderer(renderer) {
        process::exit(0);
    } else {
        process::exit(1);
    }
}

fn handle_preprocessing(pre: &dyn Preprocessor) -> Result<(), Error> {
    let (ctx, book) = CmdPreprocessor::parse_input(io::stdin())?;

    let processed_book = pre.run(&ctx, book)?;
    serde_json::to_writer(io::stdout(), &processed_book)?;

    Ok(())
}
