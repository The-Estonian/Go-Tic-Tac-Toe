package screenClear

import(
	"os/exec"
	"os"
)

func Now() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}