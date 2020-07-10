using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.GameLift.Fleet("example", new Aws.GameLift.FleetArgs
        {
            BuildId = aws_gamelift_build.Example.Id,
            Ec2InstanceType = "t2.micro",
            FleetType = "ON_DEMAND",
            RuntimeConfiguration = new Aws.GameLift.Inputs.FleetRuntimeConfigurationArgs
            {
                ServerProcess = 
                {
                    
                    {
                        { "concurrentExecutions", 1 },
                        { "launchPath", "C:\\game\\GomokuServer.exe" },
                    },
                },
            },
        });
    }

}

