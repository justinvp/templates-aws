import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const tftest = new aws.elasticbeanstalk.Application("tftest", {
    appversionLifecycle: {
        deleteSourceFromS3: true,
        maxCount: 128,
        serviceRole: aws_iam_role_beanstalk_service.arn,
    },
    description: "tf-test-desc",
});

