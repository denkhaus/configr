package web

import (
	"github.com/denkhaus/html"
	"go.marzhillstudios.com/pkg/go-html-transform/h5"
)

func H3(txt string) *html.Node {
	return (*html.Node)(h5.Element("h3", nil, h5.Text(txt)))
}

func Li() *html.Node {
	return (*html.Node)(h5.Element("li", nil))
}

func Div(id string) *html.Node {
	return (*html.Node)(h5.Div(id, nil))
}

func AccordionCtrl() *html.Node {
	return Div("").AppendClass("accordion").
		SetAttr("wix-ctrl", "Accordion")
}

func AccordionPane(heading string) *html.Node {
	return Div("").AppendClass("acc-pane").AppendChild(H3(heading))
}

func AccordionContent() *html.Node {
	return Div("").AppendClass("acc-content")
}

func Label(txt string) *html.Node {
	return Li().SetAttr("wix-label", txt)
}

func ColorPicker(modelID string, options ItemOptions) *html.Node {
	return Div("").
		SetAttr("wix-model", modelID).
		SetAttr("wix-ctrl", "ColorPicker").
		SetAttr("wix-options", options.String())
}

func LabeledColorPicker(label string, modelID string, options ItemOptions) *html.Node {
	return Label(label).AppendChild(ColorPicker(modelID, options))
}

func Input(modelID string, options ItemOptions) *html.Node {
	return Div("").
		SetAttr("wix-model", modelID).
		SetAttr("wix-ctrl", "Input").
		SetAttr("wix-options", options.String())
}

func LabeledInput(label string, modelID string, options ItemOptions) *html.Node {
	return Label(label).AppendChild(Input(modelID, options))
}

func Spinner(modelID string, options ItemOptions) *html.Node {
	return Div("").
		SetAttr("wix-model", modelID).
		SetAttr("wix-ctrl", "Spinner").
		SetAttr("wix-options", options.String())
}

func LabeledSpinner(label string, modelID string, options ItemOptions) *html.Node {
	return Label(label).AppendChild(Spinner(modelID, options))
}
