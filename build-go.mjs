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

const brothFile = fs.readFileSync("baldrick-broth.yaml", "utf8");
const brothContent = YAML.parse(brothFile);
const version = brothContent.model.project.version;
const currentDate = getBritishDate().replaceAll(' ', '-');
echo(currentDate)

echo("Linux (amd64)");
$`GOOS=linux GOARCH=amd64 go build -o build/clingy-linux-amd64 -ldflags "-X cli.Version=${version} -X cli.Date=${currentDate}"`;

echo("macOS (Intel)");
$`GOOS=darwin GOARCH=amd64 go build -o build/clingy-darwin-amd64 -ldflags "-X cli.Version=${version} -X cli.Date=${currentDate}"`;

echo("macOS (Apple Silicon)");
$`GOOS=darwin GOARCH=arm64 go build -o build/clingy-darwin-arm64 -ldflags "-X cli.Version=${version} -X cli.Date=${currentDate}"`;
