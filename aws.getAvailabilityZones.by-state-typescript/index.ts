import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const available = aws.getAvailabilityZones({
    state: "available",
});
const primary = new aws.ec2.Subnet("primary", {availabilityZone: available.then(available => available.names[0])});
// ...
const secondary = new aws.ec2.Subnet("secondary", {availabilityZone: available.then(available => available.names[1])});
// ...

