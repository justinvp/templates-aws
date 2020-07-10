using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.GlobalAccelerator.EndpointGroup("example", new Aws.GlobalAccelerator.EndpointGroupArgs
        {
            EndpointConfigurations = 
            {
                new Aws.GlobalAccelerator.Inputs.EndpointGroupEndpointConfigurationArgs
                {
                    EndpointId = aws_lb.Example.Arn,
                    Weight = 100,
                },
            },
            ListenerArn = aws_globalaccelerator_listener.Example.Id,
        });
    }

}

