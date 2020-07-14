compile:
	@echo "This is not a compilable package."

doc:
	@go doc $(pkg)

test:
ifdef pkg
	@go test ./$(pkg)
else
	@go test ./...
endif