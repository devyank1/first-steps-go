package arithmetic

import "fmt"

func arithmetic() {

	customers := 4

	packsToDeliver := 100

	alreadyDelivered := 20

	fmt.Println("Remaining packs to deliver:", packsToDeliver)

	remainingPacks := packsToDeliver - alreadyDelivered

	fmt.Println("Packs already delivered:", alreadyDelivered)

	fmt.Println("Remaining packs to deliver:", remainingPacks)

	packsByCustomers := (packsToDeliver - alreadyDelivered) / customers

	fmt.Println("Packs by Customers:", packsByCustomers)

}
