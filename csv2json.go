package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//Params for cmd line arguments
type Params struct {
	In        *string
	Out       *string
	Delimeter *string
	Verbose   *bool
}

// JSONStrHdrs defines the structure for the generated JSON with header information
type JSONStrHdrs struct {
	Headers []string `json:"headers"`
	Lines   []Line   `json:"lines"`
}

type Line struct {
	Records map[string]interface{}
}

func main() {

	// Parses parameters from command line
	params := Params{
		In:        flag.String("i", "", "input"),
		Out:       flag.String("o", "output.json", "output"),
		Delimeter: flag.String("d", ",", "delimeter in csv"),
		Verbose:   flag.Bool("v", false, "Verbose, some debug info"),
	}

	flag.Parse()

	// Read the params - debug only
	if *params.Verbose {

		fmt.Println("CSV2JSON converter - expects files as input")
		fmt.Println("File to convert:", *params.In)
		fmt.Println("File to after conversion:", *params.Out)
	}

	csvfile, err := os.Open(*params.In)
	checkError(err)

	defer csvfile.Close()

	// Read the input
	if *params.Verbose {
		fmt.Println("Opening my input file ", *params.In)
	}

	data, err := os.Open(*params.In)
	checkError(err)

	// Read the CSV
	myrdr := csv.NewReader(data)
	//myrdr.TrimLeadingSpace = true //Maybe we do not need this
	myrdr.Comma = bytes.Runes([]byte(*params.Delimeter))[0]

	// Make the JSON Structure
	counter := 0
	var headerinfo []string
	var lines []Line
	var mydata []byte

	for {
		oneline, err := myrdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if counter == 0 {
			headerinfo = oneline
		} else {
			recordMap := make(map[string]interface{})
			for i, k := range headerinfo {
				recordMap[k] = oneline[i]
			}

			lines = append(lines, Line{Records: recordMap})
		}

		counter++

	}
	if *params.Verbose {
		fmt.Println("Number of rows processed:", counter)
	}
	myjson := JSONStrHdrs{
		Headers: headerinfo,
		Lines:   lines,
	}

	mydata, err = json.MarshalIndent(myjson, "", "\t")
	checkError(err)

	ioerr := ioutil.WriteFile(*params.Out, mydata, 0644)
	checkError(ioerr)

}
