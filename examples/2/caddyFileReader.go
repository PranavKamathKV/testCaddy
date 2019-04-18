package main

import(
    "fmt"
    "io/ioutil"
    "os"

    "github.com/mholt/caddy"
    _ "github.com/mholt/caddy/caddyhttp"
)

func init(){
	fmt.Println("Inside Init")
    caddy.SetDefaultCaddyfileLoader("default",caddy.LoaderFunc(loadConfig))
}

func loadConfig(serveType string) (caddy.Input, error){
	fmt.Println("Inside LoadConfig func")
	caddy.DefaultConfigFile = "/home/testuser/dummy/testCaddy/caddyTest/examples/2/Caddyfile"
	fmt.Println("Printing cadd.DefaultConfigFile", caddy.DefaultConfigFile)
	fileContents, err:= ioutil.ReadFile(caddy.DefaultConfigFile)
	if err!=nil{
		if os.IsNotExist(err){
			fmt.Println("File not found",err)
			return nil, nil
		}
		fmt.Println("Error reading file", err)
		return nil, err
	}
	fmt.Println("Loading CaddyFile", string(fileContents))
	return caddy.CaddyfileInput{
		Contents: fileContents,
		Filepath: "./Caddyfile",
		ServerTypeName: serveType,

	}, nil

	//Filepath: caddy.DefaultConfigFile,

}


func main(){
    caddy.AppName= "Simple App"
    caddy.AppVersion= "0.1"

    caddyfile, err := caddy.LoadCaddyfile("http")
    if err!= nil{
	    panic(err)
    }

    fmt.Println("Starting the caddy server!!")

    caddyInst, err := caddy.Start(caddyfile)
    if err!=nil{
	    panic(err)
    }

    caddyInst.Wait()


}
