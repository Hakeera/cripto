// internal/worker/ticker.go
package worker

import (
	"fmt"
	"time"

	"github.com/Hakeera/cripto/internal/usecase"
)

type PriceWorker struct {
    service *usecase.PriceService
    interval time.Duration
}

func NewPriceWorker(svc *usecase.PriceService, interval time.Duration) *PriceWorker {
    return &PriceWorker{service: svc, interval: interval}
}

func (w *PriceWorker) Start() {
    ticker := time.NewTicker(w.interval)
    defer ticker.Stop()
    for range ticker.C {
        if err := w.service.UpdatePrices(); err != nil {
            fmt.Println("Erro ao atualizar preços:", err)
        } else {
            w.service.PrintPrices()
        }
    }
}

