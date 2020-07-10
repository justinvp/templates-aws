import pulumi
import pulumi_aws as aws

test_queue = aws.batch.get_job_queue(name="tf-test-batch-job-queue")

