package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/BlathanAevon/MultiSender/internal/client"
	"github.com/BlathanAevon/MultiSender/internal/wallet"
	"github.com/BlathanAevon/MultiSender/tools"
)

func MultiSender() {

	rpc, err := client.NewClient("https://bartio.rpc.berachain.com/")

	if err != nil {
		log.Fatal("Could not create an rpc")
	}

	walletsFrom, err := tools.UnpackTxt("./configs/wallets_keys_from.txt")
	if err != nil {
		log.Fatal("Could not get wallets from the file", err)
	}

	fmt.Printf("Wallets FROM loaded: %d\n", len(walletsFrom))

	if len(walletsFrom) < 1 {
		log.Fatalf("Amount of wallets is %v, load wallets!", len(walletsFrom))
	}

	walletsTo, err := tools.UnpackTxt("./configs/wallets_to.txt")
	if err != nil {
		log.Fatal("Could not get wallets from the file", err)
	}

	fmt.Printf("Wallets FROM loaded: %d\n", len(walletsTo))

	if len(walletsTo) < 1 {
		log.Fatalf("Amount of wallets is %v, load wallets!", len(walletsTo))
	}

	errCh := make(chan error)

	if len(walletsFrom) == 1 {

		walletFrom, err := wallet.NewWallet(walletsFrom[0])

		if err != nil {
			errCh <- err
		}

		for _, walletTo := range walletsTo {

			tx, err := walletFrom.SendNative(walletTo, rpc, 0.1)

			if err != nil {
				fmt.Errorf("could not send native: %v", err)
			}

			fmt.Printf("Wallet: %s | Sent %.5f of Native | Hash: %v\n", walletFrom.Address.Hex(), 0.1, tx.Hex())

			time.Sleep(1000 * time.Millisecond)
		}

		return

	}

	for i := 0; i < len(walletsFrom); i++ {

		to := walletsTo[i]

		go func(from string, to string) {
			wallet, err := wallet.NewWallet(from)

			if err != nil {
				errCh <- err
			}

			balance, err := rpc.Client.BalanceAt(context.Background(), wallet.Address, nil)

			if err != nil {
				errCh <- err
			}

			f := new(big.Float).SetInt(balance)

			divisor := new(big.Float).SetFloat64(1e18)

			f.Quo(f, divisor)

			result, _ := f.Float64()
			tx, err := wallet.SendNative(to, rpc, result*0.90)

			if err != nil {
				errCh <- fmt.Errorf("could not send native: %v", err)
			}

			fmt.Printf("Wallet: %s | Sent %.5f of Native | Hash: %v\n", wallet.Address.Hex(), result, tx.Hex())

			errCh <- nil

		}(walletsFrom[i], to)

		time.Sleep(1000 * time.Millisecond)

	}

	for i := 0; i < len(walletsFrom); i++ {
		if err := <-errCh; err != nil {
			log.Fatal("Error occured during multi-send\n", err)
		}
	}
	close(errCh)

}

func main() {

	MultiSender()

}
