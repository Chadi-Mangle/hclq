package main

import (
	"fmt"
	"log"
)

func main() {
	result, _ := ConvertHclFileToMap("test.tfvars")
	path := "fargate.redirector-ng.alb_port"

	value, err := FindByPath(result, path)
	if err != nil {
		log.Fatal(err)
	}

	v, _ := value.GetHcl()
	fmt.Printf("%v\n", string(v))

	value.Set(8000)

	v2, _ := value.GetHcl()
	fmt.Printf("%v\n", string(v2))

	// hclBytes, err := ConvertMapToHcl(result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ConvertMapToHclFile(result, "test-copy.tfvars")
}
