// Problem link: https://leetcode.com/problems/two-sum/description/

func twoSum(nums []int, target int) []int {
    intMap := make(map[int]int)

    for index, num := range nums {
        complement := target - num 

        if i, ok := intMap[complement]; ok {
            return []int{index, i}
        }

        intMap[num] = index   
    }

    return nil
}
