package main

import "fmt"
import "http"
import "os"

func main() {
	s := "http://www.michaelbernstein.info/";
	r, _, err := http.Get(s);

	if err != nil {
		fmt.Printf("Error: %v\n", err);
		return;
	}
	var buf [4096]byte;
	var n int;
	for err != os.EOF {
		n, err = r.Body.Read(buf[0:len(buf)]);	
		fmt.Printf("%s", buf[0:n]);

		if err != nil && err != os.EOF {
			fmt.Printf("Error: %v\n", err);
			return;
		}

	}
	r.Body.Close();
}