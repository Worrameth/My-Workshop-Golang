package main

import "fmt"

func main() {
	status := map[int]string{
		200: "OK",
		308: "Permission denied",
		400: "Bad Request",
		500: "Internal Server Error",
		666: "",
	}
	fmt.Printf("%#v\n", status)

	l := len(status)
	fmt.Printf("length: %#v\n", l)

	status[200] = "Okay"
	status[404] = "File Not Found"
	fmt.Printf("%#v\n", status)
	fmt.Printf("length: %#v\n", len(status))

	value := status[200]
	fmt.Printf("%#v\n", value)

	delete(status, 404)
	fmt.Printf("%#v\n", status)
	fmt.Printf("length: %#v\n", len(status))

	//v, ok := status[666]
	if v, ok := status[666]; ok{
		fmt.Printf("v: %#v\n",v)
	} else{
		fmt.Println("Not Found")
	}

	//var m map[string]string = map[string]string{}
	//var m map[string]string = make(map[string]string)
	m := map[string]string{}
	if m == nil {
		fmt.Printf("m is nil. value: %#v\n",m)
	} else{
		fmt.Printf("m is not nil. value: %#v\n",m)
	}
}
