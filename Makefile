genswag:
	@rm -rf docs
	@httpswag -api main.go -dst docs
	@cd docs && swag init -o .