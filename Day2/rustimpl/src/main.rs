use std::fs::read_to_string;

fn main() {
    println!("Hello, world!");
    cube_count()
}

fn cube_count() {
    let inputs = read_to_string("input.txt").expect("error reading files");
    let mut sum_id = 0;
    let mut totalsum = 0;

    for input in inputs.lines() {
        let (info, game) = input.split_once(": ").unwrap();
        let (_, info) = info.split_once(" ").unwrap();
        let info: i32 = info.parse().unwrap();
        let mut is_possible = true;

        let game = game.split("; ").collect::<Vec<_>>();
        let mut redsumV: Vec<i32> = Vec::new();
        let mut bluesumV: Vec<i32> = Vec::new();
        let mut greensumV: Vec<i32> = Vec::new();

        for g in game {
            let turns = g.split(", ").collect::<Vec<_>>();
            for t in turns {
                let (cube, color) = t.split_once(" ").unwrap();
                let cube: i32 = cube.parse().unwrap();

                if color == "red" {
                    redsumV.push(cube);
                } else if color == "blue" {
                    bluesumV.push(cube);
                } else if color == "green" {
                    greensumV.push(cube);
                }
            }

            let redsum = *redsumV.iter().max().unwrap_or(&0);
            let bluesum = *bluesumV.iter().max().unwrap_or(&0);
            let greensum = *greensumV.iter().max().unwrap_or(&0);

            if redsum > 12 || bluesum > 14 || greensum > 13 {
                is_possible = false;
            }

            totalsum += redsum * greensum * bluesum;
        }

        if is_possible {
            sum_id += info
        }
    }
    println!(
        "The sum of the possible game IDs is {} and the sum of the power of the minimum sets is {}",
        sum_id, totalsum
    )
}
