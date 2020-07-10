import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ni = new aws.sagemaker.NotebookInstance("ni", {
    instanceType: "ml.t2.medium",
    roleArn: aws_iam_role_role.arn,
    tags: {
        Name: "foo",
    },
});

