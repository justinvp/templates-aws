import pulumi
import pulumi_aws as aws

bar = aws.rds.Instance("bar",
    allocated_storage=10,
    backup_retention_period=0,
    engine="MySQL",
    engine_version="5.6.21",
    instance_class="db.t2.micro",
    maintenance_window="Fri:09:00-Fri:09:30",
    name="baz",
    parameter_group_name="default.mysql5.6",
    password="barbarbarbar",
    username="foo")
test = aws.rds.Snapshot("test",
    db_instance_identifier=bar.id,
    db_snapshot_identifier="testsnapshot1234")

