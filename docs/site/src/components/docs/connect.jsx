import DocPage from './doc-page'

export default function ConnectDoc() {
  return (
    <DocPage title="Connect & transfer">
      <div className="card">
        <h2>SSH</h2>
        <pre>
          <code>devbox ssh {'<id-or-name>'}</code>
        </pre>
        <p className="note">
          Default key: <code>~/.ssh/id_ed25519</code>. Pass <code>-i path/to/key</code>{' '}
          to override. Use <code>--</code> before native ssh options, e.g.{' '}
          <code>devbox ssh mybox -- -L 8080:localhost:8080</code>.
        </p>
      </div>

      <div className="card">
        <h2>Copy &amp; sync</h2>
        <ul>
          <li>
            <code>devbox cp ./file.go mybox:/home/ec2-user/app/</code> — copy a file to
            or from a box
          </li>
          <li>
            <code>devbox sync ./src mybox:/home/ec2-user/app/</code> — sync a directory
          </li>
          <li>
            <code>devbox sync --delete ./src mybox:/path</code> — sync and remove
            destination files missing from source
          </li>
        </ul>
      </div>

      <div className="card">
        <h2>Remote commands</h2>
        <ul>
          <li>
            <code>devbox exec mybox -- uname -a</code> — run a one-off command
          </li>
          <li>
            <code>devbox exec -s mybox -- &quot;cd app &amp;&amp; make&quot;</code> — run
            as a shell snippet
          </li>
          <li>
            <code>devbox exec -t mybox -- sudo apt update</code> — allocate a TTY for
            sudo or interactive commands
          </li>
          <li>
            <code>devbox forward mybox 8080</code> — forward a port from a box
          </li>
        </ul>
      </div>
    </DocPage>
  )
}
