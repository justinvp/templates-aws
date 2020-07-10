import pulumi
import pulumi_aws as aws

default = aws.redshift.Cluster("default",
    cluster_identifier="tf-redshift-cluster",
    cluster_type="single-node",
    database_name="mydb",
    master_password="Mustbe8characters",
    master_username="foo",
    node_type="dc1.large")

