##### init #####

install:
	git pull
	$(MAKE) depend
	$(MAKE) devDepend
	$(MAKE) vendoring
	$(MAKE) bindata
	$(MAKE) gen
	cd front && npm install && npm build
	$(MAKE) local

##### depend #####

depend:
	go get -u github.com/goadesign/goa
	go get -u github.com/goadesign/goa/goagen
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/Masterminds/glide

devDepend:
	go get -u github.com/alecthomas/gometalinter/...
	gometalinter --install --update

vendoring:
	rm -rf ./vendor
	glide install

deps:
	$(MAKE) depend
	$(MAKE) devDepend
	$(MAKE) vendoring

##### goa ######

REPO:=github.com/pei0804/goa-spa-sample

init: depend bootstrap
gen: clean generate

bootstrap:
	goagen bootstrap -d $(REPO)/design

main:
	goagen main -d $(REPO)/design

clean:
	rm -rf app
	rm -rf client
	rm -rf tool
	rm -rf swagger
	rm -rf schema
	rm -rf js
	rm -f build

generate:
	goagen app     -d $(REPO)/design
	goagen swagger -d $(REPO)/design
	goagen client -d $(REPO)/design


swaggerUI:
	open http://localhost:8080/swagger/index.html

run:
	go run main.go

build:
	go build -o build

lint:
	@if [ "`gometalinter ./... --config=lint_config.json | tee /dev/stderr`" ]; then \
		echo "^ - lint err" && echo && exit 1; \
	fi

bindata:
	go-bindata -ignore bindata.go -pkg front -o front/bindata.go ./front/build/...

local:
	goapp serve ./server

staging-deploy:
	goapp deploy -application enows-staging ./server

staging-rollback:
	appcfg.py rollback ./server -A enows-staging