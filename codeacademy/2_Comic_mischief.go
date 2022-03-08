package main

import "fmt"
import t "time"

func main() {
  var publisher, writer, artist, title string
  var year, pageNumber int
  var grade, cost float32

  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5

  cost = float32(pageNumber) * grade / float32(t.Now().Year() - year)

  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in the year", year, "has", pageNumber, "pages has the artist,", artist, ", and has the condition,", grade, "with the cost", cost)

  title = "Epic Vol. 1"
  writer =  "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  year = 2013
  pageNumber = 160
  grade = 9.0

  cost = float32(pageNumber) * grade / float32(t.Now().Year() - year)

  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in the year", year, "has", pageNumber, "pages has the artist,", artist, ", and has the condition,", grade, "with the cost", cost)
}