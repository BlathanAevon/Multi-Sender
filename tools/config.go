package tools

import (
	"flag"
	"fmt"
)

const (
	defaultFromFile   = "keys.txt"
	defaultToFile     = "wallets.txt"
	defaultDelayFrom  = 100
	defaultDelayTo    = 1000
	defaultShuffle    = true
	defaultAllBalance = false
	deafaultDeadline  = 30
)

type Config struct {
	RPCURL          string
	WalletsFromPath string
	WalletsToPath   string
	AmountFrom      float64
	AmountTo        float64
	DelayFrom       int
	DelayTo         int
	AllBalance      bool
	TxDeadline      int
}

func ParseFlags() (*Config, error) {
	c := &Config{}

	flag.StringVar(&c.RPCURL, "rpc", "", "RPC url of the preffered network")
	flag.StringVar(&c.WalletsFromPath, "f", defaultFromFile, "Path of the file with private keys for wallets that you want to send from")
	flag.StringVar(&c.WalletsToPath, "t", defaultToFile, "Path of the file with adresses for wallets that you want to send to")
	flag.Float64Var(&c.AmountFrom, "af", 0, "Minimum amount that will be sent")
	flag.Float64Var(&c.AmountTo, "at", 0, "Maximum amount that will be sent")
	flag.IntVar(&c.DelayFrom, "df", defaultDelayFrom, "Minimum delay between transactions")
	flag.IntVar(&c.DelayTo, "dt", defaultDelayTo, "Maximum delay between transactions")
	flag.IntVar(&c.TxDeadline, "d", deafaultDeadline, "Transaction deadline in milliseconds")
	flag.BoolVar(&c.AllBalance, "a", defaultAllBalance, "Set if you want to send the whole balance")

	help := flag.Bool("h", false, "Display usage")

	flag.Parse()

	if *help {
		return nil, nil
	}

	if c.RPCURL == "" {
		return nil, fmt.Errorf("RPC URL is required")
	}

	if !c.AllBalance {

		if c.AmountTo == 0 {
			return nil, fmt.Errorf("\"Amount to\" is required")
		}

		if c.AmountFrom == 0 {
			return nil, fmt.Errorf("\"Amount from\" is required")
		}

		if c.AmountTo < c.AmountFrom {
			return nil, fmt.Errorf("\"Amount to\" should be greater than \"Amount from\"")
		}

	}

	if c.DelayTo <= c.DelayFrom {
		return nil, fmt.Errorf("\"Delay to\" should be greater than \"Delay from\"")
	}

	return c, nil

}
