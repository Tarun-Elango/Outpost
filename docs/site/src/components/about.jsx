import { Link } from 'react-router-dom'

export default function AboutPage() {
    return (
      <>
        <div className="page-title">
          <h1>Outpost</h1>
          <a href="https://github.com/Tarun-Elango" className="byline">
            by Tarun-Elango
          </a>
        </div>
        <p className="tagline">
          Manage remote dev boxes from the CLI — provision, connect, sync, and destroy
          them using your own AWS account (BYOK).
        </p>

        <div className="card card-purpose">
          <h2>Why outpost?</h2>
          <figure className="purpose-diagram">
            <div className="purpose-flow">
              <section className="purpose-stage">
                <div className="purpose-stage-label">
                  <span>1</span>
                  Your laptop
                </div>
                <div className="purpose-terminal">
                  <div className="purpose-terminal-bar" aria-hidden="true">
                    <i />
                    <i />
                    <i />
                    <small>terminal</small>
                  </div>
                  <code>
                    <span>$ outpost create dev</span>
                    <span className="purpose-terminal-success">✓ dev box ready</span>
                    <span>$ outpost ssh dev</span>
                  </code>
                </div>
                <p>Run one command. Credentials and config stay here.</p>
              </section>

              <div className="purpose-connection" aria-hidden="true">
                <strong>OUTPOST</strong>
                <span>AWS API</span>
                <div>→</div>
              </div>

              <section className="purpose-stage purpose-stage-cloud">
                <div className="purpose-stage-label">
                  <span>2</span>
                  Your AWS account
                </div>
                <div className="purpose-boxes">
                  <div className="purpose-box">
                    <span className="purpose-status" />
                    <strong>dev</strong>
                    <small>EC2 · running</small>
                  </div>
                  <div className="purpose-box">
                    <span className="purpose-status purpose-status-idle" />
                    <strong>project-b</strong>
                    <small>EC2 · stopped</small>
                  </div>
                </div>
                <p>Private dev machines you own and control.</p>
              </section>
            </div>

            <div className="purpose-lifecycle" aria-label="Dev box lifecycle">
              <span>create</span>
              <b>→</b>
              <span>start</span>
              <b>→</b>
              <span>ssh + sync</span>
              <b>→</b>
              <span>stop</span>
              <b>→</b>
              <span>delete</span>
            </div>
            <figcaption>
              Outpost is the remote control for disposable development machines in
              your own cloud.
            </figcaption>
          </figure>
          <ul>
            <li>
              <strong>Dedicated dev machine on the cloud</strong> — your own EC2
              instance, separate from production and your daily driver
            </li>
            <li>
              <strong>Smaller blast radius</strong> — experiments, tooling, and
              dependencies stay off your main machine
            </li>
            <li>
              <strong>Fast lifecycle</strong> — create, use, and tear down boxes
              in minutes
            </li>
            <li>
              <strong>Reproducible setups</strong> — spin up consistent environments
              from templates
            </li>
            <li>
              <strong>Commands that simplify daily work</strong> —{' '}
              <code>ssh</code>, <code>sync</code>, <code>idle-stop</code>,{' '}
              <code>git-sync</code>, <code>import</code>, and more
            </li>
            <li>
              <strong>Secure by default</strong> — AWS credentials and config stored
              locally on your machine, and the code is open source and available on GitHub.
            </li>
          </ul>
        </div>

        <div className="card">
          <h2>Requirements</h2>
          <ul>
            <li>macOS or Linux</li>
            <li>Your own AWS account (BYOK)</li>
            <li>
              On <code>PATH</code> you might need: <code>ssh</code> for SSH commands, <code>scp</code> for
              copy, <code>rsync</code> for folder sync, and <code>ssh-agent</code> for
              GitHub sync between your machine and a box
            </li>
          </ul>
          <p className="note">
            outpost cli runs on your machine and uses your AWS account — no shared cloud,
            no hosted credentials. Run <code>outpost setup</code> to save keys locally in{' '}
            <code>~/.outpost/</code>.
          </p>
        </div>

        <div className="card">
          <h2>Quick install</h2>
          <p className="note">
            Every push to <code>main</code> publishes binaries to the{' '}
            <a href="https://github.com/Tarun-Elango/Outpost/releases/tag/latest">
              latest release
            </a>, run the following command:

          </p>
          <pre>
            <code>{`curl -fsSL https://raw.githubusercontent.com/Tarun-Elango/Outpost/latest/scripts/install.sh | bash`}</code>
          </pre>
          <p className="note">
            Verify with the command <code>outpost ls</code>.
          </p>
          <p className="note">
            If that worked, you&apos;re done — skip the sections below. They&apos;re
            optional alternatives for pinning a version or installing system-wide.
          </p>

          <h3>
            Pin a specific version — To install a particular release instead of{' '}
            <code>latest</code>, set <code>RELEASE_TAG</code> on <code>bash</code> (not
            on <code>curl</code>):
          </h3>
          <pre>
            <code>{`curl -fsSL https://raw.githubusercontent.com/Tarun-Elango/Outpost/latest/scripts/install.sh | RELEASE_TAG=v0.7.0 bash`}</code>
          </pre>

          <h3>
            Install system-wide — To install to <code>/usr/local/bin</code> (requires{' '}
            <code>sudo</code>, no shell config changes):
          </h3>
          <pre>
            <code>{`INSTALL_DIR=/usr/local/bin curl -fsSL https://raw.githubusercontent.com/Tarun-Elango/Outpost/latest/scripts/install.sh | sudo bash`}</code>
          </pre>

         
        </div>

        <div className="card">
          <h2>Common commands</h2>
          <ul>
            <li>
              <code>outpost setup</code> — configure AWS credentials{' '}
              <Link to="/docs/setup">(see how to get AWS credentials)</Link>
            </li>
            <li>
              <code>outpost create {'<name>'}</code> — create a box
            </li>
            <li>
              <code>outpost ls</code> — list boxes
            </li>
            <li>
              <code>outpost ssh {'<name>'}</code> — connect via SSH
            </li>
            <li>
              <code>outpost import {'<url>'}</code> — import existing instance from aws.
            </li>
          </ul>
          <p className="note">
            <Link to="/docs/commands">See all commands</Link>
          </p>
        </div>

        <div className="card">
          <h2>Links</h2>
          <ul>
            <li>
              <a href="https://github.com/Tarun-Elango/Outpost">GitHub repository</a>
            </li>
            <li>
              <a href="https://github.com/Tarun-Elango/Outpost/releases">Releases</a>
            </li>
            <li>
              <a href="https://github.com/Tarun-Elango/outpost/blob/main/readme.md">
                Full README
              </a>
            </li>
          </ul>
        </div>
      </>
    )
  }
