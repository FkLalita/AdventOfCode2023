use std::fs::read_to_string;
//use std::fs::File;
fn main() {
    println!("Hello, world!");
    cube_count()
}

fn cube_count() {
    let inputs = read_to_string("src/input.txt").expect("error reading files");
    let mut sum_id = 0;

    for input in inputs.lines() {
        let (info, game) = input.split_once(": ").unwrap();
        let (_, info) = info.split_once(" ").unwrap();
        let info: i32 = info.parse().unwrap();
        let mut is_possible = true;

        let game = game.split("; ").collect::<Vec<_>>();
        for g in game {
            let turns = g.split(", ").collect::<Vec<_>>();
            for t in turns {
                let (cube, color) = t.split_once(" ").unwrap();
                let cube: i32 = cube.parse().unwrap();

                if color == "red" && cube > 12 {
                    is_possible = false
                } else if color == "blue" && cube > 14 {
                    is_possible = false
                } else if color == "green" && cube > 13 {
                    is_possible = false
                }
            }
        }

        if is_possible {
            sum_id += info
        }
    }
    println!("{}", sum_id)
}
