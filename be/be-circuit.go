// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
)

var SpiritAddress = NewVerbAddress("*", "Spirit")

// Required matter: Index, View, Circuit
func materializeCircuit(given Reflex, matter Circuit) interface{} {

	design := matter.CircuitAt("Circuit")

	// make links
	gates := make(map[Name]Reflex)
	gates[Super] = make(Reflex)
	for name, view := range design.Flow {
		if gates[name] == nil {
			gates[name] = make(Reflex)
		}
		for vlv, vec := range view {
			if gates[name][vlv] != nil {
				continue
			}
			if gates[vec.Gate] == nil {
				gates[vec.Gate] = make(Reflex)
			}
			gates[name][vlv], gates[vec.Gate][vec.Valve] = NewSynapse()
		}
	}

	// materialize gates
	residue := New()
	spirit := make(map[Name]interface{}) // channel to pass circuit residue back to spirit gates inside the circuit
	for g, _ := range design.Gate {
		if g == Super {
			Panicf("Circuit design overwrites the “%s” gate. In design %v\n", Super, design)
		}
		gsyntax := design.At(g)
		var gresidue interface{}

		// Compute view of gate within circuit
		view := New()
		for vlv, vec := range design.Flow[g] {
			view.Grow(vlv, design.Gate[vec.Gate])
		}

		if Same(gsyntax, SpiritAddress) {
			gresidue, spirit[g] = MaterializeInstance(gates[g], newSubMatterView(matter, view), &Future{})
		} else {
			gresidue = route(gsyntax, gates[g], newSubMatterView(matter, view))
		}
		residue.Gate[g] = gresidue
	}

	// connect given synapses
	for vlv, s := range given {
		t, ok := gates[Super][vlv]
		if !ok {
			continue
		}
		delete(gates[Super], vlv)
		go Link(s, t)
	}

	// send residue of this circuit to all escher.Spirit reflexes
	res := CleanUp(residue)
	go func() {
		for _, f := range spirit {
			f.(*Future).Charge(res)
		}
	}()

	if len(gates[Super]) > 0 {
		panic("circuit valves left unconnected")
	}

	return res
}

func newSubMatterView(matter Circuit, view Circuit) Circuit {
	r := newSubMatter(matter)
	r.Include("View", view)
	return r
}
