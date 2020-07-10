import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const tftest = new aws.elasticbeanstalk.Application("tftest", {
    description: "tf-test-desc",
});
const tfTemplate = new aws.elasticbeanstalk.ConfigurationTemplate("tf_template", {
    application: tftest.name,
    solutionStackName: "64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4",
});

