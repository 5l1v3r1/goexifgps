package GoExifGPS

// Author : kurtcc on github
// This will be called parse.go
import (
	"encoding/json"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
	"log"
	"os"
	"strings"

//This is how to print a given field
)

type Exif struct {
	tif *tiff.Tiff

	main map[FieldName]*tiff.Tag
}
type FieldName string

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func TrimPrefix(s string) string {
	s = s[1:]
	return s
}

// Use like this
//LatRef, Lat, LongRef,Longd := OpenParseJson("_JEF018993_sm.jpg") 
func OpenClose(filename string) (*exif.Exif, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	f.Close()
	if err != nil {
		return nil, err
	}

	return x, nil
}

//log.Fatal(err) log.Fatal makes the program stop so we don't need this
// coz we won't be able to get false ever.
func StdinDecode() (*exif.Exif, error) {
	r := os.Stdin
	xoo, err := exif.Decode(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input: %s\n", err)
		os.Exit(0)
	}

	return xoo, nil

}

func OpenParseJson(E *exif.Exif) (string, string, string, string) {
	// I want this to return all four values each as a string.	

	b, err := E.MarshalJSON()
	if err != nil {
		panic(err) //Format die output properly
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(b, &dat); err != nil {
		panic(err)
	}

	numR := dat["GPSLatitudeRef"].(string) //Lat and LatRef
	num := dat["GPSLatitude"].([]interface{})

	num2R := dat["GPSLongitudeRef"].(string)
	num2 := dat["GPSLongitude"].([]interface{})

	//*** Latitude
	Snum := fmt.Sprintf("%s", num)
	Tnum1 := TrimPrefix(Snum)
	Tnum := TrimSuffix(Tnum1, "]")

	// *** Longitude
	Snum2 := fmt.Sprintf("%s", num2)
	Tnum2 := TrimPrefix(Snum2)
	Tnum_2 := TrimSuffix(Tnum2, "]")

	return numR, Tnum, num2R, Tnum_2
}