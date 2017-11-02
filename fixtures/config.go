package fixtures

import (
	"bytes"
	"encoding/base64"
	"fmt"
	testnet "github.com/cnupp/appssdk/testhelpers/net"
	"text/template"
)

func BuildAuth() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/v1/kv/cde/builder/token",
		Response: testnet.TestResponse{
			Status: 200,
			Body: Render(
				`[
				  {
					"CreateIndex": 100,
					"ModifyIndex": 200,
					"LockIndex": 200,
					"Key": "{{.Key}}",
					"Flags": 0,
					"Value": "{{.Value}}",
					"Session": "adf4238a-882b-9ddc-4a9d-5b6758e4159e"
				  }
				]`,
				map[string]string{"Key": "/v1/kv/cde/builder/token", "Value": string(Base64([]byte("token")))}),
		},
	}
}

func Render(temp string, data map[string]string) string {
	t, err := template.New("hooks").Parse(temp)
	if err != nil {
		panic(fmt.Errorf("Template is not valid"))
	}

	var rendered bytes.Buffer
	err = t.Execute(&rendered, data)
	if err != nil {
		panic(fmt.Errorf("Render fail"))
	}
	return rendered.String()
}

func Base64(input []byte) []byte {
	var encoded bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &encoded)
	encoder.Write(input)
	encoder.Close()
	return encoded.Bytes()
}
