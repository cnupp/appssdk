include ../../includes.mk

SOURCES := util api

test: test-unit

test-unit: ginkgo

ginkgo:
	for i in $(SOURCES); do \
		ginkgo -p $$i generic || exit 1;\
	done

generate:
