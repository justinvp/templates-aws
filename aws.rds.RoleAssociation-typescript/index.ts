import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.RoleAssociation("example", {
    dbInstanceIdentifier: aws_db_instance_example.id,
    featureName: "S3_INTEGRATION",
    roleArn: aws_iam_role_example.id,
});

