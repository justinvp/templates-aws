using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var currentRegion = Output.Create(Aws.GetRegion.InvokeAsync());
        var currentCallerIdentity = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
        var exampleContainer = new Aws.MediaStore.Container("exampleContainer", new Aws.MediaStore.ContainerArgs
        {
        });
        var exampleContainerPolicy = new Aws.MediaStore.ContainerPolicy("exampleContainerPolicy", new Aws.MediaStore.ContainerPolicyArgs
        {
            ContainerName = exampleContainer.Name,
            Policy = Output.Tuple(currentCallerIdentity, currentCallerIdentity, currentRegion, exampleContainer.Name).Apply(values =>
            {
                var currentCallerIdentity = values.Item1;
                var currentCallerIdentity1 = values.Item2;
                var currentRegion = values.Item3;
                var name = values.Item4;
                return @$"{{
	""Version"": ""2012-10-17"",
	""Statement"": [{{
		""Sid"": ""MediaStoreFullAccess"",
		""Action"": [ ""mediastore:*"" ],
		""Principal"": {{""AWS"" : ""arn:aws:iam::{currentCallerIdentity.AccountId}:root""}},
		""Effect"": ""Allow"",
		""Resource"": ""arn:aws:mediastore:{currentCallerIdentity1.AccountId}:{currentRegion.Name}:container/{name}/*"",
		""Condition"": {{
			""Bool"": {{ ""aws:SecureTransport"": ""true"" }}
		}}
	}}]
}}

";
            }),
        });
    }

}

