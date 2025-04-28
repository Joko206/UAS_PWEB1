// src/pages/Register.jsx
import { useState } from 'react'
import { useAuth } from './AuthContext'
import { Link } from 'react-router-dom'
import { UserIcon, EnvelopeIcon, LockClosedIcon } from '@heroicons/react/24/outline'

export default function Register() {
  const { register } = useAuth()
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: ''
  })

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      await register(formData.username, formData.email, formData.password)
    } catch (error) {
      // Handle error
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-900 via-blue-900 to-slate-900">
      <div className="max-w-md w-full mx-4 bg-white/5 rounded-2xl p-8 backdrop-blur-lg border border-white/10 shadow-xl transform transition-all hover:shadow-2xl">
        <div className="text-center mb-8">
          <h2 className="text-3xl font-bold text-white mb-2 animate-fade-in-up">
            Create Account
          </h2>
          <p className="text-gray-400">Join our community today</p>
        </div>

        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label className="block text-sm font-medium text-gray-300 mb-2">
              Username
            </label>
            <div className="relative">
              <UserIcon className="h-5 w-5 absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
              <input
                type="text"
                className="w-full pl-10 pr-4 py-3 bg-white/5 rounded-lg border border-white/10 
                  focus:border-blue-400 focus:ring-2 focus:ring-blue-400/30 text-white 
                  transition-all duration-300 placeholder-gray-500"
                placeholder="Enter username"
                value={formData.username}
                onChange={(e) => setFormData({...formData, username: e.target.value})}
              />
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-300 mb-2">
              Email
            </label>
            <div className="relative">
              <EnvelopeIcon className="h-5 w-5 absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
              <input
                type="email"
                className="w-full pl-10 pr-4 py-3 bg-white/5 rounded-lg border border-white/10 
                  focus:border-blue-400 focus:ring-2 focus:ring-blue-400/30 text-white 
                  transition-all duration-300 placeholder-gray-500"
                placeholder="Enter email"
                value={formData.email}
                onChange={(e) => setFormData({...formData, email: e.target.value})}
              />
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-300 mb-2">
              Password
            </label>
            <div className="relative">
              <LockClosedIcon className="h-5 w-5 absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
              <input
                type="password"
                className="w-full pl-10 pr-4 py-3 bg-white/5 rounded-lg border border-white/10 
                  focus:border-blue-400 focus:ring-2 focus:ring-blue-400/30 text-white 
                  transition-all duration-300 placeholder-gray-500"
                placeholder="Enter password"
                value={formData.password}
                onChange={(e) => setFormData({...formData, password: e.target.value})}
              />
            </div>
          </div>

          <button
            type="submit"
            className="w-full bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 
              hover:to-purple-700 text-white py-3 px-6 rounded-lg font-medium transition-all 
              duration-300 transform hover:scale-[1.02] shadow-lg hover:shadow-blue-500/30 
              flex items-center justify-center space-x-2"
          >
            <span>Sign Up</span>
          </button>
        </form>

        <p className="mt-6 text-center text-gray-400">
          Already have an account?{' '}
          <Link 
            to="/login" 
            className="text-blue-400 hover:text-blue-300 transition-colors duration-300"
          >
            Login here
          </Link>
        </p>
      </div>
    </div>
  )
}