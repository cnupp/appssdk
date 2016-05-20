include ../../includes.mk

SOURCES := util api

test: test-unit

test-unit: ginkgo

ginkgo:
	for i in $(SOURCES); do \
		ginkgo -trace -keepGoing $$i || exit 1;\
	done

generate:
