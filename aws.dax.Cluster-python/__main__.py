import pulumi
import pulumi_aws as aws

bar = aws.dax.Cluster("bar",
    cluster_name="cluster-example",
    iam_role_arn=data["aws_iam_role"]["example"]["arn"],
    node_type="dax.r4.large",
    replication_factor=1)

