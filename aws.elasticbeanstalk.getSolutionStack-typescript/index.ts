import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const multiDocker = pulumi.output(aws.elasticbeanstalk.getSolutionStack({
    mostRecent: true,
    nameRegex: "^64bit Amazon Linux (.*) Multi-container Docker (.*)$",
}, { async: true }));

