package pets

import (
	"fmt"
	"sort"
	"strings"
)

type Pet struct {
	Type           string
	Breed          string
	Name           string
	Characteristic []string
	Favorite       bool
}

// print all pets in the list
func PrintPets(list []Pet) {
	for _, p := range list {
		fmt.Printf("{Type:%s Breed:%s Name:%s Characteristic:[%s] Favorite:%t}\n",
			p.Type, p.Breed, p.Name, strings.Join(p.Characteristic, ", "), p.Favorite)
	}
}


// add new pet
func AddPet(list []Pet, p Pet) []Pet {
	return append(list, p)
}

// get favorite pets ascending ( sort by name)
func GetFavoritesAsc(list []Pet) []Pet {
	var favs []Pet
	for _, p := range list {
		if p.Favorite {
			favs = append(favs, p)
		}
	}
	sort.Slice(favs, func(i, j int) bool {
		return favs[i].Name < favs[j].Name
	})
	return favs
}

// get favorite pets descending (sort by name)
func GetFavoritesDesc(list []Pet) []Pet {
	var favs []Pet
	for _, p := range list {
		if p.Favorite {
			favs = append(favs, p)
		}
	}
	sort.Slice(favs, func(i, j int) bool {
		return favs[i].Name > favs[j].Name
	})
	return favs
}

// count pets by its type
func CountPetsByType(list []Pet) map[string]int {
	counts := make(map[string]int)
	for _, p := range list {
		// expected: [Anjing: 3, Badak: 1]
		counts[p.Type]++
	}
	return counts
}

// replace Persia with Maine Coon
func PersiaToMaineCoon(list []Pet) []Pet {
	for i, p := range list {
		if p.Type == "Kucing" && p.Breed == "Persia" {
			list[i].Breed = "Maine Coon"
		}
	}
	return list
}

// ---- palindrome ----
// ref: https://github.com/mxssl/go-palindrome
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// check palindrome by reversing the word
// if the reversed word looks the same 
// like the original one
func isPalindrome(word string) bool {
	word = strings.ToLower(word)
	return word == reverse(word)
}

// get pets with palindrome names
func GetPalindromePets(list []Pet) map[string]int {
	results := make(map[string]int)
	for _, p := range list {
		if isPalindrome(p.Name) {
			// expected: [Bob: 3]
			results[p.Name] = len(p.Name)
		}
	}
	return results
}
