# Golang Multisenderr

## CLI tool that is created (idk why) to send a native currency (maybe tokens in future) from one wallet to multiple walllets in any EVM network, send from multiple wallets to multiple wallets in any EVM network. 

# How to use:

> [!TIP]
> You can run it using `go run`, you can run it using scripts from `scripts` folder or, you can compile it using `Makefille` just type `make` and it will compile the program for any of 3 OS's. Then, you can add this binary to `PATH` so you can run it from anywhere 

<br>

### On Linux / MacOS:
**Run**
`./scripts/run.sh`

or

`go run cmd/MultiSender/main.go`
    
### On Windows:
**Run**
`./scripts/run.bat`

or

`go run cmd/MultiSender/main.go`

### Compilation to binary

run `make` and wait until compilation is completed, in `build` folder, you'll see 3 executables.


## Usage

- `-a`  
  Set if you want to send the whole balance.

- `-af <float>`  
  Minimum amount that will be sent.

- `-at <float>`  
  Maximum amount that will be sent.

- `-d <int>`  
  Transaction deadline in milliseconds (default: `30`).

- `-df <int>`  
  Minimum delay between transactions (default: `100`).

- `-dt <int>`  
  Maximum delay between transactions (default: `1000`).

- `-f <string>`  
  Path of the file with private keys for wallets that you want to send from (default: `"keys.txt"`).

- `-h`  
  Display usage.

- `-rpc <string>`  
  RPC URL of the preferred network.

- `-t <string>`  
  Path of the file with addresses for wallets that you want to send to (default: `"wallets.txt"`).


