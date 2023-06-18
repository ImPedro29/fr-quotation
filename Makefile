create_migration:
	migrate create -ext sql -dir ./db/migrations $(NAME)

migrate_up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate_down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

test:
	go clean -testcache && go test ./...