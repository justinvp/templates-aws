using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new GitLab Lightsail Instance
        var gitlabTest = new Aws.LightSail.Instance("gitlabTest", new Aws.LightSail.InstanceArgs
        {
            AvailabilityZone = "us-east-1b",
            BlueprintId = "string",
            BundleId = "string",
            KeyPairName = "some_key_name",
            Tags = 
            {
                { "foo", "bar" },
            },
        });
    }

}

