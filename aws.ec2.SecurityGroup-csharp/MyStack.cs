using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var allowTls = new Aws.Ec2.SecurityGroup("allowTls", new Aws.Ec2.SecurityGroupArgs
        {
            Description = "Allow TLS inbound traffic",
            VpcId = aws_vpc.Main.Id,
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Description = "TLS from VPC",
                    FromPort = 443,
                    ToPort = 443,
                    Protocol = "tcp",
                    CidrBlocks = 
                    {
                        aws_vpc.Main.Cidr_block,
                    },
                },
            },
            Egress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupEgressArgs
                {
                    FromPort = 0,
                    ToPort = 0,
                    Protocol = "-1",
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },
                },
            },
            Tags = 
            {
                { "Name", "allow_tls" },
            },
        });
    }

}

