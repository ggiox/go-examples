package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var codeSigningKey = []byte("ChaveSecretaParaValidarJWT")
var envSigningKey = []byte(os.Getenv("JWT_KEY_TOKEN"))

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(body))
}

// GenerateJWT retorna uma string JWT
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Fulano Beltrano"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	if len(envSigningKey) > 0 {
		tokenString, err := token.SignedString(envSigningKey)
		if err != nil {
			fmt.Printf("Ocorreu um erro usando a chave definida no ambiente: %s", err.Error())
			return "", err
		}
		fmt.Printf("Chave: %s \n", envSigningKey)
		fmt.Println("Token gerado usando chave definida no ambiente.")
		return tokenString, nil
	}

	tokenString, err := token.SignedString(codeSigningKey)
	if err != nil {
		fmt.Printf("Ocorreu um erro usando a chave definida no codigo: %s", err.Error())
		return "", err
	}
	fmt.Printf("Chave: %s \n", codeSigningKey)
	fmt.Println("Token gerado usando chave definida no codigo.")
	return tokenString, nil

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("Simple Client")

	// tokenString, err := GenerateJWT()
	// if err != nil {
	// 	fmt.Printf("Erro gerando string do token: %s", err.Error())
	// }

	// fmt.Println(tokenString)

	fmt.Println("Listener port 9001")
	handleRequests()
}
