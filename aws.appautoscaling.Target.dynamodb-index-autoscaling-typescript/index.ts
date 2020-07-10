import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dynamodbIndexReadTarget = new aws.appautoscaling.Target("dynamodb_index_read_target", {
    maxCapacity: 100,
    minCapacity: 5,
    resourceId: pulumi.interpolate`table/${aws_dynamodb_table_example.name}/index/${var_index_name}`,
    scalableDimension: "dynamodb:index:ReadCapacityUnits",
    serviceNamespace: "dynamodb",
});

