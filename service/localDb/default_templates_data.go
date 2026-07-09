package localDb

type defaultTemplate struct {
	ID          string
	Name        string
	Description string
	Script      string
	CreatedAt   string
}

// Built-in startup templates for Amazon Linux 2023 (dnf).
// Seed/sync policy lives in defaultTemplates.go.
var defaultTemplates = []defaultTemplate{
	// Languages
	{
		ID:          "00000000-0000-0000-0001-000000000001",
		Name:        "python3",
		Description: "Python 3 and pip (Amazon Linux 2023)",
		CreatedAt:   "1970-01-01 00:00:01",
		Script: `command -v python3 >/dev/null 2>&1 || dnf install -y python3 python3-pip
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000002",
		Name:        "java21",
		Description: "OpenJDK 21 (Amazon Corretto) on Amazon Linux 2023",
		CreatedAt:   "1970-01-01 00:00:02",
		Script: `command -v javac >/dev/null 2>&1 || dnf install -y java-21-amazon-corretto-devel
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000003",
		Name:        "cpp",
		Description: "GCC C/C++ toolchain on Amazon Linux 2023",
		CreatedAt:   "1970-01-01 00:00:03",
		Script: `command -v g++ >/dev/null 2>&1 || dnf install -y gcc gcc-c++ make
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000004",
		Name:        "go",
		Description: "Go toolchain from Amazon Linux 2023 repos",
		CreatedAt:   "1970-01-01 00:00:04",
		Script: `command -v go >/dev/null 2>&1 || dnf install -y golang
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000005",
		Name:        "rust",
		Description: "Rust via rustup for ec2-user",
		CreatedAt:   "1970-01-01 00:00:05",
		Script: `runuser -u ec2-user -- bash -lc 'command -v rustc >/dev/null 2>&1 || curl --proto "=https" --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y'
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000006",
		Name:        "ruby",
		Description: "Ruby from Amazon Linux 2023 repos",
		CreatedAt:   "1970-01-01 00:00:06",
		Script: `command -v ruby >/dev/null 2>&1 || dnf install -y ruby
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000007",
		Name:        "node22",
		Description: "Node.js 22 and npm from Amazon Linux 2023 repos",
		CreatedAt:   "1970-01-01 00:00:07",
		Script: `command -v node >/dev/null 2>&1 || dnf install -y nodejs22 nodejs22-npm
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000008",
		Name:        "dotnet8",
		Description: ".NET 8 SDK on Amazon Linux 2023",
		CreatedAt:   "1970-01-01 00:00:08",
		Script: `command -v dotnet >/dev/null 2>&1 || dnf install -y dotnet-sdk-8.0
`,
	},

	// Package managers
	{
		ID:          "00000000-0000-0000-0001-000000000010",
		Name:        "pip",
		Description: "Python pip (Amazon Linux 2023)",
		CreatedAt:   "1970-01-01 00:00:10",
		Script: `command -v pip3 >/dev/null 2>&1 || dnf install -y python3 python3-pip
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000011",
		Name:        "npm22",
		Description: "npm via Node.js 22 on Amazon Linux 2023",
		CreatedAt:   "1970-01-01 00:00:11",
		Script: `command -v npm >/dev/null 2>&1 || dnf install -y nodejs22 nodejs22-npm
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000012",
		Name:        "bun",
		Description: "Bun JavaScript runtime (system-wide install)",
		CreatedAt:   "1970-01-01 00:00:12",
		Script: `if ! command -v bun >/dev/null 2>&1; then
  export BUN_INSTALL=/usr/local/bun
  curl -fsSL https://bun.sh/install | bash
  ln -sf /usr/local/bun/bin/bun /usr/local/bin/bun
fi
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000013",
		Name:        "pnpm22",
		Description: "pnpm via npm for ec2-user (installs Node.js 22 if needed)",
		CreatedAt:   "1970-01-01 00:00:13",
		Script: `command -v node >/dev/null 2>&1 || dnf install -y nodejs22 nodejs22-npm
if ! runuser -u ec2-user -- bash -lc 'command -v pnpm >/dev/null 2>&1'; then
  runuser -u ec2-user -- bash -lc 'npm install -g pnpm'
  grep -q '\.local/bin' /home/ec2-user/.bashrc 2>/dev/null || echo 'export PATH="$HOME/.local/bin:$PATH"' >> /home/ec2-user/.bashrc
fi
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000014",
		Name:        "yarn22",
		Description: "Yarn via npm for ec2-user (installs Node.js 22 if needed)",
		CreatedAt:   "1970-01-01 00:00:14",
		Script: `command -v node >/dev/null 2>&1 || dnf install -y nodejs22 nodejs22-npm
if ! runuser -u ec2-user -- bash -lc 'command -v yarn >/dev/null 2>&1'; then
  runuser -u ec2-user -- bash -lc 'npm install -g yarn'
  grep -q '\.local/bin' /home/ec2-user/.bashrc 2>/dev/null || echo 'export PATH="$HOME/.local/bin:$PATH"' >> /home/ec2-user/.bashrc
fi
`,
	},

	// Tools
	{
		ID:          "00000000-0000-0000-0001-000000000015",
		Name:        "git",
		Description: "Git version control on Amazon Linux 2023",
		CreatedAt:   "1970-01-01 00:00:15",
		Script: `command -v git >/dev/null 2>&1 || dnf install -y git
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000016",
		Name:        "docker",
		Description: "Docker engine on Amazon Linux 2023 (ec2-user in docker group)",
		CreatedAt:   "1970-01-01 00:00:16",
		Script: `if ! command -v docker >/dev/null 2>&1; then
  dnf install -y docker
  systemctl enable --now docker
fi
getent group docker >/dev/null || groupadd docker
usermod -aG docker ec2-user 2>/dev/null || true
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000017",
		Name:        "uv",
		Description: "uv Python package manager for ec2-user",
		CreatedAt:   "1970-01-01 00:00:17",
		Script: `if ! runuser -u ec2-user -- bash -lc 'command -v uv >/dev/null 2>&1'; then
  runuser -u ec2-user -- bash -lc 'curl -LsSf https://astral.sh/uv/install.sh | sh'
  grep -q '\.local/bin' /home/ec2-user/.bashrc 2>/dev/null || echo 'export PATH="$HOME/.local/bin:$PATH"' >> /home/ec2-user/.bashrc
fi
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000018",
		Name:        "maven",
		Description: "Apache Maven on Amazon Linux 2023",
		CreatedAt:   "1970-01-01 00:00:18",
		Script: `command -v mvn >/dev/null 2>&1 || dnf install -y maven
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000019",
		Name:        "gradle",
		Description: "Gradle build tool on Amazon Linux 2023",
		CreatedAt:   "1970-01-01 00:00:19",
		Script: `command -v gradle >/dev/null 2>&1 || dnf install -y gradle
`,
	},

	// AI coding agents
	{
		ID:          "00000000-0000-0000-0001-000000000020",
		Name:        "claude-code",
		Description: "Claude Code CLI (Anthropic dnf repo for Amazon Linux 2023)",
		CreatedAt:   "1970-01-01 00:00:20",
		Script: `if ! command -v claude >/dev/null 2>&1; then
  tee /etc/yum.repos.d/claude-code.repo >/dev/null <<'EOF'
[claude-code]
name=Claude Code
baseurl=https://downloads.claude.ai/claude-code/rpm/stable
enabled=1
gpgcheck=1
gpgkey=https://downloads.claude.ai/keys/claude-code.asc
EOF
  dnf install -y claude-code
fi
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000021",
		Name:        "cursor",
		Description: "Cursor Agent CLI for ec2-user",
		CreatedAt:   "1970-01-01 00:00:21",
		Script: `if ! runuser -u ec2-user -- bash -lc 'command -v agent >/dev/null 2>&1'; then
  runuser -u ec2-user -- bash -lc 'curl https://cursor.com/install -fsS | bash'
  grep -q '\.local/bin' /home/ec2-user/.bashrc 2>/dev/null || echo 'export PATH="$HOME/.local/bin:$PATH"' >> /home/ec2-user/.bashrc
fi
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000022",
		Name:        "codex22",
		Description: "OpenAI Codex CLI via npm for ec2-user (installs Node.js 22 if needed)",
		CreatedAt:   "1970-01-01 00:00:22",
		Script: `command -v node >/dev/null 2>&1 || dnf install -y nodejs22 nodejs22-npm
if ! runuser -u ec2-user -- bash -lc 'command -v codex >/dev/null 2>&1'; then
  runuser -u ec2-user -- bash -lc 'npm install -g @openai/codex'
  grep -q '\.local/bin' /home/ec2-user/.bashrc 2>/dev/null || echo 'export PATH="$HOME/.local/bin:$PATH"' >> /home/ec2-user/.bashrc
fi
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000023",
		Name:        "pi",
		Description: "Pi coding agent for ec2-user (installs Node.js 22.19+ if needed)",
		CreatedAt:   "1970-01-01 00:00:23",
		Script: `command -v xz >/dev/null 2>&1 || dnf install -y xz
if ! runuser -u ec2-user -- bash -lc 'command -v pi >/dev/null 2>&1'; then
  runuser -u ec2-user -- bash -lc 'set -e
NODE_BASE="$HOME/.local/share/pi-node"
NODE_BIN="$NODE_BASE/current/bin"
if ! "$NODE_BIN/node" -e "const v=process.versions.node.split(\".\").map(Number);process.exit(v[0]>22||(v[0]===22&&(v[1]>19||(v[1]===19&&v[2]>=0)))?0:1)" 2>/dev/null; then
  case "$(uname -m)" in x86_64) ARCH=x64;; aarch64|arm64) ARCH=arm64;; *) echo "unsupported arch for pi"; exit 1;; esac
  DIST=https://nodejs.org/dist/latest-v22.x
  TMP=$(mktemp -d)
  curl -fsSL "$DIST/SHASUMS256.txt" -o "$TMP/sums"
  NODE_FILE=$(awk -v s="-linux-$ARCH.tar.xz" "index(\$2,\"node-v\")&&substr(\$2,length(\$2)-length(s)+1)==s{print \$2;exit}" "$TMP/sums")
  curl -fsSL "$DIST/$NODE_FILE" -o "$TMP/$NODE_FILE"
  grep -q " $NODE_FILE$" "$TMP/sums" || { echo "checksum entry not found for $NODE_FILE"; exit 1; }
  grep " $NODE_FILE$" "$TMP/sums" | (cd "$TMP" && sha256sum -c -)
  mkdir -p "$NODE_BASE"
  tar -xf "$TMP/$NODE_FILE" -C "$NODE_BASE"
  rm -f "$NODE_BASE/current"
  ln -s "$NODE_BASE/${NODE_FILE%.tar.xz}" "$NODE_BASE/current"
  rm -rf "$TMP"
fi
export PATH="$NODE_BIN:$HOME/.local/bin:$PATH"
npm install -g --ignore-scripts --prefix "$HOME/.local" @earendil-works/pi-coding-agent'
  grep -q 'pi-node/current/bin' /home/ec2-user/.bashrc 2>/dev/null || echo 'export PATH="$HOME/.local/share/pi-node/current/bin:$HOME/.local/bin:$PATH"' >> /home/ec2-user/.bashrc
fi
`,
	},
	{
		ID:          "00000000-0000-0000-0001-000000000024",
		Name:        "opencode",
		Description: "OpenCode AI agent CLI for ec2-user",
		CreatedAt:   "1970-01-01 00:00:24",
		Script: `if ! runuser -u ec2-user -- bash -lc 'command -v opencode >/dev/null 2>&1'; then
  runuser -u ec2-user -- bash -lc 'curl -fsSL https://opencode.ai/install | bash'
  grep -q '\.opencode/bin' /home/ec2-user/.bashrc 2>/dev/null || echo 'export PATH="$HOME/.opencode/bin:$PATH"' >> /home/ec2-user/.bashrc
fi
`,
	},
}
