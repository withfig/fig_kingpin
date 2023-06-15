# Fig Kingpin

Support for generating Fig Spec from [Kingpin CLI apps](https://github.com/alecthomas/kingpin)

## Usage

### 1. Add the `figkingpin` package and `--completion-spec-fig` to your app

```go
import (
	"github.com/alecthomas/kingpin/v2"

	// 1. Add the figkingpin package
	figkingpin "github.com/withfig/fig_kingpin"
)

var (
	app = kingpin.New("App", "A demo app")

	// 2. Add a top level flag to gen fig spec, it is hidden from the help output
	completionSpecFig = app.Flag("completion-spec-fig", "Generate completion script for fig.").Hidden().PreAction(figkingpin.GenerateFigCompletionSpec(app)).String())
```

### 2. Generate a Fig Spec via `--completion-spec-fig`

```bash
go run main.go --completion-spec-fig spec-name > spec-name.ts
```
