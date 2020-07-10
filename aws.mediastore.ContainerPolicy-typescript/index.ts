import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentRegion = pulumi.output(aws.getRegion({ async: true }));
const currentCallerIdentity = pulumi.output(aws.getCallerIdentity({ async: true }));
const exampleContainer = new aws.mediastore.Container("example", {});
const exampleContainerPolicy = new aws.mediastore.ContainerPolicy("example", {
    containerName: exampleContainer.name,
    policy: pulumi.interpolate`{
	"Version": "2012-10-17",
	"Statement": [{
		"Sid": "MediaStoreFullAccess",
		"Action": [ "mediastore:*" ],
		"Principal": {"AWS" : "arn:aws:iam::${currentCallerIdentity.accountId}:root"},
		"Effect": "Allow",
		"Resource": "arn:aws:mediastore:${currentCallerIdentity.accountId}:${currentRegion.name!}:container/${exampleContainer.name}/*",
		"Condition": {
			"Bool": { "aws:SecureTransport": "true" }
		}
	}]
}
`,
});

