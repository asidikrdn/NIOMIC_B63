// You can edit this code!
// Click here and start typing.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// membuat struct
type mahasiswa struct {
	ID    string
	Nama  string
	Nilai float64
}

// membuat fungsi untuk mengambil semua data mahasiswa dari web API
// parameternya adalah url dari web API yang ingin diakses
func fetchUsers(urlServer string) {
	var err error
	var client = &http.Client{} // menghasilkan instance http.Client. Objek ini nantinya diperlukan untuk eksekusi request.
	var dataMhs []mahasiswa

	// membuat request baru
	// http.NewRequest(method, urlYangDiRequest, formDataRequest)
	request, err := http.NewRequest("GET", urlServer+"/users", nil) // request akan berisi intance object bertipe http.Request
	if err != nil {
		panic(err)
	}

	// melakukan eksekusi request
	response, err := client.Do(request) // response akan berisi intance object bertipe http.Response
	if err != nil {
		panic(err)
	}

	// menutup body response saat sudah tidak terpakai
	defer response.Body.Close()

	// mengambil body response dan menyimpannya ke variabel dataMhs
	err = json.NewDecoder(response.Body).Decode(&dataMhs)
	if err != nil {
		panic(err)
	}

	fmt.Println(dataMhs)
}

func fetchUser(urlServer string, idMhs string) (mahasiswa, error) {
	var err error
	var client = &http.Client{} // menghasilkan instance http.Client. Objek ini nantinya diperlukan untuk eksekusi request.
	var dataMhs mahasiswa

	// Membuat payload untuk parameter ketiga dari request
	parameter := url.Values{}                            // menghasilkan object yang nantinya dapat digunakan sebagai form data request
	parameter.Set("id", idMhs)                           // men-set isi dari object yang nantinya dapat digunakan sebagai form data request
	payload := bytes.NewBufferString(parameter.Encode()) // meng-encode form data request dan mengubahnya menjadi bentuk bytes.Buffer agar bisa disisipkan pada parameter ketiga dari http.NewRequest()

	// membuat request baru
	// http.NewRequest(method, urlYangDiRequest, formDataRequest)
	request, err := http.NewRequest("GET", urlServer+"/user?"+payload.String(), payload) // request akan berisi intance object bertipe http.Request
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded") // men-set tipe header pada request

	// melakukan eksekusi request
	response, err := client.Do(request) // response akan berisi intance object bertipe http.Response
	if err != nil {
		panic(err)
	}

	// menutup body response saat sudah tidak terpakai
	defer response.Body.Close()

	// mengambil body response dan menyimpannya ke variabel dataMhs
	err = json.NewDecoder(response.Body).Decode(&dataMhs)
	if err != nil {
		panic(err)
	}

	return dataMhs, nil
}

func main() {
	var baseUrl string = "http://localhost:4135"

	fetchUsers(baseUrl)

	fmt.Println()

	mhs, err := fetchUser(baseUrl, `U001`)
	if err != nil {
		panic(err)
	}

	fmt.Println(mhs)

}
