import { Link } from 'react-router-dom'
import HowToPage from './how-to-page'

export default function VscodeSshHowTo() {
  return (
    <HowToPage title="VS Code & SSH without the CLI">
      <div className="card">
        <p>
          You can manage boxes with devbox and still connect with plain <code>ssh</code> or
          VS Code Remote-SSH. Every time you create, start, or rename a box, devbox writes
          a host entry to <code>~/.ssh/config</code> so standard SSH tools work without
          calling <code>devbox ssh</code>.
        </p>
      </div>

      <div className="card">
        <h2>What devbox adds to your SSH config</h2>
        <p>
          After <code>devbox create mybox</code>, look for a block named{' '}
          <code>devbox-mybox</code>:
        </p>
        <pre>
          <code>{`Host devbox-mybox
    HostName 203.0.113.42
    User ec2-user
    IdentityFile ~/.ssh/id_ed25519
    StrictHostKeyChecking accept-new`}</code>
        </pre>
        <p className="note">
          The <code>HostName</code> is updated automatically when a box gets a new public
          IP (for example after <code>devbox start</code>). Renaming a box rewrites the
          host alias (<code>devbox-old</code> → <code>devbox-new</code>).
        </p>
      </div>

      <div className="card">
        <h2>Prerequisites</h2>
        <ol>
          <li>
            An SSH key at <code>~/.ssh/id_ed25519</code> (devbox uses this when creating
            boxes). Generate one with <code>ssh-keygen -t ed25519</code> if needed.
          </li>
          <li>
            A running box — use <code>devbox start {'<name>'}</code> if it is stopped.
          </li>
          <li>
            The box finished provisioning. Run <code>devbox status {'<name>'}</code> or
            connect once with <code>devbox ssh {'<name>'}</code> to wait for setup.
          </li>
        </ol>
      </div>

      <div className="card">
        <h2>Connect with plain SSH</h2>
        <pre>
          <code>ssh devbox-mybox</code>
        </pre>
        <p>
          Use any SSH option your client supports — port forwarding, remote commands,{' '}
          <code>scp</code>, <code>rsync</code>, and so on:
        </p>
        <pre>
          <code>{`scp ./app.go devbox-mybox:/home/ec2-user/
ssh devbox-mybox -L 8080:localhost:8080`}</code>
        </pre>
        <p className="note">
          For copy and sync helpers built into devbox, see{' '}
          <Link to="/how-tos/transfer">Transfer data and files</Link>.
        </p>
      </div>

      <div className="card">
        <h2>Connect with VS Code</h2>
        <ol>
          <li>
            Install the{' '}
            <a
              href="https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-ssh"
              rel="noreferrer"
              target="_blank"
            >
              Remote - SSH
            </a>{' '}
            extension.
          </li>
          <li>
            Open the Command Palette (<kbd>F1</kbd> or <kbd>Ctrl+Shift+P</kbd>) →{' '}
            <strong>Remote-SSH: Connect to Host…</strong>
          </li>
          <li>
            Pick <code>devbox-mybox</code> from the list (VS Code reads{' '}
            <code>~/.ssh/config</code>).
          </li>
          <li>
            Open a folder on the remote machine, e.g.{' '}
            <code>/home/ec2-user</code>.
          </li>
        </ol>
        <p className="note">
          You still use devbox for lifecycle tasks — create, start, stop, delete, resize —
          but day-to-day editing and terminal work can stay in VS Code over SSH.
        </p>
      </div>

      <div className="card">
        <h2>When to use <code>devbox ssh</code> instead</h2>
        <ul>
          <li>
            First connection while the box is still provisioning —{' '}
            <code>devbox ssh</code> polls until the instance is ready.
          </li>
          <li>
            Passing a non-default key: <code>devbox ssh mybox -i path/to/key</code>
          </li>
          <li>
            GitHub agent forwarding via <code>devbox git-sync</code> — see{' '}
            <Link to="/how-tos/github-sync">Sync GitHub account</Link>
          </li>
        </ul>
      </div>
    </HowToPage>
  )
}
