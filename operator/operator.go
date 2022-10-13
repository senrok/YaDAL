package operator

import (
	"github.com/senrok/yadal/interfaces"
	"github.com/senrok/yadal/layer"
	"github.com/senrok/yadal/object"
)

type Operator struct {
	accessor interfaces.Accessor
}

func (o *Operator) Object(path string) object.Object {
	return object.NewObject(o.accessor, path)
}

func (o *Operator) Layer(layer layer.Layer) *Operator {
	o.accessor = layer(o.accessor)
	return o
}

// FromAccessor returns the Operator from the interfaces.Accessor
func FromAccessor(acc interfaces.Accessor) Operator {
	return Operator{
		accessor: acc,
	}
}
