package generate

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

func Generate(numberNames int, numberSurname int, language string) {
	names, surnames, err := openFilesAndGetNames(language)

	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < numberNames; i++ {
		name := mountName(numberSurname, names, surnames)
		fmt.Println(name)
	}
}

func openFilesAndGetNames(language string) (names []string, surnames []string, err error) {
	_, erro := os.Stat("data/" + language)
	if os.IsNotExist(erro) {
		return nil, nil, errors.New("names in '" + strings.ToUpper(language) + "' does not exist")
	}

	file, err := ioutil.ReadFile("data/" + language + "/names")
	if err != nil {
		return nil, nil, errors.New("error to open names files")
	}
	names = strings.Split(string(file), "\n")

	file, err = ioutil.ReadFile("data/" + language + "/surnames")
	if err != nil {
		return nil, nil, errors.New("error to open surnames files")
	}
	surnames = strings.Split(string(file), "\n")

	return names, surnames, nil
}

func mountName(numberSurnames int, listNames []string, listSurnames []string) string {
	sizeListNames := len(listNames)
	sizeListSurnames := len(listSurnames)

	name := listNames[rand.Intn(sizeListNames-1)]

	for count := 0; count < numberSurnames; count++ {
		for {
			surname := listSurnames[rand.Intn(sizeListSurnames-1)]
			if !strings.Contains(name, surname) {
				name += " " + surname
				break
			}
		}
	}

	return name
}
