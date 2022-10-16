package data

import "testing"

func TestChecksValidate(t *testing.T){
	p:= &Product{
		Name: "Kenan",
		Price: 3,
		SKU:"123",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}