// src/components/Layout.jsx
import { Link, useLocation } from 'react-router-dom'
import { useAuth } from '../pages/AuthContext'
import { HomeIcon, UserCircleIcon, ArrowRightOnRectangleIcon } from '@heroicons/react/24/outline'

export default function Layout({ children }) {
  const { token, logout } = useAuth()
  const location = useLocation()

  const navLinks = [
    { name: 'Home', path: '/', icon: HomeIcon },
    { name: 'Login', path: '/login', icon: UserCircleIcon, protected: false },
    { name: 'Register', path: '/register', icon: ArrowRightOnRectangleIcon, protected: false },
  ]

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-blue-900 to-slate-900">
      <nav className="fixed top-0 w-full bg-white/5 backdrop-blur-lg border-b border-slate-800 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex items-center justify-between h-16">
            <Link to="/" className="flex items-center space-x-2">
              <span className="text-2xl font-bold bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
                AuthApp
              </span>
            </Link>
            
            <div className="flex space-x-8">
              {navLinks.map((link) => (
                (!link.protected || token) && (
                  <Link
                    key={link.name}
                    to={link.path}
                    className={`flex items-center space-x-2 ${
                      location.pathname === link.path 
                        ? 'text-blue-400'
                        : 'text-gray-300 hover:text-blue-300'
                    } transition-all duration-300`}
                  >
                    <link.icon className="h-5 w-5" />
                    <span className="text-sm font-medium">{link.name}</span>
                  </Link>
                )
              ))}
              {token && (
                <button
                  onClick={logout}
                  className="flex items-center space-x-2 text-red-400 hover:text-red-300 transition-all"
                >
                  <ArrowRightOnRectangleIcon className="h-5 w-5" />
                  <span className="text-sm font-medium">Logout</span>
                </button>
              )}
            </div>
          </div>
        </div>
      </nav>

      <main className="pt-20 pb-12">{children}</main>
    </div>
  )
}