package usecase

// PriceStore armazena os preços das criptomoedas em memória.
type PriceStore struct {
    Prices map[string]float64
}

func NewPriceStore() *PriceStore {
    return &PriceStore{Prices: make(map[string]float64)}
}

// Update sobrescreve os preços atuais no store.
func (ps *PriceStore) Update(newPrices map[string]float64) {
    for coin, price := range newPrices {
        ps.Prices[coin] = price
    }
}

