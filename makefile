CURRENT_DIR := $(shell pwd)

serve:
	go run cmd/api/main.go
migrate:
	go run cmd/migration/main.go
storage-link:
	ln -sfn $(CURRENT_DIR)/storage/public public/storage