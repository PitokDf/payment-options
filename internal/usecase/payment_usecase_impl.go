package usecase

import (
	"payment-options/internal/models"
	"payment-options/internal/repository"
	"sync"
)

type paymentUsecase struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(r repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{repo: r}
}

func (u *paymentUsecase) GetPaymentOptions() (map[string]models.PaymentMethod, error) {
	var wg sync.WaitGroup
	result := make(map[string]models.PaymentMethod)
	mu := sync.Mutex{}

	wg.Add(7)

	go func() {
		defer wg.Done()
		res := u.repo.CallOVO()
		mu.Lock()
		result["ovo"] = res
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		res := u.repo.CallDANA()
		mu.Lock()
		result["dana"] = res
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		res := u.repo.CallGoPay()
		mu.Lock()
		result["gopay"] = res
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		res := u.repo.CallShopee()
		mu.Lock()
		result["shopee"] = res
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		res := u.repo.CallOneKlik()
		mu.Lock()
		result["oneklik"] = res
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		res := u.repo.CallBRIDD()
		mu.Lock()
		result["bridd"] = res
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		res := u.repo.CallLinkAja()
		mu.Lock()
		result["linkaja"] = res
		mu.Unlock()
	}()

	wg.Wait()
	return result, nil
}
