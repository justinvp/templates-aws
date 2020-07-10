import pulumi
import pulumi_aws as aws

selected = aws.s3.get_bucket(bucket="bucket.test.com")
test_zone = aws.route53.get_zone(name="test.com.")
example = aws.route53.Record("example",
    aliases=[{
        "name": selected.website_domain,
        "zone_id": selected.hosted_zone_id,
    }],
    name="bucket",
    type="A",
    zone_id=test_zone.id)

