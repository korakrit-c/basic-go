package main

func main() {
   println("Hello World")

   printWorkAddress()

   printHomeAddress("Bangkok")

   getCountry := getCountry();
   println(getCountry)

   getAddress, getZipCode := getHomeAddress();
   println(getAddress)
   println(getZipCode)

}

func printWorkAddress() {
   println("Work from heart")
}

func printHomeAddress(homeAddress string) {
   println(homeAddress)
}

func getCountry() string {
   return "Thailand"
}

func getHomeAddress() (string, int){
   zipCode := 10100
   address := "Bangkok"
   return address, zipCode
}
