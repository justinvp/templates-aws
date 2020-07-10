using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var available = Output.Create(Aws.GetAvailabilityZones.InvokeAsync(new Aws.GetAvailabilityZonesArgs
        {
            State = "available",
        }));
        var primary = new Aws.Ec2.Subnet("primary", new Aws.Ec2.SubnetArgs
        {
            AvailabilityZone = available.Apply(available => available.Names[0]),
        });
        // ...
        var secondary = new Aws.Ec2.Subnet("secondary", new Aws.Ec2.SubnetArgs
        {
            AvailabilityZone = available.Apply(available => available.Names[1]),
        });
        // ...
    }

}

