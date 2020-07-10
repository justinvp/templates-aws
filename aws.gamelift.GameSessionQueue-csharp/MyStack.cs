using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.GameLift.GameSessionQueue("test", new Aws.GameLift.GameSessionQueueArgs
        {
            Destinations = 
            {
                aws_gamelift_fleet.Us_west_2_fleet.Arn,
                aws_gamelift_fleet.Eu_central_1_fleet.Arn,
            },
            PlayerLatencyPolicies = 
            {
                new Aws.GameLift.Inputs.GameSessionQueuePlayerLatencyPolicyArgs
                {
                    MaximumIndividualPlayerLatencyMilliseconds = 100,
                    PolicyDurationSeconds = 5,
                },
                new Aws.GameLift.Inputs.GameSessionQueuePlayerLatencyPolicyArgs
                {
                    MaximumIndividualPlayerLatencyMilliseconds = 200,
                },
            },
            TimeoutInSeconds = 60,
        });
    }

}

