package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"log"
)

func main() {
	filename := flag.String("fn", "poem.txt", "Input file to read.")
	output := flag.String("out", "output.txt", "Optput to that file.")
	line_count := flag.Int("line", 4, "How many line of 1 poem")
	
	flag.Parse()

	fileout, err := os.Create(*output)
	if err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(fileout)

	readfile, err := os.Open(*filename)
	defer readfile.Close()
	
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(readfile)
	scanner := bufio.NewScanner(reader)

	var leftSide []string
	var rightSide []string
	var nCount int
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if nCount % 2 == 0 {
			leftSide = append(leftSide, scanner.Text())
		} else { rightSide = append(rightSide, scanner.Text()) }
		nCount++
	}

	if len(leftSide) != len(rightSide) {
		log.Fatalln("Poem is not completed.")
	}

	chapter := 1
	for i := 0; i < len(leftSide); i += *line_count {
		count := 0
		header := "บทที่ " + strconv.Itoa(chapter) + "--------\n"
		
		fileout.WriteString(header)
		
		for {
			fileout.WriteString(leftSide[i+count])
			fileout.WriteString("\n")
			count++
			if count % *line_count == 0 && count != 0 {
				break
			}
		}

		fileout.WriteString("\n")
		count = 0

		for {
			fileout.WriteString(rightSide[i+count])
			fileout.WriteString("\n")
			count++
			if count % *line_count == 0 && count != 0 {
				break
			}
		}

		fileout.WriteString("\n")
		chapter++
	}
}
