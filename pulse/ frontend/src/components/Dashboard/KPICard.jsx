import { Link } from 'react-router-dom'

function KPICard({ kpi }) {
  return (
    <Link to={`/kpi/${kpi.id}`} className="block">
      <div className="kpi-card bg-white p-4 rounded-2xl border border-gray-100 hover:shadow-md cursor-pointer">
        <div className="text-sm text-gray-500 font-medium mb-1">{kpi.title}</div>
        <div className="text-2xl md:text-3xl font-bold tracking-tight">
          {kpi.value} <span className="text-sm font-normal text-gray-400">{kpi.unit}</span>
        </div>
        <div className="flex justify-between items-center mt-2">
          <div className="text-xs text-gray-400">План: {kpi.plan} {kpi.unit}</div>
          <div className={`text-sm font-medium ${kpi.delta > 0 ? 'trend-up' : 'trend-down'}`}>
            {kpi.delta > 0 ? '▲' : '▼'} {Math.abs(kpi.delta)}%
          </div>
        </div>
        {kpi.status === 'critical' && (
          <div className="mt-2 text-[11px] badge-critical inline-block px-2 py-0.5 rounded-full">Критично</div>
        )}
        {kpi.status === 'warning' && (
          <div className="mt-2 text-[11px] badge-warning inline-block px-2 py-0.5 rounded-full">Внимание</div>
        )}
      </div>
    </Link>
  )
}

export default KPICard