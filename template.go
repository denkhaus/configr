package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/denkhaus/configr/web"
	"github.com/juju/errors"
	htmlext "golang.org/x/net/html"
)

const (
	SettingsPageTemplate = "settingspage.tmpl"
)

type SettingsContext struct {
}

func AssembleSettingsPage(ctx SettingsContext) error {
	reader, err := web.LoadTemplate(SettingsPageTemplate)
	if err != nil {
		return errors.Annotate(err, "load template")
	}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return errors.Annotate(err, "new document")
	}

	acc := web.AccordionCtrl()
	pane1 := web.AccordionPane("General Settings")
	cont1 := web.AccordionContent()
	acc.AppendChild(pane1.AppendChild(cont1))

	section := doc.Find(".settings-section")
	section.AppendNodes((*htmlext.Node)(acc))

	html, err := doc.Html()
	if err != nil {
		return errors.Annotate(err, "get html")
	}

	return nil

}
