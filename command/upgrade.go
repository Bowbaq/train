package trainCommand

import (
	"fmt"
)

const CmdBinPath = "$GOPATH/bin/train"

func Upgrade() {
	bash("go get -u github.com/Bowbaq/train")
	bash("go build -o " + CmdBinPath + " github.com/Bowbaq/train/cmd")
	fmt.Println("Installed latest train command into " + CmdBinPath)
}
