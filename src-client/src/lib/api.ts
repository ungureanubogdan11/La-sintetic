export interface PageMeta {
  page: number
  page_size: number
  total: number
  total_pages: number
}

export interface ApiSuccess<T> {
  data: T
  meta?: PageMeta
}

export interface ApiErrorBody {
  error: {
    code: string
    message: string
    details?: unknown
  }
}

export class ApiError extends Error {
  readonly status: number
  readonly code: string
  readonly details?: unknown

  constructor(status: number, code: string, message: string, details?: unknown) {
    super(message)
    this.name = 'ApiError'
    this.status = status
    this.code = code
    this.details = details
  }
}

async function request<T>(path: string, init?: RequestInit): Promise<ApiSuccess<T>> {
  const res = await fetch(path, {
    ...init,
    headers: { Accept: 'application/json', ...init?.headers },
  })

  if (res.status === 204) {
    return { data: undefined as T }
  }

  const body: unknown = await res.json().catch(() => null)

  if (!res.ok) {
    const err = (body as ApiErrorBody | null)?.error
    throw new ApiError(
      res.status,
      err?.code ?? 'unknown_error',
      err?.message ?? `request failed with status ${res.status}`,
      err?.details,
    )
  }
  return body as ApiSuccess<T>
}

export function apiGet<T>(path: string): Promise<ApiSuccess<T>> {
  return request<T>(path)
}
