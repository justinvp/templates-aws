using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new Lightsail Key Pair
        var lgKeyPair = new Aws.LightSail.KeyPair("lgKeyPair", new Aws.LightSail.KeyPairArgs
        {
        });
    }

}

