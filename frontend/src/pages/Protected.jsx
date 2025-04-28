// src/pages/Protected.jsx
import { useEffect, useState } from 'react'
import { useAuth } from './AuthContext'
import axios from 'axios'

function Protected() {
  const { token, logout } = useAuth()
  const [data, setData] = useState(null)
  const [error, setError] = useState(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('http://localhost:8080/protected', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        
        // Debugging: Log response data
        console.log('API Response:', response.data)
        
        // Pastikan struktur data sesuai
        if (response.data && response.data.user) {
          setData({
            message: response.data.message,
            user: {
              id: response.data.user.ID || response.data.user.id,
              username: response.data.user.Username || response.data.user.username,
              email: response.data.user.Email || response.data.user.email,
              createdAt: response.data.user.CreatedAt || response.data.user.createdAt
            }
          })
        } else {
          throw new Error('Invalid response structure')
        }
      } catch (error) {
        console.error('Error fetching data:', error)
        setError(error.response?.data?.error || error.message)
        logout()
      } finally {
        setLoading(false)
      }
    }
    
    if (token) {
      fetchData()
    } else {
      logout()
    }
  }, [token, logout])

  if (loading) {
    return (
      <div className="text-center mt-5">
        <div className="spinner-border text-primary" role="status">
          <span className="visually-hidden">Loading...</span>
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="alert alert-danger mt-5 mx-auto" style={{ maxWidth: '500px' }}>
        <h4 className="alert-heading">Error!</h4>
        <p>{error}</p>
        <button className="btn btn-secondary" onClick={logout}>
          Kembali ke Login
        </button>
      </div>
    )
  }

  return (
    <div className="container mt-5">
      <div className="card shadow mx-auto" style={{ maxWidth: '600px' }}>
        <div className="card-header bg-primary text-white">
          <h2 className="mb-0">Protected Page</h2>
        </div>
        <div className="card-body">
          {data ? (
            <>
              <p className="lead">{data.message}</p>
              <div className="mt-4">
                <h4>User Details:</h4>
                <ul className="list-group">
                  <li className="list-group-item">
                    <strong>ID:</strong> {data.user.id}
                  </li>
                  <li className="list-group-item">
                    <strong>Username:</strong> {data.user.username}
                  </li>
                  <li className="list-group-item">
                    <strong>Email:</strong> {data.user.email}
                  </li>
                  <li className="list-group-item">
                    <strong>Created At:</strong> 
                    {new Date(data.user.createdAt).toLocaleDateString()}
                  </li>
                </ul>
              </div>
            </>
          ) : (
            <p className="text-danger">No data available</p>
          )}
        </div>
      </div>
    </div>
  )
}

export default Protected