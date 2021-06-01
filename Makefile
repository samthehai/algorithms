test:
	@echo 'Run Go testcases'
	@go test ./...

	@echo 'Run Python testcases'
	@nosetests -v
