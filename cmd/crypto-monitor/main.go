// cmd/crypto-monitor/main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hakeera/cripto/internal/infra"
	"github.com/Hakeera/cripto/internal/usecase"
	"github.com/Hakeera/cripto/internal/worker"
)

func main() {
    // Instancia o Echo (framework web):contentReference[oaicite:19]{index=19}
    // e := echo.New()

    fmt.Println("Monitor de preços iniciado (CLI + Echo)")

    // Configura cliente CoinGecko, store em memória e serviço de preços
    client := infra.NewCoinGeckoClient()
    store := usecase.NewPriceStore()
    service := usecase.NewPriceService(client, store)

    // Inicia o worker para atualizar preços a cada 10 segundos
    w := worker.NewPriceWorker(service, 20*time.Second)
    go w.Start()

    // Aguarda interrupção (Ctrl+C) para encerrar
    fmt.Println("Pressione Ctrl+C para encerrar.")
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    <-quit
    fmt.Println("Encerrando aplicação...")
}

