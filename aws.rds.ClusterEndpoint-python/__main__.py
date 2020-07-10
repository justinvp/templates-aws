import pulumi
import pulumi_aws as aws

default = aws.rds.Cluster("default",
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    backup_retention_period=5,
    cluster_identifier="aurora-cluster-demo",
    database_name="mydb",
    master_password="bar",
    master_username="foo",
    preferred_backup_window="07:00-09:00")
test1 = aws.rds.ClusterInstance("test1",
    apply_immediately=True,
    cluster_identifier=default.id,
    identifier="test1",
    instance_class="db.t2.small")
test2 = aws.rds.ClusterInstance("test2",
    apply_immediately=True,
    cluster_identifier=default.id,
    identifier="test2",
    instance_class="db.t2.small")
test3 = aws.rds.ClusterInstance("test3",
    apply_immediately=True,
    cluster_identifier=default.id,
    identifier="test3",
    instance_class="db.t2.small")
eligible = aws.rds.ClusterEndpoint("eligible",
    cluster_endpoint_identifier="reader",
    cluster_identifier=default.id,
    custom_endpoint_type="READER",
    excluded_members=[
        test1.id,
        test2.id,
    ])
static = aws.rds.ClusterEndpoint("static",
    cluster_endpoint_identifier="static",
    cluster_identifier=default.id,
    custom_endpoint_type="READER",
    static_members=[
        test1.id,
        test3.id,
    ])

