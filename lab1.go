package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	s := bufio.NewScanner(os.Stdin)
	n := 0
	content := [][]string{}
	header := "0"

	fmt.Println("===Start===")

	var (
		iFlag = flag.String("i", "0", "Use a file with the name file-name as an input")
		oFlag = flag.String("o", "0", "Use a file with the name file-name as an output")
		rFlag = flag.String("r", "0", "Sort input lines in reverse order")
		fFlag = flag.Int("f", 1, "Sort input lines by value number N")
		hFlag = flag.String("h", "0", "The first line is a header that must be ignored during sorting but included in the output")
	)

	flag.Parse()

	if *iFlag != "0" {
		fileIn, err := os.Open(*iFlag)

		if err != nil {
			log.Fatal(err)
		}

		defer fileIn.Close()
		s = bufio.NewScanner(fileIn)
	}

	for s.Scan() {
		line := s.Text()

		if line == "" {
			break
		}

		row := strings.Split(line, ",")

		if n == 0 {
			n = len(row)

			if *fFlag > n || *fFlag < 1 {
				log.Fatalf("ERROR: The line has %d fields, but for sorting needs %d\n", n, *fFlag)
			}
		}
		if len(row) != n {
			log.Fatalf("ERROR: The length of line must be %d, but it is %d\n", n, len(row))
		}
		if *hFlag == "" && header == "0" {
			header = line
			continue
		}

		content = append(content, row)
	}

	if *rFlag == "" {
		sort.Slice(content, func(i, j int) bool { return content[i][*fFlag-1] > content[j][*fFlag-1] })
	} else {
		sort.Slice(content, func(i, j int) bool { return content[i][*fFlag-1] < content[j][*fFlag-1] })
	}

	if *oFlag != "0" {
		fileOut, err := os.OpenFile(*oFlag, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)

		if err != nil {
			log.Fatal(err)
		}

		if header != "0" {
			_, err = fileOut.WriteString(strings.Join(strings.Split(header, ","), ""))
			fileOut.WriteString("\n")

			if err != nil {
				log.Fatal(err)
			}
		}

		for _, value := range content {
			_, err = fileOut.WriteString(strings.Join(value, ""))
			fileOut.WriteString("\n")

			if err != nil {
				log.Fatal(err)
			}
		}

		defer fileOut.Close()
	}

	fmt.Printf("===Finish===\n\n")
	fmt.Printf("===Result===\n\n")

	if header != "0" {
		fmt.Printf("%s\n", strings.Join(strings.Split(header, ","), ""))
	}
	for _, value := range content {
		fmt.Printf("%s\n", strings.Join(value, ""))
	}

}
