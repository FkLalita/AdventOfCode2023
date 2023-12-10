pub fn extract_concatenated_digits(str: &str) -> Vec<i32> {
    let mut digits = Vec::new();
    for line in str.lines() {
        let line_digits: Vec<i32> = line
            .chars()
            .filter_map(|c| c.to_digit(10).map(|n| n as i32))
            .collect();

        let concatenated =
            line_digits[0].to_string() + &line_digits[line_digits.len() - 1].to_string();
        if let Ok(num) = concatenated.parse::<i32>() {
            digits.push(num)
        }
    }
    digits
}
