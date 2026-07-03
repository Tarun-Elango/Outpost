import DocPage from './doc-page'

export default function ConfigDoc() {
  return (
    <DocPage title="AWS & local config">
      <div className="card">
        <h2>Local config</h2>
        <p className="note">
          Credentials and tokens live in <code>~/.devbox/config.json</code> (mode 0600).
        </p>
        <ul>
          <li>Do not sync <code>~/.devbox</code> via dotfiles, iCloud, Dropbox, or Git</li>
          <li>Use a dedicated IAM user for AWS keys</li>
          <li>
            Run <code>devbox health</code> to verify config, credentials, region, and
            database
          </li>
        </ul>
      </div>

      <div className="card">
        <h2>AWS setup</h2>
        <ol>
          <li>
            IAM console → <strong>Users</strong> → <strong>Create user</strong> (e.g.{' '}
            <code>devbox-cli</code>)
          </li>
          <li>
            Attach <code>AmazonEC2FullAccess</code> directly and create the user
          </li>
          <li>
            Open the user → <strong>Security credentials</strong> → create an access key
            (choose <strong>Local code</strong>)
          </li>
          <li>Copy the access key ID and secret (secret shown only once)</li>
          <li>
            Save in devbox: <code>devbox setup</code>
          </li>
        </ol>
      </div>
    </DocPage>
  )
}
