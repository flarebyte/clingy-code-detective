function getBritishDate() {
  const now = new Date();

  const options = {
    year: "numeric",
    month: "long", // 'long' for full month name (e.g., June)
    day: "numeric",
  };

  const britishDate = new Intl.DateTimeFormat("en-GB", options).format(now);

  return britishDate;
}

const name = "clingy"
const projectName = "github.com/flarebyte/clingy-code-detective"

const brothFile = fs.readFileSync("baldrick-broth.yaml", "utf8");
const brothContent = YAML.parse(brothFile);
const version = brothContent.model.project.version;
const currentDate = getBritishDate().replaceAll(' ', '-');

const ldflags = `-X ${projectName}/internal/cli.Version=${version} -X ${projectName}/internal/cli.Date=${currentDate}`

echo("Linux (amd64)");
$`GOOS=linux GOARCH=amd64 go build -o build/${name}-linux-amd64 -ldflags ${ldflags}`;

echo("macOS (Intel)");
$`GOOS=darwin GOARCH=amd64 go build -o build/${name}-darwin-amd64 -ldflags ${ldflags}`;

echo("macOS (Apple Silicon)");
$`GOOS=darwin GOARCH=arm64 go build -o build/${name}-darwin-arm64 -ldflags ${ldflags}`;
