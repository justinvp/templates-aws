import pulumi
import pulumi_aws as aws

current_region = aws.get_region()
current_caller_identity = aws.get_caller_identity()
example_container = aws.mediastore.Container("exampleContainer")
example_container_policy = aws.mediastore.ContainerPolicy("exampleContainerPolicy",
    container_name=example_container.name,
    policy=example_container.name.apply(lambda name: f"""{{
	"Version": "2012-10-17",
	"Statement": [{{
		"Sid": "MediaStoreFullAccess",
		"Action": [ "mediastore:*" ],
		"Principal": {{"AWS" : "arn:aws:iam::{current_caller_identity.account_id}:root"}},
		"Effect": "Allow",
		"Resource": "arn:aws:mediastore:{current_caller_identity.account_id}:{current_region.name}:container/{name}/*",
		"Condition": {{
			"Bool": {{ "aws:SecureTransport": "true" }}
		}}
	}}]
}}

"""))

