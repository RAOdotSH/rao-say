package raosay

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

/*
Function *init* :=
  - Initializes the random number generator.
  - Seed is set to the current time.
  - This is done to ensure that the random number generated is different every time.
*/
// func init() {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)

// 	fmt.Println(r.Intn(100))
// }

/*
	Error variables :=

- These error variables are predefined errors that are returned when the provided ID does not meet the specified conditions.
1. ErrIDLength - Error in Length of the ID.
2. ErrIDHex - Not a hexadecimal ID.
*/
var (
	ErrIDLength = fmt.Errorf("ID length should be exact 5")
	ErrIDHex    = fmt.Errorf("ID must be in hexadecimal i.e. [0-9a-f]")
)

/*
	Function *split* :=

- Splits the template in to 3 parts based on the new line character.
- Returns the top, middle, and bottom part of the template.
*/
func split(template string) (string, string, string) {
	t := strings.Split(template, "\n")

	if len(t) < 7 {
		return "", "", ""
	}

	top := strings.Join(t[0:3], "\n") + "\n"
	mid := strings.Join(t[3:5], "\n") + "\n"
	bottom := strings.Join(t[5:7], "\n")

	return top, mid, bottom
}

/*
	Function *replace* :=

- Replaces the eyes and mouth of the robot with the provided template.
- It replaces the template at the specified line number.
*/
func replace(body string, replace string, x, y int) string {
	s := strings.Split(body, "\n")

	if y >= len(s) || x+len(replace) > len(s[y]) {
		return body
	}

	s[y] = s[y][:x] + replace + s[y][x+len(replace):]

	return strings.Join(s, "\n")
}

/*
	Function *Generate* :=

- Checks if the ID length is 5 and contains only hexadecimal characters.
- It then constructs the robot by selecting templates for the body, eyes, and mouth based on the corresponding characters in the ID.
*/
func Generate(id string) (string, error) {
	if len(id) != 5 {
		fmt.Println("Generated ID: ", id)
		return "", ErrIDLength
	}

	match, _ := regexp.MatchString("^[0-9a-f]{5}$", id)
	if !match {
		fmt.Println("Generated ID: ", id)
		return "", ErrIDHex
	}

	art := ""

	// Body
	top, _, _ := split(template[id[0:1]])
	_, mid, _ := split(template[id[1:2]])
	_, _, bottom := split(template[id[2:3]])

	art += top
	art += mid
	art += bottom

	// replace eyes
	art = replace(art, eye[id[3:4]], 6, 1)

	// replace mouth
	art = replace(art, mouth[id[4:5]], 7, 2)

	return art, nil
}

/*
	Function *RandomID* :=

- Create a 5-character hex ID by generating 5 random numbers between 0 to 15.
- Converts these numbers in to hexadecimal digits.
*/
func RandomID() string {
	var id strings.Builder

	for i := 0; i < 5; i++ {
		id.WriteString(fmt.Sprintf("%x", rand.Intn(16)))
	}

	return id.String()
}

/*
	Function *Random* :=

- Generates a ascii bot with random ID.
- Ignores errors as they are handled earlier.
*/
func Random() string {
	r, _ := Generate(RandomID())

	return r
}

/*
	Function *CustomGenerate* :=

- Calls *Generate* function and panics if there is an error.
- It ensures that a valid robot is always generated or the program fails immediately.
*/
func CustomGenerate(id string) string {
	result, err := Generate(id)

	if err != nil {
		panic(err)
	}

	return result
}
