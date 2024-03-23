package ui

import (
	"context"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	graphql "github.com/cli/shurcooL-graphql"

	"github.com/dlvhdr/gh-dash/data"
	"github.com/dlvhdr/gh-dash/ui/pr"
	"github.com/dlvhdr/gh-dash/ui/theme"
	"github.com/dlvhdr/gh-dash/utils"
)

// Component represents a Bubble Tea model that implements a SetSize function.
type Component interface {
	tea.Model
	SetSize(width, height int)
}

type PRModel struct {
	common pr.Common
}

func NewPRModel() PRModel {
	ctx := context.Background()
	c := pr.NewCommon(ctx, *theme.DefaultTheme, 80, 0)
	return PRModel{common: c}
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
			{
				Author: struct{ Login string }{
					Login: "tombenzera",
				},
				Body:      "Officia in veniam magna minim esse consectetur ea culpa cupidatat veniam non eiusmod velit velit elit. Adipisicing est dolore cillum esse sunt nulla excepteur veniam veniam do adipisicing in non et non.",
				UpdatedAt: time.Now().AddDate(0, 0, -1),
			},
			{
				Author: struct{ Login string }{
					Login: "dmmulroy",
				},
				Body:      "Irure magna sint officia do. Officia in veniam magna minim esse consectetur ea culpa cupidatat veniam non eiusmod velit velit elit. Adipisicing est dolore cillum esse sunt nulla excepteur veniam veniam do adipisicing in non et non. Reprehenderit nostrud ipsum amet irure ad reprehenderit dolore irure amet ullamco labore qui. Et proident cillum cupidatat amet adipisicing enim minim ex consequat laborum. Officia veniam amet enim nostrud exercitation laborum minim ut quis dolor fugiat do.",
				UpdatedAt: time.Now().AddDate(0, 0, -1),
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
				Comments: data.ReviewComments{
					Nodes: []data.ReviewComment{
						{
							Author: struct{ Login string }{
								Login: "kentcdodds",
							},
							Body:      "Eu ipsum laboris duis irure et laborum.",
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

func (m PRModel) View() string {
	content := lipgloss.NewStyle().MarginLeft(3).MarginBottom(1).Render(m.headerView())

	body := lipgloss.JoinHorizontal(lipgloss.Top, m.commentsView(), " ", m.statusesView())

	content = lipgloss.JoinVertical(lipgloss.Left, content, body, m.reviewThreads())

	return content
}

func (m *PRModel) headerView() string {
	content := ""
	s := m.common.Styles

	name := s.Common.FaintTextStyle.Render(mockPr.Repository.NameWithOwner)
	title := lipgloss.JoinHorizontal(
		lipgloss.Left,
		s.Common.MainTextStyle.Render(mockPr.Title),
		" ",
		s.Common.FaintTextStyle.Render(fmt.Sprintf("#%d", mockPr.Number)),
	)
	content = lipgloss.JoinVertical(lipgloss.Left, content, name, title)

	state := s.PrSidebar.PillStyle.Copy().
		Background(s.Colors.OpenPR).
		Render(mockPr.State)
	mergeable := s.PrSidebar.PillStyle.Copy().
		Background(s.Colors.MergedPR).
		Render(mockPr.Mergeable)

	branch := s.Common.FaintTextStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Left,
		"󰘬 ",
		mockPr.BaseRefName,
		"  ",
		mockPr.HeadRefName,
	))

	pills := lipgloss.NewStyle().MarginTop(1).Render(lipgloss.JoinHorizontal(
		lipgloss.Top,
		state,
		" ",
		mergeable,
		"  ",
		branch,
	))
	return lipgloss.JoinVertical(lipgloss.Left, content, pills)
}

func (m *PRModel) commentsView() string {
	comments := make([]string, 0, len(mockPr.Comments.Nodes))
	for i, c := range mockPr.Comments.Nodes {
		cView := m.commentView(c)
		border := lipgloss.NormalBorder()

		vLine := ""
		if i < len(mockPr.Comments.Nodes)-1 {
			vLine = lipgloss.JoinVertical(lipgloss.Left, strings.Split(strings.Repeat(border.Left, lipgloss.Height(cView)), "")...)
		}

		tl := ""
		if i == 0 {
			tl = border.TopLeft
		} else if i == len(mockPr.Comments.Nodes)-1 {
			tl = border.BottomLeft
		} else {
			tl = border.MiddleLeft
		}
		hLine := tl + border.Top + border.Top

		line := lipgloss.NewStyle().
			Foreground(m.common.Theme.FaintBorder).
			Render(
				lipgloss.JoinVertical(lipgloss.Left, hLine, vLine),
			)
		comments = append(comments, lipgloss.JoinHorizontal(lipgloss.Top, line, cView))
	}

	return lipgloss.JoinVertical(lipgloss.Left, comments...)
}

func (m *PRModel) commentView(comment data.Comment) string {
	s := m.common.Styles
	sc := s.Comment
	w := m.common.Width

	author := sc.Header.Copy().PaddingRight(1).Render(s.Common.MainTextStyle.Render(comment.Author.Login))
	time := sc.Header.Render(
		s.Common.FaintTextStyle.Render(fmt.Sprintf("commented %s", utils.TimeElapsed(comment.UpdatedAt))),
	)

	header := sc.Header.Copy().Width(w).Padding(0, 1).Render(lipgloss.JoinHorizontal(lipgloss.Top, author, time))

	body := sc.Body.Width(w - 2).Render(comment.Body)

	content := lipgloss.JoinVertical(lipgloss.Left, header, body)

	return content
}

func (m *PRModel) statusesView() string {
	statuses := make([]string, 0)
	for _, commit := range mockPr.Commits.Nodes {
		for i, context := range commit.Commit.StatusCheckRollup.Contexts.Nodes {
			status := m.statusView(context, i == len(commit.Commit.StatusCheckRollup.Contexts.Nodes)-1)
			statuses = append(statuses, status)
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, statuses...)
}

func (m *PRModel) statusView(context data.Context, isLast bool) string {
	if context.Typename == "StatusContext" {
		return m.statusContext(context.StatusContext, isLast)
	} else {
		return m.checkRun(context.CheckRun, isLast)
	}
}

func (m *PRModel) statusContext(statusContext data.StatusContext, isLast bool) string {
	var glyph string
	if statusContext.State == "SUCCESS" {
		glyph = m.common.Styles.Common.SuccessGlyph
	} else {
		glyph = m.common.Styles.Common.FailureGlyph
	}

	status := lipgloss.NewStyle().Bold(true).Render(string(statusContext.Context))
	status = lipgloss.JoinVertical(lipgloss.Left, status, string(statusContext.Description))

	status = lipgloss.JoinHorizontal(lipgloss.Top, glyph, " ", status)
	return m.applyStatusBorder(status, isLast)
}

func (m *PRModel) checkRun(checkRun data.CheckRun, isLast bool) string {
	var glyph string
	if checkRun.Conclusion == "SUCCESS" {
		glyph = m.common.Styles.Common.SuccessGlyph
	} else {
		glyph = m.common.Styles.Common.FailureGlyph
	}

	status := lipgloss.NewStyle().Bold(true).Render(string(checkRun.Name))
	status = lipgloss.JoinVertical(lipgloss.Left, status, string(checkRun.Text))

	status = lipgloss.JoinHorizontal(lipgloss.Top, glyph, " ", status)
	return m.applyStatusBorder(status, isLast)
}

func (m *PRModel) applyStatusBorder(status string, isLast bool) string {
	s := m.common.Styles.StatusContext.Root.Copy()
	if isLast {
		s = s.BorderStyle(lipgloss.NormalBorder())
	}
	return s.Render(status)
}

func (m *PRModel) reviewThreads() string {
	threads := make([]string, 0, len(mockPr.ReviewThreads.Nodes))
	for i, c := range mockPr.ReviewThreads.Nodes {
		cView := m.reviewThread(c)
		border := lipgloss.NormalBorder()

		vLine := ""
		if i < len(mockPr.Comments.Nodes)-1 {
			vLine = lipgloss.JoinVertical(lipgloss.Left, strings.Split(strings.Repeat(border.Left, lipgloss.Height(cView)), "")...)
		}

		tl := ""
		if i == 0 {
			tl = border.TopLeft
		} else if i == len(mockPr.Comments.Nodes)-1 {
			tl = border.BottomLeft
		} else {
			tl = border.MiddleLeft
		}
		hLine := tl + border.Top + border.Top

		line := lipgloss.NewStyle().
			Foreground(m.common.Theme.FaintBorder).
			Render(
				lipgloss.JoinVertical(lipgloss.Left, hLine, vLine),
			)
		threads = append(threads, lipgloss.JoinHorizontal(lipgloss.Top, line, cView))
	}

	return lipgloss.JoinVertical(lipgloss.Left, threads...)
}

func (m *PRModel) reviewThread(thread data.ReviewThread) string {
	s := m.common.Styles
	sc := s.Comment
	w := m.common.Width

	author := sc.Header.Copy().PaddingRight(1).Render(s.Common.MainTextStyle.Render(thread.Path))
	header := sc.Header.Copy().Width(w).Padding(0, 1).Render(author)

	var comments []string
	for _, c := range thread.Comments.Nodes {
		comment := m.reviewComment(c)
		comments = append(comments, comment)
	}

	body := sc.Body.Width(w - 2).Render(lipgloss.JoinVertical(lipgloss.Left, comments...))
	r := lipgloss.JoinVertical(lipgloss.Left, header, body)

	return r
}

func (m *PRModel) reviewComment(comment data.ReviewComment) string {
	r := m.common.Styles.Common.MainTextStyle.Render(comment.Author.Login)
	r = lipgloss.JoinHorizontal(lipgloss.Top, r, " ", m.common.Styles.Common.FaintTextStyle.Render(utils.TimeElapsed(comment.UpdatedAt)))

	r = lipgloss.JoinVertical(lipgloss.Left, r, comment.Body)

	return r
}
