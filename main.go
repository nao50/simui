package main

import "fmt"

func main() {
	iso7816, err := newISO7816(0)
	if err != nil {
		log("error while creating iso7816 interface: %+v", err)
		return
	}

	imsi, err := iso7816.ReadIMSI()
	if err != nil {
		log("error while reading imsi on iso7816 interface: %+v", err)
		iso7816.Close()
		return
	}

	fmt.Println("IMSI:", imsi)
}
