package main

import (
	"flag"
	"os"
	"bufio"
	"regexp"
	"fmt"
)

func main() {
	var input, output, typename, pack, vartype string
	spliter := regexp.MustCompile(`\s+`)
	flag.StringVar(&input, "i", "", "input file")
	flag.StringVar(&output, "o", "", "output go file")
	flag.StringVar(&typename, "type", "", "type name")
	flag.StringVar(&pack, "package", "", "package name")
	flag.StringVar(&vartype, "vartype", "int", "package name")
	flag.Parse()
	
	input_file, err := os.Open(input)	
	input_stream := bufio.NewScanner(input_file)
	if err != nil {
		panic(err)
	}
	output_file, err := os.Create(output)
	output_stream := bufio.NewWriter(output_file)	
	if err != nil {
		panic(err)
	}
	defer func(){
		input_file.Close()
		output_stream.Flush()
		output_file.Close()
	}()
	
	output_stream.WriteString(`package ` + pack + "\n");
	output_stream.WriteString("\n\n");
	output_stream.WriteString(`type ` + typename + " " +vartype+"\n");
	output_stream.WriteString("const ( \n");
	
	for input_stream.Scan() {
		str := input_stream.Text();
		re := spliter.Split(str,2)
		if len(re) > 1 {
			output_stream.WriteString(fmt.Sprintf("\t %s %s = %s \n",re[0],typename,re[1]))
		}
	}
	output_stream.WriteString(")\n")
	
	
}
