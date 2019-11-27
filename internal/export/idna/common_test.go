// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package idna

// This file contains code that is common between the generation code and the
// package's test code.

import (
	"log"

	"github.com/liquid-dev/text/internal/ucd"
)

func catFromEntry(p *ucd.Parser) (cat category) {
	r := p.Rune(0)
	switch s := p.String(1); s {
	case "valid":
		cat = valid
	case "disallowed":
		cat = disallowed
	case "disallowed_STD3_valid":
		cat = disallowedSTD3Valid
	case "disallowed_STD3_mapped":
		cat = disallowedSTD3Mapped
	case "mapped":
		cat = mapped
	case "deviation":
		cat = deviation
	case "ignored":
		cat = ignored
	default:
		log.Fatalf("%U: Unknown category %q", r, s)
	}
	if s := p.String(3); s != "" {
		if cat != valid {
			log.Fatalf(`%U: %s defined for %q; want "valid"`, r, s, p.String(1))
		}
		switch s {
		case "NV8":
			cat = validNV8
		case "XV8":
			cat = validXV8
		default:
			log.Fatalf("%U: Unexpected exception %q", r, s)
		}
	}
	return cat
}

var joinType = map[string]info{
	"L": joiningL,
	"D": joiningD,
	"T": joiningT,
	"R": joiningR,
}
