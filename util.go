package main

func maybePanic(e error) {
	if e != nil {
		panic(e)
	}
}
