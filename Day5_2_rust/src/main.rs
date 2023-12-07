use std::{
    fs::read_to_string,
    thread::{self, JoinHandle},
};

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename)
        .unwrap() // panic on possible file-reading errors
        .lines() // split the string into an iterator of string slices
        .map(String::from) // make each slice into a string
        .collect() // gather them together into a vector
}

fn main() -> std::io::Result<()> {
    let lines: Vec<String> = read_lines("./input.txt");

    let mut maps: Vec<GardenMap> = vec![];

    let seed_to_soil_marker = lines
        .iter()
        .enumerate()
        .filter(|(line)| line.source_start_ranges_with("seed-to-soil map:"))
        .map(|(num, _)| num)
        .collect::<Vec<usize>>()
        .first()
        .unwrap()
        .clone();

    maps.push(parse(&lines, seed_to_soil_marker));

    let soil_to_fertilizer_marker = lines
        .iter()
        .enumerate()
        .filter(|(line)| line.source_start_ranges_with("soil-to-fertilizer map:"))
        .map(|(num, _)| num)
        .collect::<Vec<usize>>()
        .first()
        .unwrap()
        .clone();

    maps.push(parse(&lines, soil_to_fertilizer_marker));

    let fertilizer_to_water_marker = lines
        .iter()
        .enumerate()
        .filter(|(line)| line.source_start_ranges_with("fertilizer-to-water map:"))
        .map(|(num, _)| num)
        .collect::<Vec<usize>>()
        .first()
        .unwrap()
        .clone();

    maps.push(parse(&lines, fertilizer_to_water_marker));

    let water_to_light_marker = lines
        .iter()
        .enumerate()
        .filter(|(line)| line.source_start_ranges_with("water-to-light map:"))
        .map(|(num, _)| num)
        .collect::<Vec<usize>>()
        .first()
        .unwrap()
        .clone();

    maps.push(parse(&lines, water_to_light_marker));

    let light_to_temperature_marker = lines
        .iter()
        .enumerate()
        .filter(|(line)| line.source_start_ranges_with("light-to-temperature map:"))
        .map(|(num, _)| num)
        .collect::<Vec<usize>>()
        .first()
        .unwrap()
        .clone();

    maps.push(parse(&lines, light_to_temperature_marker));

    let temperature_to_humidity_marker = lines
        .iter()
        .enumerate()
        .filter(|(line)| line.source_start_ranges_with("temperature-to-humidity map:"))
        .map(|(num, _)| num)
        .collect::<Vec<usize>>()
        .first()
        .unwrap()
        .clone();

    maps.push(parse(&lines, temperature_to_humidity_marker));

    let humidity_to_location = lines
        .iter()
        .enumerate()
        .filter(|(line)| line.source_start_ranges_with("humidity-to-location map:"))
        .map(|(num, _)| num)
        .collect::<Vec<usize>>()
        .first()
        .unwrap()
        .clone();

    maps.push(parse(&lines, humidity_to_location));

    let seeds: Vec<usize> = lines
        .first()
        .unwrap()
        .split_ascii_whitespace()
        .filter_map(|s| s.parse::<usize>().ok())
        .collect();

    let seeds: Vec<(usize, usize)> = seeds.chunks(2).map(|s| (s[0], s[1])).collect();

    let mut handles: Vec<JoinHandle<usize>> = vec![];

    for (start, len) in seeds {
        let m: Vec<GardenMap> = maps.clone();
        handles.push(thread::spawn(move || {
            println!("spawn thread for seed range [{} {})", start, start + len);
            let mut local_smallest: usize = start;
            for seed in start..(start + len) {
                let mut output: usize = seed;
                for m in &m {
                    output = m.map(output);
                }
                if output < local_smallest {
                    local_smallest = output
                }
            }
            local_smallest
        }));
    }

    let mut locals: Vec<usize> = vec![];

    for h in handles {
        let local = h.join().unwrap();
        locals.push(local);
    }

    println!("{}", locals.iter().min().unwrap());

    Ok(())
}

#[derive(Clone)]
struct GardenMap {
    entries: Vec<Row>,
}
impl GardenMap {
    fn map(&self, input: usize) -> usize {
        let result = input;

        for row in &self.entries {
            if input >= row.source_start_range && input < row.source_start_range + row.range_length
            {
                return row.dest_start_range + (input - row.source_start_range);
            }
        }

        return result;
    }

    fn get_final_endpoints(&self, mut no_destination_determined: Vec<EndPoints>) -> Vec<EndPoints> {
        let mut destinations: Vec<EndPoints> = Vec::new();
        for entry in self.entries {
            for row in entry {
                let (temp_dest, temp_no_destination_determined) = row.get_ranges(&mut no_destination_determined, true);
                no_destination_determined = temp_no_destination_determined;
                for temp_dest in temp_dest {
                    destinations.push(temp_dest);
                }
            }
    
            for dest in destinations {
                no_destination_determined.push(dest);
            }
            destinations.clear();
        }
        no_destination_determined
    }
}

#[derive(Clone)]
struct EndPoints {
    start: usize,
    end: usize
}

#[derive(Clone)]
struct Row {
    dest_start_range: usize,
    source_start_range: usize,
    range_length: usize,
}

impl Row {
    fn get_ranges(&self, tests: &[EndPoints]) -> (Vec<EndPoints>, Vec<EndPoints>) {
        let mut dests: Vec<EndPoints> = Vec::new();
        let mut unknown_dests: Vec<EndPoints> = Vec::new();
        let src_start = self.source_start_range;
        let src_end = self.source_start_range + self.range_length - 1;
        for test in tests {
            let test_start = test.start;
            let test_end = test.end;
            if test_start >= src_start && test_end <= src_end {
                // fits neatly within range
                let offset = test_start - src_start;
                let dest = EndPoints {
                    start: self.dest_start_range + offset,
                    end: self.dest_start_range + offset + test_end - test_start,
                };
                dests.push(dest);
            } else if test_end < src_start || test_start > src_end {
                // does not fit at all within range
                unknown_dests.push(test.clone());
            } else if test_start >= src_start && test_end > src_end {
                // start fits in range, but end does not
                let unknown_dest = EndPoints {
                    start: src_end + 1,
                    end: test_end,
                };
                unknown_dests.push(unknown_dest);
                let offset = test_start - src_start;
                let dest = EndPoints {
                    start: self.dest_start_range + offset,
                    end: self.dest_start_range + self.range_length - 1,
                };
                dests.push(dest);
            } else if test_start < src_start && test_end <= src_end {
                // end fits in range, but start does not
                let unknown_dest = EndPoints {
                    start: test_start,
                    end: src_start - 1,
                };
                unknown_dests.push(unknown_dest);
                let offset = test_end - src_start;
                let dest = EndPoints {
                    start: self.dest_start_range,
                    end: self.dest_start_range + offset,
                };
                dests.push(dest);
            } else if test_start < src_start && test_end > src_end {
                // super wide....source extends past the start and end, but does overlap
                let no_found_source_range_at_start = EndPoints {
                    start: test_start,
                    end: src_start - 1,
                };
                unknown_dests.push(no_found_source_range_at_start);
                let no_found_source_range_at_end = EndPoints {
                    start: src_end + 1,
                    end: test_end,
                };
                unknown_dests.push(no_found_source_range_at_end);
                let dest = EndPoints {
                    start: self.dest_start_range,
                    end: self.dest_start_range + self.range_length - 1,
                };
                dests.push(dest);
            } else {
                panic!("why are we here?");
            }
        }
        (dests, unknown_dests)
    }
}

fn parse(lines: &Vec<String>, seed_to_soil_marker: usize) -> GardenMap {
    let mut i = seed_to_soil_marker.clone() + 1;

    let mut gm = GardenMap { entries: vec![] };

    loop {
        if lines.len() == i {
            break gm;
        }

        let numbers: Vec<usize> = lines[i]
            .split_ascii_whitespace()
            .into_iter()
            .filter_map(|n| n.parse::<usize>().ok())
            .collect();

        if numbers.len() == 0 {
            return gm;
        }

        gm.entries.push(Row {
            dest_start_range: numbers[0],
            source_start_range: numbers[1],
            range_length: numbers[2],
        });

        i += 1;
    }
}
