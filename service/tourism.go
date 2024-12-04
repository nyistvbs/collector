package service

func (s *Service) xx() {
	val, err := s.queue.Dequeue()
	if err != nil {

	}
	_ = val

	s.dao.Tourism()
}
