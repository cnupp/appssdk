include ../includes.mk

BINARY_DEST_DIR=rootfs/usr/bin
IMAGE=$(REGISTRY_BUCKET)/builder
BUILD_WRAPPER_IMAGE=$(REGISTRY_BUCKET)/build-wrapper
VERIFY_WRAPPER_IMAGE=$(REGISTRY_BUCKET)/verify-wrapper
HELPERS := run-build
SOURCES := util api

build: check-docker
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o $(BINARY_DEST_DIR)/builder cli/builder.go || exit 1;
	@$(call check-static-binary,$(BINARY_DEST_DIR)/builder)
	for i in $(HELPERS); do \
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o $(BINARY_DEST_DIR)/$$i helpers/$$i.go || exit 1; \
	done
	@for i in $(HELPERS); do \
		$(call check-static-binary,$(BINARY_DEST_DIR)/$$i); \
	done
	@docker build -t $(IMAGE) rootfs
	@docker build -t $(BUILD_WRAPPER_IMAGE) images/build

push: check-registry
	@docker tag -f $(IMAGE) $(REGISTRY)$(IMAGE)
	@docker push $(REGISTRY)$(IMAGE)
	@docker tag -f $(BUILD_WRAPPER_IMAGE) $(REGISTRY)$(BUILD_WRAPPER_IMAGE)
	@docker push $(REGISTRY)$(BUILD_WRAPPER_IMAGE)

stop:
	@curl -X DELETE http://$(MARATHON_ENDPOINT)/v2/apps/builder?force=true


start:
	@curl -X POST http://$(MARATHON_ENDPOINT)/v2/apps?force=true -d @builder.json -H "Content-type: application/json"

restart: stop start

deploy: build push restart

test: test-unit

test-unit: ginkgo

ginkgo:
	for i in $(SOURCES); do \
		ginkgo -p $$i generic || exit 1;\
	done

generate:
