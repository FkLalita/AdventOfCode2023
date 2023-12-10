pub fn text_to_digit(input: &str) -> usize {
    let num_words = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];

    input
        .lines()
        .map(|line| {
            let mut extracted: Vec<char> = Vec::new();
            let mut current_word = String::new();

            for c in line.chars() {
                if let Some(digit) = num_words.iter().position(|&word| word == current_word) {
                    extracted.push(char::from_digit(digit as u32 + 1, 10).unwrap());
                    current_word.clear();
                } else if c.is_ascii_digit() {
                    extracted.push(c);
                } else if c.is_ascii_alphabetic() {
                    current_word.push(c);
                }
            }

            extracted.iter().collect::<String>().parse().unwrap_or(0)
        })
        .sum()
}

