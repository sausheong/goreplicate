package goreplicate

import (
	"fmt"
	"testing"
	"time"
)

var auth = "Token <token>"
var version = "a9758cbfbd5f3c2094457d996681af52552901775aa2d6dd0b17fd15df959bef"
var prompt1 = "An astronaut riding a horse in photorealistic style"
var prompt2 = "modern kids play area landscape architecture, water play area, floating kids, seating areas, perspective view, rainy weather, biopunk, cinematic photo, highly detailed, cinematic lighting, ultra-detailed, ultrarealistic, photorealism"
var prompt3 = "victorian rocking toy carousel theme park horse, overgrown, zdzisław beksiński, hr giger, mystical occult symbol in real life, high detail, green fog"
var prompt4 = "woman, beautiful, elegany, golden braided hair golden eyes, green and gold caftan, dress, smiling, fairy, shiny, realistic ,4k"

// TODO: improve on test cases
func TestClientCreateRequest(t *testing.T) {
	model := NewModel("stability-ai", "stable-diffusion", version)
	fmt.Printf("%#v", model)
	model.Input["prompt"] = prompt3
	model.Input["num_outputs"] = 1

	client := NewClient(auth, model)
	err := client.Create()
	if err != nil {
		t.Error(err)
	}

	// Only expect "starting" to be passing
	if client.Response.Status != "starting" {
		t.Errorf("expected status to be starting, got %s", client.Response.Status)
	}

	t.Logf("%#v", client.Response)
}

func TestGetRequest(t *testing.T) {
	predictId := "q3iywr5thzdsvkpscbhb6mddyu"
	fmt.Println("get request:", predictId)
	model := NewModel("stability-ai", "stable-diffusion", version)
	cl2 := NewClient(auth, model)
	t1 := time.Now()
	err := cl2.Get(predictId)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("time elapsed:", time.Since(t1))

	t.Logf("%#v\n", cl2.Response)
}

func TestClientCreateGetRequest(t *testing.T) {
	fmt.Println("Create request")
	model := NewModel("stability-ai", "stable-diffusion", version)
	fmt.Printf("%#v\n", model)
	model.Input["prompt"] = prompt4
	model.Input["num_outputs"] = 4

	cl1 := NewClient(auth, model)
	err := cl1.Create()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", cl1.Response)
	predictId := cl1.Response.ID

	fmt.Println("get request")
	cl2 := NewClient(auth, model)
	t1 := time.Now()
	err = cl2.Get(predictId)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("time elapsed:", time.Since(t1))

	t.Logf("%#v\n", cl2.Response)

}
