BUILDPATH=$(CURDIR)
BINARY=rest_body_logging

makedir:
	@if [ ! -d $(BUILDPATH)/build/bin ] ; then mkdir -p $(BUILDPATH)/build/bin ; fi

mod:
	@echo "Vendoring..."
	@go mod vendor

build: makedir 
	@echo "Compilando..."
	@go build -mod vendor -ldflags "-s -w" -o $(BUILDPATH)/build/bin/${BINARY} cmd/${BINARY}/main.go
	@echo "Binario generado en build/bin/"${BINARY}

test: 
	@echo "Ejecutando tests..."
	@go test ./... -v

coverage: 
	@echo "Coverfile..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -html=coverfile_out -o coverfile_out.html


clean:
	@echo "Limpiando binarios..."
	@if [ -d $(BUILDPATH)/build/bin ] ; then rm -rf $(BUILDPATH)/build/ ; fi
	@rm -rf coverfile_out*

.PHONY: clean build test