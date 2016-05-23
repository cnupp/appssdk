include ../../includes.mk

SOURCES := util api

test: test-unit

test-unit:
	ginkgo -r -trace -keepGoing . || exit 1;\

generate:
