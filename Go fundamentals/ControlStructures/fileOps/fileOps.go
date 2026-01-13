package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)


func Write_Float_to_file(value float64, fileName string) {
	valueText:= fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)

}

func Get_Float_from_file(fileName string)  (value float64, err error){ 
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Could not read from File")
	}
	balance_text := string(data)
	value, err = strconv.ParseFloat(balance_text, 64) //convert from float64 to string
	if err != nil {
		return 1000, errors.New("Could not read from File")
	}
	err = nil
	return 
}