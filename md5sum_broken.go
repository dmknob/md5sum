package main

import (
    //"crypto/md5"
    //"encoding/hex"
    "flag"
    "fmt"
    //"io"
    //"io/ioutil"
    //"os"
    //"path/filepath" //To slit the filename from filepath
    //"reflect" //Used to see TypeOf vars
    //"time"
)

//Check for errors and panic case detect
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	//fmt.Println(len(os.Args))
    //fileLocationPtr := flag.String("fileLocation", "md5sum", "Mandatory file to compute MD5")
    checkFlagPtr := flag.String("c", "", "To check a existing file")
    flag.Parse()
    fileLocation := flag.Args()
    fmt.Println(fileLocation)
    //fmt.Println("fileLocation:", *fileLocationPtr)
    //argsWithoutProg := os.Args[1:] //Cuts off the executable file name from args
    //fmt.Println(argsWithoutProg[0])
	fmt.Print("Checking MD5 of ")
    fmt.Println(fileLocation)
    fmt.Println(*checkFlagPtr)
    /*dat, err := ioutil.ReadFile(fileLocation) //Reads the file
    check(err)
    md5Checksum := md5.Sum(dat) //Calculates de MD5
    fmt.Printf("%x\n", md5Checksum)
    //fmt.Println(reflect.TypeOf(md5Checksum))

    /*
    md5Str := string(md5Checksum[:])
    fmt.Printf("%x\n", md5Str)
    fmt.Println(reflect.TypeOf(md5Str))
    */
/*
    fileName := filepath.Base(fileLocation)   //Get the filename
    fmt.Println(fileName)
    f, err := os.Create(fileLocation+".txt")  //Create a new file to hold de MD5 hash
    check(err)
    defer f.Close() //We must close the file when will not be in use (in this case, when main() exits)
    n, err := f.WriteString("MD5 of " + fileName + "\r\n" + hex.EncodeToString(md5Checksum[:]))
    check(err)
    fmt.Printf("wrote %d bytes\n", n)
    f.Sync()
    */
}