#
# tooling
#
gener8s:
	brew install nukleros/tap/gener8s

yot:
	brew install nukleros/tap/yot

tools: gener8s yot

#
# generated manifests/code
#
PLATFORM_CONFIG_OPERATOR_DIR ?= .source/platform-config-operator
PLATFORM_BOOTSTRAPPER_DIR ?= .source/platform-bootstrapper
download:
	vendir sync \
		--file=$(PLATFORM_CONFIG_OPERATOR_DIR)/config/vendor.yaml \
		--lock-file=$(PLATFORM_CONFIG_OPERATOR_DIR)/config/vendor.yaml.lock

overlay-operator:
	yot \
		--indent-level=2 \
		--instructions=$(PLATFORM_CONFIG_OPERATOR_DIR)/config/overlays.yaml \
		--output-directory=$(PLATFORM_CONFIG_OPERATOR_DIR) \
		--values-file=$(PLATFORM_CONFIG_OPERATOR_DIR)/config/values.yaml

overlay-bootstrapper:
	yot \
		--indent-level=2 \
		--instructions=$(PLATFORM_BOOTSTRAPPER_DIR)/config/overlays.yaml \
		--output-directory=$(PLATFORM_BOOTSTRAPPER_DIR) \
		--values-file=$(PLATFORM_BOOTSTRAPPER_DIR)/config/values.yaml

bootstrapper-manifests:
	rm -rf dist
	mkdir -p dist
	for manifest in $$(find $(PLATFORM_BOOTSTRAPPER_DIR)/manifests -type f); do \
		cat $$manifest >> dist/manifests.yaml; \
	done

MANIFESTS_DIR ?= $(PLATFORM_CONFIG_OPERATOR_DIR)/manifests
RESOURCES_FILE ?= internal/pkg/resources/generated.go
resources:
	echo 'package resources\n\nimport "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"\n\n' > $(RESOURCES_FILE)
	gener8s go -m $(MANIFESTS_DIR)/platform-config-operator.yaml >> $(RESOURCES_FILE)
	gener8s go -m $(MANIFESTS_DIR)/platform-operators.yaml >> $(RESOURCES_FILE)
	gener8s go -m $(MANIFESTS_DIR)/platform-config.yaml >> $(RESOURCES_FILE)
	go mod tidy

RBAC_FILE ?= $(PLATFORM_BOOTSTRAPPER_DIR)/static/rbac.yaml 
rbac:
	gener8s rbac yaml \
		--manifest-files=$(MANIFESTS_DIR)/*.yaml \
		--role-name=platform-bootstrapper \
		--verbs=create --verbs=update --verbs=get > $(RBAC_FILE)

generate: download overlay-operator resources fmt vet rbac overlay-bootstrapper

#
# platform-bootstrapper utility
#
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: build
build: generate fmt vet
	go build -o bin/platform-bootstrapper main.go

.PHONY: run
run: generate fmt vet
	go run ./main.go

GOLANGCI_LINT_VERSION ?= v1.55.2
install-linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

lint:
	golangci-lint run

IMG ?= quay.io/tbd-paas/platform-bootstrapper:latest
.PHONY: docker-build
docker-build:
	docker build -t $(IMG) .

.PHONY: docker-push
docker-push:
	docker push $(IMG)
