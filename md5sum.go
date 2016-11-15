package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath" //To slit the filename from filepath
    "strings"
)

//Globals
var pL = fmt.Println //Alias for print Line
var filePath = ""   //Store the path and filename of file
var checkFlag = false
var genFlag = false
var helpFlag = false
var printFlag = false
var computedMD5 [16]byte

//Check for errors and panic case detect
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func helpPrint() {
    pL("\nUse md5sum 'filename' to calculate and display the MD5sum of the file")
    pL("-c or --check 'filename' - open 'filename'.txt and compare saved MD5sum with calculated MD5sum of 'filename'")
    pL("-g ou --generate 'filename' - calculate the MD5sum of 'filename' and store it inside 'filename'.txt")
    pL("-p ou --print 'filename' - calculate the MD5sum of 'filename' and print it on screen")
    pL("-h or --help - display this help message")
}

func init() { //Runs before main(). Useful to parse os.Args (commandline params)
/*Flags:
    -h --help
    -c --check
    -g --generate
    -p --print
*/
    var n = len(os.Args) //Number of arguments passed on program start
    switch {
    case n < 2 :
        helpPrint()
        pL("\nYou must pass at least a filename to compute the MD5!")
        os.Exit(2)

    //case n == 2:

    case n == 3:
            if os.Args[1] == "-c" || os.Args[1] == "--check" {
                checkFlag = true
            } else if os.Args[1] == "-g" || os.Args[1] == "--generate"{
                genFlag = true
            } else if os.Args[1] == "-p" || os.Args[1] == "--print"{
                printFlag = true
            } else {
                helpFlag = true
            }
            filePath = os.Args[2]
    default:
        helpPrint()
        os.Exit(2)
    }
}

func fnComputeMD5() {
    pL("Computing MD5 of " + filePath)
    dat, err := ioutil.ReadFile(filePath) //Reads the file
    check(err)
    computedMD5 = md5.Sum(dat) //Calculates de MD5
    pL(hex.EncodeToString(computedMD5[:]))
}

func fnCreateFile() {
    fileName := filepath.Base(filePath)   //Get the filename
    f, err := os.Create(fileName+".txt")  //Create a new file to hold de MD5 hash
    check(err)
    defer f.Close() //We must close the file when will not be in use (in this case, when main() exits)
    _, err = f.WriteString(hex.EncodeToString(computedMD5[:]) + "  " + fileName)
    check(err)
    f.Sync()
}

func fnReadMD5() {
    //pL("Opening file: " + filePath+".txt")
    f, err := os.Open(filePath+".txt")
    check(err)
    storedHash := make([]byte, 32)  //Allocate space to stored MD5 hash
    _, err = f.Read(storedHash)
    check(err)
    fnCompareMD5(storedHash)  //Compare computedMD5 with savedMD5 and print result

}

func fnCompareMD5(savedMD5 []byte) {
    if strings.Compare(hex.EncodeToString(computedMD5[:]), string(savedMD5)) == 0 {
        pL("All OK, the MD5 hash matches the saved one.")
    } else {
        pL("Houston, you have a problem! MD5 hash of file does not match the saved one.")
    }
}

func generate() {   //Compute the MD5sum of 'filename' and store it inside new file 'filename'.txt
    //open 'filename', compute MD5 (computedMD5), store inside new file
    fnComputeMD5()  //open 'filename', compute MD5 (computedMD5)
    fnCreateFile()  //Create 'filename'.txt, store computedMD5 inside

}

func checkMD5() {   //Open 'filename'.txt, read savedMD5, open 'filename', compute MD5 (computedMD5), compare and print result
    fnComputeMD5()  //open 'filename', compute MD5 (computedMD5)
    fnReadMD5()     //open 'filename'.txt, get savedMD5, call fnCompareMD5
}

func main() {
    
	switch {
    case helpFlag:
        helpPrint()
    case genFlag:
        generate()
    case checkFlag:
        checkMD5()
    case printFlag:
        fnComputeMD5()
    }
}