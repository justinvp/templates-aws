using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var barSecurityGroup = new Aws.Ec2.SecurityGroup("barSecurityGroup", new Aws.Ec2.SecurityGroupArgs
        {
        });
        var barElasticache_securityGroupSecurityGroup = new Aws.ElastiCache.SecurityGroup("barElasticache/securityGroupSecurityGroup", new Aws.ElastiCache.SecurityGroupArgs
        {
            SecurityGroupNames = 
            {
                barSecurityGroup.Name,
            },
        });
    }

}

