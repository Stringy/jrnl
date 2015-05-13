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

        let now  = time::now();
        let year = now.tm_year + 1900;
        let dirs = format!("./{}/{}", year, now.tm_mon + 1);
        let stamp = format!("{}/{}.md", dirs, now.tm_mday);
        println!("{}\n", stamp);

        fs::create_dir_all(dirs);

        process::new("vim").arg(stamp).status().unwrap_or_else(|e| {
            panic!("failed to run vim {}", e)
        });
    }

}
