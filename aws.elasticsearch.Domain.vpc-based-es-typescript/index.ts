import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const vpc = config.require("vpc");
const domain = config.get("domain") || "tf-test";

const selectedVpc = pulumi.output(aws.ec2.getVpc({
    tags: {
        Name: vpc,
    },
}, { async: true }));
const selectedSubnetIds = selectedVpc.apply(selectedVpc => aws.ec2.getSubnetIds({
    tags: {
        Tier: "private",
    },
    vpcId: selectedVpc.id!,
}, { async: true }));
const currentRegion = pulumi.output(aws.getRegion({ async: true }));
const currentCallerIdentity = pulumi.output(aws.getCallerIdentity({ async: true }));
const esSecurityGroup = new aws.ec2.SecurityGroup("es", {
    description: "Managed by Pulumi",
    ingress: [{
        cidrBlocks: [selectedVpc.cidrBlock!],
        fromPort: 443,
        protocol: "tcp",
        toPort: 443,
    }],
    vpcId: selectedVpc.id!,
});
const esServiceLinkedRole = new aws.iam.ServiceLinkedRole("es", {
    awsServiceName: "es.amazonaws.com",
});
const esDomain = new aws.elasticsearch.Domain("es", {
    accessPolicies: pulumi.interpolate`{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Action": "es:*",
			"Principal": "*",
			"Effect": "Allow",
			"Resource": "arn:aws:es:${currentRegion.name!}:${currentCallerIdentity.accountId}:domain/${domain}/*"
		}
	]
}
`,
    advancedOptions: {
        "rest.action.multi.allow_explicit_index": "true",
    },
    clusterConfig: {
        instanceType: "m4.large.elasticsearch",
    },
    elasticsearchVersion: "6.3",
    snapshotOptions: {
        automatedSnapshotStartHour: 23,
    },
    tags: {
        Domain: "TestDomain",
    },
    vpcOptions: {
        securityGroupIds: [esSecurityGroup.id],
        subnetIds: [
            selectedSubnetIds.apply(selectedSubnetIds => selectedSubnetIds.ids[0]),
            selectedSubnetIds.apply(selectedSubnetIds => selectedSubnetIds.ids[1]),
        ],
    },
}, { dependsOn: [esServiceLinkedRole] });

