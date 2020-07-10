using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var vpc = config.RequireObject<dynamic>("vpc");
        var domain = config.Get("domain") ?? "tf-test";
        var selectedVpc = Output.Create(Aws.Ec2.GetVpc.InvokeAsync(new Aws.Ec2.GetVpcArgs
        {
            Tags = 
            {
                { "Name", vpc },
            },
        }));
        var selectedSubnetIds = selectedVpc.Apply(selectedVpc => Output.Create(Aws.Ec2.GetSubnetIds.InvokeAsync(new Aws.Ec2.GetSubnetIdsArgs
        {
            Tags = 
            {
                { "Tier", "private" },
            },
            VpcId = selectedVpc.Id,
        })));
        var currentRegion = Output.Create(Aws.GetRegion.InvokeAsync());
        var currentCallerIdentity = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
        var esSecurityGroup = new Aws.Ec2.SecurityGroup("esSecurityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Description = "Managed by Pulumi",
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    CidrBlocks = 
                    {
                        selectedVpc.Apply(selectedVpc => selectedVpc.CidrBlock),
                    },
                    FromPort = 443,
                    Protocol = "tcp",
                    ToPort = 443,
                },
            },
            VpcId = selectedVpc.Apply(selectedVpc => selectedVpc.Id),
        });
        var esServiceLinkedRole = new Aws.Iam.ServiceLinkedRole("esServiceLinkedRole", new Aws.Iam.ServiceLinkedRoleArgs
        {
            AwsServiceName = "es.amazonaws.com",
        });
        var esDomain = new Aws.ElasticSearch.Domain("esDomain", new Aws.ElasticSearch.DomainArgs
        {
            AccessPolicies = Output.Tuple(currentRegion, currentCallerIdentity).Apply(values =>
            {
                var currentRegion = values.Item1;
                var currentCallerIdentity = values.Item2;
                return @$"{{
	""Version"": ""2012-10-17"",
	""Statement"": [
		{{
			""Action"": ""es:*"",
			""Principal"": ""*"",
			""Effect"": ""Allow"",
			""Resource"": ""arn:aws:es:{currentRegion.Name}:{currentCallerIdentity.AccountId}:domain/{domain}/*""
		}}
	]
}}

";
            }),
            AdvancedOptions = 
            {
                { "rest.action.multi.allow_explicit_index", "true" },
            },
            ClusterConfig = new Aws.ElasticSearch.Inputs.DomainClusterConfigArgs
            {
                ClusterConfig = "m4.large.elasticsearch",
            },
            ElasticsearchVersion = "6.3",
            SnapshotOptions = new Aws.ElasticSearch.Inputs.DomainSnapshotOptionsArgs
            {
                SnapshotOptions = 23,
            },
            Tags = 
            {
                { "Domain", "TestDomain" },
            },
            VpcOptions = new Aws.ElasticSearch.Inputs.DomainVpcOptionsArgs
            {
                SecurityGroupIds = 
                {
                    esSecurityGroup.Id,
                },
                SubnetIds = 
                {
                    selectedSubnetIds.Apply(selectedSubnetIds => selectedSubnetIds.Ids[0]),
                    selectedSubnetIds.Apply(selectedSubnetIds => selectedSubnetIds.Ids[1]),
                },
            },
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_iam_service_linked_role.es",
            },
        });
    }

}

