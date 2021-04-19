package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	url := "https://www.bein.gg/"
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	checkError(err)
	t := time.Now()
	elapsed := t.Sub(start)
	html, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	fmt.Printf("%s\n", html)
	fmt.Printf("The Cookies for the website are %s \n",resp.Cookies())
	fmt.Printf("That Request took %s \n",elapsed)
	defer resp.Body.Close()

}




func checkError(err error){
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}