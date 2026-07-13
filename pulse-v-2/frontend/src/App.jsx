import { Routes, Route, Navigate } from 'react-router-dom'
import { useAuth } from './hooks/useAuth'
import ProtectedRoute from './components/Auth/ProtectedRoute'
import LoadingSpinner from './components/Common/LoadingSpinner'

// Страницы
import LoginPage from './pages/LoginPage'
import DashboardPage from './pages/DashboardPage'
import ProfilePage from './pages/ProfilePage'
import AlertsPage from './pages/AlertsPage'
import AdminPage from './pages/AdminPage'
import AdminUsersPage from './pages/AdminUsersPage'
import AdminRolesPage from './pages/AdminRolesPage'
import AdminIntegrationsPage from './pages/AdminIntegrationsPage'
import AdminAuditPage from './pages/AdminAuditPage'
import KPIDetailPage from './pages/KPIDetailPage'
import ReportsPage from './pages/ReportsPage'
import SettingsPage from './pages/SettingsPage'

function App() {
  const { isAuthenticated, user, loading } = useAuth()

  if (loading) {
    return <LoadingSpinner fullScreen />
  }

  return (
    <Routes>
      {/* Публичные */}
      <Route path="/login" element={
        isAuthenticated ? <Navigate to="/dashboard" /> : <LoginPage />
      } />

      {/* Защищённые */}
      <Route path="/" element={<Navigate to="/dashboard" />} />
      
      <Route path="/dashboard" element={
        <ProtectedRoute>
          <DashboardPage />
        </ProtectedRoute>
      } />

      <Route path="/profile" element={
        <ProtectedRoute>
          <ProfilePage />
        </ProtectedRoute>
      } />

      <Route path="/alerts" element={
        <ProtectedRoute>
          <AlertsPage />
        </ProtectedRoute>
      } />

      <Route path="/kpi/:id" element={
        <ProtectedRoute>
          <KPIDetailPage />
        </ProtectedRoute>
      } />

      <Route path="/reports" element={
        <ProtectedRoute requiredRole={['analyst', 'manager', 'admin']}>
          <ReportsPage />
        </ProtectedRoute>
      } />

      {/* Админ-панель (только admin) */}
      <Route path="/admin" element={
        <ProtectedRoute requiredRole={['admin']}>
          <AdminPage />
        </ProtectedRoute>
      } />

      <Route path="/admin/users" element={
        <ProtectedRoute requiredRole={['admin']}>
          <AdminUsersPage />
        </ProtectedRoute>
      } />

      <Route path="/admin/roles" element={
        <ProtectedRoute requiredRole={['admin']}>
          <AdminRolesPage />
        </ProtectedRoute>
      } />

      <Route path="/admin/integrations" element={
        <ProtectedRoute requiredRole={['admin']}>
          <AdminIntegrationsPage />
        </ProtectedRoute>
      } />

      <Route path="/admin/audit" element={
        <ProtectedRoute requiredRole={['admin']}>
          <AdminAuditPage />
        </ProtectedRoute>
      } />

      <Route path="/settings" element={
        <ProtectedRoute requiredRole={['admin']}>
          <SettingsPage />
        </ProtectedRoute>
      } />

      {/* 404 */}
      <Route path="*" element={<Navigate to="/dashboard" />} />
    </Routes>
  )
}

export default App