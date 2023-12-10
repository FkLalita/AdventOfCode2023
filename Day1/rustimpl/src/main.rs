use std::io::Read;

mod part1;
mod part2;

fn main() {
    // Open and read the file
    let mut file = std::fs::File::open("input.txt").unwrap();

    let mut content = String::new();
    file.read_to_string(&mut content).unwrap();

    // Example using file content
    let dfd_f = part2::text_to_digit(&content);
    println!(
        "Part two: print out the value of text to digit: {:?}",
        dfd_f
    );

    //    let mut contents = String::new();
    //    file.read_to_string(&mut contents).unwrap();

    let results = part1::extract_concatenated_digits(&content);

    let mut res = 0;
    for result in results {
        res += result
    }

    println!(
        "Part 1: Print out the value of digit concatenated together{}",
        res
    )
}
