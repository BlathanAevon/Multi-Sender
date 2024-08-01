package sender

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/BlathanAevon/MultiSender/internal/client"
	"github.com/BlathanAevon/MultiSender/internal/wallet"
	"github.com/BlathanAevon/MultiSender/tools"
)

func Disperse(c *tools.Config) error {

	rpc, err := client.NewClient(c.RPCURL)

	if err != nil {
		return fmt.Errorf("could not create an rpc: %v", err)
	}

	walletsFrom, err := tools.UnpackTxt(c.WalletsFromPath)

	if err != nil {
		return fmt.Errorf("could not get wallets from the file: %v", err)
	}

	fmt.Printf("Wallets FROM loaded: %d\n", len(walletsFrom))

	if len(walletsFrom) < 1 {
		return fmt.Errorf("amount of wallets is %v, load wallets", len(walletsFrom))
	}

	walletsTo, err := tools.UnpackTxt(c.WalletsToPath)

	if err != nil {
		return fmt.Errorf("could not get wallets from the file: %v", err)
	}

	fmt.Printf("Wallets TO loaded: %d\n", len(walletsTo))

	if len(walletsTo) < 1 {
		return fmt.Errorf("amount of wallets is %v, load wallets", len(walletsTo))
	}

	if len(walletsFrom) == 1 {
		if c.AllBalance {
			return fmt.Errorf("can't send whole balance when sending from 1 wallet")
		}

		wallet, err := wallet.NewWallet(walletsFrom[0])

		if err != nil {
			return fmt.Errorf("creating wallet: %w", err)
		}

		for i := 0; i < len(walletsTo); i++ {

			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			amount := c.AmountFrom + r.Float64()*(c.AmountTo-c.AmountFrom)

			tts := time.Duration(r.Intn(c.DelayTo-c.DelayFrom)+c.DelayFrom) * time.Millisecond

			time.Sleep(tts)

			_, err = wallet.SendNative(walletsTo[i], rpc, amount*0.99, c.TxDeadline)

			if err != nil {
				return fmt.Errorf("send native: %v", err)
			}

			fmt.Printf("%s | sent %.5f |\n", wallet.Address.Hex(), amount)

		}

		return nil

	}

	if len(walletsFrom) != len(walletsTo) {
		return fmt.Errorf("amount of wallets is not equal")
	}

	errCh := make(chan error)

	for i := 0; i < len(walletsTo); i++ {

		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		go func(from string, to string) {
			wallet, err := wallet.NewWallet(from)
			amount := 0.0

			if err != nil {
				errCh <- fmt.Errorf("creating wallet: %w", err)
				return
			}

			if c.AllBalance {

				balance, err := rpc.Client.BalanceAt(context.Background(), wallet.Address, nil)

				if err != nil {
					errCh <- fmt.Errorf("getting balance: %w", err)
					return
				}

				f := new(big.Float).SetInt(balance)

				divisor := new(big.Float).SetFloat64(1e18)

				f.Quo(f, divisor)

				amount, _ = f.Float64()
			} else {
				amount = c.AmountFrom + r.Float64()*(c.AmountTo-c.AmountFrom)
			}
			tts := time.Duration(r.Intn(c.DelayTo-c.DelayFrom)+c.DelayFrom) * time.Millisecond

			time.Sleep(tts)

			_, err = wallet.SendNative(to, rpc, amount*0.99, c.TxDeadline)

			if err != nil {
				errCh <- fmt.Errorf("send native: %v", err)
				return
			}

			fmt.Printf("%s | sent %.5f |\n", wallet.Address.Hex(), amount)

			errCh <- nil

		}(walletsFrom[i], walletsTo[i])

	}

	for i := 0; i < len(walletsFrom); i++ {
		if err := <-errCh; err != nil {
			fmt.Println(err)
		}
	}

	close(errCh)

	return nil

}
