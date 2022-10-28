package tui

import (
	"fmt"
	"github.com/metafates/mangal/anilist"
	"github.com/metafates/mangal/history"
	"github.com/metafates/mangal/icon"
	"github.com/metafates/mangal/provider"
	"github.com/metafates/mangal/source"
	"github.com/metafates/mangal/style"
	"strings"
)

type listItem struct {
	title       string
	description string
	internal    interface{}
	marked      bool
}

func (t *listItem) toggleMark() {
	t.marked = !t.marked
}

func (t *listItem) getMark() string {
	switch t.internal.(type) {
	case *source.Chapter:
		return icon.Get(icon.Mark)
	case *anilist.Manga:
		return icon.Get(icon.Link)
	case *provider.Provider:
		return icon.Get(icon.Search)
	default:
		return ""
	}
}

func (t *listItem) Title() (title string) {
	switch e := t.internal.(type) {
	case *source.Chapter:
		var sb = strings.Builder{}

		sb.WriteString(t.FilterValue())
		if e.Volume != "" {
			sb.WriteString(" ")
			sb.WriteString(style.Faint(e.Volume))
		}

		if e.IsDownloaded() {
			sb.WriteString(" ")
			sb.WriteString(icon.Get(icon.Downloaded))
		}

		title = sb.String()
	default:
		title = t.FilterValue()
	}

	if title != "" && t.marked {
		//title = fmt.Sprintf("%s %s", title, icon.Get(icon.Mark))
		title = fmt.Sprintf("%s %s", title, t.getMark())
	}

	return
}

func (t *listItem) Description() (description string) {
	switch e := t.internal.(type) {
	case *source.Chapter:
		description = e.URL
	case *source.Manga:
		description = e.URL
	case *history.SavedChapter:
		description = fmt.Sprintf("%s : %d / %d", e.Name, e.Index, e.MangaChaptersTotal)
	case *anilist.Manga:
		description = e.SiteURL
	default:
		description = t.description
	}

	return
}

func (t *listItem) FilterValue() string {
	switch e := t.internal.(type) {
	case *source.Chapter:
		return e.Name
	case *source.Manga:
		return e.Name
	case *history.SavedChapter:
		return e.MangaName
	case *anilist.Manga:
		return e.Name()
	default:
		return t.title
	}
}
