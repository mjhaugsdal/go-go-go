package main

import (
	"context"
	"fmt"
	"main/openapi"
)

func main() {

	brreg, err := openapi.NewClientWithResponses("https://data.brreg.no/enhetsregisteret/api")
	if err != nil {
		panic(err)
	}

	roller, err := brreg.GetRollerWithResponse(context.Background(), "926796178")
	if err != nil {
		panic(err)
	}

	responseJson := roller.JSON200
	fmt.Println(responseJson.Links.Enhet.Href)

	for _, rollegruppe := range responseJson.Rollegrupper {
		for _, rolle := range rollegruppe.Roller {
			fmt.Println(rolle.Person.Navn.Fornavn)
			fmt.Println(rolle.Person.Navn.Etternavn)
		}
	}
}
