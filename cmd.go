package main

import (
	"errors"
	"fmt"
	"strings"
)

type recipeName []string

func (r *recipeName) String() string {
	return fmt.Sprint(*r)
}

func (r *recipeName) Set(value string) error {
	if len(*r) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dn := range strings.Split(value, ",") {
		*r = append(*r, dn)
	}
	return nil
}
