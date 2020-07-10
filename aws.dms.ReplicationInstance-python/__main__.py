import pulumi
import pulumi_aws as aws

dms_assume_role = aws.iam.get_policy_document(statements=[{
    "actions": ["sts:AssumeRole"],
    "principals": [{
        "identifiers": ["dms.amazonaws.com"],
        "type": "Service",
    }],
}])
dms_access_for_endpoint = aws.iam.Role("dms-access-for-endpoint", assume_role_policy=dms_assume_role.json)
dms_access_for_endpoint__amazon_dms_redshift_s3_role = aws.iam.RolePolicyAttachment("dms-access-for-endpoint-AmazonDMSRedshiftS3Role",
    policy_arn="arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role",
    role=dms_access_for_endpoint.name)
dms_cloudwatch_logs_role = aws.iam.Role("dms-cloudwatch-logs-role", assume_role_policy=dms_assume_role.json)
dms_cloudwatch_logs_role__amazon_dms_cloud_watch_logs_role = aws.iam.RolePolicyAttachment("dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole",
    policy_arn="arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole",
    role=dms_cloudwatch_logs_role.name)
dms_vpc_role = aws.iam.Role("dms-vpc-role", assume_role_policy=dms_assume_role.json)
dms_vpc_role__amazon_dmsvpc_management_role = aws.iam.RolePolicyAttachment("dms-vpc-role-AmazonDMSVPCManagementRole",
    policy_arn="arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole",
    role=dms_vpc_role.name)
# Create a new replication instance
test = aws.dms.ReplicationInstance("test",
    allocated_storage=20,
    apply_immediately=True,
    auto_minor_version_upgrade=True,
    availability_zone="us-west-2c",
    engine_version="3.1.4",
    kms_key_arn="arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
    multi_az=False,
    preferred_maintenance_window="sun:10:30-sun:14:30",
    publicly_accessible=True,
    replication_instance_class="dms.t2.micro",
    replication_instance_id="test-dms-replication-instance-tf",
    replication_subnet_group_id=aws_dms_replication_subnet_group["test-dms-replication-subnet-group-tf"]["id"],
    tags={
        "Name": "test",
    },
    vpc_security_group_ids=["sg-12345678"])

