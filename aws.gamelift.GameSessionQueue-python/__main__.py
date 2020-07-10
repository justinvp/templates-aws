import pulumi
import pulumi_aws as aws

test = aws.gamelift.GameSessionQueue("test",
    destinations=[
        aws_gamelift_fleet["us_west_2_fleet"]["arn"],
        aws_gamelift_fleet["eu_central_1_fleet"]["arn"],
    ],
    player_latency_policies=[
        {
            "maximumIndividualPlayerLatencyMilliseconds": 100,
            "policyDurationSeconds": 5,
        },
        {
            "maximumIndividualPlayerLatencyMilliseconds": 200,
        },
    ],
    timeout_in_seconds=60)

