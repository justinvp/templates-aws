package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.NewReplicationTask(ctx, "test", &dms.ReplicationTaskArgs{
			CdcStartTime:            pulumi.String("1484346880"),
			MigrationType:           pulumi.String("full-load"),
			ReplicationInstanceArn:  pulumi.Any(aws_dms_replication_instance.Test - dms - replication - instance - tf.Replication_instance_arn),
			ReplicationTaskId:       pulumi.String("test-dms-replication-task-tf"),
			ReplicationTaskSettings: pulumi.String("..."),
			SourceEndpointArn:       pulumi.Any(aws_dms_endpoint.Test - dms - source - endpoint - tf.Endpoint_arn),
			TableMappings:           pulumi.String(fmt.Sprintf("%v%v%v%v%v", "{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"", "%", "\",\"table-name\":\"", "%", "\"},\"rule-action\":\"include\"}]}")),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("test"),
			},
			TargetEndpointArn: pulumi.Any(aws_dms_endpoint.Test - dms - target - endpoint - tf.Endpoint_arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

