package main

import (
	"fmt"

	"github.com/erhaem/test-dev-backend/pets"
	"github.com/erhaem/test-dev-backend/utils"
)

func main() {
	// no. 1
	fmt.Println("No. 1 - create Esa's pet list")
	petsData := []pets.Pet{
		{Type: "Anjing", Breed: "Golden Retriever", Name: "Otto",
			Characteristic: []string{"Energik", "Senang bermain bola"}, Favorite: true},
		{Type: "Anjing", Breed: "Siberian Husky", Name: "Max",
			Characteristic: []string{"Bulu lebat", "Mata biru"}, Favorite: true},
		{Type: "Anjing", Breed: "Beagle", Name: "Bob",
			Characteristic: []string{"Ceria", "Aktif mengajak Esa bermain di taman"}, Favorite: false},
		{Type: "Kucing", Breed: "Persia", Name: "Luna",
			Characteristic: []string{"Anggun", "Manja"}, Favorite: true},
		{Type: "Kucing", Breed: "British Short Hair", Name: "Milo",
			Characteristic: []string{"Cerdas", "Aktif"}, Favorite: true},
		{Type: "Ikan", Breed: "Koi", Name: "Nana",
			Characteristic: []string{"Indah"}, Favorite: false},
		{Type: "Ikan", Breed: "Mas", Name: "Goldie",
			Characteristic: []string{"Berwarna cerah"}, Favorite: false},
	}	
	pets.PrintPets(petsData)
	fmt.Println("--------------")

	// no. 2
	rino := pets.Pet{
		Type:           "Badak",
		Breed:          "Jawa",
		Name:           "Rino",
		Characteristic: []string{"Pekerja keras"},
		Favorite:       true,
	}
	petsData = pets.AddPet(petsData, rino)

	fmt.Println("No. 2 - add Rino to Esa's pet list")
	pets.PrintPets(petsData)
	fmt.Println("--------------")

	// no. 3
	fmt.Println("No. 3 - get favorite pets ascending")
	pets.PrintPets(pets.GetFavoritesAsc(petsData))
	fmt.Println("--------------")

	fmt.Println("No. 3 - get favorite pets descending")
	pets.PrintPets(pets.GetFavoritesDesc(petsData))
	fmt.Println("--------------")

	// no. 4
	fmt.Println("No. 4 - replace Persian cat with Maine Coon")
	fmt.Println("before replaced: ")
	pets.PrintPets(petsData)
	petsData = pets.PersiaToMaineCoon(petsData)
	fmt.Println("after replaced: ")
	pets.PrintPets(petsData)
	fmt.Println("--------------")

	// no. 5
	fmt.Println("No. 5 - count pets by type")
	for t, c := range pets.CountPetsByType(petsData) {
		fmt.Printf("%s: %d\n", t, c)
	}
	fmt.Println("--------------")

	// no. 6
	fmt.Println("No. 6 - get palindrome pets")
	for name, length := range pets.GetPalindromePets(petsData) {
		fmt.Printf("name: %s, length: %d\n", name, length)
	}
	fmt.Println("--------------")

	// no. 7
	fmt.Println("No. 7 - sum even numbers from array")
	nums := []int{15, 18, 3, 9, 6, 2, 12, 14}
	sum, evens := utils.SumEven(nums)
	fmt.Printf("numbers in array: %v\n", nums)
	fmt.Printf("even numbers: %v\n", evens)
	fmt.Printf("total sum: %d\n", sum)
	fmt.Println("--------------")

	// no. 8
	fmt.Println("No. 8 - check if two strings anagram")
	fmt.Println(`"kamu" and "muka"? =>`, utils.IsAnagram("kamu", "muka"))
	fmt.Println(`"otto" and "toto"? =>`, utils.IsAnagram("otto", "toto"))
	fmt.Println(`"max" and "milo"? =>`, utils.IsAnagram("max", "milo"))
	fmt.Println("--------------")

	// no. 9
	fmt.Println("No. 9 - reformat JSON as expectation.json")
	utils.ReformatJSON("json/case.json", "json/result.json")
	fmt.Println("--------------")
}
