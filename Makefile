.PHONY: all build
GOBIN=${GOPATH}/bin
REFLEX=${GOBIN}/reflex -v -s -R 'vendor\/|pkg\/|bin\/|build\/'

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  run         - run the program in dev mode"
	@echo ""
	@echo "  test        - run tests"
	@echo ""
	@echo "  deps        - pull and setup dependencies"
	@echo "  devdeps     - pull and setup dev dependencies"

run: build
	./bin/sixth_harmony \
			-v 1 \
			-logtostderr=true

devloop:
	${REFLEX} -- make run

test:
	${REFLEX} -- gb test

build:
	${GOBIN}/gb build

deps:
	# includes restore command
	go get -u github.com/constabulary/gb/...
	PATH=${PATH}:${GOBIN} ${GOBIN}/gb vendor restore

certs:
	test -f ./cert.pem -a -f ./key.pem || make cert-gen cert-import

cert-gen:
	go run generate_cert.go -host localhost -org-name sixth_harmony

cert-import:
	@echo 'Importing generated cert into Keychain requires sudo password:'
	sudo security add-trusted-cert -d -r trustRoot -k ${HOME}/Library/Keychains/login.keychain ./cert.pem

devdeps: certs deps
	go get github.com/cespare/reflex

deps-update:
	${GOBIN}/gb vendor update --all

