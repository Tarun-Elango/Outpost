import DocPage from './doc-page'

export default function BoxesDoc() {
  return (
    <DocPage title="Managing boxes">
      <div className="card">
        <h2>Create &amp; list</h2>
        <ul>
          <li>
            <code>devbox create {'<name>'}</code> — create a new box
          </li>
          <li>
            <code>devbox create --template {'<template>'} {'<name>'}</code> — create
            from a template
          </li>
          <li>
            <code>devbox create {'<name>'} --from {'<amiId|name>'}</code> — restore
            from a snapshot
          </li>
          <li>
            <code>devbox ls</code> — list all boxes
          </li>
          <li>
            <code>devbox status {'<id-or-name>'}</code> — show box details
          </li>
          <li>
            <code>devbox rename {'<id-or-name>'} {'<new-name>'}</code> — rename a box
          </li>
        </ul>
      </div>

      <div className="card">
        <h2>Power &amp; lifecycle</h2>
        <ul>
          <li>
            <code>devbox start {'<id-or-name>'}</code> — start a stopped box
          </li>
          <li>
            <code>devbox stop {'<id-or-name>'}</code> — stop a running box
          </li>
          <li>
            <code>devbox restart {'<id-or-name>'}</code> — reboot a running box
          </li>
          <li>
            <code>devbox delete {'<id-or-name>'}</code> — delete a box
          </li>
          <li>
            <code>devbox resize {'<id-or-name>'}</code> — resize instance type or
            root disk (box must be stopped)
          </li>
        </ul>
      </div>

      <div className="card">
        <h2>Idle stop</h2>
        <ul>
          <li>
            <code>devbox idle-stop set {'<id-or-name>'} {'<minutes>'}</code> — stop
            after inactivity
          </li>
          <li>
            <code>devbox idle-stop show {'<id-or-name>'}</code> — show idle-stop
            settings
          </li>
          <li>
            <code>devbox idle-stop update {'<id-or-name>'} {'<minutes>'}</code> — change
            timeout
          </li>
          <li>
            <code>devbox idle-stop delete {'<id-or-name>'}</code> — remove idle-stop
          </li>
        </ul>
      </div>
    </DocPage>
  )
}
