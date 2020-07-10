using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Iot.ThingType("foo", new Aws.Iot.ThingTypeArgs
        {
        });
    }

}

