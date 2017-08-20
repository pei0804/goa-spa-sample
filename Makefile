##### Convenience targets ######

REPO:=github.com/pei0804/goa-spa-sample

init: depend bootstrap
gen: clean generate
rerun: clean generate run

depend:
	@which goagen || go get -v github.com/goadesign/goa/goagen
	@go get -v ./...

bootstrap:
	@goagen bootstrap -d $(REPO)/design

main:
	@goagen main -d $(REPO)/design

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger
	@rm -rf schema
	@rm -rf js
	@rm -f build

generate:
	@goagen app     -d $(REPO)/design
	@goagen swagger -d $(REPO)/design
	@goagen client -d $(REPO)/design
	@goagen js -d $(REPO)/design
	@goagen schema -d $(REPO)/design
	@go build -o build

swaggerUI:
	@open http://localhost:8080/swagger/index.html

model:
	@rm -rf models
	@goagen --design=$(REPO)/design gen --pkg-path=github.com/goadesign/gorma

run:
	@go run main.go

build:
	@go build -o build