ZEEBE_ADDRESS=localhost:26500
POSTGRES_URL=postgres://admin:admin@localhost:5432/zeebe-example

run: run-wallet run-donation

run-purchase-api:
	ZEEBE_ADDRESS=$(ZEEBE_ADDRESS) POSTGRES_URL=$(POSTGRES_URL) go run ./purchase api

run-purchase-worker:
	ZEEBE_ADDRESS=$(ZEEBE_ADDRESS) POSTGRES_URL=$(POSTGRES_URL) go run ./purchase worker

run-approval:
	ZEEBE_ADDRESS=$(ZEEBE_ADDRESS) POSTGRES_URL=$(POSTGRES_URL) go run ./approval

run-test:
	ZEEBE_ADDRESS=$(ZEEBE_ADDRESS) POSTGRES_URL=$(POSTGRES_URL) go run ./test-broker