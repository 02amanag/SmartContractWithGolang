# Intract with Smart contract using Golang

## Setup
1. Install Ganache -> [Download](https://trufflesuite.com/ganache/)

2. Install  Solidity -> [Docs](https://docs.soliditylang.org/en/v0.8.2/installing-solidity.html)
```
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

3. Install  Geth -> [Docs](https://geth.ethereum.org/docs/install-and-build/installing-geth)
```
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install ethereum
```

## To deply contract 
```
$ make sol
```

## To intract with contract run 
```
$make run
```

## To Check and test use postman 
Import `blockChain.postman_collection.json` 
 
## Or curl commands

1. to check balance :- `curl --location --request GET 'http://localhost:1323/balance'`
2. to check admin address :- `curl --location --request GET 'http://localhost:1323/admin'`
3. to deposite 50 abount of value from account's private key in address :- 

`curl --location --request POST 'http://localhost:1323/deposite/50' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accountPrivateKey":"fd4eef6dec5575cc78f3f14d4b749094f8b88ad7883caaa8d1d24e9a01e3732d"
}'`

4. to withdrawl 10 :- 
`curl --location --request POST 'http://localhost:1323/withdrawl/50' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accountPrivateKey":"fd4eef6dec5575cc78f3f14d4b749094f8b88ad7883caaa8d1d24e9a01e3732d"
}'`
