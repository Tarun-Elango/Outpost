import { NavLink } from 'react-router-dom'

export default function SiteHeader() {
  return (
    <nav className="site-header" aria-label="Site">
      <NavLink to="/" end className={({ isActive }) => (isActive ? 'active' : undefined)}>
        About
      </NavLink>
      <NavLink to="/docs" className={({ isActive }) => (isActive ? 'active' : undefined)}>
        Docs
      </NavLink>
      <NavLink to="/how-tos" className={({ isActive }) => (isActive ? 'active' : undefined)}>
        How tos
      </NavLink>
    </nav>
  )
}
