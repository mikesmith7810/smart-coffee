import { runCoffeeFlow } from './lib/coffee.js'

// Single user, 10 iterations — quick sanity check
export const options = {
  vus: 1,
  iterations: 10,
  thresholds: {
    http_req_failed: ['rate<0.01'],
    http_req_duration: ['p(95)<500'],
  },
}

export default function () {
  runCoffeeFlow()
}
