package usecase

import (
	"fmt"

	"github.com/Hakeera/cripto/internal/infra"
	"github.com/Hakeera/cripto/internal/notifier"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type PriceService struct {
    client *infra.CoinGeckoClient
    store  *PriceStore
}

func NewPriceService(client *infra.CoinGeckoClient, store *PriceStore) *PriceService {
    return &PriceService{client: client, store: store}
}

// UpdatePrices obtém preços de Bitcoin, Ethereum e Solana via CoinGeckoClient e atualiza o store.
func (s *PriceService) UpdatePricesAndNotify(tg *notifier.TelegramClient) error {
coins := []string{"bitcoin", "ethereum", "solana"}
	prices, err := s.client.GetPrices(coins, "usd")
	if err != nil {
		return err
	}
	s.store.Update(prices)
	
	// Monta a mensagem com os preços
	message := "🪙 *Atualização de Preços*\n\n"
	caser := cases.Title(language.English)
	
	for coin, price := range prices {
		message += fmt.Sprintf("*%s*: $%.2f\n", caser.String(coin), price)
	}
	
	// Envia mensagem via Telegram
	if err := tg.SendMessage(message); err != nil {
		fmt.Println("Erro ao enviar alerta para Telegram:", err)
	}
	return nil
}

// PrintPrices imprime no console os preços armazenados.
func (s *PriceService) PrintPrices() {
    fmt.Println("===== Preços atuais =====")
    for coin, price := range s.store.Prices {
        fmt.Printf("%s: $%.2f\n", coin, price)
    }
    telegram := notifier.NewTelegramClient("SEU_BOT_TOKEN", "SEU_CHAT_ID")
    err := telegram.SendMessage("Preço de Bitcoin variou mais de 5%!")
    if err != nil {
	fmt.Println("Erro ao enviar alerta:", err)
    }

    fmt.Println("=========================")
}

