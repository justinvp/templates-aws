import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.opsworks.Stack("main", {
    customJson: `{
 "foobar": {
    "version": "1.0.0"
  }
}
`,
    defaultInstanceProfileArn: aws_iam_instance_profile_opsworks.arn,
    region: "us-west-1",
    serviceRoleArn: aws_iam_role_opsworks.arn,
    tags: {
        Name: "foobar-stack",
    },
});

