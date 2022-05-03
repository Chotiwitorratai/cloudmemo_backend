package controllers

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
        maxResults = flag.Int64("max-results", 25, "Max YouTube results")
        query      = flag.String("query", "str", "Search term")
        service    *youtube.Service
        response   *youtube.SearchListResponse
)

type YoutubeSearchResult struct {
    Title, YoutubeId string
}
const developerKey = "AIzaSyAZIfnXjE5-sxuImN31vJpwAe58nNeASfc"

func SearchMusic(c *fiber.Ctx) error {
		payload := struct {
        Data  string `json:"data"`

    }{}
	CheckToken(c)
	c.BodyParser(&payload)
        flag.Parse()
        client := &http.Client{
                Transport: &transport.APIKey{Key: developerKey},
        }

        service, err := youtube.New(client)
        if err != nil {
                log.Fatalf("Error creating new YouTube client: %v", err)
        }
		var  list = []string{"id","snippet"}
        // Make the API call to YouTube.
        call := service.Search.List(list).
                Q(payload.Data).
                MaxResults(*maxResults)

        response, err := call.Do()
        if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
		}
        result := []YoutubeSearchResult{}
        for _, item := range response.Items {
                switch item.Id.Kind {
                case "youtube#video":
                        result = append(result, YoutubeSearchResult{Title: item.Snippet.Title, YoutubeId: item.Id.VideoId})
       			}
		}
	printIDs("Videos", result)
	return c.Status(200).JSON(fiber.Map{"status": "success","message": nil,"data":result})
}



func printIDs (sectionName string, matches []YoutubeSearchResult) {
        fmt.Printf("%v:\n", sectionName)
        for _, match := range matches {
                fmt.Printf("[%v] %v\n", match.YoutubeId, match.Title)
                // fmt.Printf("[%v] %v\n", id, title)
        }
        fmt.Printf("\n\n")
}
