import { Link } from 'react-router-dom'

const topics = [
  {
    to: '/docs/install',
    title: 'Installation',
    description: 'curl script, verify install, system-wide install',
  },
  {
    to: '/docs/setup',
    title: 'Setup',
    description: 'configure AWS credentials with devbox setup',
  },
  {
    to: '/docs/boxes',
    title: 'Managing boxes',
    description: 'create, list, start, stop, delete, resize',
  },
  {
    to: '/docs/connect',
    title: 'Connect & transfer',
    description: 'ssh, cp, sync, exec, port forwarding',
  },
  {
    to: '/docs/snapshots',
    title: 'Snapshots & templates',
    description: 'save images, restore boxes, custom templates',
  },
  {
    to: '/docs/config',
    title: 'AWS & local config',
    description: 'IAM setup, ~/.devbox, health checks',
  },
]

export default function DocsIndexPage() {
  return (
    <>
      <h1>Documentation</h1>
      <p className="tagline">
        Guides for installing, configuring, and using devbox.
      </p>

      <div className="card">
        <h2>Topics</h2>
        <ul>
          {topics.map(({ to, title, description }) => (
            <li key={to}>
              <Link to={to}>{title}</Link> — {description}
            </li>
          ))}
        </ul>
      </div>
    </>
  )
}
