import http from 'k6/http'
import { check, sleep } from 'k6'

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8080'

export function runCoffeeFlow() {
  const id = String(__VU * 10000 + __ITER + 1)

  const putRes = http.put(
    `${BASE_URL}/coffee/`,
    JSON.stringify({ id: id, name: 'espresso', calories: 150 }),
    { headers: { 'Content-Type': 'application/json' } }
  )
  check(putRes, { 'PUT 200': (r) => r.status === 200 })

  sleep(0.5)

  const getRes = http.get(`${BASE_URL}/coffee/?id=${id}`)
  check(getRes, { 'GET 200': (r) => r.status === 200 })

  sleep(1)
}
