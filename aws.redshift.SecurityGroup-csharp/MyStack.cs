using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.RedShift.SecurityGroup("default", new Aws.RedShift.SecurityGroupArgs
        {
            Ingress = 
            {
                new Aws.RedShift.Inputs.SecurityGroupIngressArgs
                {
                    Cidr = "10.0.0.0/24",
                },
            },
        });
    }

}

