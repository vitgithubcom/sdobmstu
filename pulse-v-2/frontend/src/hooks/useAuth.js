import { useState, useEffect } from 'react'

export function useAuth() {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [user, setUser] = useState(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const token = localStorage.getItem('pulse_token')
    const savedUser = localStorage.getItem('pulse_user')
    
    if (token && savedUser) {
      setIsAuthenticated(true)
      setUser(JSON.parse(savedUser))
    }
    setLoading(false)
  }, [])

  const login = (token, userData) => {
    localStorage.setItem('pulse_token', token)
    localStorage.setItem('pulse_user', JSON.stringify(userData))
    setIsAuthenticated(true)
    setUser(userData)
  }

  const logout = () => {
    localStorage.removeItem('pulse_token')
    localStorage.removeItem('pulse_user')
    setIsAuthenticated(false)
    setUser(null)
  }

  return { isAuthenticated, user, loading, login, logout }
}