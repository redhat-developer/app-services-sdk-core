


module.exports ={
    go : {
        kafkamgmt: {
            definition: require("./errors_kafka_mgmt.json"),
            file: "app-services-sdk-go/kafkamgmt/apiv1/error/errors.go"
        },
        registrymgmt:  {
            definition: require("./errors_srs_mgmt.json"),
            file: "app-services-sdk-go/registrymgmt/apiv1/error/errors.go"
        },
        connectormgmt: {
            definition: require("./errors_connector_mgmt.json"),
            file: "app-services-sdk-go/connectormgmt/apiv1/error/errors.go"
        }, 
    },
    js : {
        kafkamgmt: {
            definition: require("./errors_kafka_mgmt.json"),
            file: "app-services-sdk-js/packages/kafka-management-sdk/src/errors.ts"
        },
        registrymgmt:  {
            definition: require("./errors_srs_mgmt.json"),
            file: "app-services-sdk-js/packages/registry-management-sdk/src/errors.ts"
        },
        connectormgmt: {
            definition: require("./errors_connector_mgmt.json"),
            file: "app-services-sdk-js/packages/connector-management-sdk/src/errors.ts"
        }, 
        kafkainstance: {
            definition: require("./errors_kafka_instance.json"),
            file: "app-services-sdk-js/packages/kafka-instance-sdk/src/errors.ts"
        }, 
    },
    java: {
        kafkamgmt: {
            definition: require("./errors_kafka_mgmt.json"),
            file: "app-services-sdk-java/packages/kafka-management-sdk/src/main/java/com/openshift/cloud/api/kas/ApiErrorType.java"
        },
        registrymgmt:  {
            definition: require("./errors_srs_mgmt.json"),
            file: "app-services-sdk-java/packages/registry-management-sdk/src/main/java/com/openshift/cloud/api/srs/ApiErrorType.java"
        },
        connectormgmt: {
            definition: require("./errors_connector_mgmt.json"),
            file: "app-services-sdk-java/packages/connector-management-sdk/src/main/java/com/openshift/cloud/api/connector/ApiErrorType.java"
        },
        kafkainstance: {
            definition: require("./errors_kafka_instance.json"),
            file: "app-services-sdk-java/packages/kafka-instance-sdk/src/errors.ts"
        },
    },
    python: {
        kafkamgmt: {
            definition: require("./errors_kafka_mgmt.json"),
            file: "app-services-sdk-python/sdks/kafka_mgmt_sdk/errors.py"
        },
        registrymgmt:  {
            definition: require("./errors_srs_mgmt.json"),
            file: "app-services-sdk-python/sdks/registry_mgmt_sdk/errors.py"
        },
        connectormgmt: {
            definition: require("./errors_connector_mgmt.json"),
            file: "app-services-sdk-python/sdks/connector_mgmt_sdk/errors.py"
        }, 
        kafkainstance: {
            definition: require("./errors_kafka_instance.json"),
            file: "app-services-sdk-python/sdks/kafka_instance_sdk/errors.py"
        }, 
    }
}

