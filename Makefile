# Makefile


migrate: # DB migration
	cd cmd/migrator && go run migrator.go --migrate --migration-path="../../database/migrations" --cfg-path="../../config/config.yaml"

rollback:
	cd cmd/migrator && go run migrator.go --rollback --migration-path="../../database/migrations" --cfg-path="../../config/config.yaml"
