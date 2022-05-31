package model

import "github.com/rivo/tview"

type tviewWidget struct {
	widget interface{}
}

func (tvw *tviewWidget) AsPages() *tview.Pages {

	pg, ok := tvw.widget.(*tview.Pages)
	if ok {
		return pg
	}
	return nil
}

func (tvw *tviewWidget) AsList() *tview.List {

	l, ok := tvw.widget.(*tview.List)
	if ok {
		return l
	}
	return nil
}

type Refers struct {
	Items map[string]*tviewWidget
}

func NewRefers() *Refers {
	return &Refers{
		Items: make(map[string]*tviewWidget),
	}
}

func (r *Refers) Add(name string, widget interface{}) *tviewWidget {
	tvw := &tviewWidget{widget}
	r.Items[name] = tvw
	return tvw
}

func (r *Refers) Get(name string) *tviewWidget {

	if v, ok := r.Items[name]; ok {
		return v
	}
	return nil
}
