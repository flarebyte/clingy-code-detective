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

const currentDirectory = process.cwd();
const folderName = path.basename(currentDirectory);

const projectName = `github.com/flarebyte/${folderName}`;

const brothFile = fs.readFileSync("baldrick-broth.yaml", "utf8");
const brothContent = YAML.parse(brothFile);
const version = brothContent.model.project.version;
const currentDate = getBritishDate().replaceAll(" ", "-");

const ldflags = `-X ${projectName}/internal/cli.Version=${version} -X ${projectName}/internal/cli.Date=${currentDate}`;
const platforms = [
  { label: "Linux (amd64)", os: "linux", arch: "amd64" },
  { label: "Linux (arm64)", os: "linux", arch: "arm64" },

  { label: "macOS (Intel)", os: "darwin", arch: "amd64" },
  { label: "macOS (Apple Silicon)", os: "darwin", arch: "arm64" },

  { label: "Windows (amd64)", os: "windows", arch: "amd64" },
  { label: "Windows (arm64)", os: "windows", arch: "arm64" },
];

for (const p of platforms) {
  echo(p.label);
  await $`GOOS=${p.os} GOARCH=${p.arch} go build -o build/${folderName}-${p.os}-${p.arch} -ldflags ${ldflags}`;
}

const checksum = await $`shasum -a 256 build/${folderName}-*`;
await fs.writeFile("build/checksums.txt", checksum.stdout, "utf8");
