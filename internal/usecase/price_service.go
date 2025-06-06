package usecase

import (
	"fmt"

	"github.com/Hakeera/cripto/internal/infra"
)

type PriceService struct {
    client *infra.CoinGeckoClient
    store  *PriceStore
}

func NewPriceService(client *infra.CoinGeckoClient, store *PriceStore) *PriceService {
    return &PriceService{client: client, store: store}
}

// UpdatePrices obtém preços de Bitcoin, Ethereum e Solana via CoinGeckoClient e atualiza o store.
func (s *PriceService) UpdatePrices() error {
    coins := []string{"bitcoin", "ethereum", "solana", "binancecoin", "ripple"}
    prices, err := s.client.GetPrices(coins, "usd")
    if err != nil {
        return err
    }
    s.store.Update(prices)
    return nil
}

// PrintPrices imprime no console os preços armazenados.
func (s *PriceService) PrintPrices() {
    fmt.Println("===== Preços atuais =====")
    for coin, price := range s.store.Prices {
        fmt.Printf("%s: $%.2f\n", coin, price)
    }
    fmt.Println("=========================")
}

