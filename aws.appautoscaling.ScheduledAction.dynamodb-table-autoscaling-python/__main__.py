import pulumi
import pulumi_aws as aws

dynamodb_target = aws.appautoscaling.Target("dynamodbTarget",
    max_capacity=100,
    min_capacity=5,
    resource_id="table/tableName",
    scalable_dimension="dynamodb:table:ReadCapacityUnits",
    service_namespace="dynamodb")
dynamodb_scheduled_action = aws.appautoscaling.ScheduledAction("dynamodbScheduledAction",
    resource_id=dynamodb_target.resource_id,
    scalable_dimension=dynamodb_target.scalable_dimension,
    scalable_target_action={
        "max_capacity": 200,
        "min_capacity": 1,
    },
    schedule="at(2006-01-02T15:04:05)",
    service_namespace=dynamodb_target.service_namespace)

