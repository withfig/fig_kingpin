# Fig Kingpin

Support for generating Fig Spec from [Kingpin CLI apps](https://github.com/alecthomas/kingpin)

## Usage

### 1. Add the `figkingpin` package and `--completion-script-fig` to your app

```go
import (
	"github.com/alecthomas/kingpin/v2"

	// 1. Add the figkingpin package
	figkingpin "github.com/withfig/fig_kingpin"
)

var (
	app = kingpin.New("App", "A demo app")

	// 2. Add a top level flag to gen fig spec, it is hidden from the help output
	genFig = app.Flag("completion-script-fig", "Generate completion script for fig.").Hidden().PreAction(figkingpin.GenerateFigCompletionScript(app)).Bool()
)
```

### 2. Generate a Fig Spec via `--completion-script-fig`

```bash
go run main.go --completion-script-fig > fig-spec.ts
```
