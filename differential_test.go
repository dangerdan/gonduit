package gonduit

import (
	"net/http"
	"testing"
	"time"

	"github.com/dangerdan/gonduit/core"
	"github.com/dangerdan/gonduit/entities"
	"github.com/dangerdan/gonduit/requests"
	"github.com/dangerdan/gonduit/responses"
	"github.com/dangerdan/gonduit/test/server"
	"github.com/dangerdan/gonduit/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDifferentialGetCommitPaths(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()

	s.RegisterMethod(DifferentialGetCommitPathsMethod, http.StatusOK, gin.H{
		"result": []string{
			"differential.go",
			"differential_test.go",
		},
	})

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.DifferentialGetCommitPathsRequest{RevisionID: 123}
	resp, err := c.DifferentialGetCommitPaths(req)
	assert.NoError(t, err)
	want := &responses.DifferentialGetCommitPathsResponse{
		"differential.go",
		"differential_test.go",
	}
	assert.Equal(t, want, resp)
}

func TestDifferentialGetCommitMessage(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()

	s.RegisterMethod(DifferentialGetCommitMessageMethod, http.StatusOK, gin.H{
		"result": "Commit description.",
	})

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.DifferentialGetCommitMessageRequest{
		RevisionID: 123,
	}
	resp, err := c.DifferentialGetCommitMessage(req)
	assert.NoError(t, err)
	want := responses.DifferentialGetCommitMessageResponse("Commit description.")
	assert.Equal(t, &want, resp)
}

func TestDifferentialQuery(t *testing.T) {

	tests := map[string]struct {
		apiResponse gin.H
		want        entities.DifferentialRevision
	}{
		"response_with_reviewers": {
			apiResponse: gin.H{
				"result": []gin.H{
					gin.H{
						"id": "123",
						"reviewers": map[string]string{
							"PHID-USER-123": "PHID-USER-123",
						},
					},
				},
			},
			want: entities.DifferentialRevision{
				ID: "123",
				Reviewers: entities.DifferentialRevisionReviewers{
					"PHID-USER-123": "PHID-USER-123",
				},
			},
		},
		// Phabricator returns empty slice instead of empty map when reviewers
		// do not exist. And a map when they do. This case should be handled
		// separately when unmarshaling JSON in Golang.
		"response_with_no_reviewers": {
			apiResponse: gin.H{
				"result": []gin.H{
					gin.H{
						"id":        "123",
						"reviewers": []string{},
					},
				},
			},
			want: entities.DifferentialRevision{
				ID: "123",
			},
		},
	}

	req := requests.DifferentialQueryRequest{
		IDs: []uint64{123},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := server.New()
			defer s.Close()
			s.RegisterCapabilities()

			s.RegisterMethod(
				DifferentialQueryMethod, http.StatusOK, test.apiResponse)
			c, err := Dial(s.GetURL(), &core.ClientOptions{
				APIToken: "some-token",
			})
			assert.Nil(t, err)
			resp, err := c.DifferentialQuery(req)
			assert.NoError(t, err)
			want := &responses.DifferentialQueryResponse{
				&test.want,
			}
			assert.Equal(t, want, resp)
		})
	}

}

func TestDifferentialDiffSearch(t *testing.T) {

	now := time.Now().In(time.Local)
	tests := map[string]struct {
		apiResponse gin.H
		want        responses.DifferentialDiffSearchResponse
	}{
		"response_with_reviewers": {
			apiResponse: gin.H{
				"result": gin.H{
					"data": []entities.DifferentialDiffResult{
						{
							ID: "123",
							Fields: entities.DifferentialDiffResultFields{
								DateCreated:  util.UnixTimestamp(now),
								DateModified: util.UnixTimestamp(now),
							},
						},
					},
				},
			},
			want: responses.DifferentialDiffSearchResponse{
				Data: []entities.DifferentialDiffResult{
					{
						ID: "123",
						Fields: entities.DifferentialDiffResultFields{
							DateCreated:  util.UnixTimestamp(now),
							DateModified: util.UnixTimestamp(now),
						},
					},
				},
			},
		},
	}

	req := requests.DifferentialDiffSearchRequest{
		Constraints: &requests.DifferentialDiffSearchConstraints{IDs: []int{123}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := server.New()
			defer s.Close()
			s.RegisterCapabilities()

			s.RegisterMethod(
				"differential.diff.search", http.StatusOK, test.apiResponse)
			c, err := Dial(s.GetURL(), &core.ClientOptions{
				APIToken: "some-token",
			})
			assert.Nil(t, err)
			resp, err := c.DifferentialDiffSearch(req)
			assert.NoError(t, err)
			want := &test.want
			assert.Equal(t, want, resp)
		})
	}

}
