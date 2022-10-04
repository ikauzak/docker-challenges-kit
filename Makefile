GREEN=\033[0;32m
RED=\033[0;31m
ip=$(shell sh -c "hostname -I" | awk '{print $$2}')
HOST_PORT=8000
docs-build:
	@docker build . -t docs:lab

docs-stop:
	@docker stop docs
	@echo "$(RED)Docs desativada..."

docs-start:
	@docker run --name docs -dit --rm -p $(HOST_PORT):80 docs:lab
	@echo "$(GREEN)Documentação disponível em http://${ip}:$(HOST_PORT)"

start: docs-build docs-start
stop: docs-stop
