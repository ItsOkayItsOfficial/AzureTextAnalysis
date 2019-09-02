package azuretextanalysis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Request struct for building clean interaction between analysis and request func
type Request struct {
	Key      string
	Endpoint string
	Type     string
	Text     []map[string]string
}

// Entities makes a request to the Entities API using the supplied:
//
// - API Key (Dashboard > Resources > {COGNATIVE SERVICES RESOURCE NAME} > Resource Management > Keys)
//
// - Resource Name ('http://{COGNATIVE SERVICES RESOURCE NAME}.cognitiveservices.azure.com')
//
// - Text to be analyzed for a positive/negative sentiment reading.
//
// The API returns a list of known entities and general named entities ("Person", "Location", "Organization" etc) in a given document. Known entities are returned with Wikipedia Id and Wikipedia link, and also Bing Id which can be used in Bing Entity Search API. General named entities are returned with entity types. If a general named entity is also a known entity, then all information regarding it (Wikipedia Id, Bing Id, entity type etc) will be returned.
//
// For more information see the Azure Cognative Service for Text Analytics Entities [documentation](https://westus.dev.cognitive.microsoft.com/docs/services/TextAnalytics-V2-1/operations/5ac4251d5b4ccd1554da7634).  See the [Supported Entity Types in Text Analytics API](https://docs.microsoft.com/en-us/azure/cognitive-services/text-analytics/how-tos/text-analytics-how-to-entity-linking#supported-types-for-named-entity-recognition) for the list of supported Entity Types. See the [Supported languages in Text Analytics API](https://docs.microsoft.com/en-us/azure/cognitive-services/text-analytics/text-analytics-supported-languages) for the list of enabled languages.
func Sentiment(apiKey string, resourceName string, document []map[string]string) string {

	// Define the API to make a call to
	var apiType = "entities"

	// Build a new Request struct with the inputs to pass into
	request := Request{apiKey, resourceName, apiType, document}
	output := apiRequest(request)
	return string(output)

}

// Phrases makes a request to the Key Phrase Extration API of Azure Cognative Service for Text Analytics using the supplied:
//
// - API Key (Dashboard > Resources > {COGNATIVE SERVICES RESOURCE NAME} > Resource Management > Keys)
//
// - Resource Name ('http://{COGNATIVE SERVICES RESOURCE NAME}.cognitiveservices.azure.com')
//
// - Text to be analyzed to extract key phrases within.
func Phrases(apiKey string, resourceName string, document []map[string]string) string {

	// Define the API to make a call to
	var apiType = "keyPhrases"

	// Build a new Request struct with the inputs to pass into
	request := Request{apiKey, resourceName, apiType, document}
	output := apiRequest(request)
	return string(output)

}

// Language makes a request to the Language Detection API of Azure Cognative Service for Text Analytics using the supplied:
//
// - API Key (Dashboard > Resources > {COGNATIVE SERVICES RESOURCE NAME} > Resource Management > Keys)
//
// - Resource Name ('http://{COGNATIVE SERVICES RESOURCE NAME}.cognitiveservices.azure.com')
//
// - Text to be analyzed to detect the language it's written in.
func Language(apiKey string, resourceName string, document []map[string]string) string {

	// Define the API to make a call to
	var apiType = "languages"

	// Build a new Request struct with the inputs to pass into
	request := Request{apiKey, resourceName, apiType, document}
	output := apiRequest(request)
	return string(output)

}

// Sentiment makes a request to the Sentiment Analysis API of the Azure Cognative Service for Text Analytics using the supplied:
//
// - API Key (Dashboard > Resources > {COGNATIVE SERVICES RESOURCE NAME} > Resource Management > Keys)
//
// - Resource Name ('http://{COGNATIVE SERVICES RESOURCE NAME}.cognitiveservices.azure.com')
//
// - Text to be analyzed for a positive/negative sentiment reading.
func Sentiment(apiKey string, resourceName string, document []map[string]string) string {

	// Define the API to make a call to
	var apiType = "sentiment"

	// Build a new Request struct with the inputs to pass into
	request := Request{apiKey, resourceName, apiType, document}
	output := apiRequest(request)
	return string(output)

}

func apiRequest(apiRequest Request) []byte {

	// If API Key input is blank
	if apiRequest.Key == "" {
		// Set API Key as environment variable 'TEXT_ANALYTICS_SUBSCRIPTION_KEY'
		apiRequest.Key = os.Getenv("TEXT_ANALYTICS_SUBSCRIPTION_KEY")

		// If environment variable 'TEXT_ANALYTICS_SUBSCRIPTION_KEY' is blank/does not exist
		if apiRequest.Key == "" {
			// No dice
			log.Fatal("Check API Key input or set/export the environment variable for 'TEXT_ANALYTICS_SUBSCRIPTION_KEY'.")
		}
	}

	// If Resource Name input is blank
	if apiRequest.Endpoint == "" {
		// Set Resource Name as environment variable 'TEXT_ANALYTICS_ENDPOINT'
		apiRequest.Endpoint = os.Getenv("TEXT_ANALYTICS_ENDPOINT")

		// If environment variable 'TEXT_ANALYTICS_ENDPOINT' is blank/does not exist
		if apiRequest.Endpoint == "" {
			// No dice
			log.Fatal("Check the Resource Name input or set/export the environment variable for 'TEXT_ANALYTICS_ENDPOINT'.")
		}
	}

	// Complete the definition of the API Endpoint for sentiment analysis
	var apiEndpoint = "https://" + apiRequest.Endpoint + ".cognitiveservices.azure.com/text/analytics/v2.1/" + apiRequest.Type

	// Ensuring input text to be analyzed encoded in JSON. Address pointer probably unnecessary
	documents, err := json.Marshal(&apiRequest.Text)
	if err != nil {
		log.Fatal("Error marshaling provided text into data: %v\n", err)
	}

	// Serialize text into Reader for POST request
	data := strings.NewReader("{\"documents\": " + string(documents) + "}")

	// Define HTTP request client within HTTP and timeout parameters
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	// Define HTTP request as POST with API Endpoint and Text for transmission
	request, err := http.NewRequest("POST", apiEndpoint, data)
	if err != nil {
		log.Fatal("Error creating POST request: %v\n", err)
	}

	// Add Headers to the defined HTTP request
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Ocp-Apim-Subscription-Key", apiRequest.Key)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error on request: %v\n", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	// Define throw-away interface to store the soon-to-be uncoded JSON
	var f interface{}

	// Uncoding the JSON-encoded response into our throw-away interface
	json.Unmarshal(body, &f)

	// Format the uncoded JSON into readable JSON
	jsonFormatted, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		log.Fatal("Error producing JSON: %v\n", err)
	}

	// BOOM
	return jsonFormatted
}
