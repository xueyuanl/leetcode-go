package array

func fib(N int) int {
	if (N == 0) {
		return 0;
	}
	if (N == 1) {
		return 1;
	}

	res := 0;
	v1, v2 := 0, 1;
	for i := 2; i <= N; i ++ {
		res = v1 + v2;
		v1 = v2;
		v2 = res;
	}

	return res;
}
