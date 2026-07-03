import { Link } from 'react-router-dom'

function scrollToTop(event) {
  event.preventDefault()
  window.scrollTo(0, 0)
}

export default function DocPage({ title, children }) {
  return (
    <>
      <h1>Documentation</h1>
      <p className="tagline">
        <Link to="/docs">← Topics</Link>
        {title ? ` · ${title}` : null}
      </p>
      {children}
      <p className="go-top">
        <a href="#" onClick={scrollToTop}>
          ↑ Top
        </a>
      </p>
    </>
  )
}
