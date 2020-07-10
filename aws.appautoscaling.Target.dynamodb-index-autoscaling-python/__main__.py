import pulumi
import pulumi_aws as aws

dynamodb_index_read_target = aws.appautoscaling.Target("dynamodbIndexReadTarget",
    max_capacity=100,
    min_capacity=5,
    resource_id=f"table/{aws_dynamodb_table['example']['name']}/index/{var['index_name']}",
    scalable_dimension="dynamodb:index:ReadCapacityUnits",
    service_namespace="dynamodb")

