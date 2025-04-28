// src/App.jsx
import { Routes, Route } from 'react-router-dom'
import { AuthProvider } from './pages/AuthContext'
import Home from './pages/Home'
import Login from './pages/Login'
import Register from './pages/Register'
import Protected from './pages/Protected'
import Layout from './components/Layout'

function App() {
  return (
    <AuthProvider>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/protected" element={<Protected />} />
        </Routes>
      </Layout>
    </AuthProvider>
  )
}

export default App