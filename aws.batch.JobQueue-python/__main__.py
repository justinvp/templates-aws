import pulumi
import pulumi_aws as aws

test_queue = aws.batch.JobQueue("testQueue",
    compute_environments=[
        aws_batch_compute_environment["test_environment_1"]["arn"],
        aws_batch_compute_environment["test_environment_2"]["arn"],
    ],
    priority=1,
    state="ENABLED")

