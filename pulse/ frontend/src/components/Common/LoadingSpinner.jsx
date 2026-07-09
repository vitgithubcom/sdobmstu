function LoadingSpinner({ size = 'md', color = 'blue', fullScreen = false }) {
  const sizeClasses = {
    sm: 'h-6 w-6 border-2',
    md: 'h-12 w-12 border-4',
    lg: 'h-16 w-16 border-4',
  }

  const colorClasses = {
    blue: 'border-blue-600',
    white: 'border-white',
    gray: 'border-gray-600',
  }

  const spinner = (
    <div className={`animate-spin rounded-full ${sizeClasses[size]} ${colorClasses[color]} border-t-transparent`} />
  )

  if (fullScreen) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        {spinner}
      </div>
    )
  }

  return spinner
}

export default LoadingSpinner