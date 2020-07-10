import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.mq.Broker("example", {
    brokerName: "example",
    configuration: {
        id: aws_mq_configuration_test.id,
        revision: aws_mq_configuration_test.latestRevision,
    },
    engineType: "ActiveMQ",
    engineVersion: "5.15.0",
    hostInstanceType: "mq.t2.micro",
    securityGroups: [aws_security_group_test.id],
    users: [{
        password: "MindTheGap",
        username: "ExampleUser",
    }],
});

