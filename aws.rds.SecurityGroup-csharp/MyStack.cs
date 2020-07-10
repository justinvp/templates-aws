using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Rds.SecurityGroup("default", new Aws.Rds.SecurityGroupArgs
        {
            Ingress = 
            {
                new Aws.Rds.Inputs.SecurityGroupIngressArgs
                {
                    Cidr = "10.0.0.0/24",
                },
            },
        });
    }

}

