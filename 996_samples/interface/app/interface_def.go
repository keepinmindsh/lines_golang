package app

import "996_samples/interface/domain"

type Sample struct {
}

func (s *Sample) Print(value string) {

}

func (s *Sample) Clone() domain.Interface {
	return &Sample{}
}
