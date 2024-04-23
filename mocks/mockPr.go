package mocks

import (
	"time"

	graphql "github.com/cli/shurcooL-graphql"

	"github.com/dlvhdr/gh-dash/data"
)

var Pr = data.PullRequestData{
	Number: 13261,
	Title:  "Anim anim pariatur Lorem ea sint id aliquip",
	Body:   "Voluptate culpa in non incididunt fugiat amet. Incididunt quis nostrud et eiusmod commodo reprehenderit nisi do aliquip. Proident est culpa excepteur dolore mollit id. Id excepteur commodo esse pariatur do incididunt id laborum anim est nostrud ullamco qui nostrud. Mollit ipsum incididunt tempor proident ut reprehenderit sint pariatur id. Quis non deserunt magna duis deserunt id ea et laborum.",
	Author: struct{ Login string }{
		Login: "dlvhdr",
	},
	UpdatedAt:      time.Now(),
	Url:            "https://github.com/dlvhdr/gh-dash/pull/13261",
	State:          "OPEN",
	Mergeable:      "MERGEABLE",
	ReviewDecision: "APPROVED",
	Additions:      151,
	Deletions:      126,
	HeadRefName:    "dev",
	BaseRefName:    "main",
	HeadRepository: struct{ Name string }{
		Name: "gh-dash",
	},
	HeadRef: struct{ Name string }{
		Name: "dev",
	},
	Repository: data.Repository{
		Name:          "gh-dash",
		NameWithOwner: "dlvhdr/gh-dash",
		IsArchived:    false,
	},
	Assignees: data.Assignees{
		Nodes: []data.Assignee{
			{
				Login: "dlvhdr",
			},
		},
	},
	Comments: data.Comments{
		Nodes: []data.Comment{
			{
				Author: struct{ Login string }{
					Login: "dlvhdr",
				},
				Body: `https://github.com/komodorio/mono/assets/42379604/55b247f2-516f-4bff-9fef-922863756dc8

In in ea id laborum nulla minim fugiat eiusmod voluptate nisi. Cupidatat enim sit anim excepteur magna dolor eu. Ea ipsum aute consequat laboris sint. Qui id irure aliqua aliqua cupidatat voluptate nisi incididunt dolor consectetur do cillum dolor adipisicing reprehenderit. Deserunt non Lorem voluptate quis cillum. Nulla consequat consequat Lorem aute consectetur ex sunt cillum fugiat veniam ea minim sit eu officia. Sit duis esse culpa ipsum enim dolore exercitation incididunt sunt officia anim esse.`,
				UpdatedAt: time.Now().AddDate(0, 0, -1),
			},
			{
				Author: struct{ Login string }{
					Login: "tombenzera",
				},
				Body:      "Officia in veniam magna minim esse consectetur ea culpa cupidatat veniam non eiusmod velit velit elit. Adipisicing est dolore cillum esse sunt nulla excepteur veniam veniam do adipisicing in non et non.",
				UpdatedAt: time.Now().AddDate(0, 0, -1).Add(-time.Hour),
			},
			{
				Author: struct{ Login string }{
					Login: "caarlos0",
				},
				Body:      `Aliqua proident cupidatat in qui labore consectetur ea consectetur commodo minim tempor quis nulla sint id.`,
				UpdatedAt: time.Now().AddDate(0, 0, 0),
			},
		},
	},
	LatestReviews: data.Reviews{
		Nodes: []data.Review{
			{
				Author: struct{ Login string }{
					Login: "dlvhdr",
				},
				Body:      "Labore voluptate amet enim eu cupidatat irure commodo magna anim nisi eu do exercitation consequat ad. Consequat officia culpa consequat est magna irure est tempor duis. Nostrud dolor ex ex do. Sunt dolor commodo anim.",
				State:     "",
				UpdatedAt: time.Now().AddDate(0, 0, -1).Add(-30 * time.Minute),
				Comments: data.ReviewComments{
					Nodes: []data.ReviewComment{
						{
							Id: "1",
							Author: struct{ Login string }{
								Login: "kentcdodds",
							},
							Body:      "Eu ipsum laboris duis irure et laborum.",
							UpdatedAt: time.Now().AddDate(0, 0, -1),
							StartLine: 0,
							Line:      0,
						},
					},
				},
			},
		},
	},
	ReviewThreads: data.ReviewThreads{
		Nodes: []data.ReviewThread{
			{
				Id:           "1",
				IsOutdated:   false,
				OriginalLine: 1,
				StartLine:    1,
				Line:         1,
				Path:         "ui/pr.go",
				Comments: data.ReviewComments{
					Nodes: []data.ReviewComment{
						{
							Id: "1",
							Author: struct{ Login string }{
								Login: "kentcdodds",
							},
							Body:      "Eu ipsum laboris duis irure et laborum.",
							UpdatedAt: time.Now().AddDate(0, 0, -1),
							StartLine: 0,
							Line:      0,
						},
						{
							Id: "2",
							Author: struct{ Login string }{
								Login: "dlvhdr",
							},
							Body:      "Cupidatat non pariatur nulla do incididunt id sit deserunt minim anim. Proident mollit est ad. Laborum voluptate in et incididunt ipsum velit reprehenderit quis ut laborum esse labore aliqua. Irure mollit aliqua cupidatat proident magna aute id nostrud mollit.",
							UpdatedAt: time.Now().AddDate(0, 0, -1),
							StartLine: 0,
							Line:      0,
						},
					},
					TotalCount: 0,
				},
			},
		},
	},
	IsDraft: false,
	Commits: data.Commits{
		Nodes: []struct{ Commit data.Commit }{
			{Commit: data.Commit{
				StatusCheckRollup: data.StatusCheckRollup{
					Contexts: data.Contexts{
						TotalCount: 4,
						Nodes: []data.Context{{
							Typename: "CheckRun",
							CheckRun: data.CheckRun{
								Name:       "warden/mergeBlock",
								Status:     "COMPLETED",
								Conclusion: "SUCCESS",
								CheckSuite: data.CheckSuite{
									Creator:     struct{ Login graphql.String }{Login: "dlvhdr"},
									WorkflowRun: nil,
								},
								Text: "Successful in 3s — Merge away!",
							},
						},
							{
								Typename: "StatusContext",
								StatusContext: data.StatusContext{
									Context:     "buildkite/mono",
									State:       "FAILURE",
									Creator:     struct{ Login graphql.String }{Login: "buildkite"},
									Description: "Build #64276 failed (20 minutes, 4 seconds)",
								},
							},
						},
					},
				},
			}},
		},
	},
}
