package main

import (
	"fmt"
	"sort"
)

// groupAnagrams mengelompokkan kata-kata menjadi kelompok-kelompok anagram.
func groupAnagrams(words []string) [][]string {
	// anagramGroups adalah variable yang menyimpan kelompok-kelompok kata anagram.
	anagramGroups := [][]string{}

	// Melakukan Looping melalui setiap kata dalam input.
	for _, word := range words {
		wordAnagrams := false

		// Melakukan Looping melalui kelompok-kelompok yang sudah ada.
		for i, group := range anagramGroups {
			if isAnagram(word, group[0]) {
				// Tambahkan kata ke dalam kelompok yang sudah ada.
				anagramGroups[i] = append(group, word)
				wordAnagrams = true
				break
			}
		}
		// Jika kata tidak termasuk ke dalam kelompok yang sudah ada, buat kelompok baru.
		if !wordAnagrams {
			anagramGroups = append(anagramGroups, []string{word})
		}
	}
	return anagramGroups
}

// isAnagram memeriksa apakah dua kata adalah anagram.
func isAnagram(word1, word2 string) bool {
	// Konversi kata menjadi slice dari rune.
	slice1 := []rune(word1)
	slice2 := []rune(word2)

	// Urutkan slice rune.
	sort.Slice(slice1, func(i, j int) bool { return slice1[i] < slice1[j] })
	sort.Slice(slice2, func(i, j int) bool { return slice2[i] < slice2[j] })

	// Bandingkan slice rune yang sudah diurutkan.
	return string(slice1) == string(slice2)
}
func main() {
	inputArray := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	result := groupAnagrams(inputArray)
	fmt.Println(result)
}
