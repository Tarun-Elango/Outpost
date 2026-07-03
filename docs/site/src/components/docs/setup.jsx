import { Link } from 'react-router-dom'
import DocPage from './doc-page'

export default function SetupDoc() {
  return (
    <DocPage title="Setup">
      <div className="card">
        <h2>Interactive wizard</h2>
        <p className="note">
          Run once after install to save AWS credentials and region locally.
        </p>
        <pre>
          <code>devbox setup</code>
        </pre>
        <p className="note">
          Credentials are stored in <code>~/.devbox/config.json</code> (mode 0600).
          Use a dedicated IAM user — see{' '}
          <Link to="/docs/config">AWS &amp; local config</Link> for IAM steps.
        </p>
      </div>

      <div className="card">
        <h2>Related commands</h2>
        <ul>
          <li>
            <code>devbox health</code> — check config, credentials, region, and database
          </li>
          <li>
            <code>devbox clear-creds</code> — remove saved AWS credentials
          </li>
          <li>
            <code>devbox update</code> — check for and install a newer CLI release
          </li>
          <li>
            <code>devbox version</code> — show the installed CLI version
          </li>
        </ul>
      </div>
    </DocPage>
  )
}
