from collections import Counter


class Solution:
    def largestInteger(self, nums: List[int], k: int) -> int:
        res = -1
        l = len(nums)

        # 1. 파이썬에서는 Counter를 쓰면 빈도수 해시맵을 한 줄로 만들 수 있습니다.
        # check = {값: 빈도수} 구조가 됩니다.
        check = Counter(nums)

        # 2. k == l 인 경우 (모든 숫자가 대상일 때 최댓값)
        if k == l:
            for num in nums:  # num은 인덱스가 아니라 실제 값입니다.
                if res < num:
                    res = num

        # 3. k == 1 인 경우 (빈도수가 1인 유일한 숫자 중 최댓값)
        elif k == 1:
            for num in nums:
                if check[num] == 1 and res < num:
                    res = num

        # 4. 그 외의 경우 (첫 번째 값과 마지막 값 중 빈도수가 1인 것 비교)
        else:
            first_num = nums[0]
            last_num = nums[l - 1]

            if check[first_num] == 1 and res < first_num:
                res = first_num
            if check[last_num] == 1 and res < last_num:
                res = last_num

        return res
