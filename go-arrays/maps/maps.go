package maps

import "fmt"

func main() {
	// map [KEY]VALUE
	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://www.aws.com",
		"Devyank Website":     "https://devyank.com.br",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["X"] = "https://www.x.com"
	fmt.Println(websites)

	delete(websites, "Amazon Web Services")
	fmt.Println(websites)
}
