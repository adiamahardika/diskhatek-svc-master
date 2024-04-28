package usecases

import (
	"fmt"
	"time"
)

type schedulerUsecase usecase

type SchedulerUsecase interface {
	DeleteReservedStock()
}

func (u *schedulerUsecase) DeleteReservedStock() {
	fmt.Println("Scheduler running at: ", time.Now().Format(time.DateTime))
	err := u.Options.Repository.ReservedStock.DeleteReservedStock()
	fmt.Println(err)
}
