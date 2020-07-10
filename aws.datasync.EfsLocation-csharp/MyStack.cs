using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DataSync.EfsLocation("example", new Aws.DataSync.EfsLocationArgs
        {
            Ec2Config = new Aws.DataSync.Inputs.EfsLocationEc2ConfigArgs
            {
                SecurityGroupArns = 
                {
                    aws_security_group.Example.Arn,
                },
                SubnetArn = aws_subnet.Example.Arn,
            },
            EfsFileSystemArn = aws_efs_mount_target.Example.File_system_arn,
        });
    }

}

