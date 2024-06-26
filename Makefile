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
	mkdir -p dist
	if [[ -f dist/manifests.yaml ]]; then rm -rf dist/manifests.yaml; fi
	for manifest in $$(ls $(PLATFORM_BOOTSTRAPPER_DIR)/manifests); do \
		cat $(PLATFORM_BOOTSTRAPPER_DIR)/manifests/$$manifest >> dist/manifests.yaml; \
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
		--role-name=platform-bootstrapper > $(RBAC_FILE)

generate: download overlay-operator resources fmt vet rbac overlay-bootstrapper bootstrapper-manifests

#
# platform-bootstrapper build
#
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: build
build: generate fmt vet
	export GOPRIVATE=github.com/tbd-paas/tbd-cli,github.com/tbd-paas/platform-config-operator && \
		go build -o bin/platform-bootstrapper main.go

.PHONY: run
run: generate fmt vet
	go run ./main.go

GOLANGCI_LINT_VERSION ?= v1.55.2
install-linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

lint:
	golangci-lint run

# NOTE: use of the tbd-access-token is temporary and will go away when we have all public repositories
IMG ?= quay.io/tbd-paas/platform-bootstrapper:latest
.PHONY: docker-build
docker-build:
	docker build --build-arg=GITHUB_TOKEN=$$(bw get password tbd-access-token) -t $(IMG) .

.PHONY: docker-push
docker-push:
	docker push $(IMG)

#
# platform-bootstrapper actions
#
install:
	export BOOTSTRAP_ACTION=apply && export KUBECONFIG="$${HOME}/.kube/config" && go run main.go

uninstall:
	export BOOTSTRAP_ACTION=destroy && export KUBECONFIG="$${HOME}/.kube/config" && go run main.go

install-job:
	kubectl apply -f dist/manifests.yaml
