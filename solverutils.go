package main

func (s *Server) solve(p, i int64) int64 {
	switch p {
	case 1:
		if i%3 == 0 || i%5 == 0 {
			return i
		}
	}

	return 0
}
