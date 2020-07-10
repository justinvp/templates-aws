import pulumi
import pulumi_aws as aws

ecs_mongo = aws.ecs.get_cluster(cluster_name="ecs-mongo-production")

