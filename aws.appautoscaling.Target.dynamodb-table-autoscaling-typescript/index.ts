import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dynamodbTableReadTarget = new aws.appautoscaling.Target("dynamodb_table_read_target", {
    maxCapacity: 100,
    minCapacity: 5,
    resourceId: pulumi.interpolate`table/${aws_dynamodb_table_example.name}`,
    scalableDimension: "dynamodb:table:ReadCapacityUnits",
    serviceNamespace: "dynamodb",
});

