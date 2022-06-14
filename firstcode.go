package main

import "fmt"

func main() {
	var urname string
	fmt.Println("what is your name")
	fmt.Scan(&urname)
	if urname == "vamsi" {
		vamsi := user{"vamsi", "vamsi@gmail.com", 20, true}
		myage := 30
		mymail := "vamsi@gmail.com"
		hisage := &myage
		mail := &mymail

		names := []string{"vamsi"}
		nameage := make(map[string]int)
		nameage["vamsi"] = 20
		names = append(names, "vamsi", "vinay", "vivek")
		fmt.Println(names)
		fmt.Println(nameage)
		fmt.Println(vamsi)
		fmt.Println(*hisage)
		fmt.Println(*mail)

	} else {
		fmt.Println("not allowed kid are arrgsrf gone")

	}

}

type user struct {
	name   string
	mail   string
	age    int
	active bool
}
