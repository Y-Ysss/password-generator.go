package main

import (
	"fmt"
	"strconv"
	"crypto/rand"
)

const (
	LowerAlphabets = "abcdefghijklmnopqrstuvwxyz"
	UpperAlphabets = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers = "1234567890"
	Symbols = "~!@#$%^&*()_+-={}|[]:<>?,./"
	// Avoid Ambiguity
	LowerAvoidAmbiguityAlphs = "abcdefghijkmnpqrstuvwxyz" // Remove(l, o)
	UpperAvoidAmbiguityAlphs = "ABCDEFGHJKLMNPQRSTUVWXYZ" // Remove(I, O)
)

type GeneratePassword struct {
	PassLength int
	AvoidAmbiguity bool
	LowerAlphs bool
	UpperAlphs bool
	Numbers bool
	Symbols bool
}

func scanInt(defaultVal int, message string) int {
	var str string
	fmt.Printf("\x1b[36m%s(Default:%d)\x1b[0m [Number/Enter] > ", message, defaultVal)
	fmt.Scanln(&str)
	num, _ := strconv.Atoi(str)
	if num != defaultVal && num > 0 {
		return num
	}
	return defaultVal
}

func scanBool(defaultVal bool, message string) bool {
	var str string
	fmt.Printf("\x1b[36m%s(Default:%t)\x1b[0m [Y/N/Enter] ? ", message, defaultVal)
	fmt.Scanln(&str)
	if str == "y" || str == "Y" || str == "yes" || str == "Yes" {
		return true
	} else if str == "n" || str == "N" || str == "no" || str == "No" {
		return false
	} else {
		return defaultVal
	}
}

func (genPass *GeneratePassword) defaults() {
	genPass.AvoidAmbiguity = true
	genPass.LowerAlphs = true
	genPass.UpperAlphs = false
	genPass.Numbers = true
	genPass.Symbols = false
}

func (genPass *GeneratePassword) options() {
	genPass.AvoidAmbiguity = scanBool(false, "Avoid Ambiguity")
	genPass.LowerAlphs = scanBool(true, "Lowercase Alphabets")
	genPass.UpperAlphs = scanBool(false, "Uppercase Alphabets")
	genPass.Numbers = scanBool(true, "Numbers")
	genPass.Symbols = scanBool(false, "Symbols")
}

func (genPass *GeneratePassword) setCharacters() string {
	var chars []byte
	if genPass.AvoidAmbiguity {
		if genPass.LowerAlphs {
			chars = append(chars, LowerAvoidAmbiguityAlphs...)
		}
		if genPass.UpperAlphs {
			chars = append(chars, UpperAvoidAmbiguityAlphs...)
		}
	} else {
		if genPass.LowerAlphs {
			chars = append(chars, LowerAlphabets...)
		}
		if genPass.UpperAlphs {
			chars = append(chars, UpperAlphabets...)
		}
	}
	if genPass.Numbers {
		chars = append(chars, Numbers...)
	}
	if genPass.Symbols {
		chars = append(chars, Symbols...)
	}
	return string(chars)
}

func (genPass *GeneratePassword) generate() string {
	length := genPass.PassLength
	data := genPass.setCharacters()
	buf := make([]byte, length)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < length; {
		index := int(buf[i]) % len(data)
		if index < len(data) {
			buf[i] = data[index]
			i++
		} else {
			_, err := rand.Read(buf[i:i+1])
            if err != nil {
                fmt.Println(err)
            }
        }
	}
	return string(buf)
}

func main() {
	var genPass GeneratePassword
	length := scanInt(16, "Password Length")
	if length <= 4 {
		length = 4
	}
	genPass.PassLength = length
	if scanBool(false, "Show More Options") {
		fmt.Println("----More Options--------------------")
		genPass.options()
		fmt.Println("------------------------------------")
	} else {
		genPass.defaults()
	}
	password := genPass.generate()
	fmt.Println("\n\x1b[32mPassword Generated!\x1b[0m")
	fmt.Println(password)
}