using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var available = Output.Create(Aws.GetAvailabilityZones.InvokeAsync());
        var currentRegion = Output.Create(Aws.GetRegion.InvokeAsync());
        var currentCallerIdentity = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
        var fooEip = new Aws.Ec2.Eip("fooEip", new Aws.Ec2.EipArgs
        {
            Vpc = true,
        });
        var fooProtection = new Aws.Shield.Protection("fooProtection", new Aws.Shield.ProtectionArgs
        {
            ResourceArn = Output.Tuple(currentRegion, currentCallerIdentity, fooEip.Id).Apply(values =>
            {
                var currentRegion = values.Item1;
                var currentCallerIdentity = values.Item2;
                var id = values.Item3;
                return $"arn:aws:ec2:{currentRegion.Name}:{currentCallerIdentity.AccountId}:eip-allocation/{id}";
            }),
        });
    }

}

