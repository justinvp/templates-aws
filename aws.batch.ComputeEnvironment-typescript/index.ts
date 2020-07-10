import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ecsInstanceRoleRole = new aws.iam.Role("ecs_instance_role", {
    assumeRolePolicy: `{
    "Version": "2012-10-17",
    "Statement": [
	{
	    "Action": "sts:AssumeRole",
	    "Effect": "Allow",
	    "Principal": {
		"Service": "ec2.amazonaws.com"
	    }
	}
    ]
}
`,
});
const ecsInstanceRoleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("ecs_instance_role", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role",
    role: ecsInstanceRoleRole.name,
});
const ecsInstanceRoleInstanceProfile = new aws.iam.InstanceProfile("ecs_instance_role", {
    role: ecsInstanceRoleRole.name,
});
const awsBatchServiceRoleRole = new aws.iam.Role("aws_batch_service_role", {
    assumeRolePolicy: `{
    "Version": "2012-10-17",
    "Statement": [
	{
	    "Action": "sts:AssumeRole",
	    "Effect": "Allow",
	    "Principal": {
		"Service": "batch.amazonaws.com"
	    }
	}
    ]
}
`,
});
const awsBatchServiceRoleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("aws_batch_service_role", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole",
    role: awsBatchServiceRoleRole.name,
});
const sampleSecurityGroup = new aws.ec2.SecurityGroup("sample", {
    egress: [{
        cidrBlocks: ["0.0.0.0/0"],
        fromPort: 0,
        protocol: "-1",
        toPort: 0,
    }],
});
const sampleVpc = new aws.ec2.Vpc("sample", {
    cidrBlock: "10.1.0.0/16",
});
const sampleSubnet = new aws.ec2.Subnet("sample", {
    cidrBlock: "10.1.1.0/24",
    vpcId: sampleVpc.id,
});
const sampleComputeEnvironment = new aws.batch.ComputeEnvironment("sample", {
    computeEnvironmentName: "sample",
    computeResources: {
        instanceRole: ecsInstanceRoleInstanceProfile.arn,
        instanceTypes: ["c4.large"],
        maxVcpus: 16,
        minVcpus: 0,
        securityGroupIds: [sampleSecurityGroup.id],
        subnets: [sampleSubnet.id],
        type: "EC2",
    },
    serviceRole: awsBatchServiceRoleRole.arn,
    type: "MANAGED",
}, { dependsOn: [awsBatchServiceRoleRolePolicyAttachment] });

