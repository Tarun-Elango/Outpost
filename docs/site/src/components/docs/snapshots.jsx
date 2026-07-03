import DocPage from './doc-page'

export default function SnapshotsDoc() {
  return (
    <DocPage title="Snapshots & templates">
      <div className="card">
        <h2>Snapshots</h2>
        <ul>
          <li>
            <code>devbox snapshot</code> — list all snapshots
          </li>
          <li>
            <code>devbox snapshot create {'<id-or-name>'} {'<name>'}</code> — snapshot a
            box
          </li>
          <li>
            <code>devbox snapshot ls {'<amiId-or-name>'}</code> — show snapshot details
          </li>
          <li>
            <code>devbox snapshot delete {'<amiId-or-name>'}</code> — delete a snapshot
          </li>
        </ul>
        <p className="note">
          Restore when creating a box:{' '}
          <code>devbox create {'<name>'} --from {'<amiId|name>'}</code>
        </p>
      </div>

      <div className="card">
        <h2>Templates</h2>
        <ul>
          <li>
            <code>devbox template</code> — list available templates
          </li>
          <li>
            <code>devbox template new {'<name>'} [command]</code> — create a template
          </li>
          <li>
            <code>devbox template search {'<query>'}</code> — search by name
          </li>
          <li>
            <code>devbox template rename {'<name>'} {'<new-name>'}</code> — rename
          </li>
          <li>
            <code>devbox template delete {'<name>'}</code> — delete a template
          </li>
        </ul>
        <p className="note">
          Use templates at create time:{' '}
          <code>devbox create --template {'<template>'} {'<name>'}</code>
        </p>
      </div>
    </DocPage>
  )
}
