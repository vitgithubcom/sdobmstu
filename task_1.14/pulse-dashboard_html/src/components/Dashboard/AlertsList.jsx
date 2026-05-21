import { mockAlerts } from '../../data/mockData'

function AlertsList() {
  return (
    <div className="bg-white rounded-2xl p-5 shadow-sm border border-gray-100">
      <div className="flex justify-between mb-3">
        <h3 className="font-semibold text-gray-800">⚠️ Активные тревоги</h3>
        <span className="text-xs text-blue-600">4 новых</span>
      </div>
      <div className="space-y-2 max-h-64 overflow-y-auto">
        {mockAlerts.map(alert => (
          <div key={alert.id} className="flex gap-2 text-sm p-2 rounded-lg hover:bg-gray-50 border-b border-gray-50">
            <div className={`w-1.5 h-1.5 mt-1.5 rounded-full ${
              alert.severity === 'critical' ? 'bg-red-500' : 
              alert.severity === 'warning' ? 'bg-amber-500' : 'bg-blue-400'
            }`} />
            <div className="flex-1">
              <span className="font-mono text-[11px] text-gray-400">{alert.system}</span>
              <div className="text-gray-700 text-[13px]">{alert.msg}</div>
              <span className="text-[10px] text-gray-400">{alert.time}</span>
            </div>
          </div>
        ))}
      </div>
      <button className="text-xs text-blue-600 mt-3 w-full text-center py-1 hover:bg-blue-50 rounded-full">
        Все события →
      </button>
    </div>
  )
}

export default AlertsList