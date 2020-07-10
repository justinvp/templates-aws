import pulumi
import pulumi_aws as aws

task = aws.emr.InstanceGroup("task",
    cluster_id=aws_emr_cluster["tf-test-cluster"]["id"],
    instance_count=1,
    instance_type="m5.xlarge")

