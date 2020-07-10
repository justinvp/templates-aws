import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dynamodbTarget = new aws.appautoscaling.Target("dynamodb", {
    maxCapacity: 100,
    minCapacity: 5,
    resourceId: "table/tableName",
    scalableDimension: "dynamodb:table:ReadCapacityUnits",
    serviceNamespace: "dynamodb",
});
const dynamodbScheduledAction = new aws.appautoscaling.ScheduledAction("dynamodb", {
    resourceId: dynamodbTarget.resourceId,
    scalableDimension: dynamodbTarget.scalableDimension,
    scalableTargetAction: {
        maxCapacity: 200,
        minCapacity: 1,
    },
    schedule: "at(2006-01-02T15:04:05)",
    serviceNamespace: dynamodbTarget.serviceNamespace,
});

