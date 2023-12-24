use std::fs;
use std::collections::HashMap;


enum Signal {
    High,
    Low
}
enum Kind {
    Button,
    FlipFlopper,
    Inverter,
    Unknown
}
struct StateMachine {
    kinds: Kind,
    signals: HashMap<String, Signal>,
    workflows: HashMap<String, Workflow>,
}

struct Workflow {
    name: String,
    next_workflows: Vec<String>,
    // Define the fields of the Workflow struct here
}
fn parse_file(file_name: &str) -> StateMachine {
    let mut state_machine: StateMachine = StateMachine {
        kinds: Kind::Unknown,
        signals: HashMap::new(),
        workflows: HashMap::new(),
    };
    let mut v: Vec<i32> = Vec::new();
    let contents = fs::read_to_string(file_name)
        .expect("Something went wrong reading the file");
    for line in contents.lines() {
        let mut workflow: Workflow = Workflow {
            name: String::from(""),
            next_workflows: Vec::new(),
        };
        let mut tokens: Vec<String> = line.split_whitespace()
            .filter(|s| *s != "->")
            .map(|s| s.to_string())
            .collect();
        let raw_name = tokens.remove(0);
        let name = raw_name.trim_start_matches(&['%', '&'][..]);
        let next_workflows = tokens;
        let workflow = Workflow {
            name: name.to_string(),
            next_workflows,
        };
        let kind: Kind = match raw_name.as_str() {
            "broadcaster" => Kind::Button,
            _ if raw_name.starts_with('%') => Kind::FlipFlopper,
            _ if raw_name.starts_with('&') => Kind::Inverter,
            _ => panic!("Unknown kind")    
        };
        state_machine.kinds = kind;
        state_machine.workflows.insert(name.to_string(), workflow);
        state_machine.signals.insert(name.to_string(), Signal::Low);
    }
    return state_machine;
}

fn part_1(state_machine: &StateMachine) -> i32 {
    let mut sum: i32 = 0;

    return sum;
}

fn part_2(state_machine: &StateMachine) -> i32 {
    let mut sum: i32 = 0;

    return sum;
}

fn main() {
    let state_machine: StateMachine = parse_file("input.txt");
    println!("part1: {}", part_1(&state_machine));
    println!("part2: {}", part_2(&state_machine));
}
