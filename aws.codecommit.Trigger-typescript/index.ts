import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testRepository = new aws.codecommit.Repository("test", {
    repositoryName: "test",
});
const testTrigger = new aws.codecommit.Trigger("test", {
    repositoryName: testRepository.repositoryName,
    triggers: [{
        destinationArn: aws_sns_topic_test.arn,
        events: ["all"],
        name: "all",
    }],
});

