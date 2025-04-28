// src/context/AuthContext.jsx
import { createContext, useContext, useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'

const AuthContext = createContext()

export function AuthProvider({ children }) {
  const [token, setToken] = useState(localStorage.getItem('token'))
  const navigate = useNavigate()

  const login = async (username, password) => {
    try {
      const response = await axios.post('http://localhost:8080/login', {
        username,
        password
      })
      
      localStorage.setItem('token', response.data.token)
      setToken(response.data.token)
      navigate('/protected')
    } catch (error) {
      console.error('Login error:', error.response?.data)
      throw error
    }
  }

  const register = async (username, email, password) => {
    try {
      await axios.post('http://localhost:8080/register', {
        username,
        email,
        password
      })
      navigate('/login')
    } catch (error) {
      console.error('Registration error:', error.response?.data)
      throw error
    }
  }

  const logout = () => {
    localStorage.removeItem('token')
    setToken(null)
    navigate('/login')
  }

  return (
    <AuthContext.Provider value={{ token, login, register, logout }}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => useContext(AuthContext)