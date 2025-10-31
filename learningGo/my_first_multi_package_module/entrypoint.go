package main

import (
	"oh_wow_a_multi_package_module/package1"
	"oh_wow_a_multi_package_module/package2"
)

func main() {
	package1.Echo1("Hi")
	package1.Echo2("Bye")
	package2.Echo1("Greetings")
	package2.Echo2("Farewell")
}
