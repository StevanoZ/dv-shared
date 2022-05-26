export POSTGRESQL_URL='postgresql://postgres:postgres_password@localhost:5454/dwiz_vent?sslmode=disable'

migrateInit:
	migrate create -ext sql -dir db/migration -seq alter_user

createDb:
	docker exec -it postgres_1 createdb -U postgres dv_test_db

dropDb:
	docker exec -it postgres_1 dropdb -U postgres dwiz_vent --force

migrateUp:
	migrate -path db/migration -database ${POSTGRESQL_URL} -verbose up

migrateUp1:
	migrate -path db/migration -database ${POSTGRESQL_URL} -verbose up 1

migrateDown:
	migrate -path db/migration -database ${POSTGRESQL_URL} -verbose down

migrateDown1:
	migrate -path db/migration -database ${POSTGRESQL_URL} -verbose down 1

mockToken:
	mockgen -package shrd_mock_token -destination token/mock/maker_mock.go -source=token/maker.go

mockSvc:
	mockgen -package shrd_mock_svc -destination service/mock/pubsub_service_mock.go -source=service/pubsub_service.go

mockPubSub:
	mockgen -package mock_pubsub -destination pubsub/mock/pubsub_mock.go -source=pubsub/pubsub.go

test:
	go test -v -covermode=atomic -race -coverpkg=./... ./... \
	-coverprofile coverage.out.tmp && cat coverage.out.tmp | grep -v "_mock.go" | grep -v "main.go" | grep -v "helper/" | grep -v "_gen.go" > coverage.out && rm coverage.out.tmp && \
	go tool cover -func coverage.out
	
q2c:
	sqlc generate

.PHONY: migrateInit createDb migrateUp migrateUp1 migrateDown migrateDown1 q2c mockSvc mockToken mockPubSub test

