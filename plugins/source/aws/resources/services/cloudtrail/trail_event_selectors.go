package cloudtrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudtrail/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func trailEventSelectors() *schema.Table {
	tableName := "aws_cloudtrail_trail_event_selectors"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_EventSelector.html`,
		Resolver:    fetchCloudtrailTrailEventSelectors,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudtrail"),
		Transform:   transformers.TransformWithStruct(&types.EventSelector{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "trail_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchCloudtrailTrailEventSelectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.CloudTrailWrapper)
	c := meta.(*client.Client)
	svc := c.Services().Cloudtrail
	response, err := svc.GetEventSelectors(ctx, &cloudtrail.GetEventSelectorsInput{TrailName: r.TrailARN}, func(options *cloudtrail.Options) {
		options.Region = *r.HomeRegion
	})
	if err != nil {
		return err
	}
	res <- response.EventSelectors
	return nil
}
