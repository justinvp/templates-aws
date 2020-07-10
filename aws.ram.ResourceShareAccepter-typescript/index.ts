import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const alternate = new aws.Provider("alternate", {
    profile: "profile1",
});
const senderShare = new aws.ram.ResourceShare("sender_share", {
    allowExternalPrincipals: true,
    tags: {
        Name: "tf-test-resource-share",
    },
}, { provider: alternate });
const receiver = pulumi.output(aws.getCallerIdentity({ async: true }));
const senderInvite = new aws.ram.PrincipalAssociation("sender_invite", {
    principal: receiver.accountId,
    resourceShareArn: senderShare.arn,
}, { provider: alternate });
const receiverAccept = new aws.ram.ResourceShareAccepter("receiver_accept", {
    shareArn: senderInvite.resourceShareArn,
});

