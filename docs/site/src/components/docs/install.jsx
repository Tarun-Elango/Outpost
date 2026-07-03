import DocPage from './doc-page'

export default function InstallDoc() {
  return (
    <DocPage title="Installation">
      <div className="card">
        <h2>Quick install</h2>
        <pre>
          <code>{`curl -fsSL https://raw.githubusercontent.com/Tarun-Elango/devbox-cli/main/scripts/install.sh | bash`}</code>
        </pre>
        <p className="note">
          Detects your OS and CPU, downloads the matching binary, installs to{' '}
          <code>~/.local/bin</code>, and adds that directory to your shell config if
          needed. Restart your shell, then verify:
        </p>
        <pre>
          <code>devbox ls</code>
        </pre>
      </div>

      <div className="card">
        <h2>System-wide install</h2>
        <p className="note">
          Installs to <code>/usr/local/bin</code> for all users. Requires{' '}
          <code>sudo</code> and skips shell config changes.
        </p>
        <pre>
          <code>{`INSTALL_DIR=/usr/local/bin curl -fsSL https://raw.githubusercontent.com/Tarun-Elango/devbox-cli/main/scripts/install.sh | sudo bash`}</code>
        </pre>
      </div>

      <div className="card">
        <h2>Build from source</h2>
        <pre>
          <code>{`git clone https://github.com/Tarun-Elango/devbox-cli.git
cd devbox-cli
go build -o devbox .`}</code>
        </pre>
        <p className="note">
          Install to <code>$GOPATH/bin</code> with <code>go install .</code> (binary
          name is <code>devbox-cli</code> from the module name).
        </p>
      </div>
    </DocPage>
  )
}
