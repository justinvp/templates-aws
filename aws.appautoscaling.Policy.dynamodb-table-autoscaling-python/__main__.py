import pulumi
import pulumi_aws as aws

dynamodb_table_read_target = aws.appautoscaling.Target("dynamodbTableReadTarget",
    max_capacity=100,
    min_capacity=5,
    resource_id="table/tableName",
    scalable_dimension="dynamodb:table:ReadCapacityUnits",
    service_namespace="dynamodb")
dynamodb_table_read_policy = aws.appautoscaling.Policy("dynamodbTableReadPolicy",
    policy_type="TargetTrackingScaling",
    resource_id=dynamodb_table_read_target.resource_id,
    scalable_dimension=dynamodb_table_read_target.scalable_dimension,
    service_namespace=dynamodb_table_read_target.service_namespace,
    target_tracking_scaling_policy_configuration={
        "predefinedMetricSpecification": {
            "predefinedMetricType": "DynamoDBReadCapacityUtilization",
        },
        "targetValue": 70,
    })

