# Go Replicate

[![Go Reference](https://pkg.go.dev/badge/github.com/sausheong/goreplicate.svg)](https://pkg.go.dev/github.com/sausheong/goreplicate)

This is a simple Go package for interacting with the Replicate (https://replicate.com) HTTP APIs. Replicate is an API service that allows developers to use machine learning models easily through calling APIs.

This package is based on the HTTP APIs -- https://replicate.com/docs/reference/http. Please read the APIs to understand how it works first.

## How to use

First, create a model. Go to https://replicate.com/explore to find the model you want to use. Then set up the input parameters for the model.

````go
model := NewModel("stability-ai", "stable-diffusion", version)
model.Input["prompt"] = "An astronaut riding a horse in photorealistic style"
model.Input["num_outputs"] = 4
````

Then create a client, passing into it the API authentication token (you can get it here once you registered an account -- https://replicate.com/account) and the model you just created.

Using the client you can send a create prediction call to the API to start the prediction.

````go
client := NewClient(auth, model)
err := client.Create()
if err != nil {
    // resolve the error here
}
````

The client should have a `Response`, with an ID after calling the `Create` function. This is the prediction ID that you can use to retrieve the results of the prediction. At this point in time, the `Response` struct will have nothing in the `Output`

To get the results of the prediction, you can use the same client, or you can create a new client and call the `Get` function. This will return you the results of the prediction in the `Output`.

````go
err := client.Get(predictId)
if err != nil {
    // resolve the error here
}    
````

