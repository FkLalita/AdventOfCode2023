use std::fs::read_to_string;

fn main() {
    println!("Hello, world!");
    cube_count();
}

fn cube_count() {
    let inputs = read_to_string("input.txt").expect("error reading files");
    let mut sum_id = 0;
    let mut totalsum = 0;

    // Initialize maximum counts to the minimum possible value

    for input in inputs.lines() {
        let mut red_max = 0;
        let mut blue_max = 0;
        let mut green_max = 0;
        let (info, game) = input.split_once(": ").unwrap();
        let (_, info) = info.split_once(" ").unwrap();
        let info: i32 = info.parse().unwrap();
        let mut is_possible = true;

        let game = game.split("; ").collect::<Vec<_>>();
        for g in game {
            let turns = g.split(", ").collect::<Vec<_>>();
            for t in turns {
                let (_, color) = t.split_once(" ").unwrap();

                // Update maximum counts
                let cube: i32 = t.split_once(" ").unwrap().0.parse().unwrap();
                match color {
                    "red" => red_max = red_max.max(cube),
                    "blue" => blue_max = blue_max.max(cube),
                    "green" => green_max = green_max.max(cube),
                    _ => (),
                }

                // Check if the current cube count violates the conditions
                if (color == "red" && cube > 12)
                    || (color == "blue" && cube > 14)
                    || (color == "green" && cube > 13)
                {
                    is_possible = false;
                }
            }
        }
        totalsum += red_max * green_max * blue_max;

        if is_possible {
            // Calculate the product of maximum counts for each color

            sum_id += info;
        }
    }

    println!(
        "The sum of the possible game IDs is {}.\nThe sum of the powers of the minimum sets is {}.",
        sum_id, totalsum
    );
}
