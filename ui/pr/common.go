package pr

import (
	"context"

	"github.com/charmbracelet/log"

	ghctx "github.com/dlvhdr/gh-dash/ui/context"
	"github.com/dlvhdr/gh-dash/ui/keys"
	"github.com/dlvhdr/gh-dash/ui/theme"
)

type Common struct {
	Ctx           context.Context
	Width, Height int
	Styles        *ghctx.Styles
	Theme         *theme.Theme
	KeyMap        *keys.KeyMap
	Logger        *log.Logger
}

// NewCommon returns a new Common struct.
func NewCommon(ctx context.Context, theme theme.Theme, width, height int) Common {
	if ctx == nil {
		ctx = context.TODO()
	}
	styles := ghctx.InitStyles(theme)
	return Common{
		Ctx:    ctx,
		Width:  width,
		Height: height,
		Theme:  &theme,
		Styles: &styles,
		KeyMap: &keys.Keys,
		Logger: log.FromContext(ctx).WithPrefix("ui"),
	}
}
