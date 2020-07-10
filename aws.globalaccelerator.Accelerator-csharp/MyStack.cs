using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.GlobalAccelerator.Accelerator("example", new Aws.GlobalAccelerator.AcceleratorArgs
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
    }

}

