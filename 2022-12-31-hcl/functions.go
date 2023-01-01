package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"
)

func NewEvalContext() *hcl.EvalContext {

	functions := map[string]function.Function{
		"upper":     stdlib.UpperFunc,
		"lower":     stdlib.LowerFunc,
		"min":       stdlib.MinFunc,
		"max":       stdlib.MaxFunc,
		"strlen":    stdlib.StrlenFunc,
		"substr":    stdlib.SubstrFunc,
		"csvdecode": CSVDecodeFunc,
	}

	variables := map[string]cty.Value{
		"description": cty.StringVal("description"),
	}

	return &hcl.EvalContext{
		Functions: functions,
		Variables: variables,
	}
}

var CSVDecodeFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "str",
			Type: cty.String,
		},
		{
			Name: "comma",
			Type: cty.String,
		},
	},
	Type: func(args []cty.Value) (cty.Type, error) {
		str := args[0]
		if !str.IsKnown() {
			return cty.DynamicPseudoType, nil
		}

		r := strings.NewReader(str.AsString())
		cr := csv.NewReader(r)
		if comma := args[1]; !comma.IsNull() {
			cr.Comma = bytes.Runes([]byte(comma.AsString()))[0]
		}
		headers, err := cr.Read()
		if err == io.EOF {
			return cty.DynamicPseudoType, fmt.Errorf("missing header line")
		}
		if err != nil {
			return cty.DynamicPseudoType, csvError(err)
		}

		atys := make(map[string]cty.Type, len(headers))
		for _, name := range headers {
			if _, exists := atys[name]; exists {
				return cty.DynamicPseudoType, fmt.Errorf("duplicate column name %q", name)
			}
			atys[name] = cty.String
		}
		return cty.List(cty.Object(atys)), nil
	},
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		ety := retType.ElementType()
		atys := ety.AttributeTypes()
		str := args[0]
		r := strings.NewReader(str.AsString())
		cr := csv.NewReader(r)
		cr.FieldsPerRecord = len(atys)
		if comma := args[1]; !comma.IsNull() {
			cr.Comma = bytes.Runes([]byte(comma.AsString()))[0]
		}

		// Read the header row first, since that'll tell us which indices
		// map to which attribute names.
		headers, err := cr.Read()
		if err != nil {
			return cty.DynamicVal, err
		}

		var rows []cty.Value
		for {
			cols, err := cr.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return cty.DynamicVal, csvError(err)
			}

			vals := make(map[string]cty.Value, len(cols))
			for i, str := range cols {
				name := headers[i]
				vals[name] = cty.StringVal(str)
			}
			rows = append(rows, cty.ObjectVal(vals))
		}

		if len(rows) == 0 {
			return cty.ListValEmpty(ety), nil
		}
		return cty.ListVal(rows), nil
	},
})

func csvError(err error) error {
	switch err := err.(type) {
	case *csv.ParseError:
		return fmt.Errorf("CSV parse error on line %d: %w", err.Line, err.Err)
	default:
		return err
	}
}
