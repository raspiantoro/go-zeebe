ZEEBE_ADDRESS=localhost:26500
POSTGRES_URL=postgres://admin:admin@localhost:5432/zeebe-example

run-purchase-api:
	ZEEBE_ADDRESS=$(ZEEBE_ADDRESS) POSTGRES_URL=$(POSTGRES_URL) go run ./purchase api

run-purchase-worker:
	ZEEBE_ADDRESS=$(ZEEBE_ADDRESS) POSTGRES_URL=$(POSTGRES_URL) go run ./purchase worker

run-approval-api:
	ZEEBE_ADDRESS=$(ZEEBE_ADDRESS) POSTGRES_URL=$(POSTGRES_URL) go run ./approval api

run-test:
	ZEEBE_ADDRESS=$(ZEEBE_ADDRESS) POSTGRES_URL=$(POSTGRES_URL) go run ./test-broker