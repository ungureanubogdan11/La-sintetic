import { useEffect, useState } from 'react'
import { ApiError, apiGet } from './lib/api'

interface Health {
  status: string
  database: string
}

type State =
  | { kind: 'loading' }
  | { kind: 'ok'; health: Health }
  | { kind: 'error'; message: string }

function StatusDot({ up }: { up: boolean }) {
  return (
    <span
      className={`inline-block size-2.5 rounded-full ${up ? 'bg-emerald-500' : 'bg-red-500'}`}
      aria-hidden
    />
  )
}

export default function App() {
  const [state, setState] = useState<State>({ kind: 'loading' })

  useEffect(() => {
    apiGet<Health>('/api/health')
      .then((res) => setState({ kind: 'ok', health: res.data }))
      .catch((err: unknown) =>
        setState({
          kind: 'error',
          message: err instanceof ApiError ? err.message : 'API unreachable',
        }),
      )
  }, [])

  return (
    <main className="flex min-h-screen items-center justify-center bg-slate-50 p-6 text-slate-800">
      <section className="w-full max-w-sm rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
        <h1 className="text-2xl font-semibold text-emerald-700">La Sintetic</h1>
        <p className="mt-1 text-sm text-slate-500">
          Organizarea meciurilor de fotbal pe teren sintetic
        </p>

        <dl className="mt-6 space-y-2 text-sm">
          {state.kind === 'loading' && <p className="text-slate-500">Se verifică API-ul…</p>}

          {state.kind === 'error' && (
            <div className="flex items-center gap-2">
              <StatusDot up={false} />
              <span>{state.message}</span>
            </div>
          )}

          {state.kind === 'ok' && (
            <>
              <div className="flex items-center justify-between">
                <dt>API</dt>
                <dd className="flex items-center gap-2">
                  <StatusDot up={state.health.status === 'ok'} />
                  {state.health.status}
                </dd>
              </div>
              <div className="flex items-center justify-between">
                <dt>Bază de date</dt>
                <dd className="flex items-center gap-2">
                  <StatusDot up={state.health.database === 'up'} />
                  {state.health.database}
                </dd>
              </div>
            </>
          )}
        </dl>
      </section>
    </main>
  )
}
