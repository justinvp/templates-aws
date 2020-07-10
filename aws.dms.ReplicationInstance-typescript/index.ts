import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dmsAssumeRole = pulumi.output(aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            identifiers: ["dms.amazonaws.com"],
            type: "Service",
        }],
    }],
}, { async: true }));
const dms_access_for_endpoint = new aws.iam.Role("dms-access-for-endpoint", {
    assumeRolePolicy: dmsAssumeRole.json,
});
const dms_access_for_endpoint_AmazonDMSRedshiftS3Role = new aws.iam.RolePolicyAttachment("dms-access-for-endpoint-AmazonDMSRedshiftS3Role", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role",
    role: dms_access_for_endpoint.name,
});
const dms_cloudwatch_logs_role = new aws.iam.Role("dms-cloudwatch-logs-role", {
    assumeRolePolicy: dmsAssumeRole.json,
});
const dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole = new aws.iam.RolePolicyAttachment("dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole",
    role: dms_cloudwatch_logs_role.name,
});
const dms_vpc_role = new aws.iam.Role("dms-vpc-role", {
    assumeRolePolicy: dmsAssumeRole.json,
});
const dms_vpc_role_AmazonDMSVPCManagementRole = new aws.iam.RolePolicyAttachment("dms-vpc-role-AmazonDMSVPCManagementRole", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole",
    role: dms_vpc_role.name,
});
// Create a new replication instance
const test = new aws.dms.ReplicationInstance("test", {
    allocatedStorage: 20,
    applyImmediately: true,
    autoMinorVersionUpgrade: true,
    availabilityZone: "us-west-2c",
    engineVersion: "3.1.4",
    kmsKeyArn: "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
    multiAz: false,
    preferredMaintenanceWindow: "sun:10:30-sun:14:30",
    publiclyAccessible: true,
    replicationInstanceClass: "dms.t2.micro",
    replicationInstanceId: "test-dms-replication-instance-tf",
    replicationSubnetGroupId: aws_dms_replication_subnet_group_test_dms_replication_subnet_group_tf.id,
    tags: {
        Name: "test",
    },
    vpcSecurityGroupIds: ["sg-12345678"],
});

