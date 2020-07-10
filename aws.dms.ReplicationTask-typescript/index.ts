import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new replication task
const test = new aws.dms.ReplicationTask("test", {
    cdcStartTime: "1.48434688e+09",
    migrationType: "full-load",
    replicationInstanceArn: aws_dms_replication_instance_test_dms_replication_instance_tf.replicationInstanceArn,
    replicationTaskId: "test-dms-replication-task-tf",
    replicationTaskSettings: "...",
    sourceEndpointArn: aws_dms_endpoint_test_dms_source_endpoint_tf.endpointArn,
    tableMappings: "{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"%\",\"table-name\":\"%\"},\"rule-action\":\"include\"}]}",
    tags: {
        Name: "test",
    },
    targetEndpointArn: aws_dms_endpoint_test_dms_target_endpoint_tf.endpointArn,
});

