use std::collections::HashMap;

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        let mut char_map = HashMap::new();
        s.chars().for_each(|c| *char_map.entry(c).or_insert(0) += 1);
        t.chars().for_each(|c| *char_map.entry(c).or_insert(0) -= 1);
        char_map.into_values().all(|count| count == 0)
    }
}