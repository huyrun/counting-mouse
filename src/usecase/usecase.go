//go:generate mockgen -source usecase.go -destination ../test/mock/usecase.go -package mock

package usecase

import (
	"context"
	"github.com/huyrun/counting-mouse/src/mouses"
)

type UseCase interface {
	CountAliveMouse(ctx context.Context, originalAmount, lifeSpan, n int) int
}

type useCase struct {
}

func NewUseCase() *useCase {
	u := &useCase{}
	return u
}

func (u *useCase) CountAliveMouse(ctx context.Context, originalAmount, lifeSpan, n int) int {
	return mouses.CountAliveMouse(originalAmount, lifeSpan, n)
}
