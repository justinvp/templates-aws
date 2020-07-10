import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cloudwatch = new aws.ses.EventDestination("cloudwatch", {
    cloudwatchDestinations: [{
        defaultValue: "default",
        dimensionName: "dimension",
        valueSource: "emailHeader",
    }],
    configurationSetName: aws_ses_configuration_set_example.name,
    enabled: true,
    matchingTypes: [
        "bounce",
        "send",
    ],
});

