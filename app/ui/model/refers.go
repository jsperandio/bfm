package model

type Refers struct {
	Items map[string]interface{}
}

func NewRefers() *Refers {
	return &Refers{
		Items: make(map[string]interface{}),
	}
}

func (r *Refers) Add(name string, widget interface{}) interface{} {
	r.Items[name] = widget
	return &widget
}

func (r *Refers) Get(name string) interface{} {
	if v, ok := r.Items[name]; ok {
		return v
	}
	return nil
}
