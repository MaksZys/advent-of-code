use std::fs::File;
use std::io::{BufRead, BufReader};

enum Direction {
    INCREASING,
    DECREASING,
    IDLE,
}

struct AnalysisResult {
    increasing: bool,
    direction: Direction
}


fn is_increasing(a: u64, b: u64) -> Direction {
   if a < b {
       return Direction::INCREASING;
   }

   Direction::DECREASING
}

fn analyse_record(previous_direction: &Direction, value1: &u64, value2: &u64) -> AnalysisResult {
    let direction = is_increasing(*value1, *value2);
    if check_direction(&previous_direction, &direction) {
        return AnalysisResult {
            direction,
            increasing: false,
        }
    }

    let diff = ( *value1 ).abs_diff(*value2);
    if diff > 3 || diff < 1 {
        return AnalysisResult {
            direction,
            increasing: false,
        }
    }

    AnalysisResult {
        direction,
        increasing: true,
    }
}

fn check_direction(direction: &Direction, previous: &Direction) -> bool {
    match (previous, direction) {
        (Direction::INCREASING, Direction::DECREASING) => {
           true
        }
        (Direction::DECREASING, Direction::INCREASING) => {
           true
        }
        _ => false
    }
}

fn main() {
    let file = File::open("test.txt").expect("Failed to open file");
    let reader = BufReader::new(file);

    let mut counter = 0;
    for line in reader.lines() {
        let record = line
            .unwrap()
            .split_whitespace()
            .collect::<Vec<_>>()
            .iter()
            .map(|s| s.parse::<u64>())
            .filter_map(Result::ok)
            .collect::<Vec<_>>();

        let mut previous_direction = Direction::IDLE;
        let mut increase_by = 1;

        for i in 1..record.len() {
            let direction = is_increasing(record[i - 1], record[i]);
            if check_direction(&previous_direction, &direction) {
                increase_by = 0;
            }
            previous_direction = direction;

            let diff = record[i - 1].abs_diff(record[i]);
            if diff > 3 || diff < 1 {
                increase_by = 0;
            }
        }

        counter += increase_by;
    }

    println!("Part 1: {}", counter);


    // Part 2
    let file = File::open("test.txt").expect("Failed to open file");
    let reader = BufReader::new(file);
    let mut counter = 0;

    let mut fallback_array: Vec<Vec<u64>> = Vec::new();
    for line in reader.lines() {
        let record = line
            .unwrap()
            .split_whitespace()
            .collect::<Vec<_>>()
            .iter()
            .map(|s| s.parse::<u64>())
            .filter_map(Result::ok)
            .collect::<Vec<_>>();

        let mut previous_direction = Direction::IDLE;
        let mut increase_by = 1;
        let mut fallback_record: Vec<u64> = Vec::new();
        fallback_record.push(record[0]);

        for i in 1..record.len() {
            let analysis_result = analyse_record(&previous_direction, &record[i], &record[i - 1]);
            previous_direction = analysis_result.direction;

            if increase_by == 1 && !analysis_result.increasing {
                increase_by = 0;
                continue;
            }

            // if !analysis_result.increasing || increase_by == 0 {
            //     // if increase_by != 0 {
            //     //     fallback_record.push(record[i]);
            //     // }
            //
            //     increase_by = 0;
            //     continue;
            // }

            fallback_record.push(record[i]);
        }

        counter += increase_by;
        if fallback_record.len() == record.len() - 1 {
            fallback_array.push(fallback_record);
        }
    }


    println!("Part 2 precalculation: {}", counter);
    print!("fallback array: {:?} \n", fallback_array);

    for i in 1..fallback_array.len() {
        let record = &fallback_array[i];
        let mut previous_direction = Direction::IDLE;
        let mut increase_by = 1;

        for j in 1..record.len() {
            let analysis_result = analyse_record(&previous_direction, &record[j], &record[j - 1]);
            previous_direction = analysis_result.direction;
            if !analysis_result.increasing || increase_by == 0 {
                increase_by = 0;
                break;
            }
        }

        counter += increase_by;
    }

    println!("Part 2: {}", counter);
}
