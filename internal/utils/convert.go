package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

func ConvertHclToMap(content []byte) (map[string]any, error) {
	parser := hclparse.NewParser()

	file, diags := parser.ParseHCL(content, "")
	if diags.HasErrors() {
		return nil, fmt.Errorf("error parsing HCL: %v", diags)
	}

	ctx := &hcl.EvalContext{}
	attrs, diags := file.Body.JustAttributes()
	if diags.HasErrors() {
		return nil, fmt.Errorf("error during evaluation of the HCL attributes: %v", diags)
	}

	result := make(map[string]any)
	for name, attr := range attrs {
		val, diags := attr.Expr.Value(ctx)
		if diags.HasErrors() {
			return nil, fmt.Errorf("error evaluating %s: %v", name, diags)
		}

		jsonBytes, err := ctyjson.Marshal(val, val.Type())
		if err != nil {
			log.Printf("JSON marshal error for %s: %v", name, err)
			result[name] = val.GoString()
			continue
		}

		var convertVal any
		err = json.Unmarshal(jsonBytes, &convertVal)
		if err != nil {
			result[name] = val.GoString()
			continue
		}

		result[name] = convertVal
	}

	return result, nil
}

func ConvertHclFileToMap(filename string) (map[string]any, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return ConvertHclToMap(content)
}

func ConvertMapToHcl(data map[string]any) ([]byte, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error converting to JSON: %v", err)
	}

	parser := hclparse.NewParser()
	jsonFile, diags := parser.ParseJSON(jsonBytes, "")
	if diags.HasErrors() {
		return nil, fmt.Errorf("error parsing JSON to HCL: %v", diags)
	}

	hclFile := hclwrite.NewEmptyFile()
	body := hclFile.Body()

	ctx := &hcl.EvalContext{}
	jsonAttrs, diags := jsonFile.Body.JustAttributes()
	if diags.HasErrors() {
		return nil, fmt.Errorf("error evaluating JSON attributes: %v", diags)
	}

	for name, attr := range jsonAttrs {
		val, diags := attr.Expr.Value(ctx)
		if diags.HasErrors() {
			return nil, fmt.Errorf("error evaluating %s: %v", name, diags)
		}
		body.SetAttributeValue(name, val)
	}

	return hclFile.Bytes(), nil
}

func ConvertMapToHclFile(data map[string]any, filename string) error {
	hclBytes, err := ConvertMapToHcl(data)
	if err != nil {
		return err
	}

	err = saveBytesToFile(hclBytes, filename)
	if err != nil {
		return err
	}

	return nil
}
