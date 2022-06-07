help:
	@echo 'Uso:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run: inicia o servidor de desenvolvimento backend na porta :4000 caso nenhum outro valor seja passado como parâmetro 
run:
	go run ./cmd/api/*.go

## install: instala as dependências tanto do front quanto backend 
install:
	go mod tidy
	cd ./frontend/ && npm install

## dev: inicia o servidor de desenvolvimento do svelte na porta 3000 
dev:
	cd ./frontend/ && npm run dev 

