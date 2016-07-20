// Copyright 2016 Palm Stone Games, Inc. All rights reserved.

package elements

import (
	"code.palmstonegames.com/polymer"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var Router *RouterType

func InitRouter(defaultRoute string) {
	Router = &RouterType{Route: defaultRoute}

	template := polymer.GetDocument().GetElementByID("tmpl-base")
	template.(*polymer.AutoBindGoTemplate).Bind(Router)
}

type RouterType struct {
	*polymer.BindProto

	Route string `polymer:"bind"`
}

func (r *RouterType) SetRoute(route string) {
	r.Route = route
	r.Notify("route")
}

func (r *RouterType) ComputeIsActive(section string, active string) bool {
	return section == active
}

func (r *RouterType) HandleAnchorClick(e *polymer.Event) {
	e.PreventDefault()
	r.SetPath(r.getAnchorPath(e.Path))
}

func (r *RouterType) SetPath(path string) {
	if path != r.Route {
		r.Route = path
		r.Notify("route")
		js.Global.Get("history").Call("pushState", "", "", path) // TODO: Use the strongly typed polymer.GetWindow().History() API once that gets implemented
	}
}

func (r *RouterType) getAnchorPath(path []polymer.Element) string {
	var target polymer.Element
	for _, el := range path {
		if el.TagName() == "A" {
			target = el
			break
		}
	}

	if target == nil {
		panic("Could not find target")
	}

	return target.(*polymer.WrappedElement).UnwrappedElement.(*dom.HTMLAnchorElement).Pathname
}
