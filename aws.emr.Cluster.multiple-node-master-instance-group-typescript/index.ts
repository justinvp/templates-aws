import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Map public IP on launch must be enabled for public (Internet accessible) subnets
const exampleSubnet = new aws.ec2.Subnet("example", {
    mapPublicIpOnLaunch: true,
});
const exampleCluster = new aws.emr.Cluster("example", {
    // core_instance_group must be configured
    coreInstanceGroup: {},
    ec2Attributes: {
        subnetId: exampleSubnet.id,
    },
    masterInstanceGroup: {
        // Master instance count must be set to 3
        instanceCount: 3,
    },
    // EMR version must be 5.23.0 or later
    releaseLabel: "emr-5.24.1",
    // Termination protection is automatically enabled for multiple masters
    // To destroy the cluster, this must be configured to false and applied first
    terminationProtection: true,
});

