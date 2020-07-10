import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dynamodbTableReadTarget = new aws.appautoscaling.Target("dynamodb_table_read_target", {
    maxCapacity: 100,
    minCapacity: 5,
    resourceId: "table/tableName",
    scalableDimension: "dynamodb:table:ReadCapacityUnits",
    serviceNamespace: "dynamodb",
});
const dynamodbTableReadPolicy = new aws.appautoscaling.Policy("dynamodb_table_read_policy", {
    policyType: "TargetTrackingScaling",
    resourceId: dynamodbTableReadTarget.resourceId,
    scalableDimension: dynamodbTableReadTarget.scalableDimension,
    serviceNamespace: dynamodbTableReadTarget.serviceNamespace,
    targetTrackingScalingPolicyConfiguration: {
        predefinedMetricSpecification: {
            predefinedMetricType: "DynamoDBReadCapacityUtilization",
        },
        targetValue: 70,
    },
});

