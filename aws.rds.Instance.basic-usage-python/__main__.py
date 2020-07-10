import pulumi
import pulumi_aws as aws

default = aws.rds.Instance("default",
    allocated_storage=20,
    engine="mysql",
    engine_version="5.7",
    instance_class="db.t2.micro",
    name="mydb",
    parameter_group_name="default.mysql5.7",
    password="foobarbaz",
    storage_type="gp2",
    username="foo")

