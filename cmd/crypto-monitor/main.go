// cmd/crypto-monitor/main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hakeera/cripto/internal/infra"
	"github.com/Hakeera/cripto/internal/notifier"
	"github.com/Hakeera/cripto/internal/usecase"
	"github.com/Hakeera/cripto/internal/worker"

	"github.com/joho/godotenv"
)

func main() {
    // Carrega o .env
    if err := godotenv.Load(); err != nil {
        fmt.Println("Erro ao carregar .env:", err)
    }

    fmt.Println("Monitor de preços iniciado (CLI + Echo)")

    // Inicializa cliente CoinGecko e store em memória
    client := infra.NewCoinGeckoClient()
    store := usecase.NewPriceStore()
    service := usecase.NewPriceService(client, store)

    // Inicializa cliente Telegram
    telegram := notifier.NewTelegramClient(
        os.Getenv("TELEGRAM_BOT_TOKEN"),
        os.Getenv("TELEGRAM_CHAT_ID"),
    )

    // Inicia o worker para atualizar preços a cada 20 segundos
    w := worker.NewPriceWorker(service, 20*time.Second, telegram)
    go w.Start()

    // Aguarda interrupção (Ctrl+C) para encerrar
    fmt.Println("Pressione Ctrl+C para encerrar.")
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    <-quit
    fmt.Println("Encerrando aplicação...")
}
