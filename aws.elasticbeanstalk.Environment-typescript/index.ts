import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const tftest = new aws.elasticbeanstalk.Application("tftest", {
    description: "tf-test-desc",
});
const tfenvtest = new aws.elasticbeanstalk.Environment("tfenvtest", {
    application: tftest.name,
    solutionStackName: "64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4",
});

