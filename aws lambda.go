package main

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	gitUrl, _ := url.QueryUnescape(r.QueryStringParameters[`giturl`])

	fileExt := strings.Split(gitUrl, `.`)

	if fileExt == nil {
		fmt.Println(`Failed to parse URL path: `, gitUrl)
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}
	fileMime := mime.TypeByExtension(`.` + fileExt[len(fileExt)-1])

	fileData, err := http.Get(`https://raw.githubusercontent.com` + gitUrl)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	defer fileData.Body.Close()

	fileBytes, err := ioutil.ReadAll(fileData.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	fmt.Println(`Serving `, gitUrl)

	headers := map[string]string{`Content-Type`: fileMime}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(fileBytes), Headers: headers}, nil
}

func main() {
	lambda.Start(Handler)
}
