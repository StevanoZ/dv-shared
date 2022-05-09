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
	mockgen -package shrd_mock_svc -destination service/mock/message_broker_service_mock.go -source=service/message_broker_service.go

test:
	go test -v -cover ./...
	
q2c:
	sqlc generate

.PHONY: migrateInit createDb migrateUp migrateUp1 migrateDown migrateDown1 q2c mockSvc mockToken test
