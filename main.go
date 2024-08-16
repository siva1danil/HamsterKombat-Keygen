package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

func CreateClientId(seed int64) string {
	random := rand.New(rand.NewSource(seed))

	minTime, maxTime := int64(1721926800000), int64(1723820773000)
	minNumber, maxNumber := big.NewInt(0), big.NewInt(0)
	minNumber.SetString("100000000000000000", 10)
	maxNumber.SetString("9999999999999999999", 10)

	tmpNumber := big.NewInt(0)
	tmpNumber.Sub(maxNumber, minNumber)

	timestamp := minTime + random.Int63n(maxTime-minTime)
	number := big.NewInt(0)
	number.Rand(random, tmpNumber)
	number.Add(number, minNumber)

	id := fmt.Sprintf("%d-%d", timestamp, number)

	return id
}

func LoginClient(ctx context.Context, appToken string, clientId string) (string, error) {
	reqBody := fmt.Sprintf("{\"appToken\": \"%s\", \"clientId\": \"%s\", \"clientOrigin\": \"android\"}", appToken, clientId)
	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.gamepromo.io/promo/login-client", strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Debug(string(respBody))

	decoder := json.NewDecoder(strings.NewReader(string(respBody)))
	decoder.DisallowUnknownFields()
	decoded := struct {
		ClientToken string `json:"clientToken"`
	}{}
	err = decoder.Decode(&decoded)
	if err != nil {
		return "", err
	}

	return decoded.ClientToken, nil
}

func RegisterEvent(ctx context.Context, clientToken string, promoId string, eventId string, eventType string, eventOrigin string) (bool, error) {
	reqBody := fmt.Sprintf("{\"eventId\": \"%s\", \"eventType\": \"%s\", \"eventOrigin\": \"%s\", \"promoId\": \"%s\"}", eventId, eventType, eventOrigin, promoId)
	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.gamepromo.io/promo/register-event", strings.NewReader(reqBody))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+clientToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	log.Debug(string(respBody))

	decoder := json.NewDecoder(strings.NewReader(string(respBody)))
	decoder.DisallowUnknownFields()
	decoded := struct {
		HasCode      bool   `json:"hasCode"`
		ErrorCode    string `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	}{}
	err = decoder.Decode(&decoded)
	if err != nil {
		return false, err
	}

	return decoded.HasCode && decoded.ErrorCode == "" && decoded.ErrorMessage == "", nil
}

func CreateCode(ctx context.Context, clientToken string, promoId string) (string, error) {
	reqBody := fmt.Sprintf("{\"promoId\": \"%s\"}", promoId)
	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.gamepromo.io/promo/create-code", strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+clientToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	decoder := json.NewDecoder(strings.NewReader(string(respBody)))
	decoder.DisallowUnknownFields()
	decoded := struct {
		Code string `json:"promoCode"`
	}{}
	err = decoder.Decode(&decoded)
	if err != nil {
		return "", err
	}

	return decoded.Code, nil
}

func run(ctx context.Context, config ConfigEntry, clientId string, clientToken string) error {
	reader := bufio.NewReader(os.Stdin)

	if clientToken == "" {
		log.Info("Logging in client...")
		_clientToken, err := LoginClient(ctx, config.AppToken, clientId)
		if err != nil {
			return err
		}
		clientToken = _clientToken
		log.Info("Obtained client token", "clientToken", clientToken)
	}

	for {
		log.Info("Sleeping...")
		select {
		case <-time.After(60 * time.Second):
		case <-ctx.Done():
		}
		if ctx.Err() != nil {
			break
		}

		log.Info("Registering event...")
		hasCode, err := RegisterEvent(ctx, clientToken, config.PromoId, uuid.NewString(), config.EventType, config.EventOrigin)
		if err != nil {
			return err
		}

		if hasCode {
			log.Info("Creating code...")
			code, err := CreateCode(ctx, clientToken, config.PromoId)
			if err != nil {
				return err
			}
			log.Infof("Code: [ %s ]", code)

			log.Info("Generate a new code? (y/N)")
			line, _, err := reader.ReadLine()
			if err != nil {
				return err
			} else if line[0] != 'y' && line[0] != 'Y' {
				break
			}
		}
	}

	return nil
}

func main() {
	clientSeed := flag.Int("client-seed", 0, "Random seed used when generating client ID")
	app := flag.String("app", "", "App name (one of BIKE|CUBE|CLONE|TRAIN|MERGE|TWERK)")
	token := flag.String("token", "", "Client token (if available)")
	flag.Parse()

	if clientSeed == nil || app == nil {
		flag.PrintDefaults()
		os.Exit(1)
	}
	config, ok := Config[*app]
	if !ok {
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.SetLevel(log.DebugLevel)
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		cancel()
	}()

	clientId := CreateClientId(int64(*clientSeed))
	log.Info("Starting...", "clientSeed", *clientSeed, "app", *app, "token", *token, "clientId", clientId)
	run(ctx, config, clientId, *token)
	log.Info("Exiting")
}
