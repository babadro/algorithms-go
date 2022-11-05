use std::cmp::Ordering;
use std::collections::BinaryHeap;

#[derive(Eq, PartialEq)]
struct Num {
    num: i32,
    arr_idx: i32,
}

impl Ord for Num {
    fn cmp(&self, other: &Self) -> Ordering {
        other.num.cmp(&self.num)
    }
}

impl PartialOrd for Num {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

pub fn find_kth_smallest(lists: Vec<Vec<i32>>, k: i32) -> i32 {
    let mut num_idxes = vec![0; lists.len()];

    let mut heap = BinaryHeap::new();
    for arr_idx in 0..num_idxes.len() {
        let mut arr = &lists[arr_idx];
        if arr.len() > 0 {
            heap.push(Num{num: arr[0], arr_idx: arr_idx as i32 });
            num_idxes[arr_idx] = 1;
        }
    }

    if heap.len() == 0 {
        return i32::MAX;
    }

    let mut i: i32 = 0;
    while let Some(Num{num, arr_idx}) = heap.pop() {
        if i+1 == k {
            return num
        }

        let arr_idx: usize = arr_idx as usize;

        let num_idx = num_idxes[arr_idx];
        if num_idx < lists[arr_idx].len() {
            heap.push(Num{num: lists[arr_idx][num_idx as usize], arr_idx: arr_idx as i32});
            num_idxes[arr_idx] += 1;
        }

        i += 1
    }

    return i32::MAX
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_find_kth_smallest() {
        struct TestCase {
            lists: Vec<Vec<i32>>,
            k: i32,
            want: i32,
        }

        let test_cases = vec![
            TestCase{
                lists: vec![vec![2,6,8], vec![3,6,7],vec![1,3,4]],
                k: 5,
                want: 4
            },
            TestCase{
                lists: vec![vec![5,8,9], vec![1,7]],
                k: 3,
                want: 7,
            },
        ];

        for tc in test_cases {
            let got = find_kth_smallest(tc.lists, tc.k);
            assert_eq!(got, tc.want);
        }
    }
}
