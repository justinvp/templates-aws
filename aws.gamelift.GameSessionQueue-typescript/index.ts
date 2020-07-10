import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.gamelift.GameSessionQueue("test", {
    destinations: [
        aws_gamelift_fleet_us_west_2_fleet.arn,
        aws_gamelift_fleet_eu_central_1_fleet.arn,
    ],
    playerLatencyPolicies: [
        {
            maximumIndividualPlayerLatencyMilliseconds: 100,
            policyDurationSeconds: 5,
        },
        {
            maximumIndividualPlayerLatencyMilliseconds: 200,
        },
    ],
    timeoutInSeconds: 60,
});

