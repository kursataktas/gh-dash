package ui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/dlvhdr/gh-dash/data"
)

type PRModel struct {
}

func NewPRModel() PRModel {
	return PRModel{}
}

func (m PRModel) Init() tea.Cmd {
	return nil
}

func (m PRModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

var mockPr = data.PullRequestData{
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
				Body:      "In in ea id laborum nulla minim fugiat eiusmod voluptate nisi. Cupidatat enim sit anim excepteur magna dolor eu. Ea ipsum aute consequat laboris sint. Qui id irure aliqua aliqua cupidatat voluptate nisi incididunt dolor consectetur do cillum dolor adipisicing reprehenderit. Deserunt non Lorem voluptate quis cillum. Nulla consequat consequat Lorem aute consectetur ex sunt cillum fugiat veniam ea minim sit eu officia. Sit duis esse culpa ipsum enim dolore exercitation incididunt sunt officia anim esse.",
				UpdatedAt: time.Now().AddDate(0, 0, -2),
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
				UpdatedAt: time.Now().AddDate(0, 0, -2),
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
				Comments:     data.ReviewComments{},
			},
		},
	},
	IsDraft: false,
	Commits: data.Commits{},
}

func (m PRModel) View() string {
	content := strings.Builder{}
	content.WriteString(mockPr.Title + "\n")
	content.WriteString("2nd line\n")
	content.WriteString("3nd line\n")
	content.WriteString("4nd line\n")
	return content.String()
}
