
use std::env;

mod commands;
use commands::Command;

fn main() {
    let cmds = vec![
        commands::AddCommand {
            name: "add",
            usage: "journal add <filename>",
            desc: "add a new entry for the current day",
            longuse: ""
        }
    ];

    for cmd in cmds.iter() {
        if cmd.name() == env::args().nth(1).unwrap() {
            cmd.run(env::args().collect());
        } else {
            println!("No valid command found");
        }
    }
}

