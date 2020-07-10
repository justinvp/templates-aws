using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var defaultRouteTable = new Aws.Ec2.DefaultRouteTable("defaultRouteTable", new Aws.Ec2.DefaultRouteTableArgs
        {
            DefaultRouteTableId = aws_vpc.Foo.Default_route_table_id,
            Routes = 
            {
                ,
            },
            Tags = 
            {
                { "Name", "default table" },
            },
        });
    }

}

