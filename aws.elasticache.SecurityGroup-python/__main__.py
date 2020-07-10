import pulumi
import pulumi_aws as aws

bar_security_group = aws.ec2.SecurityGroup("barSecurityGroup")
bar_elasticache_security_group_security_group = aws.elasticache.SecurityGroup("barElasticache/securityGroupSecurityGroup", security_group_names=[bar_security_group.name])

