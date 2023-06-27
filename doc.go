/*
The package epctds is a helper to to parse EPC hex messages into EPC Tag Data Standard objects

	Resources

	https://ref.gs1.org/standards/tds/

Example:

	func main() {
		tag, err := epctds.ParseEpcTagData("3118E511C46699F387000000")
		if err != nil {
			panic(err)
		}
		switch tag.(type) {
		case epctds.SSCC96:
			// do something
		}
	}
*/
package epctds
