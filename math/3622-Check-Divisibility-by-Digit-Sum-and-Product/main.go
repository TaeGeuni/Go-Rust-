package main

func checkDivisibility(n int) bool {
	sumRes, mulRes := 0, 1
	nBackup := n

	for n > 0 {
		num := n % 10
		sumRes += num
		mulRes *= num

		n = n / 10
	}

	return nBackup%(sumRes+mulRes) == 0
}
