import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.outposts.getOutposts({
    siteId: data.aws_outposts_site.id,
});

