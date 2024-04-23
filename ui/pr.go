package ui

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/dlvhdr/gh-dash/data"
	"github.com/dlvhdr/gh-dash/mocks"
	"github.com/dlvhdr/gh-dash/ui/common"
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
	common   pr.Common
	viewport viewport.Model
}

func NewPRModel() PRModel {
	ctx := context.Background()
	c := pr.NewCommon(ctx, *theme.DefaultTheme, 80, 0)
	return PRModel{common: c, viewport: viewport.Model{}}
}

func (m PRModel) Init() tea.Cmd {
	return nil
}

func (m PRModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height
		m.viewport.SetContent(m.content())
	case tea.KeyMsg:
		switch {

		case msg.Type == tea.KeyCtrlC, msg.String() == "q":
			return m, tea.Quit

		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m PRModel) content() string {
	content := lipgloss.NewStyle().MarginLeft(3).MarginBottom(1).Render(m.headerView())

	activities := m.activitiesView()
	statuses := m.statusesView()
	body := lipgloss.JoinHorizontal(lipgloss.Top, activities, strings.Repeat(" ", m.viewport.Width-lipgloss.Width(activities)-lipgloss.Width(statuses)), statuses)

	content = lipgloss.JoinVertical(lipgloss.Left, content, body)
	return content
}

func (m PRModel) View() string {

	return m.viewport.View()
}

func (m *PRModel) headerView() string {
	content := ""
	s := m.common.Styles

	name := s.Common.FaintTextStyle.Render(mocks.Pr.Repository.NameWithOwner)
	title := lipgloss.JoinHorizontal(
		lipgloss.Left,
		s.Common.MainTextStyle.Render(mocks.Pr.Title),
		" ",
		s.Common.FaintTextStyle.Render(fmt.Sprintf("#%d", mocks.Pr.Number)),
	)
	content = lipgloss.JoinVertical(lipgloss.Left, content, name, title)

	state := s.PrSidebar.PillStyle.Copy().
		Background(s.Colors.OpenPR).
		Render(mocks.Pr.State)
	mergeable := s.PrSidebar.PillStyle.Copy().
		Background(s.Colors.MergedPR).
		Render(mocks.Pr.Mergeable)

	branch := s.Common.FaintTextStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Left,
		"󰘬 ",
		mocks.Pr.BaseRefName,
		"  ",
		mocks.Pr.HeadRefName,
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

type Activity interface {
	UpdatedAt() time.Time
	Icon() *string
	View(m *PRModel) string
}

type commentModel struct {
	data.Comment
}

func (c commentModel) UpdatedAt() time.Time {
	return c.Comment.UpdatedAt
}

func (c commentModel) Icon() *string {
	return nil
}

func (c commentModel) View(m *PRModel) string {
	return m.commentView(c.Comment)
}

type reviewModel struct {
	data.Review
	data.ReviewThread
}

func (r reviewModel) UpdatedAt() time.Time {
	return r.Review.UpdatedAt
}

func (r reviewModel) Icon() *string {
	icon := "󰈈"
	return &icon
}

func (r reviewModel) View(m *PRModel) string {
	s := m.common.Styles
	sc := s.Comment
	w := m.common.Width

	header := sc.Header.Copy().Width(w-1).Padding(0, 1).Render(
		fmt.Sprintf(
			"%s reviewed %s",
			r.Review.Author.Login,
			utils.TimeElapsed(r.Review.UpdatedAt),
		),
	)
	body := sc.Body.Width(w - 3).Render(r.Review.Body)
	body = lipgloss.JoinVertical(lipgloss.Left, body, "", m.reviewThread(r.ReviewThread))

	return lipgloss.JoinVertical(lipgloss.Left, header, body)
}

type reviewThreadModel struct {
	data.ReviewThread
}

func (rt reviewThreadModel) UpdatedAt() time.Time {
	if len(rt.ReviewThread.Comments.Nodes) == 0 {
		return time.Time{}
	}
	return rt.ReviewThread.Comments.Nodes[0].UpdatedAt
}

func (rt reviewThreadModel) View(m *PRModel) string {
	return m.reviewThread(rt.ReviewThread)
}

func (rt reviewThreadModel) Icon() *string {
	icon := "󰈈"
	return &icon
}

func (m *PRModel) activitiesView() string {
	nodes := make([]string, 0, len(mocks.Pr.Comments.Nodes)+len(mocks.Pr.LatestReviews.Nodes))
	sortedActivities := make([]Activity, 0, len(mocks.Pr.Comments.Nodes)+len(mocks.Pr.LatestReviews.Nodes))
	for _, c := range mocks.Pr.Comments.Nodes {
		sortedActivities = append(sortedActivities, commentModel{c})
	}
	for _, r := range mocks.Pr.LatestReviews.Nodes {
		thread := mocks.Pr.ReviewThreads.Nodes[0]
		sortedActivities = append(sortedActivities, reviewModel{r, thread})
	}
	sort.Slice(sortedActivities, func(i, j int) bool {
		return sortedActivities[i].UpdatedAt().After(sortedActivities[j].UpdatedAt())
	})

	for i, activity := range sortedActivities {
		view := activity.View(m)
		border := lipgloss.NormalBorder()

		vLine := ""
		if i < len(sortedActivities)-1 {
			vLine = lipgloss.JoinVertical(lipgloss.Left, strings.Split(strings.Repeat(border.Left, lipgloss.Height(view)), "")...)
		}

		tl := ""
		if activity.Icon() != nil {
			tl = lipgloss.NewStyle().Foreground(m.common.Theme.PrimaryText).Render(*activity.Icon() + " ")
		} else if i == 0 {
			tl = border.TopLeft
		} else if i == len(sortedActivities)-1 {
			tl = border.BottomLeft
		} else {
			tl = border.MiddleLeft
		}
		hLine := tl + lipgloss.NewStyle().Foreground(m.common.Theme.FaintBorder).Render(border.Top+border.Top)

		line := lipgloss.NewStyle().
			Foreground(m.common.Theme.FaintBorder).
			Render(
				lipgloss.JoinVertical(lipgloss.Left, hLine, vLine),
			)
		nodes = append(nodes, lipgloss.JoinHorizontal(lipgloss.Top, line, view))
	}

	return lipgloss.JoinVertical(lipgloss.Left, nodes...)
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
	s := m.common.Styles.StatusContext.Root.Copy().BorderTop(true).Bold(true)
	header := s.Render("󰝖 Checks")
	statuses := make([]string, 0)
	for _, commit := range mocks.Pr.Commits.Nodes {
		for i, context := range commit.Commit.StatusCheckRollup.Contexts.Nodes {
			status := m.statusView(context, i == len(commit.Commit.StatusCheckRollup.Contexts.Nodes)-1)
			statuses = append(statuses, status)
		}
	}

	rStatuses := lipgloss.JoinVertical(lipgloss.Left, statuses...)
	return lipgloss.JoinVertical(lipgloss.Left, header, rStatuses)
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
	threads := make([]string, 0, len(mocks.Pr.ReviewThreads.Nodes))
	for i, c := range mocks.Pr.ReviewThreads.Nodes {
		cView := m.reviewThread(c)
		border := lipgloss.NormalBorder()

		vLine := ""
		if i < len(mocks.Pr.Comments.Nodes)-1 {
			vLine = lipgloss.JoinVertical(lipgloss.Left, strings.Split(strings.Repeat(border.Left, lipgloss.Height(cView)), "")...)
		}

		tl := ""
		if i == 0 {
			tl = border.TopLeft
		} else if i == len(mocks.Pr.Comments.Nodes)-1 {
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

	header := sc.Header.Copy().UnsetBackground().Border(common.ThinBorder).BorderBottom(false).BorderForeground(
		sc.Body.GetBorderBottomForeground(),
	).Copy().Width(w-2).Padding(0, 1).Bold(true).Render(fmt.Sprintf(" %s", thread.Path))

	var comments []string
	for i, c := range thread.Comments.Nodes {
		comment := m.reviewComment(c)
		comments = append(comments, comment)
		if i < len(thread.Comments.Nodes)-1 {
			comments = append(comments, "")
		}
	}

	body := sc.Body.Width(w - 2).Render(lipgloss.JoinVertical(lipgloss.Left, comments...))
	r := lipgloss.JoinVertical(lipgloss.Left, header, body)

	return lipgloss.NewStyle().PaddingLeft(2).Render(r)
}

func (m *PRModel) reviewComment(comment data.ReviewComment) string {
	r := m.common.Styles.Common.MainTextStyle.Copy().Underline(true).Render(comment.Author.Login)
	r = lipgloss.JoinHorizontal(lipgloss.Top, r, " ", m.common.Styles.Common.FaintTextStyle.Render(utils.TimeElapsed(comment.UpdatedAt)))

	r = lipgloss.JoinVertical(lipgloss.Left, r, comment.Body)

	return r
}
