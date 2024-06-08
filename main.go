package main

import "bytes"

type Person struct {
	Name  string
	Phone string
}

func normalize(phone string) string {
	//will be good for processing large sets of data
	var buf bytes.Buffer
	for _, ch := range phone {
		//just scan through the phone number and if it is a digit, write it to the buffer
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	//gives us the string representation of the buffer written to buf
	return buf.String()
	//we want real digits only  (123) 456-7893 -> 1234567893

}

// using regular expressions to normalize the phone number
// func normalize(phone string) string {
// 	//Here "D" is a non-digit character "
// 	re := regexp.MustCompile("\\D")
// 	//replace all non-digit characters with an empty string
// 	return re.ReplaceAllString(phone, "")

// }
