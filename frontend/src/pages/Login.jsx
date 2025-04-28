// src/pages/Login.jsx
import { useState } from 'react'
import { useAuth } from './AuthContext'
import { EnvelopeIcon, LockClosedIcon } from '@heroicons/react/24/outline'

export default function Login() {
  const { login } = useAuth()
  const [formData, setFormData] = useState({
    username: '',
    password: ''
  })

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      await login(formData.username, formData.password)
    } catch (error) {
      alert("Something went wrong!")
      console.error('Error:', error)
    }
  }

  return (
    <div className="max-w-md mt-10 mx-auto bg-white/5 rounded-2xl p-8 backdrop-blur-lg border border-white/10 shadow-xl">
      <div className="text-center mb-8">
        <h2 className="text-3xl font-bold text-white mb-2 animate-fade-in-up">Welcome Back</h2>
        <p className="text-gray-400">Please sign in to continue</p>
      </div>
      
      <form onSubmit={handleSubmit} className="space-y-6">
        <div>
          <label className="block text-sm font-medium text-gray-300 mb-2">Username</label>
          <div className="relative">
            <EnvelopeIcon className="h-5 w-5 absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
            <input
              type="text"
              className="w-full pl-10 pr-4 py-3 bg-white/5 rounded-lg border border-white/10 
                focus:border-blue-400 focus:ring-2 focus:ring-blue-400/30 text-white 
                transition-all duration-300"
              value={formData.username}
              onChange={(e) => setFormData({...formData, username: e.target.value})}
            />
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-300 mb-2">Password</label>
          <div className="relative">
            <LockClosedIcon className="h-5 w-5 absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
            <input
              type="password"
              className="w-full pl-10 pr-4 py-3 bg-white/5 rounded-lg border border-white/10 
                focus:border-blue-400 focus:ring-2 focus:ring-blue-400/30 text-white 
                transition-all duration-300"
              value={formData.password}
              onChange={(e) => setFormData({...formData, password: e.target.value})}
            />
          </div>
        </div>

        <button
          type="submit"
          className="w-full bg-blue-600 hover:bg-blue-700 text-white py-3 px-6 rounded-lg
            font-medium transition-all duration-300 transform hover:scale-[1.02] shadow-lg
            hover:shadow-blue-500/30 flex items-center justify-center space-x-2"
        >
          <span>Sign In</span>
        </button>
      </form>
    </div>
  )
}