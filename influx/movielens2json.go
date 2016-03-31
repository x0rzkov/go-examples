package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
   JSON Schema

   [{
   	"measurement": "ratings",
   	"tags": {
   		"userid": 1,
   		"movieid": 1193
   	},
   	"fields": {
   		"rating": 5
   	},
   	"time": 978300760
   }]

*/

type Tags struct {
	Userid  int `json:"userid"`
	Movieid int `json:"movieid"`
}

type Fields struct {
	Rating int
}

type Point struct {
	Measurement string  `json:"measurement"`
	Tag         *Tags   `json:"tag"`
	Field       *Fields `json:"field"`
	Time        int     `json:"time"`
}

type points []Point

var t *Tags
var fld *Fields
var p Point

func main() {

	f, _ := os.Open("ratings.dat")

	scanner := bufio.NewScanner(f)

	i := 0

	slcPoints := make(points, 0)

	for scanner.Scan() {
		i++

		line := scanner.Text()

		result := strings.Split(line, "::")

		var uid, mid, r int
		var tm int
		cnt := 0

		for _ = range result {

			//Example: 1::1193::5::978300760
			//userid::movieid::rating::time

			if cnt == 3 {

				uid, _ = strconv.Atoi(result[0])
				mid, _ = strconv.Atoi(result[1])
				r, _ = strconv.Atoi(result[2])
				tm, _ = strconv.Atoi(result[3])

				t = &Tags{Userid: uid, Movieid: mid}

				fld = &Fields{Rating: r}

				p = Point{Measurement: "movielens", Tag: t, Field: fld, Time: tm}

				_, err := json.Marshal(t)
				if err != nil {
					log.Println(err)
				}

				slcPoints = append(slcPoints, p)
			}
			cnt++
		}

		/*
			if i == 5 {
				break
			}
		*/
	}

	//log.Println(slcPoints)

	jason, err := json.Marshal(slcPoints)
	if err != nil {
		log.Println(err)
	}
	//log.Println(string(jason))

	file, err := os.Create("json.txt")
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	_, err = file.Write(jason)
	if err != nil {
		log.Println(err)
	}
	log.Println("ok")
}
