package pets

import (
	"reflect"
	"testing"
)

var samplePets = []Pet{
	{Type: "Anjing", Breed: "Golden Retriever", Name: "Otto", Characteristic: "Energik", Favorite: true},
	{Type: "Anjing", Breed: "Siberian Husky", Name: "Max", Characteristic: "Bulu lebat", Favorite: true},
	{Type: "Anjing", Breed: "Beagle", Name: "Bob", Characteristic: "Ceria", Favorite: false},
	{Type: "Kucing", Breed: "Persia", Name: "Luna", Characteristic: "Manja", Favorite: true},
	{Type: "Kucing", Breed: "British Short Hair", Name: "Milo", Characteristic: "Cerdas", Favorite: true},
}

func TestAddPet(t *testing.T) {
	list := samplePets
	newPet := Pet{Type: "Badak", Breed: "Jawa", Name: "Rino", Characteristic: "Pekerja keras", Favorite: true}
	list = AddPet(list, newPet)

	if list[len(list)-1].Name != "Rino" {
		t.Errorf("AddPet failed, want last pet = Rino, got %s", list[len(list)-1].Name)
	}
}

func TestGetFavoritesAsc(t *testing.T) {
	got := GetFavoritesAsc(samplePets)
	want := []string{"Luna", "Max", "Milo", "Otto"}

	for i, p := range got {
		if p.Name != want[i] {
			t.Errorf("GetFavoritesAsc wrong at index %d: got %s, want %s", i, p.Name, want[i])
		}
	}
}

func TestGetFavoritesDesc(t *testing.T) {
	got := GetFavoritesDesc(samplePets)
	want := []string{"Otto", "Milo", "Max", "Luna"}

	for i, p := range got {
		if p.Name != want[i] {
			t.Errorf("GetFavoritesDesc wrong at index %d: got %s, want %s", i, p.Name, want[i])
		}
	}
}

func TestCountPetsByType(t *testing.T) {
	got := CountPetsByType(samplePets)
	want := map[string]int{"Anjing": 3, "Kucing": 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("CountPetsByType = %v; want %v", got, want)
	}
}

func TestPersiaToMaineCoon(t *testing.T) {
	list := PersiaToMaineCoon(samplePets)
	found := false
	for _, p := range list {
		if p.Breed == "Maine Coon" {
			found = true
		}
	}
	if !found {
		t.Errorf("PersiaToMaineCoon failed, want at least one Maine Coon")
	}
}

func TestGetPalindromePets(t *testing.T) {
	got := GetPalindromePets(samplePets)
	want := map[string]int{"Otto": 4, "Bob": 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetPalindromePets = %v; want %v", got, want)
	}
}
