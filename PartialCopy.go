/**
* Author: WuLongyue(wly65535@163.com)
* Date: 2018-05-06
*/

package main

import "fmt"
import "flag"
import "io"
import "os"

func main() {
	var inputFileName, outputFileName string
	var inputFileSkip, outputFileSeek int64
	var outputFileCount int64
	var outputFileAppend bool

	flag.StringVar(&inputFileName, "input", "", "input filename. (position parameter is also ok)")
	flag.StringVar(&outputFileName, "output", "", "output filename. (position parameter is also ok)")
	flag.Int64Var(&inputFileSkip, "skip", 0, "seek bytes of input file. (negative number is supported)")
	flag.Int64Var(&outputFileSeek, "seek", 0, "seek bytes of output file. (negative number is supported)")
	flag.Int64Var(&outputFileCount, "count", 0, "copy max count bytes.")
	flag.BoolVar(&outputFileAppend, "append", false, "append if file exists")
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println("\nExample:")
		fmt.Printf("\t1. %s src.txt dst.txt\n", os.Args[0])
		fmt.Println("\t   #copy src.txt to dst.txt, use position parameter")
		fmt.Printf("\n\t2. %s -output dst.txt src.txt\n", os.Args[0])
		fmt.Println("\t   #copy src.txt to dst.txt, use position and key parameter")
		fmt.Println("\t   #unspecified parammeter will use position parameter")
		fmt.Printf("\n\t3. %s -skip 100 src.txt dst.txt\n", os.Args[0])
		fmt.Println("\t   #seek 100 bytes from beginning of src.txt before copy")
		fmt.Printf("\n\t4. %s -skip -100 src.txt dst.txt\n", os.Args[0])
		fmt.Println("\t   #seek 100 bytes from ending of src.txt before copy")
		fmt.Printf("\n\t5. %s -seek 100 src.txt dst.txt\n", os.Args[0])
		fmt.Println("\t   #seek 100 bytes from beginning of dst.txt before copy")
		fmt.Printf("\n\t6. %s -seek -100 src.txt dst.txt\n", os.Args[0])
		fmt.Println("\t   #seek 100 bytes from ending of dst.txt before copy")
		fmt.Printf("\n\t7. %s -count 50 src.txt dst.txt\n", os.Args[0])
		fmt.Println("\t   #copy 50 bytes from src.txt to dst.txt")
		fmt.Printf("\n\t8. %s -append src.txt dst.txt\n", os.Args[0])
		fmt.Println("\t   #use append mode instead of truncate if file exists")
	}
	flag.Parse()

	if inputFileName != "" && outputFileName != "" && flag.NArg() > 0 || (inputFileName == "" && outputFileName != "" || outputFileName == "" && inputFileName != "") && flag.NArg() > 1 {
		fmt.Println("invalid parameter count")
		return
	}

	if inputFileName == "" && outputFileName == "" {
		inputFileName = flag.Arg(0)
		outputFileName = flag.Arg(1)
	} else if inputFileName == "" {
		inputFileName = flag.Arg(0)
	} else if outputFileName == "" {
		outputFileName = flag.Arg(0)
	}

	if inputFileName != "" {
		_, err := os.Stat(inputFileName)
		if err != nil {
			fmt.Println("os.Stat", err)
			return	
		}
	}

	var inputFile, outputFile *os.File
	if inputFileName != "" {
		var err error
		inputFile, err = os.Open(inputFileName)
		if err != nil {
			fmt.Println("os.Open", err)
			return	
		}
		defer inputFile.Close()
	}
	if outputFileName != "" {
		var err error
		if outputFileAppend || outputFileSeek != 0 {
			openFlag := os.O_CREATE | os.O_RDWR
			if outputFileAppend {
				openFlag |= os.O_APPEND
			}
			outputFile, err = os.OpenFile(outputFileName, openFlag, 0755)
		} else {
			outputFile, err = os.Create(outputFileName)
		}
		if err != nil {
			fmt.Println("os.Create", err)
			return	
		}
		defer outputFile.Close()
	}

	if inputFile == nil {
		inputFile = os.Stdin	
	}
	if outputFile == nil {
		outputFile = os.Stdout	
	}

	if inputFileSkip != 0 {
		var err error
		if inputFileSkip > 0 {
			_, err = inputFile.Seek(inputFileSkip, os.SEEK_SET)
		} else if inputFileSkip < 0 {
			_, err = inputFile.Seek(inputFileSkip, os.SEEK_END)
		}
		if err != nil {
			fmt.Println("inputFile.Seek", err)
			return	
		}
	}
	if outputFileSeek != 0 {
		var err error
		if outputFileSeek > 0 {
			_, err = outputFile.Seek(outputFileSeek, os.SEEK_SET)
			} else if outputFileSeek < 0 {
				_, err = outputFile.Seek(outputFileSeek, os.SEEK_END)
		}
		if err != nil {
			fmt.Println("outputFile.Seek", err)
			return	
		}
	}

	if outputFileCount == 0 {
		writen, err := io.Copy(outputFile, inputFile)
		if err != nil {
			fmt.Println("io.Copy", err)
			return	
		}
		fmt.Println(writen, "bytes copied")
	} else {
		writen, err := io.CopyN(outputFile, inputFile, outputFileCount)
		if err != nil {
			fmt.Println("io.CopyN", err)
			return	
		}
		fmt.Println(writen, "bytes copied")
	}
}