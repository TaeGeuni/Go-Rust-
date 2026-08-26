class Solution:
    def checkDivisibility(self, n: int) -> bool:
        sum_num = 0
        mul_num = 1
        n_backup = n

        while n > 0:
            num = n % 10
            sum_num += num
            mul_num *= num
            n //= 10

        return n_backup % (sum_num + mul_num) == 0


# "/" and "//"
