package main

func (s *Server) solve4(i, t int64) int64 {
	num1 := i / t
	num2 := i % t

	if isPalindrome(num1 * num2) {
		return num1 * num2
	}

	return 0
}

func (s *Server) solve5(v, t int64) int64 {
	for i := int64(2); i <= t; i++ {
		if v%i != 0 {
			return 0
		}
	}

	return v
}

func (s *Server) solve6(i int64) int64 {
	return i * i
}
