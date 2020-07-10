using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleAccelerator = new Aws.GlobalAccelerator.Accelerator("exampleAccelerator", new Aws.GlobalAccelerator.AcceleratorArgs
        {
            Attributes = new Aws.GlobalAccelerator.Inputs.AcceleratorAttributesArgs
            {
                FlowLogsEnabled = true,
                FlowLogsS3Bucket = "example-bucket",
                FlowLogsS3Prefix = "flow-logs/",
            },
            Enabled = true,
            IpAddressType = "IPV4",
        });
        var exampleListener = new Aws.GlobalAccelerator.Listener("exampleListener", new Aws.GlobalAccelerator.ListenerArgs
        {
            AcceleratorArn = exampleAccelerator.Id,
            ClientAffinity = "SOURCE_IP",
            PortRanges = 
            {
                new Aws.GlobalAccelerator.Inputs.ListenerPortRangeArgs
                {
                    FromPort = 80,
                    ToPort = 80,
                },
            },
            Protocol = "TCP",
        });
    }

}

