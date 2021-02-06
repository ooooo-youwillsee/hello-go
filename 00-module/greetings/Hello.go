package hello

import (
	"errors"
)


func Hello(str string) (string,error) {
    if str == "" {
        return "", errors.New("error !!!")
    }
    println("hello " +  str)
    
    return str,nil
}
