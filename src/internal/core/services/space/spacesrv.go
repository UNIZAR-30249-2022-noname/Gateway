package space

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
)

type SpaceServiceImp struct {
	spaceRepository ports.SpaceRepository
}

func New(spaceRepository ports.SpaceRepository) *SpaceServiceImp {
	return &SpaceServiceImp{spaceRepository: spaceRepository}
}

func (svc *SpaceServiceImp) RequestInfoSlots(req domain.ReqInfoSlot) (domain.AllInfoSlot, error) {
	if(req.Name != "" && req.Date != ""){
		return svc.spaceRepository.RequestInfoSlots(req)
	} 
	return domain.AllInfoSlot{}, apperrors.ErrInvalidInput
}

func (svc *SpaceServiceImp) Reserve(sp domain.Space, init, end domain.Hour, date, person string) (string, error) {
	return svc.spaceRepository.Reserve(sp, init, end, date, person)
}

func (svc *SpaceServiceImp) ReserveBatch(spaces []domain.Space, init, end domain.Hour, dates []string, person string) (string, error) {
	return svc.spaceRepository.ReserveBatch(spaces, init, end, dates, person)
}
func (svc *SpaceServiceImp) FilterBy(params domain.SpaceFilterParams) ([]domain.Spaces, error) {
	return svc.spaceRepository.FilterBy(params)
}
