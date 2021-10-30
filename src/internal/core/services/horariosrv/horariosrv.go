package horariosrv

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
)

type service struct {
	horarioRepositorio ports.HorarioRepositorio
}

func New(horarioRepositorio ports.HorarioRepositorio) *service {
	return &service{horarioRepositorio: horarioRepositorio}
}

func (srv *service) GetAvailableHours(terna domain.Terna) ([]domain.AvailableHours, error) {
	res, err := srv.horarioRepositorio.GetAvailableHours(terna)
	if err != nil {
		return []domain.AvailableHours{}, err
	}

	return res, nil
}

func (srv *service) CreateNewEntry(entry domain.Entry) (string, error) {
	return "", nil
}
