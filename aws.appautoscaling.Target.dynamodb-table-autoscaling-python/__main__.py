import pulumi
import pulumi_aws as aws

dynamodb_table_read_target = aws.appautoscaling.Target("dynamodbTableReadTarget",
    max_capacity=100,
    min_capacity=5,
    resource_id=f"table/{aws_dynamodb_table['example']['name']}",
    scalable_dimension="dynamodb:table:ReadCapacityUnits",
    service_namespace="dynamodb")

