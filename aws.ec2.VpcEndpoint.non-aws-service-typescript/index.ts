import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ptfeServiceVpcEndpoint = new aws.ec2.VpcEndpoint("ptfe_service", {
    privateDnsEnabled: false,
    securityGroupIds: [aws_security_group_ptfe_service.id],
    serviceName: var_ptfe_service,
    subnetIds: [local_subnet_ids],
    vpcEndpointType: "Interface",
    vpcId: var_vpc_id,
});
const internal = pulumi.output(aws.route53.getZone({
    name: "vpc.internal.",
    privateZone: true,
    vpcId: var_vpc_id,
}, { async: true }));
const ptfeServiceRecord = new aws.route53.Record("ptfe_service", {
    name: pulumi.interpolate`ptfe.${internal.name!}`,
    records: [ptfeServiceVpcEndpoint.dnsEntries.apply(dnsEntries => (<any>dnsEntries[0])["dns_name"])],
    ttl: 300,
    type: "CNAME",
    zoneId: internal.zoneId!,
});

