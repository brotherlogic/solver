package main

func (s *Server) solve(p, i int64) int64 {
	switch p {
	case 1:
		if i%3 == 0 || i%5 == 0 {
			return i
		}
	case 2:
		return s.solve2(i)
	}

	return 0
}

func (s *Server) solve2(i int64) int64 {
	s1 := int64(1)
	s2 := int64(2)
	t := int64(0)
	sum := int64(3)
	for s2 < i {
		t = s2 + s1
		s1 = s2
		s2 = t
		sum += s2
	}
	return sum - s2
}
