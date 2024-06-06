download:
	vendir sync --file platform-config-operator/config/vendor.yaml --lock-file platform-config-operator/config/vendor.yaml.lock

overlay:
	yot -I 2 -i platform-config-operator/config/overlays.yaml -o platform-config-operator -f platform-config-operator/config/values.yaml

gener8s:
	brew install nukleros/tap/gener8s

resources:
	echo 'package resources\n\nimport "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"\n\n' > ../internal/pkg/resources/generated.go
	gener8s go -m platform-config-operator/manifests/platform-config-operator.yaml >> ../internal/pkg/resources/generated.go
	gener8s go -m platform-config-operator/manifests/platform-operators.yaml >> ../internal/pkg/resources/generated.go
	gener8s go -m platform-config-operator/manifests/platform-config.yaml >> ../internal/pkg/resources/generated.go
	cd .. && go mod tidy
