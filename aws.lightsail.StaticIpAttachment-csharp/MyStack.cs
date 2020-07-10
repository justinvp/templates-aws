using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testStaticIp = new Aws.LightSail.StaticIp("testStaticIp", new Aws.LightSail.StaticIpArgs
        {
        });
        var testInstance = new Aws.LightSail.Instance("testInstance", new Aws.LightSail.InstanceArgs
        {
            AvailabilityZone = "us-east-1b",
            BlueprintId = "string",
            BundleId = "string",
            KeyPairName = "some_key_name",
        });
        var testStaticIpAttachment = new Aws.LightSail.StaticIpAttachment("testStaticIpAttachment", new Aws.LightSail.StaticIpAttachmentArgs
        {
            InstanceName = testInstance.Id,
            StaticIpName = testStaticIp.Id,
        });
    }

}

