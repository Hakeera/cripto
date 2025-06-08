// internal/worker/ticker.go
package worker

import (
	"fmt"
	"time"

	"github.com/Hakeera/cripto/internal/notifier"
	"github.com/Hakeera/cripto/internal/usecase"
)

type PriceWorker struct {
	service  *usecase.PriceService
	interval time.Duration
	telegram *notifier.TelegramClient
}

func NewPriceWorker(service *usecase.PriceService, interval time.Duration, telegram *notifier.TelegramClient) *PriceWorker {
	return &PriceWorker{
		service:  service,
		interval: interval,
		telegram: telegram,
	}
}

func (w *PriceWorker) Start() {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()
	for range ticker.C {
		if err := w.service.UpdatePricesAndNotify(w.telegram); err != nil {
			fmt.Println("Erro ao atualizar preços:", err)
		} else {
			w.service.PrintPrices()
		}
	}
}

