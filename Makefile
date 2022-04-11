ZEEBE_ADDRESS=localhost:26500

run: run-wallet run-donation

run-wallet:
	@ZEEBE_ADDRESS=$(ZEEBE_ADDRESS) go run ./wallet

run-donation:
	@ZEEBE_ADDRESS=$(ZEEBE_ADDRESS) go run ./donation