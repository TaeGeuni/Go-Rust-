class Solution:
    def firstStableIndex(self, nums: list[int], k: int) -> int:
        l = len(nums)
        min_check = [0] * l
        min_num = nums[l - 1]
        max_check = [0] * l
        max_num = nums[0]

        for i in range(l):
            min_num = min(min_num, nums[l - 1 - i])
            min_check[l - 1 - i] = min_num

            max_num = max(max_num, nums[i])
            max_check[i] = max_num

        for i in range(l):
            if max_check[i] - min_check[i] <= k:
                return i

        return -1
