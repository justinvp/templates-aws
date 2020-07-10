import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sns = new aws.ses.EventDestination("sns", {
    configurationSetName: aws_ses_configuration_set_example.name,
    enabled: true,
    matchingTypes: [
        "bounce",
        "send",
    ],
    snsDestination: {
        topicArn: aws_sns_topic_example.arn,
    },
});

