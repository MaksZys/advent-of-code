use std::collections::HashMap;
use std::fs::File;
use std::io;
use std::io::{BufRead, BufReader};

fn part1(row1: &Vec<u32>, row2: &Vec<u32>) -> u32 {
    let mut sum: u32 = 0;

    if row2.len() == row1.len() {
        for i in 0..row1.len() {
            sum += row1[i].abs_diff(row2[i]);
        }
    }

    return sum;
}

fn part2(row1: &Vec<u32>, row2: &Vec<u32>) -> u64 {
    let mut number_map: HashMap<u32, u32> = HashMap::new();

    for i in 0..row1.len() {
        if !number_map.contains_key(&row1[i]) {
            number_map.insert(row1[i].clone(), 0);
        }
    }

    for i in 0..row1.len() {
        if !number_map.contains_key(&row2[i]) {
            continue;
        }

        let row2_value = number_map.get(&row2[i]).unwrap();
        let increased_row2_value = *row2_value + 1;
        number_map.insert(row2[i].clone(), increased_row2_value);
    }

    // let sum: u64 =0;
    // for (key, value) in &number_map {
    //     println!("Key: {}, Value: {}", key, value);
    // }

    let mut sum: u64 = 0;
    for i in 0..row1.len() {
        let multiplier = *number_map.get(&row1[i]).unwrap() as u64;

        sum += row1[i] as u64 * multiplier;
    }

    sum
}

fn main() -> io::Result<()> {
    let file_path = "input.txt";

    let file = File::open(file_path)?;
    let reader = BufReader::new(file);

    let mut row1: Vec<u32> = vec![];
    let mut row2: Vec<u32> = vec![];

    let separator = "   ";
    for line in reader.lines() {
        let line = line.unwrap();
        let parts: Vec<&str> = line.split(separator).collect();

        if parts.len() >= 2 {
            row1.push(parts[0].parse::<u32>().unwrap_or(0));
            row2.push(parts[1].parse::<u32>().unwrap_or(0));
        }
    }

    row1.sort();
    row2.sort();

    println!("Part 1: {} \n", part1(&row1, &row2));
    println!("Part 2: {} \n", part2(&row1, &row2));

    Ok(())
}
