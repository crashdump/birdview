package sync

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/health"
	"github.com/crashdump/birdview/internal/model"
	"log"
	"regexp"
)

func (c *clients) GetHealthEvents() model.HealthEvents {
	input := &health.DescribeEventsForOrganizationInput{}

	// DescribeEventsForOrganizationOutput
	events, err := c.HealthClient.DescribeEventsForOrganization(context.TODO(), input)
	if err != nil {
		log.Fatalf("Error DescribeEventsForOrganization: %v\n", err)
		return nil
	}

	reAccount := regexp.MustCompile(`[0-9]{12}`)
	var results model.HealthEvents
	for _, event := range events.Events {
		results = append(results, model.HealthEvent{
			Name:        *event.EventTypeCode,
			Status:      string(event.StatusCode),
			Account:     reAccount.FindString(*event.Arn),
			Region:      *event.Region,
			Service:     *event.Service,
			StartTime:   *event.StartTime,
			UpdatedTime: *event.LastUpdatedTime,
		})
	}

	return results
}
