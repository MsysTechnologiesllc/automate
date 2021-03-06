//
//  Author:: Salim Afiune <afiune@chef.io>, Gina Peers <gpeers@chef.io>
//  Copyright:: Copyright 2018, Chef Software Inc.
//

package integration_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"

	event "github.com/chef/automate/components/event-service/config"

	"github.com/chef/automate/api/interservice/event_feed"
	"github.com/chef/automate/components/event-feed-service/pkg/feed"
	"github.com/chef/automate/components/event-feed-service/pkg/persistence"
	"github.com/chef/automate/components/event-feed-service/pkg/server"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"

	"github.com/chef/automate/lib/grpc/grpctest"
)

var (
	millisPerSecond     = int64(1000)
	nanosPerMillisecond = int32(time.Millisecond)
)

func TestEventFeedReturnErrorWithWrongParameters(t *testing.T) {
	var (
		ctx  = context.Background()
		date = time.Now().UTC()
	)
	cases := []struct {
		description string
		request     event_feed.FeedRequest
	}{
		{
			description: "The Start date is after the End date",
			request: event_feed.FeedRequest{
				End:   date.AddDate(0, 0, -6).Unix() * 1000,
				Start: date.Unix() * 1000,
				Size:  10,
			},
		},
		{
			description: "Before and After parameters should not both be set",
			request: event_feed.FeedRequest{
				End:    date.Unix() * 1000,
				Start:  date.AddDate(0, 0, -6).Unix() * 1000,
				Size:   10,
				Before: date.AddDate(0, 0, -3).Unix() * 1000,
				After:  date.AddDate(0, 0, -1).Unix() * 1000,
			},
		},
		{
			description: "If the Before param is set the Cursor param needs to be set also",
			request: event_feed.FeedRequest{
				End:    date.Unix() * 1000,
				Start:  date.AddDate(0, 0, -6).Unix() * 1000,
				Size:   10,
				Before: date.AddDate(0, 0, -3).Unix() * 1000,
			},
		},
		{
			description: "When the After is set without the Cursor, After must be equal to End",
			request: event_feed.FeedRequest{
				End:   date.Unix() * 1000,
				Start: date.AddDate(0, 0, -6).Unix() * 1000,
				Size:  10,
				After: date.AddDate(0, 0, -3).Unix() * 1000,
			},
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("with parameters filters=%v it %s should return an error",
			test.request, test.description), func(t *testing.T) {

			_, err := testSuite.feedClient.GetFeed(ctx, &test.request)

			grpctest.AssertCode(t, codes.InvalidArgument, err)
		})
	}
}

func TestEventFeedReturnOnlyEventsWithinDateRange(t *testing.T) {
	var (
		ctx          = context.Background()
		totalEntries = 10
		pageSize     = int32(totalEntries)
		startDate    = time.Now().UTC()
		entries      = []*feed.FeedEntry{}
	)

	for i := 0; i < totalEntries; i++ {

		var (
			published = startDate.AddDate(0, 0, i*-1)
			created   = published
		)

		data := feed.FeedEntry{
			ID:                 uuid.Must(uuid.NewV4()).String(),
			ProducerID:         "urn:mycompany:user:violet",
			ProducerName:       "Violet",
			ProducerObjectType: "user",
			ProducerTags:       []string{"mycompany", "engineering department", "compliance team"},
			FeedType:           "event",
			EventType:          event.ScanJobUpdatedEventName,
			Tags:               []string{"mygroup", "compliance", "scan"},
			Published:          published,
			ActorID:            "urn:mycompany:user:violet",
			ActorObjectType:    "User",
			ActorName:          "Violet",
			Verb:               "update",
			ObjectID:           "urn:chef:compliance:scan-job",
			ObjectObjectType:   "ScanJob",
			ObjectName:         "Scan Job",
			TargetID:           "urn:mycompany:environment:production",
			TargetObjectType:   "Environment",
			TargetName:         "Production",
			Created:            created,
		}
		entries = append(entries, &data)
		// TODO: write bulk load
		testSuite.feedBackend.CreateFeedEntry(&data)
	}

	testSuite.RefreshIndices(persistence.IndexNameFeeds)
	defer GetTestSuite().DeleteAllDocuments()

	expectedEntries, err := server.FromInternalFormatToList(entries)
	if err != nil {
		log.Warnf("Couldn't set up expected result set... error converting feed entries from internal format to rpc; error: %v", err)
	}

	cases := []struct {
		description string
		request     event_feed.FeedRequest
		expected    []*event_feed.FeedEntry
	}{
		{
			description: "should return all 10 events (default)",
			request: event_feed.FeedRequest{
				Size: pageSize,
			},
			expected: expectedEntries,
		},
		{
			description: "should return all 10 events",
			request: event_feed.FeedRequest{
				Start: startDate.AddDate(0, 0, -10).Unix() * 1000,
				Size:  pageSize,
			},
			expected: expectedEntries,
		},
		{
			description: "should return first 5 events",
			request: event_feed.FeedRequest{
				Start: startDate.AddDate(0, 0, -5).Unix() * 1000,
				Size:  pageSize,
			},
			expected: expectedEntries[0:6],
		},
		{
			description: "should return last 5 events",
			request: event_feed.FeedRequest{
				End:  startDate.AddDate(0, 0, -5).Unix() * 1000,
				Size: pageSize,
			},
			expected: expectedEntries[5:10],
		},
		{
			description: "should return one event",
			request: event_feed.FeedRequest{
				Start: startDate.AddDate(0, 0, -5).Unix() * 1000,
				End:   startDate.AddDate(0, 0, -5).Unix() * 1000,
				Size:  pageSize,
			},
			expected: expectedEntries[5:6],
		},
		{
			description: "dates after should return zero events",
			request: event_feed.FeedRequest{
				Start: startDate.AddDate(0, 0, 1).Unix() * 1000,
				End:   startDate.AddDate(0, 0, 5).Unix() * 1000,
				Size:  pageSize,
			},
			expected: []*event_feed.FeedEntry{},
		},
		{
			description: "dates before should return zero events",
			request: event_feed.FeedRequest{
				Start: startDate.AddDate(0, 0, -11).Unix() * 1000,
				End:   startDate.AddDate(0, 0, -10).Unix() * 1000,
				Size:  pageSize,
			},
			expected: []*event_feed.FeedEntry{},
		},
	}

	// Run all the cases!
	for _, test := range cases {
		t.Run(fmt.Sprintf("with request '%v' it %s", test.request, test.description),
			func(t *testing.T) {
				res, err := testSuite.feedClient.GetFeed(ctx, &test.request)
				if assert.Nil(t, err) {
					assert.Equal(t, len(test.expected), len(res.FeedEntries))
					for index, expectedEvent := range test.expected {
						assert.Equal(t, expectedEvent.SourceEventPublished, res.FeedEntries[index].SourceEventPublished)
					}
				}
			})
	}
}

func TestEventFeedFilterEventType(t *testing.T) {
	var (
		ctx          = context.Background()
		totalEntries = 12
		pageSize     = int32(totalEntries)
		entries      = []*feed.FeedEntry{}
		eventTypes   = []string{"profile", "scanjobs"}
	)

	for i := 0; i < totalEntries; i++ {
		var (
			name             = "Fred"
			userURN          = "urn:mycompany:user:fred"
			eventType        = event.ScanJobUpdatedEventName
			tags             = []string{"org_1", "compliance", eventTypes[1]}
			verb             = "update"
			objectID         = "urn:chef:compliance:scan-job"
			objectObjectType = eventTypes[1]
			objectName       = "Scan Job"
		)

		if i > 5 {
			name = "Violet"
			userURN = "urn:mycompany:user:violet"
			eventType = event.ProfileCreatedEventName
			tags = []string{"org_2", "compliance", eventTypes[0]}
			verb = "create"
			objectID = "urn:chef:compliance:profile"
			objectObjectType = eventTypes[0]
			objectName = "Profile"
		}

		data := feed.FeedEntry{
			ID:                 uuid.Must(uuid.NewV4()).String(),
			ProducerID:         userURN,
			ProducerName:       name,
			ProducerObjectType: "user",
			ProducerTags:       []string{"mycompany", "engineering department", "compliance team"},
			FeedType:           "event",
			EventType:          eventType,
			Tags:               tags,
			Published:          time.Now().UTC(),
			ActorID:            userURN,
			ActorObjectType:    "User",
			ActorName:          name,
			Verb:               verb,
			ObjectID:           objectID,
			ObjectObjectType:   objectObjectType,
			ObjectName:         objectName,
			TargetID:           "urn:mycompany:environment:production",
			TargetObjectType:   "Environment",
			TargetName:         "Production",
			Created:            time.Now().UTC(),
		}

		entries = append(entries, &data)
		// TODO: write bulk load
		testSuite.feedBackend.CreateFeedEntry(&data)
	}

	testSuite.RefreshIndices(persistence.IndexNameFeeds)
	defer testSuite.DeleteAllDocuments()

	expectedEntries, err := server.FromInternalFormatToList(entries)
	if err != nil {
		log.Warnf("Couldn't set up expected result set... error converting feed entries from internal format to rpc; error: %v", err)
		expectedEntries = []*event_feed.FeedEntry{}
	}

	// reverse the list of expected entries to put them in desc order
	for i, j := 0, len(expectedEntries)-1; i < j; i, j = i+1, j-1 {
		expectedEntries[i], expectedEntries[j] = expectedEntries[j], expectedEntries[i]
	}

	cases := []struct {
		description string
		request     event_feed.FeedRequest
		expected    []*event_feed.FeedEntry
	}{
		{
			description: "should return all 12 events (default)",
			request: event_feed.FeedRequest{
				Size: pageSize,
			},
			expected: expectedEntries,
		},
		{
			description: "should return only 6 profile type events",
			request: event_feed.FeedRequest{
				Filters: []string{"event-type:" + eventTypes[0]},
				Size:    pageSize,
			},
			expected: expectedEntries[0:6],
		},
		{
			description: "should return only 6 scan job type events",
			request: event_feed.FeedRequest{
				Filters: []string{"event-type:" + eventTypes[1]},
				Size:    pageSize,
			},
			expected: expectedEntries[6:12],
		},
		{
			description: "filter requestor name of 'fred'",
			request: event_feed.FeedRequest{
				Filters: []string{"requestor_name:Fred"},
				Size:    pageSize,
			},
			expected: expectedEntries[6:12],
		},
		{
			description: "filter requestor name of 'Violet'",
			request: event_feed.FeedRequest{
				Filters: []string{"requestor_name:Violet"},
				Size:    pageSize,
			},
			expected: expectedEntries[0:6],
		},
	}

	// Run all the cases!
	for _, test := range cases {
		t.Run(fmt.Sprintf("with request '%v' it %s", test.request, test.description),
			func(t *testing.T) {
				res, err := testSuite.feedClient.GetFeed(ctx, &test.request)
				if assert.Nil(t, err) {
					assert.Equal(t, int64(len(test.expected)), res.TotalEntries)
					assert.ElementsMatch(t, test.expected, res.FeedEntries)
				}
			})
	}
}

func toMilliseconds(timestamp *google_protobuf.Timestamp) int64 {
	return timestamp.GetSeconds()*millisPerSecond +
		int64(timestamp.GetNanos()/nanosPerMillisecond)
}
