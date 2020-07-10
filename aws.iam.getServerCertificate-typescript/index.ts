import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const my_domain = pulumi.output(aws.iam.getServerCertificate({
    latest: true,
    namePrefix: "my-domain.org",
}, { async: true }));
const elb = new aws.elb.LoadBalancer("elb", {
    listeners: [{
        instancePort: 8000,
        instanceProtocol: "https",
        lbPort: 443,
        lbProtocol: "https",
        sslCertificateId: my_domain.arn,
    }],
});

