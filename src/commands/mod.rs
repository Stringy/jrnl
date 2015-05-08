use std::process::Command as process;
use std::fs;

extern crate time;

pub trait Command {
    fn run(&self, args: Vec<String>);
    fn name(&self) -> &str;
}

pub struct NewCommand {
    pub name: &'static str,
}

impl Command for NewCommand {

    fn run(&self, args: Vec<String>) {
        println!("Running {} with args {:?}\n", self.name, args);
    }

    fn name(&self) -> &str {
        self.name
    }

}

pub struct AddCommand {
    pub name: &'static str,
    pub desc: &'static str,
    pub usage: &'static str,
    pub longuse: &'static str,
}

impl Command for AddCommand {

    fn name(&self) -> &str {
        self.name 
    }

    fn run(&self, args: Vec<String>) {
        let mut filename: &str;
        if args.len() == 1 {
            filename = &args[0];
        } else {
            filename = "entry";
        }

        let now = time::now();
        let stamp = format!("{}/{}/{}/{}.md", now.tm_year, now.tm_mon, now.tm_mday, filename);
        println!("{}\n", stamp);

        process::new("vim").arg(stamp).status().unwrap_or_else(|e| {
            panic!("failed to run vim {}", e)
        });

        // let attr = try!(fs::metadata(stamp));
    }

}
