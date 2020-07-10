import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myInstance = new aws.opsworks.RdsDbInstance("my_instance", {
    dbPassword: "somePass",
    dbUser: "someUser",
    rdsDbInstanceArn: aws_db_instance_my_instance.arn,
    stackId: aws_opsworks_stack_my_stack.id,
});

