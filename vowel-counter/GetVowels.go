package main

func GetVowels(input string) string {

	var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

	var output = ""

	for i := 0; i < len(input); i++ {

		for j, _ := range vowels {

			if rune(input[i]) == vowels[j] {
				output = output + string(vowels[j])
				vowels = append(vowels[:j],vowels[j+1:] ...)
				break
			}
		}
	} 

	return output


}


