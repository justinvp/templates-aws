import pulumi
import pulumi_aws as aws

example = aws.ecs.get_service(cluster_arn=data["aws_ecs_cluster"]["example"]["arn"],
    service_name="example")

