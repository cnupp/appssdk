include ../../includes.mk

SOURCES := util api
test: deps 
	ginkgo -r -trace -keepGoing . || exit 1;\

deps:
	@go get github.com/onsi/ginkgo/ginkgo
 
generate:
