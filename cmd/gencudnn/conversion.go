package main

import (
	"fmt"

	bg "github.com/gorgonia/bindgen"
)

// Param represents a parameter in the signature
type Param struct {
	Name        string
	Type        string
	IsPtr       bool
	IsRetVal    bool // if it's a RetVal, the address is taken instead
	Convertible bool // there exists a conversion function from Go->C and vice versa
}

func MakeParam(name, typ string, isPtr bool) Param {
	return Param{
		Name:  name,
		Type:  typ,
		IsPtr: isPtr,
	}
}

// GoSignature represents a function signature in Go
type GoSignature struct {
	Name     string
	Receiver Param
	Params   []Param
	RetVals  []Param
	CSig     *bg.CSignature
}

func (s GoSignature) Format(f fmt.State, c rune) {
	fmt.Fprint(f, "func ")
	if s.Receiver.Name != "" {
		fmt.Fprintf(f, "(%v %v) ", s.Receiver.Name, s.Receiver.Type)
	}
	fmt.Fprintf(f, "%v(", s.Name)
	for _, p := range s.Params {
		fmt.Fprintf(f, "%v %v, ", p.Name, p.Type)
	}
	fmt.Fprint(f, ") ")
	if len(s.RetVals) > 0 {
		fmt.Fprint(f, "(")
		for _, r := range s.RetVals {
			fmt.Fprintf(f, "%v %v, ", r.Name, r.Type)
		}
		fmt.Fprint(f, ")")
	}
}

func csig2gosig(s *bg.CSignature) *GoSignature {
	retVal := &GoSignature{
		CSig: s,
	}

	// var receiver Param
	// var name string
	if isContextual(s.Name) {

	}
	return retVal
}