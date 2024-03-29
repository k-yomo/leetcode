impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        for (i, a) in nums.iter().enumerate() {
            for (j, b) in nums[(i+1)..].iter().enumerate() {
                if a + b == target {
                    return vec![i as i32, (i+1+j) as i32]
                }
            }
        }
        panic!()
    }
}