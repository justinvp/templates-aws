import pulumi
import pulumi_aws as aws

# Create a new replication subnet group
test = aws.dms.ReplicationSubnetGroup("test",
    replication_subnet_group_description="Test replication subnet group",
    replication_subnet_group_id="test-dms-replication-subnet-group-tf",
    subnet_ids=["subnet-12345678"],
    tags={
        "Name": "test",
    })

