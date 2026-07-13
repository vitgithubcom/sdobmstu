import { Component } from 'react'

class ErrorBoundary extends Component {
  constructor(props) {
    super(props)
    this.state = { hasError: false, error: null, errorInfo: null }
  }

  static getDerivedStateFromError(error) {
    return { hasError: true }
  }

  componentDidCatch(error, errorInfo) {
    console.error('ErrorBoundary caught an error:', error, errorInfo)
    this.setState({ error, errorInfo })
  }

  handleReload = () => {
    window.location.reload()
  }

  handleGoBack = () => {
    window.history.back()
  }

  render() {
    if (this.state.hasError) {
      return (
        <div className="min-h-screen flex items-center justify-center bg-gray-50 px-4">
          <div className="max-w-md w-full bg-white rounded-2xl shadow-xl p-8 text-center">
            <div className="text-6xl mb-4">😵</div>
            <h1 className="text-2xl font-bold text-gray-800 mb-2">Что-то пошло не так</h1>
            <p className="text-gray-600 text-sm mb-6">
              Произошла ошибка при загрузке страницы. Попробуйте перезагрузить или вернуться назад.
            </p>
            
            {process.env.NODE_ENV === 'development' && this.state.error && (
              <div className="mb-4 text-left bg-red-50 rounded-xl p-4 text-xs font-mono text-red-700 overflow-auto max-h-40">
                <div className="font-bold mb-1">Ошибка:</div>
                <div>{this.state.error.toString()}</div>
                {this.state.errorInfo && (
                  <>
                    <div className="font-bold mt-2 mb-1">Стек:</div>
                    <div>{this.state.errorInfo.componentStack}</div>
                  </>
                )}
              </div>
            )}

            <div className="flex gap-3 justify-center">
              <button
                onClick={this.handleGoBack}
                className="px-4 py-2 border border-gray-300 rounded-xl text-sm text-gray-700 hover:bg-gray-50 transition"
              >
                ← Назад
              </button>
              <button
                onClick={this.handleReload}
                className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-xl text-sm transition"
              >
                ⟳ Перезагрузить
              </button>
            </div>
          </div>
        </div>
      )
    }

    return this.props.children
  }
}

export default ErrorBoundary