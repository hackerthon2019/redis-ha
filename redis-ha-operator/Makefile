GOPATH=${HOME}/go

.DEFAULT_GOAL := all

.PHONY: build-redis-st-image
build-redis-st-image:
	GOOS=linux GOARCH=amd64 go get github.com/RedisLabs/sentinel_tunnel
	cp -f ${GOPATH}/bin/linux_amd64/sentinel_tunnel helm-charts/redis-ha-st/docker/redis-st/sentinel_tunnel
	chmod +x helm-charts/redis-ha-st/docker/redis-st/sentinel_tunnel
	docker build --no-cache --rm -t redis-st:latest helm-charts/redis-ha-st/docker/redis-st/

.PHONY: test-chart
test-chart:
	helm install --name demo helm-charts/redis-ha-st

.PHONY: install-crd
install-crd:
	kubectl apply -f deploy/crds/hackathon_v1alpha1_redisha_crd.yaml

.PHONY: install-operator
install-operator:
	kubectl apply -f deploy/service_account.yaml
	kubectl apply -f deploy/role.yaml
	kubectl apply -f deploy/role_binding.yaml
	kubectl apply -f deploy/operator.yaml

.PHONY: install-cr
install-cr:
	kubectl apply -f deploy/crds/hackathon_v1alpha1_redisha_cr.yaml

.PHONY: all
all: build-redis-st-image install-crd install-operator install-cr