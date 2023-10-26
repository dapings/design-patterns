package prototype

type Obj1 struct {
	name string
}

func (o *Obj1) Clone() Prototype {
	oc := *o
	return &oc
}

type Obj2 struct {
	name string
}

func (o *Obj2) Clone() Prototype {
	return &Obj2{o.name}
}
