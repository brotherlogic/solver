package main

func (s *Server) solve4(i, t int64) int64 {
	num1 := i / t
	num2 := i % t

	if isPalindrome(num1 * num2) {
		return num1 * num2
	}

	return 0
}
