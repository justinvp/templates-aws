using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new replication task
        var test = new Aws.Dms.ReplicationTask("test", new Aws.Dms.ReplicationTaskArgs
        {
            CdcStartTime = "1484346880",
            MigrationType = "full-load",
            ReplicationInstanceArn = aws_dms_replication_instance.Test_dms_replication_instance_tf.Replication_instance_arn,
            ReplicationTaskId = "test-dms-replication-task-tf",
            ReplicationTaskSettings = "...",
            SourceEndpointArn = aws_dms_endpoint.Test_dms_source_endpoint_tf.Endpoint_arn,
            TableMappings = "{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"%\",\"table-name\":\"%\"},\"rule-action\":\"include\"}]}",
            Tags = 
            {
                { "Name", "test" },
            },
            TargetEndpointArn = aws_dms_endpoint.Test_dms_target_endpoint_tf.Endpoint_arn,
        });
    }

}

