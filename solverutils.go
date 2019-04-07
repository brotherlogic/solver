package main

func (s *Server) solve(p, i, g int64) int64 {
	switch p {
	case 1:
		if i%3 == 0 || i%5 == 0 {
			return i
		}
	case 2:
		return s.solve2(i)
	case 3:
		return s.solve3(i, g)
	}

	return 0
}

func (s *Server) solve2(i int64) int64 {
	s1 := int64(1)
	s2 := int64(2)
	t := int64(0)
	sum := int64(2)
	found := int64(0)
	for s2 < i {
		t = s2 + s1
		s1 = s2
		s2 = t
		if s2%2 == 0 {
			found = int64(1)
			sum += s2
		} else {
			found = int64(0)
		}
	}

	return sum - (s2 * found)
}

func (s *Server) solve3(i int64, t int64) int64 {
	if isPrime(i) && t%i == 0 {
		return i
	}
	return 0
}
