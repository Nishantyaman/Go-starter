package main

import(
	"net/http"

	"github.com/gin-gonic/gin"
)
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID:"1", Title:"Test", Artist: "John", Price: 57.11},
	{ID:"2",Title: "Tum Mile",Artist: "Arijit", Price: 18.11},
	{ID:"3",Title: "Sarah", Artist: "Sara vegan", Price: 40.22},
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK,albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err !=nil{
		return
	}

	albums = append(albums,newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main(){
	router := gin.Default()
	router.GET("/albums",getAlbums)
	router.POST("/albums",postAlbums)

	router.Run("localhost:8080")
}