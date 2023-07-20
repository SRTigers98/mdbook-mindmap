use clap::{Arg, Command};

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

    if let Some(sub_args) = matches.subcommand_matches("supports") {
        print!("supports command called!")
    } else {
        print!("root command called!")
    }
}
