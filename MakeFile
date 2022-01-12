run:
	go run main.go

sol:
	solc --optimize --abi ./contracts/MySmartContract.sol -o build
	solc --optimize --bin ./contracts/MySmartContract.sol -o build
	abigen --abi=./build/MySmartContract.abi --bin=./build/MySmartContract.bin --pkg=api --out=./api/MySmartContract.go