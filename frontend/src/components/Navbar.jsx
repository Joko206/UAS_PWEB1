// src/components/Navbar.jsx
import { Link } from 'react-router-dom'
import { useAuth } from '../pages/AuthContext'

function Navbar() {
  const { token, logout } = useAuth()

  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light">
      <div className="container">
        <Link className="navbar-brand" to="/">My App</Link>
        <div className="collapse navbar-collapse">
          <div className="navbar-nav">
            {!token && (
              <>
                <Link className="nav-link" to="/login">Login</Link>
                <Link className="nav-link" to="/register">Register</Link>
              </>
            )}
            {token && (
              <>
                <Link className="nav-link" to="/protected">Protected</Link>
                <button className="nav-link btn btn-link" onClick={logout}>Logout</button>
              </>
            )}
          </div>
        </div>
      </div>
    </nav>
  )
}

export default Navbar