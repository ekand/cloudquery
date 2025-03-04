# Table: aws_route53_hosted_zones

This table shows data for Amazon Route 53 Hosted Zones.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_HostedZone.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_route53_hosted_zones:
  - [aws_route53_hosted_zone_query_logging_configs](aws_route53_hosted_zone_query_logging_configs.md)
  - [aws_route53_hosted_zone_resource_record_sets](aws_route53_hosted_zone_resource_record_sets.md)
  - [aws_route53_hosted_zone_traffic_policy_instances](aws_route53_hosted_zone_traffic_policy_instances.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|caller_reference|`utf8`|
|id|`utf8`|
|name|`utf8`|
|config|`json`|
|linked_service|`json`|
|resource_record_set_count|`int64`|
|tags|`json`|
|delegation_set_id|`utf8`|
|delegation_set|`json`|
|vpcs|`json`|