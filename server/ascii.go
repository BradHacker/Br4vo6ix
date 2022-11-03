package main

import "fmt"

const CoverArtColor = "\033[31;1m"
const CoverArtResetColor = "\033[0m"

const CoverArt = `
    ____       __ __             _____ _     
   / __ )_____/ // /_   ______  / ___/(_)  __
  / __  / ___/ // /| | / / __ \/ __ \/ / |/_/
 / /_/ / /  /__  __/ |/ / /_/ / /_/ / />  <  
/_____/_/     /_/  |___/\____/\____/_/_/|_|
`

func PrintCoverArt() {
	fmt.Print(string(CoverArtColor), CoverArt, string(CoverArtResetColor))
}
