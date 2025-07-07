package services

import "ewallet-ums/internal/interfaces"

type Healcheck struct {
	HealcheckRepository interfaces.IHealthcheckRepository
}

func (s *Healcheck) HealtcheckService() (string, error) {
	return "service healty", nil
}
