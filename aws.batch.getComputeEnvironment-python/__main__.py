import pulumi
import pulumi_aws as aws

batch_mongo = aws.batch.get_compute_environment(compute_environment_name="batch-mongo-production")

