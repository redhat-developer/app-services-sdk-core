.PHONY: generate-sdk
generate-sdk:
	kiota generate -l typescript -d ../.openapi/account-management.json -o ./packages/account-management-sdk
	kiota generate -l typescript -d ../.openapi/connector_mgmt.yaml -o ./packages/connector-management-sdk
	kiota generate -l typescript -d ../.openapi/kafka-admin-rest.yaml -o ./packages/kafka-instance-sdk
	kiota generate -l typescript -d ../.openapi/srs-fleet-manager.json -o ./packages/registry-management-sdk
	kiota generate -l typescript -d ../.openapi/kafka-admin-rest.yaml -o ./packages/kafka-management-sdk
	kiota generate -l typescript -d ../.openapi/registry-instance.json -o ./packages/registry-instance-sdk
	kiota generate -l typescript -d ../.openapi/smartevents_mgmt_v2.yaml -o ./packages/smart-events-management-sdk
	kiota generate -l typescript -d ../.openapi/service-accounts.yaml -o ./packages/service-accounts-sdk

.PHONY: build
build:
	(cd packages/account-management-sdk && npx tsc apiClient.ts --outDir dist/ -d)