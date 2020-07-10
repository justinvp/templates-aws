import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dbInstance = pulumi.output(aws.getArn({
    arn: "arn:aws:rds:eu-west-1:123456789012:db:mysql-db",
}, { async: true }));

