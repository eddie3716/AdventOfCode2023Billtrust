use core::panic;
use std::fs;
use std::collections::HashMap;
use std::collections::HashSet;
use std::collections::VecDeque;
use std::collections::hash_map::Entry;


#[derive(PartialEq, Clone, Copy, Debug)]
enum Signal {
    High,
    Low
}

#[derive(PartialEq, Debug)]
enum Kind {
    Button,
    FlipFlopper,
    Inverter
}
struct QueueItem<'a> {
    signal: Signal,
    from_name: &'a String,
    to_workflow: &'a Workflow,
}

struct Workflow {
    name: String,
    kind: Kind,
    senders: Vec<String>,
    next_workflows: Vec<String>,
    // Define the fields of the Workflow struct here
}
fn parse_file(file_name: &str) -> HashMap::<String, Workflow> {
    let mut workflows: HashMap<String, Workflow> = HashMap::<String, Workflow>::new();
    let contents = fs::read_to_string(file_name)
        .expect("Something went wrong reading the file");
    for line in contents.lines() {
        let mut tokens: Vec<String> = line.split(|c: char| c.is_whitespace() || c == ',')
            .filter(|s| *s != "->" && *s !="")
            .map(|s| s.to_string())
            .collect();
        let raw_name = tokens.remove(0);
        let name = raw_name.trim_start_matches(&['%', '&'][..]);
        let next_workflows = tokens;
        
        let kind: Kind = match raw_name.as_str() {
            "broadcaster" => Kind::Button,
            _ if raw_name.starts_with('%') => Kind::FlipFlopper,
            _ if raw_name.starts_with('&') => Kind::Inverter,
            _ => panic!("Unknown kind")    
        };
        let workflow = Workflow {
            name: name.to_string(),
            kind: kind,
            senders: Vec::new(),
            next_workflows,
        };
        workflows.insert(name.to_string(), workflow);
    }

    for (name, &mut workflow) in workflows.iter_mut() {
        if workflow.kind == Kind::Inverter {
            for (other_name, other_workflow) in workflows.iter_mut() {
                if other_workflow.next_workflows.contains(&name) {
                    workflow.senders.push(other_name.clone());
                }
            }
        }
    }

    return workflows;
}

fn part_1(workflows: &mut HashMap::<String, Workflow>) -> i32 {
    let mut low_pulses: i32 = 0;
    let mut high_pulses: i32 = 0;
    let mut queue: VecDeque<QueueItem> = VecDeque::new();

    let mut on_or_off: HashMap::<String, bool> = HashMap::<String, bool>::new();
    let mut inverter_signals: HashMap::<String, Signal> = HashMap::<String, Signal>::new();
    
    for (name, workflow) in workflows.iter() {
        if workflow.kind == Kind::FlipFlopper {
            on_or_off.insert(name.clone(), false);
        }
    }

    let button: String = "button".to_string();
    
    for _i in 1..=1000 {
        low_pulses += 1;

        let mut visited: HashSet::<String> = HashSet::<String>::new();

        if let Some(broadcaster) = workflows.get("broadcaster") {
            let queue_item = QueueItem {
                signal: Signal::Low,
                from_name: &button,
                to_workflow: broadcaster,
            };
            queue.push_back(queue_item);
        } else {
            panic!("No broadcaster found");
        }
        println!("{} -> {:?} -> ({})", &button, Signal::Low, "broadcaster");
        while let Some(queue_item) = queue.pop_front() {
            let to_workflow = queue_item.to_workflow;
            let from_name = queue_item.from_name;
            let current_signal = queue_item.signal;
            let current_kind = &to_workflow.kind;
            let current_workflow_name = to_workflow.name.clone();
            let new_signal: Signal;
            if current_kind == &Kind::FlipFlopper {
                if let Entry::Occupied(mut entry) = on_or_off.entry(current_workflow_name.clone()) {
                    if current_signal == Signal::Low {
                        *entry.get_mut() = !*entry.get();
                    }
                    if *entry.get() {
                        new_signal = Signal::High;
                    } else {
                        new_signal = Signal::Low;
                    }
                } else {
                    panic!("No signal found for '{}'..we've reached a termination", current_workflow_name);
                }                   
            } else if current_kind == &Kind::Inverter {
                let inverter_key = to_workflow.name.clone() + "_" + from_name;
                if let Some(inverter_signal) = inverter_signals.get(&inverter_key) {
                    if *inverter_signal == Signal::Low {
                        new_signal = Signal::Low;
                    } else {
                        new_signal = Signal::High;
                    }
                } else {
                    new_signal = Signal::Low;
                }
                if current_signal == Signal::Low {
                    inverter_signals.insert(inverter_key.clone(), Signal::Low);
                } else {
                    inverter_signals.insert(inverter_key.clone(), Signal::High);
                }
            } else if current_workflow_name == "broadcaster" {
                new_signal = Signal::Low;              
            } else {
                if current_signal == Signal::Low {
                    new_signal = Signal::Low;   
                } else {
                    new_signal = Signal::High;
                }
            }
            for next_workflow_name in &to_workflow.next_workflows {

                let new_signal_str = match new_signal {
                    Signal::Low => "low",
                    Signal::High => "high",
                };
                let visited_key = current_workflow_name.clone() + "_" + next_workflow_name + "_" + new_signal_str;
                if visited.contains(&visited_key) {
                    continue;
                } else {
                    visited.insert(visited_key);
                }
                if new_signal == Signal::Low {
                    low_pulses += 1;
                } else {
                    high_pulses += 1;
                }

                if let Some(next_workflow) = workflows.get(next_workflow_name) {
                    println!("{} -> {:?} -> ({})", current_workflow_name, new_signal, next_workflow_name);
                    let new_queue_item:QueueItem<'_> = QueueItem {
                        signal: new_signal,
                        from_name: &to_workflow.name,
                        to_workflow: next_workflow,
                    };
                    queue.push_back(new_queue_item);
                } else {
                    //println!("No workflow found for '{}'", next_workflow_name);
                }
            }
        }
    }
    println!("low_pulses: {}", low_pulses);
    println!("high_pulses: {}", high_pulses);
    return low_pulses*high_pulses;
}

fn part_2(_workflows: &mut HashMap::<String, Workflow>) -> i32 {
    let sum: i32 = 0;

    return sum;
}

fn main() {
    let mut workflows:HashMap<String, Workflow>  = parse_file("testinput.txt");
    println!("part1: {}", part_1(&mut workflows));
    println!("part2: {}", part_2(&mut workflows));
}
