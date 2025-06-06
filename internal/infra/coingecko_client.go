package infra

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type CoinGeckoClient struct {
    httpClient *http.Client
}

func NewCoinGeckoClient() *CoinGeckoClient {
    return &CoinGeckoClient{httpClient: &http.Client{}}
}

// GetPrices consulta a API pública do CoinGecko (endpoint /simple/price) e retorna um mapa coin->USD.
func (c *CoinGeckoClient) GetPrices(ids []string, vsCurrency string) (map[string]float64, error) {
    url := fmt.Sprintf(
        "https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s",
        strings.Join(ids, ","), vsCurrency,
    )
    resp, err := c.httpClient.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()


    var raw map[string]map[string]FloatOrString
    if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
    	return nil, err
    }

    prices := make(map[string]float64)
    for coin, m := range raw {
	    prices[coin] = float64(m[vsCurrency])
    }

    return prices, nil
}

type FloatOrString float64

func (f *FloatOrString) UnmarshalJSON(data []byte) error {
	var num float64
	if err := json.Unmarshal(data, &num); err == nil {
		*f = FloatOrString(num)
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		n, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		*f = FloatOrString(n)
		return nil
	}
	return fmt.Errorf("valor inválido: %s", string(data))
}

