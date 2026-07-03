import { Link } from 'react-router-dom'

function scrollToTop(event) {
  event.preventDefault()
  window.scrollTo(0, 0)
}

export default function HowToPage({ title, children }) {
  return (
    <>
      <h1>How tos</h1>
      <p className="tagline">
        <Link to="/how-tos">← Guides</Link>
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
