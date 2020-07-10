import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const elasticbeanstalk = new aws.iam.ServiceLinkedRole("elasticbeanstalk", {
    awsServiceName: "elasticbeanstalk.amazonaws.com",
});

