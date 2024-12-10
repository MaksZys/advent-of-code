use std::fs::File;
use std::io::{BufRead, BufReader};
use regex::Regex;

fn part1() {
    let file = File::open("input.txt").unwrap();
    let reader = BufReader::new(file);

    let mut sum = 0;
    for line in reader.lines() {
        let line = line.unwrap();

        let reg = Regex::new(r"mul\(\d+,\d+\)");
        for found in reg.unwrap().find_iter(&line) {
            let found = found.as_str()
                .replace("mul(", "")
                .replace(")", "")
                .split(",")
                .map(|x| x.parse::<u32>().unwrap_or_else(|_| 1))
                .collect::<Vec<u32>>();

            sum += found.iter().fold(1, |acc, x| acc * x);
        }
    }

    println!("{}", sum);
}

fn part2() {
    let file = File::open("input.txt").unwrap();
    let reader = BufReader::new(file);

    let mut sum: u64 = 0;
    let mut multiplayer = 1;
    for line in reader.lines() {
        let line = line.unwrap();

        let reg = Regex::new(r"mul\(\d+,\d+\)|don't\(\)|do\(\)");

        for found in reg.unwrap().find_iter(&line) {
            let found = found.as_str();
            if found == "do()" {
                multiplayer = 1;
                continue;
            } else if found == "don\'t()" {
                multiplayer = 0;
                continue;
            }

            if multiplayer == 1 {
                let value = found
                    .replace("mul(", "")
                    .replace(")", "")
                    .split(",")
                    .map(|x| x.parse::<u64>().unwrap_or_else(|_| 1))
                    .collect::<Vec<u64>>();

                sum += value.iter().fold(1, |acc, x| acc * x);
            }
        }
    }

    println!("{}", sum);
}

fn main() {
    // part1();
    part2();
}
